package main

import (
	"distributed-transaction/cmd/app"
	dao2 "distributed-transaction/dao"
	"distributed-transaction/db"
	"distributed-transaction/handlers"
	service2 "distributed-transaction/service"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "distributed-transaction/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.ProductServiceServer
}

func main() {
	db := db.DatabaseConn()
	productDAO := dao2.NewProductDao(db)
	productService := service2.NewProductService(productDAO)
	productHandler := handlers.NewProductHandler(productService)

	app := app.NewServer(productHandler)

	app.Start()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
