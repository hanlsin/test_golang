package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Example of gRPC")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("error: net.Listen:", err)
		return
	}

	server := grpc.NewServer()
}
