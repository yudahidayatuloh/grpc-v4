package main

import (
	"context"
	"grpc-course-protobuf/pb/chat"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientconn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to create client ", err)
	}

	chatClient :=  chat.NewChatServiceClient(clientconn)
	stream, err := chatClient.SendMassage(context.Background())
	if err != nil {
		log.Fatal("Failed to send message", err)
	}

	err = stream.Send(&chat.ChatMassage{
		UserId: 123,
		Content: "Hallo from client",
	})
	if err != nil {
		log.Fatal("Failed to send via stream", err)
	}
	err = stream.Send(&chat.ChatMassage{
		UserId: 123,
		Content: "Hallo again",
	})
	if err != nil {
		log.Fatal("Failed to send via stream", err)
	}
	err = stream.Send(&chat.ChatMassage{
		UserId: 123,
		Content: "Hallo abayy",
	})
	if err != nil {
		log.Fatal("Failed to send via stream", err)
	}
	time.Sleep(5 * time.Second)
	err = stream.Send(&chat.ChatMassage{
		UserId: 123,
		Content: "Hallo cokk",
	})
	if err != nil {
		log.Fatal("Failed to send via stream", err)
	}
	
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Failed close", err)
	}
	log.Println("Connection is closed. Message: ", res.Message)
}
