package route

import (
	"net/http"

	"github.com/scofieldpeng/golibs/gin/response"
	"github.com/gin-gonic/gin"
)

// NoFound 为没有找到对应路由的handler
func NoFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, response.NewErrMsg(104, "not found"))
	return
}
