package rdsop

import (
	"context"
	"encoding/base64"

	"github.com/go-redis/redis/v9"
	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"

	"server/common"
	"server/common/log"
	"server/common/rds"
)

type GmMsg struct {
	Msg  string
	Data string
}

func AddOfflineGMCmd(shortId int64, msg gogo.Message) {
	msgType := common.ProtoType(msg)
	bytes, _ := gogo.Marshal(msg)
	gm := GmMsg{
		Msg:  msgType,
		Data: base64.StdEncoding.EncodeToString(bytes),
	}

	jsonStr := common.JsonMarshal(gm)
	rds.Ins.RPush(context.Background(), GMCmdListKey(shortId), jsonStr)
}

func GetOfflineGMCmd(shortId int64, sys *actor.System) []gogo.Message {
	ctx := context.Background()
	key := GMCmdListKey(shortId)

	defer rds.Ins.Del(ctx, key)

	var result []gogo.Message
	for {
		item, err := rds.Ins.LPop(ctx, key).Result()
		if err == redis.Nil {
			return result
		}
		if err != nil {
			log.Errorw("GetOfflineGMCmd failed", "err", err, "shrotId", shortId)
			return result
		}

		var gm GmMsg
		common.JsonUnmarshal(item, &gm)
		v, ok := sys.ProtoIndex().FindMsgByName(gm.Msg)
		if ok {
			b, _ := base64.StdEncoding.DecodeString(gm.Data)
			if err := gogo.Unmarshal(b, v); err == nil {
				result = append(result, v)
			}
		}
	}

}
