package door

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"server/common/log"
)

// http接口签名流程：
// 1.获取当前系统UTC时间戳 timestamp (秒)
// 2.将timestamp转换成字符串和secretKey拼接在一起获得sum    sum=timestamp+secretKey
// 3.token=md5(sum)
// 4.将timestamp插入http头部X-Time中
// 5.将token插入http头部X-Signature中
// 封装以上步骤，每次http请求执行

const secretKey = "62f1onGbkK8iIQBvaC4n*2#oK4rwj&aOa5"
const timeDiffHour = 2 * time.Hour

func checkToken(ctx *gin.Context) {
	clientTime := ctx.GetHeader("X-Time")           // 当前utc时间，(秒)
	clientSignature := ctx.GetHeader("X-Signature") // 签名
	if clientTime == "" || clientSignature == "" {
		info := gin.H{
			"error":     "time or signature is empty",
			"time":      clientTime,
			"signature": clientSignature,
		}
		log.Warnf("sign failed info:%v", info)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, info)
		return
	}
	//clientSignature := ctx.GetHeader("Authorization")

	// 获取请求体内容
	//body, err := io.ReadAll(ctx.Request.Body)
	//if err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	//	return
	//}
	//
	//// 重新设置请求体内容，因为之前读取了一次，需要再放回去
	//ctx.Request.Body = io.NopCloser(strings.NewReader(string(body)))

	clientNow := time.Unix(cast.ToInt64(clientTime), 0)
	now := time.Now()
	// 请求者的时间，不在当前时间，前后2小时内，视为无效
	if clientNow.Before(now) && now.Sub(clientNow) > timeDiffHour ||
		clientNow.After(now) && clientNow.Sub(now) > timeDiffHour {
		log.Warnf("sign time invalid client now:%v  now:%v", clientNow, now)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid client time"})
		return
	}

	// 计算签名
	calculatedSignature := calculateSignature(clientTime, secretKey)

	// 验证签名
	if clientSignature != calculatedSignature {
		log.Warnf("sign failed client signature:%v signature:%v", clientSignature, calculatedSignature)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
		return
	}

	// 继续处理请求
	ctx.Next()
}

func calculateSignature(time, secretKey string) string {
	str := time + secretKey
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}
