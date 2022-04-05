package main

import (
	"fmt"
	"log"
	"net"

	pb "grpc-go/pb"
	service "grpc-go/pkg/service"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTestServer(grpcServer, service.NewTestServer())

	fmt.Println("Now listening...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)

	}

}
