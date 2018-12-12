### Structure
1) User sends direct messages to `receiver` using grpc
2) `receiver` translates message to Frensh and sends result to `streamer`
3) User subscribes to messages directly from `streamer` using grpc or via web-proxy using http and receives translations


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
    -d '{ "text": "привет" }' \
    -plaintext \
    localhost:55000 \
    internal.receiver.Receiver/Publish

grpcurl \
    -plaintext \
    localhost:55001 \
    internal.streamer.Streamer/Subscribe
```

### Debug
docker build -t gcr.io/peak-orbit-214114/istio-go:grpc -f Dockerfile.grpc .
docker push gcr.io/peak-orbit-214114/istio-go:grpc
