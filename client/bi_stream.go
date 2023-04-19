package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/roshanlc/go-grpc/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			msg, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while streaming %v", err)
			}

			log.Println(msg)

		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending request: %v", err)
		}

		time.Sleep(time.Second)
	}

	stream.CloseSend()
	<-waitc
	log.Printf("bidirectional streaming finished!")
}
