package main

import (
	"go-grpc-api-gateway/pkg/auth"
	"go-grpc-api-gateway/pkg/config"
	"go-grpc-api-gateway/pkg/order"
	"go-grpc-api-gateway/pkg/product"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config:", err)
	}

	router := gin.Default()

	authSvc := *auth.RegisterRoutes(router, &config)
	product.RegisterRoutes(router, &config, &authSvc)
	order.RegisterRoutes(router, &config, &authSvc)

	router.Run(":" + config.Port)
}
