package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ynigoreyes/super-calculator/api/Calculator/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (server) Add(ctx context.Context, nums *pb.PairOfNumbers) (*pb.Results, error) {
	log.Printf("Recieved %v and %v!\n", nums.Number1, nums.Number2)
	return &pb.Results{Value: nums.Number1 + nums.Number2}, nil
}

func (server) Subtract(ctx context.Context, nums *pb.PairOfNumbers) (*pb.Results, error) {
	log.Printf("Recieved %v and %v\n", nums.Number1, nums.Number2)
	return &pb.Results{Value: nums.Number1 - nums.Number2}, nil
}

func (server) Multiply(ctx context.Context, nums *pb.PairOfNumbers) (*pb.Results, error) {
	log.Printf("Recieved %v and %v\n", nums.Number1, nums.Number2)
	return &pb.Results{Value: nums.Number1 * nums.Number2}, nil
}

func (server) Divide(ctx context.Context, nums *pb.PairOfNumbers) (*pb.Results, error) {
	log.Printf("Recieved %v and %v\n", nums.Number1, nums.Number2)
	return &pb.Results{Value: nums.Number1 / nums.Number2}, nil
}

func main() {
	port := 50051

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Could not listen to port %d: %v", port, err)
	}

	log.Printf("Listening to port %d", port)

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
