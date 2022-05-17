package services

import (
	"context"
	"go-grpc-order-svc/pkg/client"
	"go-grpc-order-svc/pkg/db"
	"go-grpc-order-svc/pkg/models"
	"go-grpc-order-svc/pkg/pb"
	"log"
	"net/http"
)

type Server struct {
	Handler    db.Handler
	ProductSvc client.ProductServiceClient
}

func (server *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := server.ProductSvc.FindOne(req.ProductId)

	log.Println("product:", product.Data)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	if product.Status == http.StatusNotFound {
		return &pb.CreateOrderResponse{
			Status: product.Status,
			Error:  product.Error,
		}, nil
	}

	if product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  "not enough stock",
		}, nil
	}

	order := models.Order{
		Price:     product.Data.Price,
		ProductId: product.Data.Id,
		UserId:    req.UserId,
	}

	server.Handler.DB.Create(&order)

	res, err := server.ProductSvc.DecreaseStock(req.ProductId, order.Id)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	if res.Status == http.StatusConflict {
		server.Handler.DB.Delete(&models.Order{}, order.Id)

		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  res.Error,
		}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
