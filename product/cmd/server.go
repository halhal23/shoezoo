package main

import (
	"fmt"
	"log"
	"net"

	handler "github.com/halhal23/shoezoo/product/interface/grpc"
	pb "github.com/halhal23/shoezoo/product/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Welcome to product service for shoezoo")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Printf("Faild to listen: %v\n", err)
	}
	server := grpc.NewServer()
	reflection.Register(server)
	h := handler.NewProductGrpcHandler()
	pb.RegisterProductServiceServer(server, h)
	if err := server.Serve(lis); err != nil {
		log.Printf("Faild to serve: %v\n", err)
	}
}
