package route

import (
	"net/http"

	"github.com/scofieldpeng/golibs/gin/response"
	"github.com/gin-gonic/gin"
)

// NoMethod 没有找到对应路由的 method 的 handler
func NoMethod(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, response.NewErrMsg(104, "not found"))
}
