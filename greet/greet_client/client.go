package main

import (
	"fmt"
	"log"

	"github.com/vataraxia/grpc-go-coba/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Connecting to server...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server")
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}
