package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scofieldpeng/golibs/gin/response"
)

// NoFound 为没有找到对应路由的handler
func NoFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, response.NewErrData(104, "not found"))
	return
}
