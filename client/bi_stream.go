package main

import (
	"context"
	pb "github.com/choyalpramod/gRPCinGo/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	waitChannel := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("could not receive response: %v", err)
			}
			log.Printf("Received message: %v\n", message.Message)
		}
		close(waitChannel)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if streamErr := stream.Send(req); streamErr != nil {
			log.Fatalf("could not send names: %v", streamErr)
		}
		time.Sleep(time.Second * 2)
	}
	if sendErr := stream.CloseSend(); sendErr != nil {
		log.Fatalf("could not close send: %v", sendErr)
	}
	<-waitChannel
	log.Printf("Bidirectional streaming finished")
}
