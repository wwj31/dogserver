package player

import (
	"context"
	"math/rand"
	"reflect"
	"strings"
	"time"

	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"server/common"
	"server/common/log"
	"server/common/mongodb"
	"server/proto/innermsg/inner"
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
		keepAlive   time.Time
	}
)

func (s *Player) OnInit() {
	s.sender = common.NewSendTools(s)

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
	message, msgName, gSession, err := common.UnwrapperGateMsg(msg.RawMsg())
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

	if msgName == "" {
		msgName = s.System().ProtoIndex().MsgName(message.(gogo.Message))
	}
	handle, ok := controller.GetHandler(msgName)
	if !ok {
		log.Errorw("player undefined route ", "msg", msgName)
		return
	}
	log.Debugw("player handle msg", "player", s.ID(), "msg", msgName)
	handle(s, message)

	pt, ok := message.(gogo.Message)
	if ok {
		msgName := s.System().ProtoIndex().MsgName(pt)
		outer.Put(msgName, message)
	}

	s.keepAlive = tools.Now()
}

func (s *Player) RID() string {
	return s.roleId
}

func (s *Player) Send2Client(pb gogo.Message) {
	if pb == nil || !s.Online() {
		return
	}
	if err := s.sender.Send2Client(s.gSession, pb); err != nil {
		log.Errorw("player send faild", "err", err)
	}
}

func (s *Player) Login(first bool) {
	log.Infow("player login", "id", s.roleId)
	for _, mod := range s.models {
		mod.OnLogin(first)
	}

	s.CancelTimer(s.exitTimerId)
}

func (s *Player) Logout() {
	log.Infow("player logout", "id", s.roleId)
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

			str := strings.Split(common.ProtoType(doc), ".")
			if len(str) < 2 {
				log.Errorw("msg name get failed", "type", reflect.TypeOf(doc).String())
				continue
			}

			if reflect.ValueOf(doc).IsNil() {
				log.Errorw("doc is nil interface{}", "type", str[1])
				continue
			}

			if _, err := mongodb.Ins.Collection(str[1]).UpdateByID(
				context.Background(),
				s.roleId,
				bson.M{"$set": doc},
				options.Update().SetUpsert(true)); err != nil {
				log.Errorw("player store failed", "collection", str[1], "err", err)
			}
		}
	}
	log.Infow("player stored model", "RID", s.roleId)
}

const aliveDuration = 5 * time.Second // 24*time.Hour
func (s Player) checkAlive() {
	duration := tools.Now().Sub(s.keepAlive)
	if duration > aliveDuration && !s.Online() {
		//s.Exit()
	}
}
