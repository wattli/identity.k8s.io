package uds

import (
	"context"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	i, err := handler(ctx, req)
	glog.Infof("%q: req: %+v ctx: %+v err: %v", info.FullMethod, req, ctx, err)
	return i, err
}

func recursiveInterceptor(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if len(interceptors) == 0 {
			return handler(ctx, req)
		}
		if len(interceptors) == 1 {
			return interceptors[0](ctx, req, info, handler)
		}
		return interceptors[0](ctx, req, info, func(ctx context.Context, req interface{}) (interface{}, error) {
			return recursiveInterceptor(interceptors[1:]...)(ctx, req, info, handler)
		})
	}
}
