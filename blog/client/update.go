package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"log"
)

func updateBlog(client pb.BlogServiceClient, id string) {

	log.Printf("<--updateBlog Function has been invoked->>")

	_, err := client.UpdateBlog(context.Background(), &pb.Blog{
		Id:       id,
		AuthorId: "2",
		Title:    "learning Go is Crazy",
		Content:  "it is kind of fun and crazy in the same time",
	})

	if err != nil {
		log.Fatalf("Unexpected Error: %v \n", err)
	}

	log.Printf("Blog Post with ID %s has been Updated successfully\n", id)

}
