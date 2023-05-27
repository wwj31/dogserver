package player

import (
	"context"
	"math/rand"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/golang/protobuf/proto"

	"server/common/actortype"
	"server/rdsop"

	"server/common/router"
	"server/mgo"

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

		currentMsg actor.Message
	}
)

func (p *Player) OnInit() {
	p.observer = common.NewObserver()

	// 初始化玩家所有功能模块
	p.initModule()
	p.load()

	// 定时回存
	p.storeTicker()

	router.Result(p, p.responseHandle)
	log.Infow("player actor OnInit", "id", p.ID())
}

// 所有消息,处理完统一返回流程
func (p *Player) responseHandle(resultMsg any) {
	msg, ok := resultMsg.(proto.Message)
	if !ok {
		return
	}

	// 网关消息，直接将消息转发给session, 其他服务消息，走内部通讯接口
	if actortype.IsActorOf(p.currentMsg.GetSourceId(), actortype.GatewayActor) {
		p.Send2Client(msg)
	} else {
		var err error
		if p.currentMsg.GetRequestId() != "" {
			err = p.Response(p.currentMsg.GetRequestId(), msg)
		} else {
			err = p.Send(p.currentMsg.GetSourceId(), msg)
		}

		if err != nil {
			log.Warnw("response to actor failed",
				"source", p.currentMsg.GetSourceId(), "msg name", p.currentMsg.GetMsgName())
		}
	}
}

func (p *Player) OnStop() bool {
	p.store()
	log.Infow("player OnStop", "rid", p.roleId)
	return true
}

func (p *Player) OnHandle(msg actor.Message) {
	message, _, gSession, err := common.UnwrappedGateMsg(msg.Payload())
	expect.Nil(err)

	p.currentMsg = msg

	// 重连的情况，除了EnterGame消息，其他都不处理
	if p.gSession != gSession && gSession != "" {
		if _, ok := message.(*outer.EnterGameReq); !ok {
			log.Warnw("recv message from the old session",
				"local session", p.gSession,
				"new session", gSession,
			)
			return
		}
		p.SetGateSession(gSession)
	}

	pt, ok := message.(proto.Message)
	if !ok {
		log.Warnw("unknown msg", "msg", reflect.TypeOf(message).String())
		return
	}

	log.Infow("input",
		"rid", p.roleId,
		"gSession", gSession,
		"msg", reflect.TypeOf(pt),
		"data", pt.String())
	router.Dispatch(p, pt)
}

func (p *Player) Send2Client(pb proto.Message) {
	log.Infow("output", "rid", p.roleId, "gSession", p.gSession, "online", p.Online(), "msg", reflect.TypeOf(pb), "data", pb.String())
	if pb == nil || !p.Online() {
		return
	}
	p.gSession.SendToClient(p, pb)
}

func (p *Player) RID() string                             { return p.roleId }
func (p *Player) ShortId() int64                          { return p.shortId }
func (p *Player) Observer() *common.Observer              { return p.observer }
func (p *Player) GateSession() common.GSession            { return p.gSession }
func (p *Player) SetGateSession(gSession common.GSession) { p.gSession = gSession }
func (p *Player) Online() bool                            { return p.GateSession().Valid() }

func (p *Player) Login(first bool, enterGameRsp *outer.EnterGameRsp) {
	for _, mod := range p.models {
		mod.OnLogin(first, enterGameRsp)
	}

	if first {
		p.store()
	}

	p.CancelTimer(p.exitTimerId)
	p.UpdateInfoToRedis()
}

func (p *Player) Logout() {
	for _, mod := range p.models {
		mod.OnLogout()
	}

	p.SetGateSession("")
	p.UpdateInfoToRedis()
	p.exitTimerId = p.AddTimer(tools.XUID(), tools.Now().Add(time.Minute), func(dt time.Duration) {
		p.Exit()
	})
}

func (p *Player) PlayerInfo() *inner.PlayerInfo {
	return &inner.PlayerInfo{
		RID:        p.roleId,
		ShortId:    p.shortId,
		Name:       p.Role().Name(),
		Icon:       p.Role().Icon(),
		Gender:     p.Role().Gender(),
		UpShortId:  p.Role().UpShortId(),
		AllianceId: p.Alliance().AllianceId(),
		Position:   p.Alliance().Position(),
		LoginAt:    tools.TimeFormat(p.Role().LoginAt()),
		LogoutAt:   tools.TimeFormat(p.Role().LogoutAt()),
		GSession:   p.gSession.String(),
	}
}
func (p *Player) UpdateInfoToRedis() {
	rdsop.SetPlayerInfo(p.PlayerInfo())
}

func (p *Player) load() {
	for _, mod := range p.models {
		doc := mod.Data()
		if doc != nil {
			if reflect.ValueOf(doc).IsNil() {
				log.Errorw("doc is nil interface", "type", reflect.TypeOf(doc).Name())
				continue
			}

			coll := mgo.GoGoCollectionType(doc)
			result := mongodb.Ins.Collection(coll).FindOne(context.Background(), bson.M{"_id": p.roleId})
			if result.Err() == mongo.ErrNoDocuments {
				if _, ok := doc.(*inner.RoleInfo); ok {
					// 新玩家直接跳过
					break
				}

				// 老玩家找不到新添加的表，不做处理
				continue
			} else if result.Err() != nil {
				log.Errorw("player store failed", "collection", coll, "err", result.Err())
				return
			}

			if err := result.Decode(doc); err != nil {
				log.Errorw("player store failed", "collection", coll, "err", result.Err())
				return
			}
		}
		mod.OnLoaded()
	}
}

func (p *Player) storeTicker() {
	randDur := func() time.Duration {
		return time.Duration(rand.Intn(int(30*time.Second))) + (30 * time.Second)
	}

	execAt := tools.Now().Add(randDur())
	p.saveTimerId = p.AddTimer(tools.XUID(), execAt, func(dt time.Duration) {
		p.store()
		if p.Online() {
			p.storeTicker()
		}
	})
}

func (p *Player) store() {
	for _, mod := range p.models {
		doc := mod.Data()
		if doc != nil {
			collType := mgo.GoGoCollectionType(doc)
			_, err := mongodb.Ins.Collection(collType).UpdateByID(context.Background(), p.roleId,
				bson.M{"$set": doc},
				options.Update().SetUpsert(true),
			)

			if err != nil {
				log.Errorw("store failed", "rid", p.roleId, "mod", reflect.TypeOf(mod))
			}
			//mgo.Store(collType, p.roleId, gogo.Clone(doc))
		}
	}
	log.Infow("player stored model", "rid", p.roleId)
}
