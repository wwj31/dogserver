package player

import (
	"context"
	"math/rand"
	"reflect"
	"server/common/router"
	"server/mgo"
	"strings"
	"time"

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

func New(roleId string, gamer iface.Gamer) *Player {
	p := &Player{
		roleId: roleId,
		gamer:  gamer,
	}
	return p
}

type (
	Player struct {
		actor.Base
		gamer    iface.Gamer
		gSession common.GSession       // 网络session
		models   [allmod]iface.Modeler // 玩家所有功能模块
		observer *common.Observer

		roleId      string
		saveTimerId string
		exitTimerId string
		keepAlive   time.Time
	}
)

func (s *Player) OnInit() {
	s.observer = common.NewObserver()

	// 初始化玩家所有功能模块
	s.initModule()
	s.load()

	// 定时回存
	s.storeTicker()
	log.Infow("player actor activated", "id", s.ID())
}

func (s *Player) OnStop() bool {
	s.store()
	return true
}

func (s *Player) OnHandle(msg actor.Message) {
	defer func() {
		s.keepAlive = tools.Now()
	}()

	message, msgName, gSession, err := common.UnwrappedGateMsg(msg.Payload())
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

	pt, ok := message.(gogo.Message)
	if !ok {
		log.Warnw("unknown msg", "msg", reflect.TypeOf(message).String())
		return
	}

	if msgName == "" {
		msgName = msg.GetMsgName()
	}
	log.Infow("input", "player", s.roleId, "msg", reflect.TypeOf(pt), "data", pt.String())
	router.Dispatch(s, pt)
}

func (s *Player) RID() string {
	return s.roleId
}

func (s *Player) Observer() *common.Observer {
	return s.observer
}

func (s *Player) Send2Client(pb gogo.Message) {
	log.Infow("output", "player", s.roleId, "online", s.Online(), "msg", reflect.TypeOf(pb), "data", pb.String())
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

	s.CancelTimer(s.exitTimerId)
}

func (s *Player) Logout() {
	log.Infow("player logout", "id", s.roleId)
	for _, mod := range s.models {
		mod.OnLogout()
	}
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
