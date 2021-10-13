package main

import (
	"fmt"
	"log"
	"net"
	"prime/src/prime_proto"

	"google.golang.org/grpc"
)

type server struct {
	prime_proto.UnimplementedDecompositionServer
}

func (*server) PrimeNumberDecomposition(req *prime_proto.PrimeNumberDecompositionRequest, stream prime_proto.Decomposition_PrimeNumberDecompositionServer) error {
	fmt.Println("Decompondo número:")
	number := req.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&prime_proto.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor foi incrementado: %v\n", divisor)
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v\n", err)
	}
	log.Printf("Deu Bom !\n")
	s := grpc.NewServer()
	prime_proto.RegisterDecompositionServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to set the server: %v", err)
	}
}
