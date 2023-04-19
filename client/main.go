package main

import (
	"log"

	pb "github.com/roshanlc/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":9000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to start client %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{Names: []string{
		"Ram",
		"Shyam",
		"Hari",
		"Geeta",
	},
	}

	// callSayHello(client)
	// callSayHelloServerStreaming(client, names)
	callSayHelloClientStreaming(client, names)
}
