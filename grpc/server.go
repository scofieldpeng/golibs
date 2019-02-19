package grpc

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/scofieldpeng/golibs/grpc/interceptor"
	"google.golang.org/grpc"
	"time"
)

// create a new grpc server
func NewServer(opt ...grpc.ServerOption) *grpc.Server {
	if len(opt) == 0 {
		opt = make([]grpc.ServerOption, 0, 2)
	}
	opt = append(opt, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_recovery.UnaryServerInterceptor(),
		interceptor.ServerLoggerInterceptor(),
	), ), grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_recovery.StreamServerInterceptor(),
		interceptor.ServerStreamLoggerInterceptor(),
	)))

	opt = append(opt, grpc.ConnectionTimeout(time.Second*time.Duration(10)))

	return grpc.NewServer(opt...)
}
