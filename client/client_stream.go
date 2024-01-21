package main

import (
	"context"
	pb "github.com/choyalpramod/gRPCinGo/proto"
	"log"
	"time"
)

func callSayHellClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client Streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("could not send names: %v", err)
		}
		log.Printf("Sent name: %v", name)
		time.Sleep(time.Second * 2)
	}

	res, errorReceived := stream.CloseAndRecv()
	if errorReceived != nil {
		log.Fatalf("could not receive response: %v", errorReceived)
	}
	log.Printf("Client streaming finished with message: %v", res)
}
