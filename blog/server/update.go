package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateBlog(ctx context.Context, blog *pb.Blog) (*empty.Empty, error) {

	objectId, err := primitive.ObjectIDFromHex(blog.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot Parse ID")
	}

	data := &BlogItem{
		AuthorId: blog.AuthorId,
		Title:    blog.Title,
		Content:  blog.Content,
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": data})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot update Document")
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot Find Blog with Id")
	}

	return &empty.Empty{}, nil
}
