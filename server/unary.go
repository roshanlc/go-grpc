package main

import (
	"context"
	"log"

	pb "github.com/roshanlc/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	log.Printf("Received request: %s", req.String())
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
