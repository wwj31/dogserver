package door

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wwj31/dogactor/actor"
	"go.mongodb.org/mongo-driver/bson"

	"server/common/log"
	"server/common/mongodb"
	"server/service/login/account"
)

type Door struct {
	actor.Base
	ginEngine *gin.Engine
}

func New() *Door {
	return &Door{}
}

func (s *Door) OnInit() {
	s.ginEngine = gin.Default()
	gin.SetMode(gin.ReleaseMode)

	s.ginEngine.Use(func(ctx *gin.Context) {
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		log.Infow("request body", "path", ctx.Request.URL.Path, "data", string(body))

		// 重新设置请求体内容，因为之前读取了一次，需要再放回去
		ctx.Request.Body = io.NopCloser(strings.NewReader(string(body)))
		ctx.Next()
	})

	mail := s.ginEngine.Group("/health")
	mail.GET("/alive", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	alliance := s.ginEngine.Group("/alliance")
	alliance.Use(checkToken)
	alliance.POST("/position", setPosition)

	go func() {
		log.Infow("gin startup ", "port", 9999)
		if err := s.ginEngine.Run(fmt.Sprintf(":%v", 9999)); err != nil {
			log.Errorw("gin startup failed ", "err", err)
		}
	}()
	log.Infow("door OnInit")
}

func getRIDByShortId(shortId int64) string {
	cur, err := mongodb.Ins.Collection(account.Collection).Aggregate(context.Background(), []bson.M{
		{
			"$match": bson.M{account.LastShortId: shortId},
		},
		{
			"$project": bson.M{
				"_id":                1,
				account.LastLoginRId: 1,
			},
		},
	})

	defer cur.Close(context.Background())
	if err != nil {
		log.Errorf("getRIDByShortId aggregate failed", "err", err)
		return ""
	}

	if cur.Next(context.Background()) {
		acc := &account.Account{}
		if err := cur.Decode(acc); err != nil {
			log.Errorw("getRIDByShortId decode failed", "shortId", shortId, "err", err)
			return ""
		}
		return acc.LastLoginRID
	}

	log.Warnw("can not find rid ", "shortid", shortId)
	return ""
}

func (s *Door) OnStop() bool {
	log.Debugw("door stop", "id", s.ID())
	return true
}

func (s *Door) OnHandle(m actor.Message) {
	//payload := m.Payload()
	//v, _, gSession, err := common.UnwrappedGateMsg(payload)
	//expect.Nil(err)
	//
	//switch msg := v.(type) {
	//case *outer.LoginReq:
	//	err = s.LoginReq(m.GetSourceId(), gSession, msg)
	//default:
	//	err = fmt.Errorf("undefined localmsg type %v", msg)
	//}
	//
	//if err != nil {
	//	log.Errorw("handle outer error", "err", err)
	//}
}
