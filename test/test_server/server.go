package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-go-course/test/testpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	fmt.Printf("Testing function was invoked with %v\n", req)
	firstName := req.GetTesting().GetFirstName()
	lastName := req.GetTesting().GetLastName()
	result := "Hello " + firstName + " " + lastName

	res := &pb.TestResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Hello world")

	// 50051 is the default port for grpc
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to Listen %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	// register test server
	pb.RegisterTestServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Listen %v", err)
	}

}
