package client

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/log"
	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
	"server/proto/message"
	"time"
)

var logger = log.New(log.TAG_DEBUG_I)

const addr = "127.0.0.1:9001"

type Client struct {
	actor.Base
	cli       network.INetClient
	msgParser *tools.ProtoParser
	UID       uint64
	RID       uint64
}

func (s *Client) OnInit() {
	s.cli = network.NewTcpClient(addr, func() network.ICodec { return &network.StreamCodec{} })
	s.cli.AddLast(func() network.INetHandler { return &SessionHandler{client: s} })
	expect.Nil(s.cli.Start(false))

	s.InitCmd()
	s.msgParser = tools.NewProtoParser().Init("message", "MSG")

	// 心跳
	s.AddTimer(tools.UUID(), 20*time.Second, func(dt int64) {
		s.SendToServer(message.MSG_PING.Int32(), &message.Ping{})
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
	case *message.LoginRsp:
		logger.KV("msg", msg.String()).Info("login success!")
		s.UID = msg.UID
		s.RID = msg.RID
	case *message.EnterGameRsp:
		logger.KV("msg", msg.String()).Info("enter success!")
	case *message.Pong:
		logger.Info("aliving~")
	}
}
