### Structure
1) User send direct messages to `consumer` using GRPC or via web-proxy.
2) User receive message streams from `streamer` usding GRPC or via web-proxy
3) `consumer` resend all incoming messages to `streamer`


###

docker run -p 5672:5672 rabbitmq:alpine

grpcurl -plaintext localhost:55000 list
grpcurl -plaintext localhost:55001 list

grpcurl \
    -d '{ "message": "hello" }' \
    -plaintext \
    localhost:55000 \
    test.Chat/Send

###
grpcurl \
    -d '{}' \
    -plaintext \
    localhost:55001 \
    test.Chat/Listen

grpcurl \
    -plaintext \
    -d '{ "message": "hello", "receiverVersion": "rcv" }' \
    localhost:55002 \
    internal.Sender/NewMessage