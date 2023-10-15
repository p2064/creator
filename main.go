package main

import (
	"fmt"
	"log"
	"net"

	"os"

	"github.com/p2064/creator/handlers"
	"github.com/p2064/creator/proto"
	"github.com/p2064/pkg/logs"
	"google.golang.org/grpc"

	"github.com/p2064/pkg/config"
)

func main() {
	logs.InfoLogger.Print("Start creator")
	if config.Status != config.GOOD {
		logs.ErrorLogger.Fatalf("failed to get config")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("CREATOR_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterCreatorServiceServer(grpcServer, &handlers.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		logs.ErrorLogger.Fatalf("failed to serve: %s", err)
	}
}
