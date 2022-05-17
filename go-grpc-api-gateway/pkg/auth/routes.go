package auth

import (
	"go-grpc-api-gateway/pkg/auth/routes"
	"go-grpc-api-gateway/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, config *config.Config) *ServiceClient {
	service := &ServiceClient{
		Client: InitServiceClient(config),
	}

	routes := router.Group("/auth")
	routes.POST("/register", service.Register)
	routes.POST("/login", service.Login)

	return service
}

func (service *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, service.Client)
}

func (service *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, service.Client)
}
