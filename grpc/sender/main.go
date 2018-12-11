//go:generate protoc --proto_path=../../proto/chat --go_out=plugins=grpc:./proto-chat sender.proto
//go:generate protoc --proto_path=../../proto/internal --go_out=plugins=grpc:./proto-internal sender.proto
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbChat "./proto-chat"
	pbInternal "./proto-internal"
)

var (
	port         = flag.Int("port", 55001, "API port")
	internalPort = flag.Int("internal_port", 55002, "internal port")
	version      = flag.String("version", "v1", "Service version")
)

// internal API
type internal struct{}

func (s *internal) NewMessage(ctx context.Context, input *pbInternal.Message) (*pbInternal.EmptyResponse, error) {
	log.Printf("incoming message: %s", input.Message)
	return &pbInternal.EmptyResponse{}, nil
}

// external API
type chat struct{}

func (s *chat) Listen(_ *pbChat.Empty, stream pbChat.Chat_ListenServer) error {
	log.Printf("someone connected")
	for {
		res := &pbChat.Message{Message: "qwe", ReceiverVersion: "r1", SenderVersion: "s1"}
		log.Printf("echoed...")
		if err := stream.Send(res); err != nil {
			log.Printf("error: %v", err)
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}

func runService(wg *sync.WaitGroup, name string, server *grpc.Server, port int, version string) {
	defer wg.Done()
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting %s on port %d (version %s)", name, port, version)
	server.Serve(lis)
}

func main() {
	flag.Parse()

	chatServer := grpc.NewServer()
	pbChat.RegisterChatServer(chatServer, &chat{})
	reflection.Register(chatServer)

	internalServer := grpc.NewServer()
	pbInternal.RegisterSenderServer(internalServer, &internal{})
	reflection.Register(internalServer)

	var wg sync.WaitGroup
	wg.Add(1)
	go runService(&wg, "api", chatServer, *port, *version)
	go runService(&wg, "internal", internalServer, *internalPort, *version)
	wg.Wait()
}
