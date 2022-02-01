package common

import (
	"fmt"
	"regexp"
	"strings"

	"server/common/log"
)

type ActorId = string

const (
	World_Actor   = "world"
	Login_Actor   = "login"
	Game_Actor    = "game"
	Player_Actor  = "player"
	GateWay_Actor = "gateway"

	Client = "client"
	Robot  = "robot"
)

func GameName(id int32) ActorId {
	return fmt.Sprintf("%v%v_Actor", Game_Actor, id)
}

func PlayerId(id uint64) ActorId {
	return fmt.Sprintf("%v%v_Actor", Player_Actor, id)
}

func WorldName(id int32) ActorId {
	return fmt.Sprintf("%v%v_Actor", World_Actor, id)
}

func GatewayName(id int32) ActorId {
	return fmt.Sprintf("%v%v_Actor", GateWay_Actor, id)
}

func AId(actorId ActorId, typ string) (str string) {
	t := strings.Trim(actorId, typ)
	s := strings.Split(t, "_")
	if len(s) == 2 {
		str = s[0]
	}
	return
}

// 匹配actor类型 按照固定格式匹配
func IsActorOf(actorId, typ string) bool {
	str := typ + "([0-9]+)_Actor"
	match, e := regexp.MatchString(str, actorId)
	if e != nil {
		log.Errorw("IsActorOf regexp error", "err", e)
		return false
	}
	return match
}
