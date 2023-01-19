package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(client pb.GreetServiceClient) {
	log.Printf("doGreet Many Times Function has been invoked")

	stream, err := client.GreetManyTimes(context.Background(), &pb.GreetRequest{
		ClientName: "Te7a",
	})

	if err != nil {
		log.Fatalf("Could not Greet: %v \n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v \n", err)
		}

		log.Printf("Server Said: %s \n", msg.Result)
	}
}
