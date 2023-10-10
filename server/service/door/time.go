package door

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/tools"
)

func setTime(ctx *gin.Context) {
	if door == nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": fmt.Errorf("door is not nil").Error()})
		return
	}

	req := gin.H{}

	var bDate, bDay, bHour, bMin, bSec bool
	if ctx.Request.Method == "GET" {
		req["date"], bDate = ctx.GetQuery("date")
		req["d"], bDay = ctx.GetQuery("d")
		req["h"], bHour = ctx.GetQuery("h")
		req["m"], bMin = ctx.GetQuery("m")
		req["s"], bSec = ctx.GetQuery("s")
	} else if ctx.Request.Method == "POST" {
		_ = ctx.BindJSON(&req)
	}
	_ = ctx.BindJSON(&req)

	var (
		err         error
		addDuration time.Duration
		now         = tools.Now()
	)
	switch {
	// 设置时间到某个值
	case bDate:
		var t time.Time
		t, err = time.ParseInLocation(tools.StdTimeSimpleLayout, cast.ToString(req["date"]), time.Local)
		if t.Before(now) {
			err = fmt.Errorf("设置的日期只能是将来某个时间 当前时间:%v", tools.Now().Local().String())
			break
		}
		addDuration = t.Sub(now)

	case bDay:
		addDuration = cast.ToDuration(req["d"]) * tools.Day
	case bHour:
		addDuration = cast.ToDuration(req["h"]) * time.Hour
	case bMin:
		addDuration = cast.ToDuration(req["m"]) * time.Minute
	case bSec:
		addDuration = cast.ToDuration(req["s"]) * time.Second
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tools.ModifyTimeOffset(int64(addDuration))

	ctx.JSON(http.StatusOK, gin.H{"info": "时间设置成功，当前时间:" + tools.Now().Local().String()})
}

func clearTime(ctx *gin.Context) {
	tools.TimeOffset = 0
	tools.ModifyTimeOffset(0)
	ctx.JSON(http.StatusOK, gin.H{"info": "时间重置:" + tools.Now().Local().String()})
}
