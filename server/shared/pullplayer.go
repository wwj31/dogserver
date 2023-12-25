package shared

import (
	"context"

	"github.com/wwj31/dogactor/actor"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/rds"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
)

// 确保玩家处于激活状态

func PullPlayer(sender actor.Messenger, targetShortId int64, targetRID string, defaultGameId ...string) (outer.ERROR, string) {
	// 确保玩家处于激活状态
	// 找玩家最近登录过的game节点，如果没找到优先用传入的默认节点，否则统一交给1节点处理
	var dispatchGameId string
	gameNodeId, _ := rds.Ins.Get(context.Background(), rdsop.GameNodeKey(targetShortId)).Result()
	if gameNodeId != "" {
		dispatchGameId = gameNodeId
	} else {
		if len(defaultGameId) > 0 {
			dispatchGameId = defaultGameId[0]
		} else {
			dispatchGameId = actortype.GameName(1)
		}
	}

	v, pullErr := sender.RequestWait(dispatchGameId, &inner.PullPlayer{RID: targetRID})
	if yes, errCode := common.IsErr(v, pullErr); yes {
		log.Warnw("SetScoreForDownReq pull player failed", "err code", errCode)
		return outer.ERROR_FAILED, dispatchGameId
	}
	return outer.ERROR_OK, dispatchGameId
}
