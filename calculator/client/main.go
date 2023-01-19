package main

import (
	pb "github.com/mohamedFatehy/go-grpc/calculator/proto"
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

	client := pb.NewCalculatorServiceClient(conn)
	getSum(client)
	getPrime(client)
}
