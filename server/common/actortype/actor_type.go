package actortype

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

type ActorId = string

const (
	World_Actor   = "world"
	Login_Actor   = "login"
	Game_Actor    = "game"
	Player_Actor  = "player"
	GateWay_Actor = "gateway"
	Chat_Actor    = "chat"

	Client = "client"
	Robot  = "robot"
)

func GameName(id int32) ActorId {
	return fmt.Sprintf("%v_%v_Actor", Game_Actor, id)
}

func PlayerId(id string) ActorId {
	return fmt.Sprintf("%v_%v_Actor", Player_Actor, id)
}

func WorldName(id int32) ActorId {
	return fmt.Sprintf("%v_%v_Actor", World_Actor, id)
}

func GatewayName(id int32) ActorId {
	return fmt.Sprintf("%v_%v_Actor", GateWay_Actor, id)
}
func ChatName(id uint16) ActorId {
	return fmt.Sprintf("%v_%v_Actor", Chat_Actor, id)
}

func AId(actorId ActorId, typ string) (str string) {
	t := strings.Trim(actorId, typ)
	s := strings.Split(t, "_")
	if len(s) == 2 {
		str = s[0]
	}
	return
}

// NumAndType get actor's number and type of actor
func NumAndType(actorId ActorId) (int, string) {
	str := strings.Split(actorId, "_")
	if len(str) != 3 {
		return -1, ""
	}
	return cast.ToInt(str[1]), str[0]
}

// 匹配actor类型 按照固定格式匹配
func IsActorOf(actorId, typ string) bool {
	str := strings.Split(actorId, "_")
	if len(str) != 3 {
		return false
	}
	return typ == str[0]
}
