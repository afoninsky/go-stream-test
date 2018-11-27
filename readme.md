###
> Test GRPC service with reflection and keepalive support. Implemented to test bug with stream timeouts in kubernetes environments

```
docker build -t gcr.io/peak-orbit-214114/sl-grpc-proxy:grpc-debug .
docker push gcr.io/peak-orbit-214114/sl-grpc-proxy:grpc-debug
kubectl apply -f kubernetes
```


```
time grpcurl \
    -d '{ "item": "echo" }' \
    grpc.internal.streamlayer.io:443 \
    proto.Debug/Echo
```

```
time GODEBUG=http2debug=2 grpcurl \
    -d '{ "item": "silence" }' \
    -keepalive-time 3 \
    grpc.internal.streamlayer.io:443 \
    proto.Debug/Silence
```