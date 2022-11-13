package main

import (
	"context"
	"fmt"
	"github.com/LiveAlone/GoRpc/lib"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

const ServerPort = 7070

type server struct {
	lib.UnimplementedHelloWorldServiceServer
}

func (s *server) SayHello(ctx context.Context, in *lib.HelloRequest) (*lib.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &lib.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloStreamReturn(in *lib.HelloRequest, stream lib.HelloWorldService_SayHelloStreamReturnServer) error {
	log.Printf("Stream Return Received: %v", in.GetName())
	// stream 分批返回数据
	for i := 0; i < 10; i++ {
		err := stream.Send(&lib.HelloReply{
			Message: "Hello " + in.GetName() + " Stream Return",
		})
		if err != nil {
			log.Printf("stream return get err %v", err)
			return err
		}

	}
	return nil
}

func (s *server) SayHelloStreamParam(stream lib.HelloWorldService_SayHelloStreamParamServer) error {
	// 批量数据请求
	for i := 0; i < 5; i++ {
		rq, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("stream param get req %v", rq)
	}
	err := stream.SendAndClose(&lib.HelloReply{
		Message: "SayHelloStreamParam return",
	})
	if err != nil {
		log.Printf("params request error %v", err)
	}
	return err
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opt := grpc.ChainUnaryInterceptor(firstServerInterceptor)
	s := grpc.NewServer(opt)
	lib.RegisterHelloWorldServiceServer(s, &server{})

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
