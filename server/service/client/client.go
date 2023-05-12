package client

import (
	"reflect"
	"time"

	"server/common/log"
	"server/proto/outermsg/outer"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
)

// const addr = "ws://1.14.17.15:7001/"
const addr = "ws://localhost:7001/"

type Client struct {
	actor.Base
	cli       *WsClient
	UID       string
	RID       string
	NewPlayer bool
	mails     []*outer.Mail
	DeviceID  string
	Token     string
}

func (s *Client) OnInit() {
	s.cli = Dial(addr, &SessionHandler{client: s})
	s.cli.Startup()

	if s.DeviceID != "" {
		s.login()
	}

	// 心跳
	s.AddTimer(tools.XUID(), tools.Now().Add(20*time.Second), func(dt time.Duration) {
		s.SendToServer(outer.Msg_IdHeartReq.Int32(), &outer.HeartReq{})
	}, -1)
}

func (s *Client) SendToServer(msgId int32, pb proto.Message) {
	bytes, err := proto.Marshal(pb)
	expect.Nil(err)

	data := network.CombineMsgWithId(msgId, bytes)
	s.cli.SendMsg(data)
}

func (s *Client) OnHandle(m actor.Message) {
	switch msg := m.Payload().(type) {
	case *outer.HeartRsp:
		log.Infow("aliving~")
	case *outer.Fail:
		log.Infow("msg respones fail", "err:", msg.String())
	// 登录
	case *outer.LoginRsp:
		log.Infow("login success!", "msg", msg.String())
		s.RID = msg.RID
		s.NewPlayer = msg.NewPlayer
		s.Token = msg.Token
		s.enter()
	case *outer.EnterGameRsp:
		log.Infow("EnterGameRsp!", "msg", msg.String())
		s.cli.Close()

		s.AddTimer(tools.XUID(), tools.Now().Add(3*time.Second), func(dt time.Duration) {
			s.cli = Dial(addr, &SessionHandler{client: s})
			s.cli.Startup()

			if s.DeviceID != "" {
				s.login()
			}
		})
	case *outer.RoleInfo:
		log.Infow("RoleInfoPush!", "msg", msg.String())

	// 邮件
	case *outer.MailListRsp:
		s.mails = append(s.mails, msg.Mails...)
		log.Infow("MailListRsp!", "msg", msg.String())
	case *outer.ReadMailRsp:
		log.Infow("ReadMailRsp!", "msg", msg.String())
	case *outer.ReceiveMailItemRsp:
		log.Infow("ReceiveMailItemRsp!", "msg", msg.String())

	default:
		log.Infow("unknown type!", "type", reflect.TypeOf(msg).String(), "msg", msg)
	}
}
