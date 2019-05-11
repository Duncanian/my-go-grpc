package main

import (
	"context"
	"log"

	pb "github.com/Duncanian/my-go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}

	defer conn.Close()

	client := pb.NewSimpleHelloClient(conn)

	myGreeting, err := client.SayHello(context.Background(), &pb.MyTarget{Target: "world"})
	if err != nil {
		log.Fatalf("Failed to send greeting: %v", err)
	}
	log.Println(myGreeting)
}
