package main

import (
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "localhost:9999"

func main() {

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to Connect: %v \n", address)
	}
	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)
	blogId := createBlog(client)
	readBlog(client, blogId)
	updateBlog(client, blogId)
	readBlog(client, blogId)
	listBlog(client)
}
