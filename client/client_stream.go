package main

import (
	"context"
	"log"
	"time"

	pb "github.com/roshanlc/go-grpc/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending name: %v", err)
		}

		log.Printf("Sent request with name: %v", name)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()

	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("error while receiving: %v", err)
	}

	log.Printf("%v", res.Messages)
}
