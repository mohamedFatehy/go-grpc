package main

import (
	"context"
	"fmt"
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

const address = "0.0.0.0:9999"

type Server struct {
	pb.BlogServiceServer
}

var collection *mongo.Collection

func main() {

	// set mongo client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	// connect to mongo db
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// connect to mongo database
	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen to on: %v ", address)
	}

	fmt.Printf("listening on server: %v \n", address)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
