package main

import (
	"context"
	"fmt"
	"io"
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
	doServerStreaming(c)
}

func doServerStreaming(c prime_proto.DecompositionClient) {
	fmt.Println("Inicialização da função")
	req := &prime_proto.PrimeNumberDecompositionRequest{
		Number: 1200000,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Erro durante verificação de primo: %v \n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error ao processar primo %v\n", err)
		}
		log.Printf("Número é divisivel por: %v\n", res.GetPrimeFactor())
	}
}
