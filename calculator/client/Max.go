package main

import (
	"context"
	"fmt"
	pb "github.com/mohamedFatehy/go-grpc/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(client pb.CalculatorServiceClient) {
	log.Printf("Max Function has been invoked")

	numbersList := []uint32{1, 5, 3, 6, 2, 20}

	waitc := make(chan struct{})

	req, err := client.Max(context.Background())

	if err != nil {
		log.Fatalf("Could not Get Prime due : %v \n", err)
	}

	go func() {
		for _, number := range numbersList {
			fmt.Printf("sending no : %d \n", number)
			req.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}
		req.CloseSend()
	}()

	go func() {
		for {
			res, err := req.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("failed to get resonse due to %v \n", err)
				break
			}

			fmt.Printf("server said max is:  %d\n", res.MaxResult)
		}
		close(waitc)
	}()

	<-waitc
}
