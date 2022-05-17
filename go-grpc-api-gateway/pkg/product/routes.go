package product

import (
	"go-grpc-api-gateway/pkg/auth"
	"go-grpc-api-gateway/pkg/config"
	"go-grpc-api-gateway/pkg/product/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, config *config.Config, authSvc *auth.ServiceClient) {
	authMiddleware := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(config),
	}

	routes := router.Group("/product")
	routes.Use(authMiddleware.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}
