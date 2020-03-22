package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/scofieldpeng/golibs/gin/response"
	"github.com/scofieldpeng/golibs/gin/route"
)

func New() *gin.Engine {
	engine := gin.New()

	// 开启健康检查
	engine.Any(`/healthcheck`, func(ctx *gin.Context) {
		response.OK(ctx)
		return
	})

	engine.NoMethod(route.NoMethod)
	engine.NoRoute(route.NoFound)

	engine.Use(Logger())
	engine.Use(Recovery())

	return engine
}
