package main

import (
	"fmt"
	"log"
	"prime/src/prime_proto"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client!!")
	//1# - criar uma conexão com servidor
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect : %v", err)
	}
	defer conn.Close()
	c := prime_proto.NewDecompositionClient(conn)
	if c != nil {
		fmt.Println("Client created")
	}
	//doServerStreaming(c)
}
