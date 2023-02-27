package main

import (
	"context"
	pb "distributed-transaction/proto"
	"log"
)

func (s *Server) AddProductToCart(ctx context.Context, in *pb.ProductRequest) (*pb.ProductResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.ProductResponse{
		Name:   in.Name,
		Price:  in.Price,
		TxInfo: in.TxInfo,
	}, nil
}
