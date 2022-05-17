package main

import (
	"fmt"
	"go-grpc-auth-svc/pkg/config"
	"go-grpc-auth-svc/pkg/db"
	"go-grpc-auth-svc/pkg/pb"
	"go-grpc-auth-svc/pkg/services"
	"go-grpc-auth-svc/pkg/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config:", err)
	}

	db := db.Init(config.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       config.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	webServer, err := net.Listen("tcp", ":"+config.Port)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	fmt.Println("Auth Service listening on port", config.Port)

	server := services.Server{
		Handler: db,
		Jwt:     jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(webServer); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
