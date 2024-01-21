package main

import (
	"context"
	pb "github.com/choyalpramod/gRPCinGo/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Print("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive response: %v", err)
		}
		log.Println(message)
	}
	log.Println("Streaming finished with message")
}
