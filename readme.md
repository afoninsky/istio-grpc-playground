### Structure
1) User send direct messages to `consumer` using GRPC or via web-proxy.
2) User receive message streams from `streamer` usding GRPC or via web-proxy
3) `consumer` resend all incoming messages to `streamer`


###

docker run -p 5672:5672 rabbitmq:alpine

grpcurl -plaintext localhost:55000 describe
grpcurl -plaintext localhost:55001 describe

grpcurl \
    -d '{ "text": "hello" }' \
    -plaintext \
    localhost:55000 \
    internal.receiver.Receiver/Publish

grpcurl \
    -plaintext \
    localhost:55001 \
    internal.streamer.Streamer/Subscribe