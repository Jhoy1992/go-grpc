package routes

import (
	"context"
	"go-grpc-api-gateway/pkg/product/pb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindOne(ctx *gin.Context, client pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := client.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
