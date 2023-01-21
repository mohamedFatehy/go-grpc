package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"log"
)

func createBlog(client pb.BlogServiceClient) string {

	log.Printf("<--createBlog Function has been invoked-->")

	blog := &pb.Blog{
		AuthorId: "1",
		Title:    "Learning Go",
		Content:  "Learning GoLang is Fun, you will enjoy it a lot",
	}

	res, err := client.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected Error: %v \n", err)
	}

	log.Printf("Blog Post has been created with ID :%s \n", res.Id)

	return res.Id
}
