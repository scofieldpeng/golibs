package interceptor

import (
	"context"
	"github.com/scofieldpeng/golibs/log"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

// LoggerInterceptor用来跟踪grpc server的每次的日志
func ServerLoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.GetLogger().WithFields(logrus.Fields{
			"method": info.FullMethod,
			"req":    req,
			"time":   time.Now().UnixNano(),
		}).Info("request receive")

		startTime := time.Now()

		resp, err = handler(ctx, req)

		log.GetLogger().WithFields(logrus.Fields{
			"method":     info.FullMethod,
			"req":        req,
			"resp":       resp,
			"time":       time.Now().UnixNano(),
			"spend_time": time.Now().Sub(startTime).Nanoseconds(),
		}).Info()

		return resp, err
	}
}

// ClientLoggerInterceptor 用来跟踪grpc client每次的请求日志
func ClientLoggerInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		log.GetLogger().WithFields(logrus.Fields{
			"method": method,
			"time":   time.Now().UnixNano(),
			"args":   req,
		}).Info("request begin")

		startTime := time.Now()

		err = invoker(ctx, method, req, reply, cc, opts...)
		log.GetLogger().WithFields(logrus.Fields{
			"method":     method,
			"req":        req,
			"resp":       reply,
			"time":       time.Now().UnixNano(),
			"spend_time": time.Now().Sub(startTime).Nanoseconds(),
		}).Info()

		return err
	}
}
