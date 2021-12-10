package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/halhal23/shoezoo/product/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("hello product service for shoezoo 1")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Printf("Faild to listen: %v\n", err)
	}
	server := grpc.NewServer()
	service := NewProductService()
	reflection.Register(server)
	pb.RegisterProductServiceServer(server, service)
	if err := server.Serve(lis); err != nil {
		log.Printf("Faild to serve: %v\n", err)
	}
}

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) ShowProduct(c context.Context, req *pb.ShowProductRequest) (*pb.ShowProductResponse, error) {
	return &pb.ShowProductResponse{
		Product: &pb.Product{
			Id:    23,
			Name:  "ジョーダン",
			Price: 12000,
		},
	}, nil
}
