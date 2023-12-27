package shared

import (
	"context"
	"time"

	"github.com/wwj31/dogactor/actor"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/rds"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
)

// 激活玩家对象，如果玩家已经存在某个节点无需处理，
// 否则，将玩家在指定节点激活

func PullPlayer(sender actor.Messenger, targetRID string, gameNode ...string) (outer.ERROR, string) {
	defer rds.Lock("lock:pull:" + targetRID)() // 防止玩家同时被多个节点拉起，这里需要加锁

	var (
		dispatchGameNode string
		ctx              = context.Background()
		gameNodeKey      = rdsop.GameNodeKey(targetRID)
	)

	// 优先找玩家最近登录过的game节点,如果没找到用gameNode
	dispatchGameNode, _ = rds.Ins.Get(ctx, gameNodeKey).Result()
	if dispatchGameNode == "" {
		if len(gameNode) > 0 {
			dispatchGameNode = gameNode[0]
		} else {
			dispatchGameNode = actortype.GameName(1) // 没有传入指定节点，默认1节点
		}
	}

	v, pullErr := sender.RequestWait(dispatchGameNode, &inner.PullPlayer{RID: targetRID})
	if yes, errCode := common.IsErr(v, pullErr); yes {
		log.Warnw("pull player failed", "err code", errCode, "RID", targetRID, "gameNode", gameNode)
		return outer.ERROR_FAILED, dispatchGameNode
	}

	rds.Ins.Set(ctx, gameNodeKey, dispatchGameNode, 7*24*time.Hour)
	return outer.ERROR_OK, dispatchGameNode
}
