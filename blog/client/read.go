package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"log"
)

func readBlog(client pb.BlogServiceClient, id string) *pb.Blog {

	log.Printf("<--readBlog Function has been invoked->>")
	res, err := client.ReadBlog(context.Background(), &pb.BlogId{
		Id: id,
	})

	if err != nil {
		log.Fatalf("Unexpected Error: %v \n", err)
	}

	log.Printf("Blog Post with ID %s has been Read and get data :%s \n", id, res)

	return res
}
