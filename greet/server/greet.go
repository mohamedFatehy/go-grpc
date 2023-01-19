package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet fun was invoked with %v", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.ClientName,
	}, nil
}
