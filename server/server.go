package main

import (
	"context"
	"fmt"
	"github.com/LiveAlone/GoRpc/lib"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	lib.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *lib.HelloRequest) (*lib.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &lib.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *lib.HelloRequest) (*lib.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &lib.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7070))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opt := grpc.ChainUnaryInterceptor(firstServerInterceptor)
	s := grpc.NewServer(opt)
	lib.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func firstServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Printf("服务端一元过滤器1：开始")
	res, err := handler(ctx, req)
	log.Printf("服务端一元过滤器1：完成")
	return res, err
}
