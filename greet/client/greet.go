package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/greet/proto"
	"log"
)

func doGreet(client pb.GreetServiceClient) {
	log.Printf("doGreet Function has been invoked")

	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		ClientName: "mohammed",
	})

	if err != nil {
		log.Fatalf("Could not Greet: %v \n", err)
	}

	log.Printf("Server Said: %s \n", res.Result)
}
