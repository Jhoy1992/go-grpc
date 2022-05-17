package auth

import (
	"context"
	"go-grpc-api-gateway/pkg/auth/pb"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMidlewareConfig struct {
	service *ServiceClient
}

func InitAuthMiddleware(service *ServiceClient) AuthMidlewareConfig {
	return AuthMidlewareConfig{service}
}

func (config *AuthMidlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := config.service.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}
