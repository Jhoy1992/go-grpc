package services

import (
	"context"
	"go-grpc-product-svc/pkg/db"
	"go-grpc-product-svc/pkg/models"
	"go-grpc-product-svc/pkg/pb"
	"log"
	"net/http"
)

type Server struct {
	Handler db.Handler
}

func (server *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Price

	log.Println(req)

	if result := server.Handler.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	}, nil
}

func (server *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product models.Product

	if result := server.Handler.DB.First(&product, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    product.Id,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (server *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	if result := server.Handler.DB.First(&product, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock is not enough",
		}, nil
	}

	var log models.StockDecreaseLog

	if result := server.Handler.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error == nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		}, nil
	}

	product.Stock = product.Stock - 1

	server.Handler.DB.Save(&product)

	log.OrderId = req.OrderId
	log.ProductRefer = product.Id

	server.Handler.DB.Save(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil

}
