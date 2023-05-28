package door

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wwj31/dogactor/actor"
	"io"
	"net/http"
	"server/common/log"
	"strings"
)

type Door struct {
	actor.Base
	ginEngine *gin.Engine
}

func New() *Door {
	return &Door{}
}

var door *Door

func (s *Door) OnInit() {
	door = s
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
	alliance.POST("/setmaster", setMaster)

	go func() {
		log.Infow("gin startup ", "port", 9999)
		if err := s.ginEngine.Run(fmt.Sprintf(":%v", 9999)); err != nil {
			log.Errorw("gin startup failed ", "err", err)
		}
	}()
	log.Infow("door OnInit")
}

func (s *Door) OnStop() bool {
	log.Debugw("door stop", "id", s.ID())
	return true
}
