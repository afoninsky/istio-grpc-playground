//go:generate protoc --proto_path=../../proto --go_out=plugins=grpc:./proto/receiver ../../proto/receiver.proto
//go:generate protoc --proto_path=../../proto --go_out=plugins=grpc:./proto/streamer ../../proto/streamer.proto
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

	pb "./proto/receiver"
	pbStreamer "./proto/streamer"
)

var (
	port            = flag.Int("port", 55000, "API port")
	version         = flag.String("version", "receiver-1", "Service version")
	streamerAddress = flag.String("streamer", "localhost:55001", "Streamer endpoint")
)

type service struct {
	client pbStreamer.StreamerClient
}

// redirect incoming message to the sender service
func (s *service) Publish(ctx context.Context, input *pb.Message) (*pb.Empty, error) {
	log.Printf("Publishing -> %s", input.Text)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := s.client.Receive(ctx, &pbStreamer.Message{
		Text: input.Text,
	})
	if err != nil {
		log.Printf("failed to send: %v", err)
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}

func main() {
	flag.Parse()

	// init connection to another service
	conn, err := grpc.Dial(*streamerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cant connect to %s: %v", *streamerAddress, err)
	}
	defer conn.Close()
	client := pbStreamer.NewStreamerClient(conn)

	// create local grpc service
	lis, listenErr := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if listenErr != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterReceiverServer(grpcServer, &service{
		client: client,
	})
	reflection.Register(grpcServer)
	log.Printf("Starting API on port %d (version %s)", *port, *version)
	grpcServer.Serve(lis)
}
