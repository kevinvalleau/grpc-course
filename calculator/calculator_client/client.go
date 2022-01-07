package main

import (
	"context"
	"fmt"
	"grpc-course/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	cs := calculatorpb.NewCalculatorServiceClient(cc)
	doUnarySum(cs)

}

func doUnarySum(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a unary RPC Sum...")
	req := &calculatorpb.SumRequest{
		Number1: 3,
		Number2: 10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Calculator RPC: %v", err)
	}
	log.Printf("Response from Calculator: %v\n", res.Result)
}
