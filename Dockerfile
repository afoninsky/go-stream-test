FROM golang:1.11 AS builder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOBIN /bin
WORKDIR /src/server
COPY . /src
RUN go get
RUN go build -a -installsuffix nocgo -o /server .

FROM scratch
COPY --from=builder /server ./
ENTRYPOINT ["./server"]