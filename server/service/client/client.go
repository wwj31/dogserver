package client

import (
	"reflect"
	"server/proto/outermsg/outer"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
)

const addr = "127.0.0.1:9001"

type Client struct {
	actor.Base
	cli       network.Client
	UID       uint64
	RID       uint64
	NewPlayer bool
	mails     []*outer.Mail
	ACC       string
}

func (s *Client) OnInit() {
	s.cli = network.NewTcpClient(addr, func() network.DecodeEncoder { return &network.StreamCode{} })
	s.cli.AddLast(func() network.NetSessionHandler { return &SessionHandler{client: s} })
	expect.Nil(s.cli.Start(false))

	if s.ACC != "" {
		s.login(s.ACC)
		s.listMail("0")
		s.item("123", "1")
	} else {
		s.InitCmd()
	}

	// 心跳
	s.AddTimer(tools.XUID(), tools.Now().Add(20*time.Second), func(dt time.Duration) {
		s.SendToServer(outer.MSG_PING.Int32(), &outer.Ping{})
	}, -1)
}

func (s *Client) SendToServer(msgId int32, pb proto.Message) {
	bytes, err := proto.Marshal(pb)
	expect.Nil(err)

	data := network.CombineMsgWithId(msgId, bytes)
	err = s.cli.SendMsg(data)
	expect.Nil(err)
}

func (s *Client) OnHandleMessage(sourceId, targetId string, v interface{}) {
	switch msg := v.(type) {
	case *outer.Pong:
		logger.Infow("aliving~")
	case *outer.Fail:
		logger.Infow("msg respones fail", "err:", msg.String())
	// 登录
	case *outer.LoginResp:
		logger.Infow("login success!", "msg", msg.String())
		s.UID = msg.UID
		s.RID = msg.RID
		s.NewPlayer = msg.NewPlayer
		s.enter()
	case *outer.EnterGameResp:
		logger.Infow("EnterGameResp!", "msg", msg.String())
	case *outer.RoleInfoPush:
		logger.Infow("RoleInfoPush!", "msg", msg.String())
	case *outer.ItemInfoPush:
		logger.Infow("ItemInfoPush!", "msg", msg.String())

	// 道具
	case *outer.UseItemResp:
		logger.Infow("UseItemResp!", "msg", msg.String())
	case *outer.ItemChangeNotify:
		logger.Infow("ItemChangeNotify!", "msg", msg.String())

	// 邮件
	case *outer.MailListResp:
		s.mails = append(s.mails, msg.Mails...)
		logger.Infow("MailListResp!", "msg", msg.String())
	case *outer.ReadMailResp:
		logger.Infow("ReadMailResp!", "msg", msg.String())
	case *outer.ReceiveMailItemResp:
		logger.Infow("ReceiveMailItemResp!", "msg", msg.String())

		// 聊天
	case *outer.ChatNotify:
		logger.Infow("ChatNotify!", "msg", msg.String())

	default:
		logger.Infow("unknown type!", "type", reflect.TypeOf(msg).String(), "msg", msg)
	}
}
