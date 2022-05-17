package auth

import (
	"fmt"
	"go-grpc-api-gateway/pkg/auth/pb"
	"go-grpc-api-gateway/pkg/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(config *config.Config) pb.AuthServiceClient {
	conn, err := grpc.Dial(config.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect to auth service: ", err)
	}

	return pb.NewAuthServiceClient(conn)
}
