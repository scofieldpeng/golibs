package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// 成功时返回的code值
	SuccessCode = 0
)

// Forbidden 非法访问错误
func Forbidden(ctx *gin.Context, code int, errorMsg string, data ...interface{}) {
	ctx.JSON(http.StatusForbidden, NewErrData(code, errorMsg, data...))
}

// UnAuth 认证错误
func UnAuth(ctx *gin.Context, code int, errorMsg string, data ...interface{}) {
	ctx.JSON(http.StatusUnauthorized, NewErrData(code, errorMsg, data...))
}

// 操作成功
func OK(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, NewData(SuccessCode, map[string]interface{}{}, "ok"))
}

// 返回成功,第二个参数为返回的数据,如果msg需要传值,可以传到第三个参数,否则默认"ok"
func Success(ctx *gin.Context, data interface{}, msg ...string) {
	if len(msg) == 0 {
		msg = make([]string, 1)
		msg[0] = "ok"
	}
	ctx.JSON(http.StatusOK, NewData(SuccessCode, data, msg[0]))
}

// 返回错误,第二个参数为http code,第三个参数为gin.H的返回值,建议使用NewErrorData
func Fail(ctx *gin.Context, httpCode int, data gin.H) {
	ctx.JSON(httpCode, data)
}

// 新建数据结构体
func NewData(code int, data interface{}, msg string) gin.H {
	return gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	}
}

// 新建错误消息
func NewErrData(code int, msg string, data ...interface{}) map[string]interface{} {
	if len(data) == 0 {
		data = make([]interface{}, 1)
		data[0] = make(map[string]interface{})
	}

	return NewData(code, data[0], msg)
}
