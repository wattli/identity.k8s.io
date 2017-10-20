package uds

import (
	"context"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	glog.Infof("%q: %#v", info.FullMethod, req)
	return handler(ctx, req)
}

func ServiceAccountInterceptor(serviceAccount string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, "service-account", serviceAccount), req)
	}
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
