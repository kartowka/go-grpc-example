package main

import (
	"context"
	"log"
	"time"

	chatpb "github.com/antfley/go-grpc-example/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()
	c := chatpb.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.SendMessage(ctx, &chatpb.Message{Body: "Hello from client"})
	if err != nil {
		log.Fatalf("Error when calling SendMessage: %v", err)
	}
	log.Printf("Response from server: %v", response)

}
