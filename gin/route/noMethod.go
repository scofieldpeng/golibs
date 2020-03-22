package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scofieldpeng/golibs/gin/response"
)

// NoMethod 没有找到对应路由的 method 的 handler
func NoMethod(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, response.NewErrData(104, "not found"))
}
