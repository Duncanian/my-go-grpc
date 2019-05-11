package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Duncanian/my-go-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type server struct {
	myserver pb.HelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.MyTarget) (*pb.MyGreeting, error) {
	if req.Target == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Your target cannot be empty")
	}
	return &pb.MyGreeting{
		Greeting: fmt.Sprintf("Hello %s!", req.Target),
	}, nil
}

func main() {
	myPort := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "localhost", myPort))
	if err != nil {
		log.Fatalf("Error listening %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServer(grpcServer, &server{})
	grpcServer.Serve(lis)
	log.Println(grpcServer)
}
