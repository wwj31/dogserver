package rdskey

import (
	"context"
	"encoding/json"
	"server/common/log"
	"server/common/rds"
	"server/proto/innermsg/inner"
)

func SetPlayerInfo(info *inner.PlayerInfo) {
	b, err := json.Marshal(info)
	if err != nil {
		log.Errorw("SetPlayerInfo json marshal failed", "err", err, "info", info.String())
		return
	}

	rds.Ins.Set(context.Background(), PlayerInfoKey(info.ShortId), string(b), 0)
}

func PlayerInfo(shortId int64) (info inner.PlayerInfo) {
	str, err := rds.Ins.Get(context.Background(), PlayerInfoKey(shortId)).Result()
	if err != nil {
		log.Errorw("PlayerInfo redis get failed", "err", err, "key", PlayerInfoKey(shortId))
		return
	}

	err = json.Unmarshal([]byte(str), &info)
	if err != nil {
		log.Errorw("PlayerInfo json unmarshal failed", "err", err, "key", PlayerInfoKey(shortId))
		return
	}
	return info
}
