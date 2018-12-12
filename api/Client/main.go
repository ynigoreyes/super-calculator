package main

import (
	"log"
	"time"

	"golang.org/x/net/context"

	pb "github.com/ynigoreyes/super-calculator/api/Client/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("api:50051", grpc.WithInsecure())
	if err != nil {
		log.Println("Error with Dial")
		log.Fatalln(err)
	}

	defer conn.Close()

	for {
		client := pb.NewCalculatorClient(conn)
		nums := &pb.PairOfNumbers{Number1: 1, Number2: 2}
		res, err := client.Add(context.Background(), nums)

		if err != nil {
			log.Println("Error with Add")
			log.Fatalln(err)
		}

		time.Sleep(5)
		log.Println(res)
	}
}
