FROM golang:1.11 AS builder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOBIN /bin
WORKDIR /src/server
RUN apt-get update && apt-get install gogoprotobuf -y
RUN go get -u github.com/golang/protobuf/protoc-gen-go
COPY . /src
RUN go generate
RUN go get -v
RUN go build -a -installsuffix nocgo -o /server .

FROM scratch
COPY --from=builder /server ./
ENTRYPOINT ["./server"]
