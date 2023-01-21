package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"io"
	"log"
)

func listBlog(client pb.BlogServiceClient) {

	log.Printf("<--listBlog Function has been invoked-->")
	stream, err := client.ListBlogs(context.Background(), &empty.Empty{})

	if err != nil {
		log.Fatalf("Unexpected Error: %v \n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Unexpected Error: %v \n", err)
		}

		log.Print(res)
	}
}
