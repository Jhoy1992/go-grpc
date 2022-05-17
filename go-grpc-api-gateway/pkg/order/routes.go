package order

import (
	"go-grpc-api-gateway/pkg/auth"
	"go-grpc-api-gateway/pkg/config"
	"go-grpc-api-gateway/pkg/order/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, config *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(config),
	}

	routes := router.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
