package handler

import (
	"context"
	"log"

	pb "github.com/halhal23/shoezoo/product/proto/v1"
)

type ProductGrpcHandler struct{}

func NewProductGrpcHandler() *ProductGrpcHandler {
	return &ProductGrpcHandler{}
}

func (s *ProductGrpcHandler) ShowProduct(c context.Context, req *pb.ShowProductRequest) (*pb.ShowProductResponse, error) {
	id := req.GetId()
	log.Printf("Get Request ShowProduct: id=%v", id)
	return &pb.ShowProductResponse{
		Product: &pb.Product{
			Id:    id,
			Name:  "ハリーウィンストン",
			Price: 34000,
		},
	}, nil
}
