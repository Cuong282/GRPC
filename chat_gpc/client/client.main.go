package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"test/OneDrive/Desktop/GRPC/chat_gpc/chatpb"
)

func main() {
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := chatpb.NewChatServiceClient(conn)

	msg := &chatpb.Message{
		Text:     "1",
		Sender:   "Client 1",
		Receiver: "client 2",
	}
	response, err := client.SendMessage(context.Background(), msg)
	if err != nil {
		log.Println(err)
	}
	log.Println("Received response:", response.Text)

	// Client 2 sends a message to Client 1
	msg = &chatpb.Message{
		Text:     "2",
		Sender:   "Client 2",
		Receiver: "client 1",
	}
	response, err = client.SendMessage(context.Background(), msg)
	if err != nil {
		log.Println(err)
	}
	log.Println("Received response:", response.Text)
}
