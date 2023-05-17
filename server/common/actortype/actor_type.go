package actortype

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

type ActorId = string

const (
	WorldActor   = "world"
	LoginActor   = "login"
	GameActor    = "game"
	PlayerActor  = "player"
	GatewayActor = "gateway"
	ChatActor    = "chat"
)

func GameName(id int32) ActorId {
	return fmt.Sprintf("%v_%v_Actor", GameActor, id)
}

func PlayerId(id string) ActorId {
	return fmt.Sprintf("%v_%v_Actor", PlayerActor, id)
}

func WorldName(id int32) ActorId {
	return fmt.Sprintf("%v_%v_Actor", WorldActor, id)
}

func GatewayName(id int32) ActorId {
	return fmt.Sprintf("%v_%v_Actor", GatewayActor, id)
}
func ChatName(id int32) ActorId {
	return fmt.Sprintf("%v_%v_Actor", ChatActor, id)
}

func RID(actorId ActorId) (str string) {
	s := strings.Split(actorId, "_")
	if len(s) == 3 {
		str = s[1]
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

func IsActorOf(actorId, typ string) bool {
	str := strings.Split(actorId, "_")
	if len(str) != 3 {
		return false
	}
	return typ == str[0]
}
