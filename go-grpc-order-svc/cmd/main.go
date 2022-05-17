package main

import (
	"fmt"
	"go-grpc-order-svc/pkg/client"
	"go-grpc-order-svc/pkg/config"
	"go-grpc-order-svc/pkg/db"
	"go-grpc-order-svc/pkg/pb"
	"go-grpc-order-svc/pkg/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("could not load config:", err)
	}

	handler := db.Init(config.DBUrl)

	listener, err := net.Listen("tcp", ":"+config.Port)

	if err != nil {
		log.Fatalln("could not listen:", err)
	}

	productSvc := client.InitProductServiceClient(config.ProductSvcUrl)

	if err != nil {
		log.Fatalln("could not connect to product service:", err)
	}

	fmt.Println("Oder service listening on port", config.Port)

	server := services.Server{
		Handler:    handler,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln("could not serve:", err)
	}
}
