package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	ls, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(ls); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
