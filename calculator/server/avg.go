package main

import (
	pb "github.com/mohamedFatehy/go-grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Avg fun was invoked \n")

	var sum uint32 = 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				AvgResult: float32(sum) / float32(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading from the client steam: %s \n", err)
		}

		// increase the count and the sum
		count++
		sum += req.Number
	}

	return nil
}
