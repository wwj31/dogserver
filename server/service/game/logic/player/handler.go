package player

import (
	"github.com/wwj31/dogactor/log"
	"reflect"
	"server/proto/message"
	"server/service/game/iface"
)

var route = map[string]func(player iface.Player, msg interface{}){
	reflect.TypeOf(MsgLogin{}).String(): Login,
}

func (s *Player) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	name := reflect.TypeOf(msg).String()
	handler, ok := route[name]
	if !ok {
		log.KV("name", name).Error("player undefined route ")
		return
	}
	handler(s, msg)
}

func Login(player iface.Player, v interface{}) {
	msg := v.(MsgLogin)
	player.SetGateSession(msg.GSession)

	//// 新号处理
	if player.IsNewRole() {
		// todo ...
		player.Item().Add(map[int64]int64{123: 999})
	}

	player.Login()

	player.Send2Client(&message.LoginRsp{
		UID:     player.Role().UUId(),
		RID:     player.Role().RoleId(),
		Cryptic: "",
	})
}
