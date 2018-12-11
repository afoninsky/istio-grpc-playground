//go:generate protoc --proto_path=../../proto/chat --go_out=plugins=grpc:./proto-chat receiver.proto
//go:generate protoc --proto_path=../../proto/internal --go_out=plugins=grpc:./proto-internal sender.proto
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbChat "./proto-chat"
	pbInternal "./proto-internal"
)

var (
	port          = flag.Int("port", 55000, "API port")
	version       = flag.String("version", "rcv-v1", "Service version")
	senderAddress = flag.String("sender_address", "localhost:55002", "Sender address")
)

type chat struct {
	client  pbInternal.SenderClient
	version string
}

// redirect incoming message to the sender service
func (c *chat) Send(ctx context.Context, input *pbChat.Message) (*pbChat.EmptyResponse, error) {
	log.Printf("sending to the receiver -> %s", input.Message)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.client.NewMessage(ctx, &pbInternal.Message{
		Message:         input.Message,
		ReceiverVersion: c.version,
	})
	if err != nil {
		log.Printf("failed to send: %v", err)
	}
	return &pbChat.EmptyResponse{}, err
}

func main() {
	flag.Parse()

	// init connection to another service
	conn, err := grpc.Dial(*senderAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbInternal.NewSenderClient(conn)

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// _, sendErr := client.NewMessage(ctx, &pbInternal.Message{
	// 	Message:         "qwe",
	// 	ReceiverVersion: "asd",
	// })
	// if sendErr != nil {
	// 	log.Fatalf("failed to send: %v", sendErr)
	// }

	// create local grpc service
	lis, listenErr := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if listenErr != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pbChat.RegisterChatServer(grpcServer, &chat{
		client:  client,
		version: *version,
	})
	reflection.Register(grpcServer)
	log.Printf("Starting API on port %d (version %s)", *port, *version)
	grpcServer.Serve(lis)
}
