package door

import (
	"fmt"
	"net/http"
	"time"

	"server/common"
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

// 玩家上下分操作
func addGold(ctx *gin.Context) {
	req := gin.H{}

	if ctx.Request.Method == "GET" {
		req["shortId"], _ = ctx.GetQuery("shortId")
		req["gold"], _ = ctx.GetQuery("gold")
	} else if ctx.Request.Method == "POST" {
		_ = ctx.BindJSON(&req)
	}

	shortId := cast.ToInt64(req["shortId"])
	gold := cast.ToInt64(req["gold"])

	if shortId == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "shortId is empty"})
		return
	}

	if gold == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "the gold is zero"})
		return
	}

	playerInfo := rdsop.PlayerInfo(shortId)
	if playerInfo.RID == "" {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Errorf("找不到玩家:[%v]", shortId).Error()})
		return
	}

	if door == nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": fmt.Errorf("door is not nil").Error()})
		return
	}

	if playerInfo.RoomId != 0 {
		v, err := door.RequestWait(actortype.RoomName(playerInfo.RoomId), &inner.RoomCanSetGoldReq{ShortId: playerInfo.ShortId})
		if yes, _ := common.IsErr(v, err); yes {
			ctx.JSON(http.StatusBadRequest,
				gin.H{"error": fmt.Errorf("request room failed playerInfo:%v", playerInfo.String()).Error()})
			return
		}
		rsp := v.(*inner.RoomCanSetGoldRsp)
		if !rsp.Ok {
			ctx.JSON(http.StatusBadRequest,
				gin.H{"error": fmt.Errorf("the player's room does not allow to set gold:%v", playerInfo.String()).Error()})
			return
		}
	}

	gmMsg := &inner.ModifyGoldReq{Gold: gold}
	result, err := door.RequestWait(actortype.PlayerId(playerInfo.RID), gmMsg, time.Second)
	if err != nil {
		rdsop.AddOfflineGMCmd(playerInfo.ShortId, gmMsg)
		ctx.JSON(http.StatusOK, gin.H{"info": "玩家不在线，命令保存等待玩家上线执行"})
		return

	}

	rsp, ok := result.(*inner.ModifyGoldRsp)
	if !ok {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": "设置分失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"info": fmt.Sprintf("设置分数成功，玩家[%v]当前分数:[%v]", rsp.Info.ShortId, rsp.Info.Gold)})
}
