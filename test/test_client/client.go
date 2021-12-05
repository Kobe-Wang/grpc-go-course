package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-go-course/test/testpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello world: I am client")

	// client connects the grpc server localhost:50051, and without SSL in the local
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed %v", err)
	}
	defer conn.Close()

	// create client for connection with protocol buffer
	c := pb.NewTestServiceClient(conn)
	// fmt.Printf("Create client: %f", c)

	doUnary(c)
}

func doUnary(c pb.TestServiceClient) {
	fmt.Println("Starting to do Unary RPC...")

	req := &pb.TestRequest{
		Testing: &pb.Testing{
			FirstName: "Kobe",
			LastName:  "Wang",
		},
	}

	resp, err := c.Test(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Test RPC %v", err)
	}

	log.Printf("Response from Test: %v", resp.Result)
}
