package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"test/OneDrive/Desktop/GRPC/chat_gpc/chatpb"

	"google.golang.org/grpc"
)

var (
	conns   []net.Conn
	connCh  = make(chan net.Conn)
	closeCh = make(chan net.Conn)
	msgCh   = make(chan string)
)

type chatService struct {
	chatpb.ChatServiceServer
}

func (s *chatService) ChatUser1(ctx context.Context, req *chatpb.User1) (*chatpb.Token, error) {
	fmt.Printf("chat success1")

	return &chatpb.Token{Token: "login_success"}, nil
}
func (s *chatService) ChatUser2(ctx context.Context, req *chatpb.User2) (*chatpb.Token, error) {
	fmt.Printf("chat success2")
	return &chatpb.Token{Token: "login_success"}, nil
}

func (*chatService) SendMessage(ctx context.Context, req *chatpb.Message) (*chatpb.Message, error) {
	log.Println("Received message:", req.Text)

	// Send a response back to the client
	response := &chatpb.Message{
		Text:     "Received message from " + req.Text,
		Sender:   "server",
		Receiver: req.Sender,
	}
	return response, nil
}
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:3001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	chatpb.RegisterChatServiceServer(s, &chatService{})
	log.Printf("gRPC server listening on port 3001")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
	fmt.Println(">>>>>>>>>:")
}
