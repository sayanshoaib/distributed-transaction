package main

import (
	"context"
	pb "distributed-transaction/proto"
	"log"
)

func addProductToCart(c pb.ProductServiceClient) {
	log.Println("addProductToCart was invoked")
	res, err := c.AddProductToCart(context.Background(), &pb.ProductRequest{
		Name:  "Potato",
		Price: 25,
		TxInfo: &pb.TransactionInfo{
			ServiceName:    "Product Service",
			TxName:         "Add Product To Cart",
			IsTxSuccessful: true,
		},
	})

	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}

	log.Printf("Product Name: %s, Product Price: %v\n", res.Name, res.Price)
	log.Printf("Transaction info: ")
	log.Printf("Service Name: %s", res.TxInfo.ServiceName)
	log.Printf("Transaction Name: %s", res.TxInfo.TxName)
	log.Printf("Is Transaction Successful: %t", res.TxInfo.IsTxSuccessful)
}
