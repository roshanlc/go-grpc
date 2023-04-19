package main

import (
	"log"
	"net"

	pb "github.com/roshanlc/go-grpc/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

const (
	port = ":9000"
)

func main() {
	listen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to start the server %v", err)
	}

	gRPCServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(gRPCServer, &helloServer{})

	log.Printf("Started server at port %v", port)
	if err := gRPCServer.Serve(listen); err != nil {
		log.Fatalf("failed to start gRPC Server %v", err)
	}

}
