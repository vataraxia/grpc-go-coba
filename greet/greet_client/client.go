package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vataraxia/grpc-go-coba/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Connecting to server... ")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server")
	}
	defer cc.Close()

	fmt.Printf("Connected!!\n\n")

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Iman",
			LastName:  "Suhardiman",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
