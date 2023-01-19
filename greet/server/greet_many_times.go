package main

import (
	"fmt"
	pb "github.com/mohamedFatehy/go-grpc/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes fun was invoked with %v \n", in)

	for i := 0; i < 100; i++ {
		res := fmt.Sprintf("Hello %s no %d \n", in.ClientName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
