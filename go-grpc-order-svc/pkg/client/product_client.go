package client

import (
	"context"
	"fmt"
	"go-grpc-order-svc/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("could not connect to grpc server:", err)
	}

	client := ProductServiceClient{
		Client: pb.NewProductServiceClient(conn),
	}

	return client
}

func (client *ProductServiceClient) FindOne(productId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}

	return client.Client.FindOne(context.Background(), req)
}

func (client *ProductServiceClient) DecreaseStock(productId int64, orderId int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}

	return client.Client.DecreaseStock(context.Background(), req)
}
