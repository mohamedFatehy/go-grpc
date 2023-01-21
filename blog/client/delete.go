package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"log"
)

func deleteBlog(client pb.BlogServiceClient, id string) {

	log.Printf("<--deleteBlog Function has been invoked->>")
	_, err := client.DeleteBlog(context.Background(), &pb.BlogId{
		Id: id,
	})

	if err != nil {
		log.Fatalf("Unexpected Error: %v \n", err)
	}

	log.Printf("Blog Post with ID %s has been Deleted \n", id)
}
