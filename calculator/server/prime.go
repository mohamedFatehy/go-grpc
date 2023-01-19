package main

import (
	pb "github.com/mohamedFatehy/go-grpc/calculator/proto"
	"log"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Sum fun was invoked with %v", in)

	k := uint32(2)
	n := in.PrimeNumber
	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{
				PrimeResult: k,
			})
			n = n / k
		} else {
			k++
		}
	}

	return nil
}
