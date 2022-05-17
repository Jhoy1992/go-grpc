package product

import (
	"fmt"
	"go-grpc-api-gateway/pkg/config"
	"go-grpc-api-gateway/pkg/product/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(config *config.Config) pb.ProductServiceClient {
	conn, err := grpc.Dial(config.ProductSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Failed to connect to product service:", err)
	}

	return pb.NewProductServiceClient(conn)
}
