//go:generate protoc --proto_path=../../proto --go_out=plugins=grpc:./proto/streamer streamer.proto
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "./proto/streamer"
)

const keyName = "message"

var (
	port         = flag.Int("port", 55001, "API port")
	version      = flag.String("version", "send-v1", "Service version")
	amqpURL      = flag.String("amqp_url", "amqp://localhost:5672", "AMQP url")
	exchangeName = flag.String("exchange", "grpc-tests", "Exchange name")
)

type service struct {
	channel *amqp.Channel
}

// receive new message from "sender" service
// publish it into queue
func (s *service) Receive(ctx context.Context, input *pb.Message) (*pb.Empty, error) {
	// log.Printf("Message from %s: %s", input.ReceiverVersion, input.Message)
	err := s.channel.Publish(*exchangeName, keyName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(input.Text),
	})
	return &pb.Empty{}, err
}

// got connection from client
// subscribe it on messages from the queue
func (s *service) Subscribe(_ *pb.Empty, stream pb.Streamer_SubscribeServer) error {
	log.Printf("Client is connected")
	queue, queueErr := s.channel.QueueDeclare("", false, true, false, true, nil)
	failOnError(queueErr, "Failed to declare a queue")
	bindErr := s.channel.QueueBind(
		queue.Name,
		keyName,
		*exchangeName,
		false,
		nil,
	)
	failOnError(bindErr, "Failed to bind a queue")
	messages, err := s.channel.Consume(queue.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	for d := range messages {
		log.Printf("Received a message: %s", d.Body)
		res := &pb.Message{Text: string(d.Body)}
		if err := stream.Send(res); err != nil {
			log.Printf("Error while sending to the client: %v", err)
			break
		}
	}
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	flag.Parse()

	// connect to AMQP service
	amqpConnection, amqpErr := amqp.Dial(*amqpURL)
	failOnError(amqpErr, "Failed to connect to RabbitMQ")
	defer amqpConnection.Close()
	amqpChannel, chanErr := amqpConnection.Channel()
	// defer amqpChannel.Close()
	failOnError(chanErr, "Failed to open a channel")
	exchErr := amqpChannel.ExchangeDeclare(*exchangeName, "direct", false, true, false, true, nil)
	failOnError(exchErr, "Failed to declare the Exchange")

	// create GRPC server
	lis, listenErr := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if listenErr != nil {
		log.Fatalf("failed to listen: %v", listenErr)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStreamerServer(grpcServer, &service{
		channel: amqpChannel,
	})
	reflection.Register(grpcServer)
	log.Printf("Starting on port %d (version %s)", *port, *version)
	grpcServer.Serve(lis)
}
