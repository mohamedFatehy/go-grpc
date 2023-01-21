package main

import (
	"context"
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, blogId *pb.BlogId) (*pb.Blog, error) {

	objectId, err := primitive.ObjectIDFromHex(blogId.Id)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Cannot Parse ID")
	}

	data := &BlogItem{}

	res := collection.FindOne(ctx, bson.M{"_id": objectId})

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			"cannot find blog with the id provided")
	}

	return documentToBlog(data), nil
}
