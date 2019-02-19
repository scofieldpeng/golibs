package grpc

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/scofieldpeng/golibs/grpc/interceptor"
	"google.golang.org/grpc"
)

// NewClient to create a new grpc.ClientConn object
func NewClient(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if len(opts) == 0 {
		opts = make([]grpc.DialOption, 0, 3)
	}

	opts = append(opts, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
		interceptor.ClientLoggerInterceptor(),
	)), grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
		interceptor.ClientStreamLoggerInterceptor(),
	)))

	return grpc.Dial(target, opts...)
}

// NewInSecureClient to create a new grpc.clientConn object,but with insecure mechanism
func NewInSecureClient(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if len(opts) == 0 {
		opts = make([]grpc.DialOption, 0, 3)
	}

	opts = append(opts, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
		interceptor.ClientLoggerInterceptor(),
	)), grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
		interceptor.ClientStreamLoggerInterceptor(),
	)))
	opts = append(opts, grpc.WithInsecure())

	return grpc.Dial(target, opts...)
}
