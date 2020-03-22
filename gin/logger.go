package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/scofieldpeng/golibs/log"
	"github.com/sirupsen/logrus"
	"time"
)

// Logger logger中间件
func Logger(noLogPath ...string) gin.HandlerFunc {
	if len(noLogPath) == 0 {
		noLogPath = make([]string, 0)
	}

	skipLogPathMap := make(map[string]bool)
	skipLogPathMap["/healthcheck"] = true
	for _, v := range noLogPath {
		skipLogPathMap[v] = true
	}

	return func(ctx *gin.Context) {
		// write json log
		requestTime := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.Query().Encode()

		ctx.Next()

		if _, ok := skipLogPathMap[path]; !ok {
			log.GetLogger().WithFields(logrus.Fields{
				"status_code":  ctx.Writer.Status(),
				"method":       ctx.Request.Method,
				"client_ip":    ctx.ClientIP(),
				"request_time": requestTime.Format("2006-01-02 15:04:05"),
				"latency":      time.Now().Sub(requestTime),
				"path":         path,
				"query":        query,
			}).Infoln("capture request: " + ctx.Request.URL.String())
		}
	}
}
