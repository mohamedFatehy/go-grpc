package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mohamedFatehy/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBlogs(in *empty.Empty, stream pb.BlogService_ListBlogsServer) error {

	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(codes.NotFound,
			fmt.Sprintf("unknown error due to: %v", err))
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal,
				fmt.Sprintf("Error While Decoding data from mongo: %v", err))
		}
		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal,
			fmt.Sprintf("unknown internal error: %v", err))
	}

	return nil
}
