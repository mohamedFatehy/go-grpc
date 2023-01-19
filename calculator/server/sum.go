package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum fun was invoked with %v", in)

	return &pb.SumResponse{
		SumResult: in.FirstNumber + in.SecondNumber,
	}, nil
}
