package main

import (
	"context"
	"fmt"
	pb "github.com/mohamedFatehy/go-grpc/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(client pb.GreetServiceClient) {

	log.Printf("doGreetEveryone has been invoked")

	stream, err := client.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("failed to call GreetEveryone method")
	}

	reqs := []*pb.GreetRequest{
		{ClientName: "mohamed"},
		{ClientName: "ahmed"},
		{ClientName: "abdo"},
		{ClientName: "kareem"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			fmt.Printf("you introduced:  %s\n", req.ClientName)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("failed to get resonse due to %v \n", err)
				break
			}

			fmt.Printf("server responsed %s\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
