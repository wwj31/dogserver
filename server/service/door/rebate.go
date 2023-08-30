package door

import (
	"fmt"
	"net/http"

	"server/rdsop"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// 设置成员返利点
func setRebate(ctx *gin.Context) {
	req := gin.H{}
	_ = ctx.BindJSON(&req)
	shortId := cast.ToInt64(req["shortId"])
	if shortId == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "shortId is empty"})
		return
	}

	rebatePoint := cast.ToInt32(req["rebate"])
	if rebatePoint == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "rebate is zero"})
		return
	}

	playerInfo := rdsop.PlayerInfo(shortId)
	if playerInfo.RID == "" {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Errorf("找不到玩家:[%v]", shortId).Error()})
		return
	}

	if playerInfo.AllianceId == 0 {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Errorf("玩家:[%v]没有联盟", shortId).Error()})
		return
	}

	if door == nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": fmt.Errorf("door is not nil").Error()})
		return
	}

	// 设置盟主返利点位
	info := rdsop.SetRebateInfoByDoor(shortId, rebatePoint)
	str := fmt.Sprintf("玩家返利点位:%v 所有下级点位：%v ", info.Point, info.DownPoints)

	ctx.JSON(http.StatusOK, gin.H{"info": str})
}
