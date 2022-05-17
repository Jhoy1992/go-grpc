package order

import (
	"fmt"
	"go-grpc-api-gateway/pkg/config"
	"go-grpc-api-gateway/pkg/order/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(config *config.Config) pb.OrderServiceClient {
	conn, err := grpc.Dial(config.OrderSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Failed to connect to order service:", err)
	}

	return pb.NewOrderServiceClient(conn)
}
