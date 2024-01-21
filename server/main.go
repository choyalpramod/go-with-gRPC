package main

import (
	pb "github.com/choyalpramod/gRPCinGo/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
