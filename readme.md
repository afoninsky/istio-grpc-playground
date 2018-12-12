### Structure
1) User send direct messages to `receiver` using grpc
2) `receiver` proxies messages to `streamer`
3) User subscribe to messages directly from `streamer` using grpc or via web-proxy using http


### Run locally
```
docker-compose build
docker-compose up
```

### Test
```
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
```