package player

import (
	"strings"

	"github.com/spf13/cast"

	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

// 玩家离线
var _ = router.Reg(func(player *player.Player, msg *outer.GMReq) any {
	var (
		command string
		args    []string
	)

	strs := strings.Split(msg.Cmd, " ")

	if len(strs) < 1 {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	command = strs[0]
	for i := 1; i < len(strs); i++ {
		args = append(args, strs[i])
	}

	switch command {
	case "gold":
		if len(args) <= 0 {
			return outer.ERROR_MSG_REQ_PARAM_INVALID
		}

		val := cast.ToInt64(args[0])
		player.Role().AddGold(val)
	}
	return &outer.GMRsp{}
})
