package player

import (
	"context"
	"github.com/wwj31/dogactor/expect"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/tools"
	"go.mongodb.org/mongo-driver/bson"

	"server/common"
	"server/common/log"
	"server/common/mongodb"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
	"server/service/game/logic/player/controller"
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
		gSession common.GSession // 网络session
		sender   common.SendTools
		models   [allmod]iface.Modeler // 玩家所有功能模块

		roleId      string
		saveTimerId string
		exitTimerId string
		keepAlive   int64
	}
)

func (s *Player) OnInit() {
	s.sender = common.NewSendTools(s)

	// 初始化玩家所有功能模块
	s.initModule()
	s.load()

	// 定时回存
	randTime := tools.Now().Add(time.Second)
	s.saveTimerId = s.AddTimer(tools.XUID(), randTime, func(dt time.Duration) {
		s.store()
		s.checkAlive()
	}, -1)
	log.Infow("player actor activated", "id", s.ID())
}

func (s *Player) OnHandle(msg actor.Message) {
	message, msgName, gSession, err := common.UnwrapperGateMsg(msg.RawMsg())
	expect.Nil(err)

	handle, ok := controller.GetHandler(msgName)
	if !ok {
		log.Errorw("player undefined route ", "name", msgName)
		return
	}

	// 重连的情况，除了EnterGame消息，其他都不处理
	if s.gSession != gSession {
		if _, ok := message.(outer.EnterGameReq); !ok {
			log.Warnw("recv message of old session:%v new session:%v", s.gSession, gSession)
			return
		}
		s.SetGateSession(gSession)
	}

	handle(s, message)

	pt, ok := message.(proto.Message)
	if ok {
		msgName := s.System().ProtoIndex().MsgName(pt)
		outer.Put(msgName, message)
	}

	s.keepAlive = tools.NowTime()
}

func (s *Player) Send2Client(pb proto.Message) {
	if pb == nil || !s.Online() {
		return
	}
	if err := s.sender.Send2Client(s.gSession, pb); err != nil {
		log.Errorw("player send faild", "err", err)
	}
}

func (s *Player) Login(first bool) {
	for _, mod := range s.models {
		mod.OnLogin(first)
	}

	s.CancelTimer(s.exitTimerId)
}

func (s *Player) Logout() {
	for _, mod := range s.models {
		mod.OnLogout()
	}
}

func (s *Player) GateSession() common.GSession            { return s.gSession }
func (s *Player) SetGateSession(gSession common.GSession) { s.gSession = gSession }
func (s *Player) Online() bool                            { return s.Role().LoginAt().After(s.Role().LogoutAt()) }

func (s *Player) load() {
	for _, mod := range s.models {
		doc := mod.Data()
		if doc != nil {
			str := strings.Split(controller.MsgName(doc), ".")
			if len(str) < 2 {
				log.Errorw("msg name get failed", "type", reflect.TypeOf(doc).String())
				continue
			}

			result := mongodb.Ins.Collection(str[1]).FindOne(context.Background(), bson.M{"_id": s.roleId})
			if result.Err() != nil {
				log.Errorw("player store failed", "collection", str[1], "doc", doc.String())
				return
			}

			if err := result.Decode(doc); err != nil {
				log.Errorw("player store failed", "collection", str[1], "doc", doc.String())
				return
			}
		}
		mod.OnLoaded()
	}
	log.Infow("player loaded model", "RID", s.roleId)
}

func (s *Player) store() {
	for _, mod := range s.models {
		doc := mod.Data()
		if doc != nil {
			str := strings.Split(controller.MsgName(doc), ".")
			if len(str) < 2 {
				log.Errorw("msg name get failed", "type", reflect.TypeOf(doc).String())
				continue
			}

			if _, err := mongodb.Ins.Collection(str[1]).UpdateByID(
				context.Background(),
				s.roleId,
				doc,
				options.Update().SetUpsert(true)); err != nil {
				log.Errorw("player store failed", "collection", str[1], "doc", doc.String())
			}
		}
	}
	log.Infow("player stored model", "RID", s.roleId)
}

const aliveDuration = 5 * time.Second // 24*time.Hour
func (s Player) checkAlive() {
	now := tools.NowTime()
	duration := now - s.keepAlive
	if duration > int64(aliveDuration) && !s.Online() {
		//s.Exit()
	}
}
