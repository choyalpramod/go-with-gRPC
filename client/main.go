package main

import (
	pb "github.com/choyalpramod/gRPCinGo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	port = ":8080"
)

func main() {
	connection, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer closeConnection(connection)
	client := pb.NewGreetServiceClient(connection)
	names := &pb.NamesList{
		Names: []string{"Pramod", "Ivan", "Michael"},
	}

	//callSayHello(client)
	//callSayHelloServerStream(client, names)
	//callSayHellClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}

func closeConnection(connection *grpc.ClientConn) {
	if err := connection.Close(); err != nil {
		log.Fatalf("failed to close connection: %v", err)
	}
}
