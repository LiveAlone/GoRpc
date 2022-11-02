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

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7070))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	lib.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
