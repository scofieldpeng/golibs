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
			"method":  info.FullMethod,
			"req":     req,
			"resp":    resp,
			"time":    time.Now().UnixNano(),
			"process": time.Now().Sub(startTime).Nanoseconds(),
		}).Info("request processed")

		return resp, err
	}
}

// ServerStreamLoggerInterceptor server端的stream流日志记录
func ServerStreamLoggerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		beginTime := time.Now()
		log.GetLogger().WithFields(logrus.Fields{
			"srv":  srv,
			"info": info,
			"time": beginTime.UnixNano(),
		}).Info("stream request receive")
		err := handler(srv, ss)

		log.GetLogger().WithFields(logrus.Fields{
			"srv":     srv,
			"info":    info,
			"time":    beginTime.UnixNano(),
			"process": time.Now().Sub(beginTime).Nanoseconds(),
		}).Info("stream request finish")

		return err
	}
}

// ClientLoggerInterceptor 用来跟踪grpc client每次的请求日志
func ClientLoggerInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		log.GetLogger().WithFields(logrus.Fields{
			"method": method,
			"time":   time.Now().UnixNano(),
			"req":    req,
			"target": cc.Target(),
		}).Info("request begin[client]")

		startTime := time.Now()

		err = invoker(ctx, method, req, reply, cc, opts...)
		log.GetLogger().WithFields(logrus.Fields{
			"method":  method,
			"time":    time.Now().UnixNano(),
			"req":     req,
			"resp":    reply,
			"target":  cc.Target(),
			"process": time.Now().Sub(startTime).Nanoseconds(),
		}).Info("request finish[client]")

		return err
	}
}

func ClientStreamLoggerInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		startTime := time.Now()
		log.GetLogger().WithFields(logrus.Fields{
			"time":           startTime.UnixNano(),
			"method":         method,
			"streamName":     desc.StreamName,
			"isClientStream": desc.ClientStreams,
			"isServerStream": desc.ServerStreams,
			"target":         cc.Target(),
		}).Info("stream request begin[client]")

		cs, err := streamer(ctx, desc, cc, method, opts...)

		log.GetLogger().WithFields(logrus.Fields{
			"time":           startTime.UnixNano(),
			"method":         method,
			"streamName":     desc.StreamName,
			"isClientStream": desc.ClientStreams,
			"isServerStream": desc.ServerStreams,
			"target":         cc.Target(),
			"process":        time.Now().Sub(startTime).Nanoseconds(),
		}).Info("stream request finish[client]")

		return cs, err
	}
}
