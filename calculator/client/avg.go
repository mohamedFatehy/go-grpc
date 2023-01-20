package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/calculator/proto"
	"log"
	"time"
)

func doAvg(client pb.CalculatorServiceClient) {
	log.Printf("doAvg Function has been invoked")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := client.Avg(context.Background())

	if err != nil {
		log.Fatalf("Could not doAvg : %v \n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending Request with Number %d \n", req.Number)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while recieving res from Long Greet: %v \n", err)
	}

	log.Printf("Avg : %.2f \n", res.AvgResult)
}
