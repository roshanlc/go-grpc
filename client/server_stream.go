package main

import (
	"context"
	"io"
	"log"

	pb "github.com/roshanlc/go-grpc/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) {

	log.Printf("streaming started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send name list: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error while streaming : %s", err)
		}
		log.Println(msg)

	}

	log.Printf("streaming finished")
}
