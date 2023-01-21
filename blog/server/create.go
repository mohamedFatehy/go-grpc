package main

import (
	"context"
	"fmt"
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, blog *pb.Blog) (*pb.BlogId, error) {

	data := BlogItem{
		AuthorId: blog.AuthorId,
		Title:    blog.Title,
		Content:  blog.Content,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v\n", err))
	}

	objectId, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Error(
			codes.Internal,
			"Cannot Convert to ObjectID")
	}

	return &pb.BlogId{
		Id: objectId.Hex(),
	}, nil
}
