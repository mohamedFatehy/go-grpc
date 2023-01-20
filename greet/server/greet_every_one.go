package main

import (
	"fmt"
	pb "github.com/mohamedFatehy/go-grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	log.Printf("GreetEveryone fun was invoked with\n")
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("An error has been occurred: %s \n", err)
		}
		err = stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s!", req.ClientName),
		})

		if err != nil {
			log.Fatalf("Failed to Send Response: %s \n", err)
		}
	}

	return nil
}
