package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(client pb.CalculatorServiceClient, n int32) {

	log.Printf("Sqrt was invoked for: %v \n", n)

	res, err := client.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error Message from server : %v \n", e.Message())
			log.Printf("Error Code from server : %v \n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("You Probably send negative number")
			}

			return
		} else {
			log.Printf("Non gRPC error : %v \n", err.Error())
		}
	}

	log.Printf("the sqrt of %d is %f\n", n, res.Result)
}
