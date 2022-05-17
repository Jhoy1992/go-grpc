package main

import (
	"fmt"
	"go-grpc-product-svc/pkg/config"
	"go-grpc-product-svc/pkg/db"
	"go-grpc-product-svc/pkg/pb"
	"go-grpc-product-svc/pkg/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config:", err)
	}

	handler := db.Init(config.DBUrl)

	listener, err := net.Listen("tcp", ":"+config.Port)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	fmt.Println("Product server is running on port", config.Port)

	server := services.Server{
		Handler: handler,
	}

	grpcsServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcsServer, &server)

	if err := grpcsServer.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
