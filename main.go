package main

import (
	"log"
	"net"

	"github.com/p2064/creator/handlers"
	"github.com/p2064/creator/proto"
	"google.golang.org/grpc"

	"github.com/p2064/pkg/config"
)

func main() {
	log.Print("Start creator")
	if config.Status != config.GOOD {
		log.Print("failed to get config")
	}
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterCreatorServiceServer(grpcServer, &handlers.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
