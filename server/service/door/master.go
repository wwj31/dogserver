package door

import (
	"fmt"
	"net/http"
	"server/common/actortype"
	"server/proto/innermsg/inner"
	"server/rdsop"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// 设置成员职位
func setMaster(ctx *gin.Context) {
	req := gin.H{}
	_ = ctx.BindJSON(&req)
	shortId := cast.ToInt64(req["shortId"])
	if shortId == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "shortId is empty"})
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

	if _, ok := result.(*inner.Ok); !ok {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": "创建联盟失败"})
		return
	}
}
