package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/LiveAlone/GoRpc/lib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:7070", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := lib.NewHelloWorldServiceClient(conn)

	// 单个消费
	//r, err := c.SayHello(context.Background(), &lib.HelloRequest{Name: *name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetMessage())

	// stream return
	//stream, err := c.SayHelloStreamReturn(context.Background(), &lib.HelloRequest{Name: *name})
	//for {
	//	r, err := stream.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalf("could not greet: %v", err)
	//		break
	//	}
	//	log.Printf("Greeting: %s", r.GetMessage())
	//}

	// stream 参数流数据
	stream, err := c.SayHelloStreamParam(context.Background())
	if err != nil {
		log.Printf("say hello param request error, %v", err)
	}
	for i := 0; i < 100; i++ {
		err = stream.Send(&lib.HelloRequest{Name: fmt.Sprintf("ids value is %d", i)})
		if err != nil {
			log.Printf("send err %v", err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("response recv error %v", err)
		return
	}
	log.Printf("reply content is %v", reply)
}
