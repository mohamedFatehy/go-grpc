package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/calculator/proto"
	"io"
	"log"
)

func getPrime(client pb.CalculatorServiceClient) {
	log.Printf("getPrime Function has been invoked")

	primeNo := 120

	stream, err := client.Prime(context.Background(), &pb.PrimeRequest{
		PrimeNumber: uint32(primeNo),
	})

	if err != nil {
		log.Fatalf("Could not Get Prime due : %v \n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v \n", err)
		}

		log.Printf("Prime Factor of %d is : %d \n", primeNo, msg.PrimeResult)
	}
}
