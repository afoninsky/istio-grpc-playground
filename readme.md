grpcurl -plaintext localhost:55000 list
grpcurl -plaintext localhost:55001 list

grpcurl \
    -d '{ "message": "hello" }' \
    -plaintext \
    localhost:55000 \
    proto.Chat/Send

grpcurl \
    -plaintext \
    localhost:55000 \
    example.Chat/Listen