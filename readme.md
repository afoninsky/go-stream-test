###
> Test GRPC service with reflection and keepalive support. Implemented to test bug with stream timeouts in kubernetes environments

```
docker build -t gcr.io/peak-orbit-214114/sl-grpc-proxy:grpc-debug .
docker push gcr.io/peak-orbit-214114/sl-grpc-proxy:grpc-debug
kubectl apply -f kubernetes
```


```
# debug
grpcurl \
  -d '{ "item": "echo" }' \
  -plaintext \
  35.227.27.52:50051 \
  proto.Debug/Echo

# default
grpcurl \
  -d '{ "item": "echo" }' \
  -plaintext \
  35.227.22.162:50051 \
  proto.Debug/Echo

# istio plain
grpcurl \
  -d '{ "item": "echo" }' \
  -plaintext \
  grpc.testing.streamlayer.io:31400 \
  proto.Debug/Echo
```

```
time GODEBUG=http2debug=2 grpcurl \
    -d '{ "item": "silence" }' \
    -keepalive-time 3 \
    grpc.internal.streamlayer.io:443 \
    proto.Debug/Silence
```
