package main

import (
	"context"
	"fmt"
	pb "github.com/mohamedFatehy/go-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadline(client pb.GreetServiceClient, timeout time.Duration) {

	log.Printf("doGreetWithDeadline has been invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := client.GreetWithDeadline(ctx, &pb.GreetRequest{
		ClientName: "Te7a",
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Request Deadline Exceeded")
			}
			return
		} else {
			log.Printf("Non gRPC error : %v \n", err.Error())
		}
	}

	fmt.Printf("Server responded: %v \n", req.Result)

}
