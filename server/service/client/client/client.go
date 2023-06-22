package client

import (
	"reflect"
	"sync/atomic"
	"time"

	"server/common/log"
	"server/proto/outermsg/outer"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
)

type Client struct {
	actor.Base
	Addr      string
	cli       *WsClient
	Reconnect int64
	UID       string
	RID       string
	NewPlayer bool
	mails     []*outer.Mail
	DeviceID  string
	Token     string
	Phone     string
	PWD       string
	UpShortId int64
	EnterGame atomic.Bool
	waiter    chan proto.Message
}

func (s *Client) OnInit() {
	s.cli = Dial(s.Addr, &SessionHandler{client: s}).Startup()

	s.login(1)

	// 心跳
	s.AddTimer(tools.XUID(), tools.Now().Add(20*time.Second), func(dt time.Duration) {
		s.SendToServer(outer.Msg_IdHeartReq.Int32(), &outer.HeartReq{})
	}, -1)

	if s.Reconnect != -1 {
		s.AddTimer(tools.XUID(), tools.Now().Add(time.Duration(s.Reconnect)*time.Millisecond), func(dt time.Duration) {
			s.cli.Close()
			time.Sleep(5 * time.Millisecond)
			s.cli = Dial(s.Addr, &SessionHandler{client: s}).Startup()
			s.login(1)
		}, -1)
	}
}
func (s *Client) Close() {
	s.cli.Close()
}

func (s *Client) Req(msgId outer.Msg, pb proto.Message) proto.Message {
	for !s.EnterGame.Load() {
		time.Sleep(200 * time.Millisecond)
	}

	s.waiter = make(chan proto.Message, 1)
	s.Send(s.ID(), func() {
		s.SendToServer(msgId.Int32(), pb)
	})
	return <-s.waiter
}

func (s *Client) SendToServer(msgId int32, pb proto.Message) {
	bytes, err := proto.Marshal(pb)
	expect.Nil(err)

	data, _ := proto.Marshal(&outer.Base{
		MsgId: msgId,
		Data:  bytes,
	})
	s.cli.SendMsg(data)
}

func (s *Client) OnHandle(m actor.Message) {
	switch msg := m.Payload().(type) {
	case *outer.HeartRsp:
		log.Infow("aliving~")
	case *outer.FailRsp:
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
		s.EnterGame.Store(true)
		//s.SendToServer(outer.Msg_IdAgentMembersReq.Int32(), &outer.AgentMembersReq{})
		//s.AddTimer(tools.XUID(), tools.Now().Add(3*time.Second), func(dt time.Duration) {
		//	s.cli.Close()
		//	s.cli = Dial(s.Addr, &SessionHandler{client: s})
		//	s.cli.Startup()
		//	s.login(2)
		//})
	case *outer.AllianceInfoNtf:
		log.Infow("alliance info notify!", "msg", msg.String())
	case *outer.RoomPlayerOnlineNtf:
		log.Infow("room player online!", "msg", msg.String())
	case *outer.RoomPlayerEnterNtf:
		log.Infow("room player enter!", "msg", msg.String())
	case *outer.RoomPlayerLeaveNtf:
		log.Infow("room player leave!", "msg", msg.String())
	case *outer.RoomPlayerReadyNtf:
		log.Infow("room player ready!", "msg", msg.String())
	default:
		pb := msg.(proto.Message)
		log.Infow("msg", "type", reflect.TypeOf(pb), "data", pb.String())
		s.waiter <- pb
	}
}
