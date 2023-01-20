package main

import (
	"fmt"
	pb "github.com/mohamedFatehy/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const address = "localhost:8888"

func main() {

	var opts []grpc.DialOption
	tls := true

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			fmt.Printf("Error while loading CA trust certificates: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))

	}

	conn, err := grpc.Dial(address, opts...)

	if err != nil {
		log.Fatalf("Failed to Connect: %v \n", address)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	doGreet(client)
	//doGreetManyTimes(client)
	//longGreet(client)
	//doGreetEveryone(client)
	//doGreetWithDeadline(client, 5*time.Second)
	//doGreetWithDeadline(client, 1*time.Second)

}
