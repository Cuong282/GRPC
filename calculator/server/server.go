package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"example.com/m/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.CalculatorServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("sum called...")
	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1() - req.GetNum2(),
	}

	return resp, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PNDRequest,
	stream grpc.ServerStreamingServer[calculatorpb.PNDResponse]) error {
	fmt.Println("primeNumberDecomposition called...")
	k := int32(2)
	N := req.GetNumber()
	for N > 1 {
		if N%k == 0 {
			N = N / k
			stream.Send(&calculatorpb.PNDResponse{
				Result: k,
			})
		} else {
			k++
			fmt.Println("k increase to &v", k)
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(s, &server{})

	fmt.Println("server is running...")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
	fmt.Println(">>>>>>>>>:")
}
