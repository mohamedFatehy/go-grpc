package main

import (
	"fmt"
	pb "github.com/mohamedFatehy/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const address = "0.0.0.0:8888"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen to on: %v ", address)
	}

	fmt.Printf("listening on server: %v \n", address)

	var opts []grpc.ServerOption
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			fmt.Printf("Invalid server credentials: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))

	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
