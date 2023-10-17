package player

import (
	"server/common/log"
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/rdsop"
	"server/service/game/logic/player"
)

// 获取金币变动记录
var _ = router.Reg(func(p *player.Player, msg *outer.GoldRecordsReq) any {
	if msg.StartIndex < 0 || msg.StartIndex > msg.EndIndex {
		log.Warnw("GoldRecordsReq param err", "msg", msg.String())
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	records, totalLen := rdsop.GetUpdateGoldRecord(p.Role().ShortId(), msg.StartIndex, msg.EndIndex)
	var pbRecords []*outer.GoldRecords
	for _, record := range records {
		pbRecords = append(pbRecords, &outer.GoldRecords{
			GoldUpdateType: record.Type.ToPB(),
			Gold:           record.Gold,
			UpShortId:      record.UpShortId,
			DownShortId:    record.DownShortId,
			GameType:       record.GameType,
			OccurAt:        record.OccurAt.UnixMilli(),
		})
	}

	return &outer.GoldRecordsRsp{
		TotalLen: totalLen,
		Records:  pbRecords,
	}
})
