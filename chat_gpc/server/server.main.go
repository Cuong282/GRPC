package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"test/OneDrive/Desktop/my_project/GRPC/chat_gpc/chatpb"

	"google.golang.org/grpc"
)

type server struct {
	chatpb.ChatServiceServer
}

func (s *server) ChatUser1(ctx context.Context, req *chatpb.User1) (*chatpb.Token, error) {
	fmt.Printf("chat success1")
	return &chatpb.Token{Token: "login_success"}, nil
}
func (s *server) ChatUser2(ctx context.Context, req *chatpb.User2) (*chatpb.Token, error) {
	fmt.Printf("chat success2")
	return &chatpb.Token{Token: "login_success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	chatpb.RegisterChatServiceServer(s, &server{})
	log.Printf("gRPC server listening on port 8080")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
	fmt.Println(">>>>>>>>>:")
}
