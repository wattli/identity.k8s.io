package main

import (
	"fmt"
	"net"
	"os"

	"github.com/golang/glog"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/health"
)

func newUnixServer(path string, register func(s *grpc.Server), interceptors ...grpc.UnaryServerInterceptor) (*UnixServer, error) {
	os.Remove(path)
	lis, err := net.Listen("unix", path)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(recursiveInterceptor(interceptors...)))
	//	healthapi.RegisterHealthServer(s, health.NewServer())
	register(s)
	return &UnixServer{lis: lis, server: s}, nil
}

type UnixServer struct {
	lis    net.Listener
	server *grpc.Server
}

func (uds *UnixServer) Serve() error {
	err := uds.server.Serve(uds.lis)
	if err != nil {
		glog.Errorf("server stopped: %v", err)
	}
	return err
}

func (uds *UnixServer) Stop() {
	uds.server.Stop()
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	glog.Infof("%q: %#v", info.FullMethod, req)
	return handler(ctx, req)
}

func serviceAccountInterceptor(serviceAccount string) grpc.UnaryServerInterceptor {
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
