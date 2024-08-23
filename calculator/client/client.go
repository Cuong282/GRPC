package main

import (
	"context"
	"io"
	"log"

	"example.com/m/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Printf("err while dial %v", err)
	}

	defer cc.Close()

	client := calculatorpb.NewCalculatorClient(cc)

	// log.Printf("service client %f", client)

	// callSum(client)
	callPND(client)
}

func callSum(c calculatorpb.CalculatorClient) {
	log.Println("call sum api:")
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 51,
		Num2: 6,
	})
	if err != nil {
		log.Fatalf("call sum api error %v", err)
	}
	log.Printf("sum api response %v\n", resp.GetResult())

}

func callPND(c calculatorpb.CalculatorClient) {
	log.Println("call sum api:")
	stream, err := c.PrimeNumberDecomposition(context.Background(), &calculatorpb.PNDRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("calPND err %v", err)
	}
	for {
		resp, recvErr := stream.Recv()
		if recvErr == io.EOF {
			log.Fatalf("server finish streaming")
			return
		}
		log.Printf("prime number %v", resp.GetResult())
	}
}
