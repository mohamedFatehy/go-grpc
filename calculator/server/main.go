package main

import (
	"fmt"
	pb "github.com/mohamedFatehy/go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const address = "0.0.0.0:8888"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen to on: %v ", address)
	}

	fmt.Printf("listening on server: %v \n", address)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
