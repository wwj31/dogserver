package client

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/log"
	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
	"server/proto/message"
)

var logger = log.New(log.TAG_DEBUG_I)

const addr = "127.0.0.1:9001"

type Client struct {
	actor.Base
	cli       network.INetClient
	msgParser *tools.ProtoParser
}

func (s *Client) OnInit() {
	s.cli = network.NewTcpClient(addr, func() network.ICodec { return &network.StreamCodec{} })
	s.cli.AddLast(func() network.INetHandler { return &SessionHandler{client: s} })
	expect.Nil(s.cli.Start(false))

	s.msgParser = tools.NewProtoParser().Init("message", "MSG")
}

func (s *Client) SendToServer(pb proto.Message) {
	data, err := proto.Marshal(pb)
	expect.Nil(err)

	err = s.cli.SendMsg(data)
	expect.Nil(err)
}

func (s *Client) OnHandleMessage(sourceId, targetId string, v interface{}) {
	switch msg := v.(type) {
	case message.LoginRsp:
		logger.KV("msg", msg.String()).Info("login success")
	}
}
