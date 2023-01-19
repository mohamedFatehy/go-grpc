package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/calculator/proto"
	"log"
)

func getSum(client pb.CalculatorServiceClient) {
	log.Printf("getSum Function has been invoked")

	firstNo, secondNo := 10, 18

	res, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  uint32(firstNo),
		SecondNumber: uint32(secondNo),
	})

	if err != nil {
		log.Fatalf("Could not Sum: %v \n", err)
	}

	log.Printf("Server Summed %d + %d = %d \n", firstNo, secondNo, res.SumResult)
}
