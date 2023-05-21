package player

import (
	"context"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"reflect"
	"server/common/actortype"
	"server/common/rdskey"
	"strings"
	"time"

	"server/common/router"
	"server/mgo"

	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"server/common"
	"server/common/log"
	"server/common/mongodb"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
)

func New(account *inner.Account, info *inner.LoginRoleInfo, gamer iface.Gamer) *Player {
	p := &Player{
		roleId:      info.RID,
		shortId:     info.ShortId,
		accountInfo: account,
		gamer:       gamer,
	}
	return p
}

type (
	Player struct {
		actor.Base
		gamer    iface.Gamer
		gSession common.GSession       // 网络session
		models   [allMod]iface.Modeler // 玩家所有功能模块
		observer *common.Observer

		roleId      string
		shortId     int64
		accountInfo *inner.Account

		saveTimerId string
		exitTimerId string
		keepAlive   time.Time

		currentMsg actor.Message
	}
)

func (s *Player) OnInit() {
	s.observer = common.NewObserver()

	// 初始化玩家所有功能模块
	s.initModule()
	s.load()

	// 定时回存
	s.storeTicker()

	router.Result(s, s.responseHandle)
	log.Infow("player actor activated", "id", s.ID())
}

// 所有消息,处理完统一返回流程
func (s *Player) responseHandle(resultMsg any) {
	msg, ok := resultMsg.(proto.Message)
	if !ok {
		return
	}

	// 网关消息，直接将消息转发给session, 其他服务消息，走内部通讯接口
	if actortype.IsActorOf(s.currentMsg.GetSourceId(), actortype.GatewayActor) {
		s.Send2Client(msg)
	} else {
		var err error
		if s.currentMsg.GetRequestId() != "" {
			err = s.Response(s.currentMsg.GetRequestId(), msg)
		} else {
			err = s.Send(s.currentMsg.GetSourceId(), msg)
		}

		if err != nil {
			log.Warnw("response to actor failed",
				"source", s.currentMsg.GetSourceId(), "msg name", s.currentMsg.GetMsgName())
		}
	}
}

func (s *Player) OnStop() bool {
	s.store()
	return true
}

func (s *Player) OnHandle(msg actor.Message) {
	s.currentMsg = msg
	defer func() {
		s.keepAlive = tools.Now()
	}()

	message, _, gSession, err := common.UnwrappedGateMsg(msg.Payload())
	expect.Nil(err)

	// 重连的情况，除了EnterGame消息，其他都不处理
	if s.gSession != gSession && gSession != "" {
		if _, ok := message.(*outer.EnterGameReq); !ok {
			log.Warnw("recv message from the old session",
				"local session", s.gSession,
				"new session", gSession,
			)
			return
		}
		s.SetGateSession(gSession)
	}

	pt, ok := message.(proto.Message)
	if !ok {
		log.Warnw("unknown msg", "msg", reflect.TypeOf(message).String())
		return
	}

	log.Infow("input", "rid", s.roleId, "msg", reflect.TypeOf(pt), "data", pt.String())
	router.Dispatch(s, pt)
}

func (s *Player) RID() string {
	return s.roleId
}
func (s *Player) ShortId() int64 {
	return s.shortId
}

func (s *Player) Observer() *common.Observer {
	return s.observer
}

func (s *Player) Send2Client(pb proto.Message) {
	log.Infow("output", "rid", s.roleId, "online", s.Online(), "msg", reflect.TypeOf(pb), "data", pb.String())
	if pb == nil || !s.Online() {
		return
	}
	s.gSession.SendToClient(s, pb)
}

func (s *Player) GateSession() common.GSession            { return s.gSession }
func (s *Player) SetGateSession(gSession common.GSession) { s.gSession = gSession }
func (s *Player) Online() bool                            { return s.Role().LoginAt().After(s.Role().LogoutAt()) }

func (s *Player) Login(first bool, enterGameRsp *outer.EnterGameRsp) {
	for _, mod := range s.models {
		mod.OnLogin(first, enterGameRsp)
	}

	if first {
		s.store()
	}

	s.CancelTimer(s.exitTimerId)
	s.UpdateInfoToRedis()
}

func (s *Player) Logout() {
	log.Infow("player logout", "id", s.roleId)

	for _, mod := range s.models {
		mod.OnLogout()
	}
	s.store()
	s.UpdateInfoToRedis()
	s.exitTimerId = tools.XUID()

	s.AddTimer(s.exitTimerId, tools.Now().Add(3*time.Second), func(dt time.Duration) {
		s.Exit()
	})
}

func (s *Player) UpdateInfoToRedis() {
	rdskey.SetPlayerInfo(&inner.PlayerInfo{
		RID:        s.roleId,
		ShortId:    s.shortId,
		Name:       s.Role().Name(),
		Icon:       s.Role().Icon(),
		Gender:     s.Role().Gender(),
		AllianceId: s.Alliance().AllianceId(),
		LoginAt:    tools.TimeFormat(s.Role().LoginAt()),
		LogoutAt:   tools.TimeFormat(s.Role().LogoutAt()),
		GSession:   s.gSession.String(),
	})
}

func (s *Player) load() {
	for _, mod := range s.models {
		doc := mod.Data()
		if doc != nil {
			if reflect.ValueOf(doc).IsNil() {
				log.Errorw("doc is nil interface", "type", reflect.TypeOf(doc).Name())
				continue
			}

			str := strings.Split(common.ProtoType(doc), ".")
			if len(str) < 2 {
				log.Errorw("msg name get failed", "type", reflect.TypeOf(doc).String())
				continue
			}

			result := mongodb.Ins.Collection(str[1]).FindOne(context.Background(), bson.M{"_id": s.roleId})
			if result.Err() == mongo.ErrNoDocuments {
				if _, ok := doc.(*inner.RoleInfo); ok {
					// 新玩家直接跳过
					break
				}

				// 老玩家找不到新添加的表，不做处理
				continue
			} else if result.Err() != nil {
				log.Errorw("player store failed", "collection", str[1], "err", result.Err())
				return
			}

			if err := result.Decode(doc); err != nil {
				log.Errorw("player store failed", "collection", str[1], "err", result.Err())
				return
			}
		}
		mod.OnLoaded()
	}
	log.Infow("player loaded model", "RID", s.roleId)
}

func (s *Player) storeTicker() {
	randDur := func() time.Duration {
		return time.Duration(rand.Intn(int(30*time.Second))) + (30 * time.Second)
	}

	execAt := tools.Now().Add(randDur())
	s.saveTimerId = s.AddTimer(tools.XUID(), execAt, func(dt time.Duration) {
		s.store()
		s.checkAlive()
		s.storeTicker()
	})
}

func (s *Player) store() {
	for _, mod := range s.models {
		doc := mod.Data()
		if doc != nil {
			collType := mgo.GoGoCollectionType(doc)
			mgo.Store(collType, s.roleId, gogo.Clone(doc))
		}
	}
	log.Infow("player stored model", "RID", s.roleId)
}

const aliveDuration = 5 * time.Second // 24*time.Hour
func (s *Player) checkAlive() {
	duration := tools.Now().Sub(s.keepAlive)
	if duration > aliveDuration && !s.Online() {
		//s.Exit()
	}
}
