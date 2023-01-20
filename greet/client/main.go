package main

import (
	pb "github.com/mohamedFatehy/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "localhost:8888"

func main() {

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to Connect: %v \n", address)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	doGreet(client)
	doGreetManyTimes(client)
	longGreet(client)
	doGreetEveryone(client)
}
