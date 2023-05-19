package door

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func setPosition(ctx *gin.Context) {
	req := gin.H{}
	_ = ctx.BindJSON(&req)
	shortId, position := cast.ToInt64(req["shortId"]), cast.ToString(req["position"])
	if shortId == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "shortId is empty"})
		return
	}

	rid := getRIDByShortId(cast.ToInt64(shortId))
	if rid == "" {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": fmt.Errorf("can not find rid by shortId:%v", shortId).Error()})
		return
	}

	_ = position
	//gSession := common.GateSession()
	//gSession.SendToClient()
}
