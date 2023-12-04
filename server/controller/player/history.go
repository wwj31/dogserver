package player

import (
	"sort"

	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
	"server/service/game/logic/player"
)

// 记录游戏
var _ = router.Reg(func(p *player.Player, msg *inner.GameHistoryInfoReq) any {
	p.Room().AddGamblingHistory(msg.Info)
	return nil
})

// 获取游戏记录
var _ = router.Reg(func(p *player.Player, msg *outer.RoomGamblingHistoryReq) any {
	history := p.Room().GamblingHistory()

	var list []*outer.HistoryInfo
	for _, v := range history {
		for _, info := range v.List {
			list = append(list, &outer.HistoryInfo{
				GameType:    info.GameType,
				RoomId:      info.RoomId,
				GameStartAt: info.GameStartAt,
				GameOverAt:  info.GameOverAt,
				WinGold:     info.WinGold,
			})
		}
	}

	sort.Slice(list, func(i, j int) bool { return list[i].GameOverAt > list[j].GameOverAt })
	return &outer.RoomGamblingHistoryRsp{List: list}
})

// 获取游戏记录
var _ = router.Reg(func(p *player.Player, msg *outer.RoomRecordingReq) any {
	recording := rdsop.GetRoomRecording(msg.RoomId, msg.GameStartAt)
	return &outer.RoomRecordingRsp{GameRecordData: &outer.Recording{
		GameStartAt: recording.GameStartAt,
		GameOverAt:  recording.GameOverAt,
		Room:        recording.Room,
		Messages:    recording.Messages,
	}}
})
