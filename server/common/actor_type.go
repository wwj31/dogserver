package common

import (
	"fmt"
	"regexp"

	"github.com/wwj31/dogactor/log"
)

type ActorId = string

// 静态Actor类型,全局唯一
const (
	Login_Actor  = "login"
	Center_Actor = "centerWorld"
	Client       = "client"
)

// 动态增删的actor,会有多个 game1、game2
const (
	Game_Actor    = "game"
	Player_Actor  = "player"
	GateWay_Actor = "gateway"
	World_Actor   = "world"
	DB_Actor      = "dB"
	MYSQL_Actor   = "mysql"
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

// 匹配actor类型 按照固定格式匹配
func IsActorOf(actorId, typ string) bool {
	str := typ + "([0-9]+)_Actor"
	match, e := regexp.MatchString(str, actorId)
	if e != nil {
		log.KV("actorId", actorId).ErrorStack(3, "error")
		return false
	}
	return match
}
