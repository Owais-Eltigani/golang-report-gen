package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// create a listener for users requests
	lis, err := net.Listen("tcp", "5000")

	if err != nil {
		log.Fatal("server: error while creating the listener", err)
		return
	}

	grpcServer := grpc.NewServer()

	//
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("server: error while init server", err)
		return
	}

}
