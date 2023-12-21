package door

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"server/common/actortype"
	"server/proto/innermsg/inner"
	"server/rdsop"
)

// 玩家列表
func playerList(ctx *gin.Context) {
	req := gin.H{}

	if ctx.Request.Method == "GET" {
		req["shortId"], _ = ctx.GetQuery("shortId")
	} else if ctx.Request.Method == "POST" {
		_ = ctx.BindJSON(&req)
	}

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

	if playerInfo.AllianceId != 0 {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Errorf("玩家:[%v]已有联盟", shortId).Error()})
		return
	}

	if playerInfo.UpShortId != 0 {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Errorf("玩家:[%v]已有上级:[%v]",
				shortId, playerInfo.UpShortId).Error()})
		return
	}

	if door == nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": fmt.Errorf("door is not nil").Error()})
		return
	}

	result, err := door.RequestWait(actortype.AllianceMgrName(), &inner.CreateAllianceReq{
		MasterShortId: playerInfo.ShortId,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	if _, ok := result.(*inner.CreateAllianceRsp); !ok {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": "创建联盟失败"})
		return
	}

	// 设置盟主返利点位
	rdsop.SetRebateInfoByDoor(shortId, rebatePoint)

	ctx.JSON(http.StatusOK, gin.H{"info": "联盟创建成功"})
}
