package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/greet/proto"
	"log"
	"time"
)

func longGreet(client pb.GreetServiceClient) {
	log.Printf("longGreet Function has been invoked")

	reqs := []*pb.GreetRequest{
		{ClientName: "mohamed"},
		{ClientName: "ahmed"},
		{ClientName: "abdelrahman"},
		{ClientName: "kareem"},
	}

	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Could not Long Greet: %v \n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending Request with client name %s \n", req.ClientName)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while recieving res from Long Greet: %v \n", err)
	}

	log.Printf("Long Greet:\n%s\n", res.Result)
}
