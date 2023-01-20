package main

import (
	pb "github.com/mohamedFatehy/go-grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {

	var max uint32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("there is an err %s \n", err)
		}

		if max < req.Number {
			max = req.Number
		}

		stream.Send(&pb.MaxResponse{
			MaxResult: max,
		})
	}

	return nil
}
