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

	translate "cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "./proto/receiver"
	pbStreamer "./proto/streamer"
)

var (
	port            = flag.Int("port", 55000, "API port")
	version         = flag.String("version", "receiver-1", "Service version")
	streamerAddress = flag.String("streamer", "localhost:55001", "Streamer endpoint")
	apiKey          = flag.String("api-key", "AIzaSyB9jxPQVweX0VSYdiiZnYwYcRvluO-P-a0", "Google translate API key")
)

type service struct {
	streamer   pbStreamer.StreamerClient
	translator *translate.Client
}

// redirect incoming message to the sender service
func (s *service) Publish(ctx context.Context, input *pb.Message) (*pb.Empty, error) {
	log.Printf("--->")
	log.Printf("message: %s", input.Text)

	translations, translateErr := s.translator.Translate(ctx,
		[]string{input.Text}, language.French, nil)
	failOnError(translateErr, "Failed to translate")
	log.Printf("translate: %s", translations[0].Text)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := s.streamer.Receive(ctx, &pbStreamer.Message{
		Text: translations[0].Text,
	})
	if err != nil {
		log.Printf("failed to send: %v", err)
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	flag.Parse()

	// init translate api
	ctx := context.Background()
	tClient, tErr := translate.NewClient(ctx, option.WithAPIKey(*apiKey))
	failOnError(tErr, "Failed to create google translate API")
	defer tClient.Close()

	// init connection to another service
	conn, grpcErr := grpc.Dial(*streamerAddress, grpc.WithInsecure())
	failOnError(grpcErr, "Failed to create grpc client")
	defer conn.Close()
	gClient := pbStreamer.NewStreamerClient(conn)

	// create local grpc service
	lis, listenErr := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	failOnError(listenErr, "Failed to listen")

	grpcServer := grpc.NewServer()
	pb.RegisterReceiverServer(grpcServer, &service{
		streamer:   gClient,
		translator: tClient,
	})
	reflection.Register(grpcServer)
	log.Printf("Starting API on port %d (version %s)", *port, *version)
	grpcServer.Serve(lis)
}
