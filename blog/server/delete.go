package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteBlog(ctx context.Context, blogId *pb.BlogId) (*empty.Empty, error) {

	objectId, err := primitive.ObjectIDFromHex(blogId.Id)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Cannot Parse ID")
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": objectId})

	if err != nil {
		return nil, status.Errorf(codes.Internal,
			fmt.Sprintf("cannot delete blog due to: %v", err))
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound,
			"Blog Not found")
	}

	return &empty.Empty{}, nil
}
