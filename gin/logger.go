package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/scofieldpeng/golibs/log"
)

// Logger logger中间件
func Logger() gin.HandlerFunc {
	return gin.LoggerWithWriter(log.GetWriter())
}
