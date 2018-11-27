//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/debug.proto
package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "../proto"
)

type server struct{}

func (s *server) Echo(in *pb.Item, out pb.Debug_EchoServer) error {
	log.Printf("incoming request: %s", in.Item)

	for {
		res := &pb.Item{Item: in.Item}
		log.Printf("echoed...")
		if err := out.Send(res); err != nil {
			log.Printf("error: %v", err)
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (s *server) Silence(in *pb.Item, out pb.Debug_SilenceServer) error {
	log.Printf("incoming request: %s", in.Item)

	for {
		time.Sleep(time.Duration(1) * time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// grpcServer := grpc.NewServer(grpc.KeepaliveParams(
	// 	keepalive.ServerParameters{
	// 		MaxConnectionAgeGrace: (time.Duration(5) * time.Second),
	// 	},
	// ))
	grpcServer := grpc.NewServer()
	pb.RegisterDebugServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	log.Printf("Starting service on port 50051 (v2)")
	grpcServer.Serve(lis)
}
