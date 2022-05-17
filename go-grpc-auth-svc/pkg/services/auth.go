package services

import (
	"context"
	"go-grpc-auth-svc/pkg/db"
	"go-grpc-auth-svc/pkg/models"
	"go-grpc-auth-svc/pkg/pb"
	"go-grpc-auth-svc/pkg/utils"
	"net/http"
)

type Server struct {
	Handler db.Handler
	Jwt     utils.JwtWrapper
}

func (server *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if result := server.Handler.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "User with this email already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	server.Handler.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (server *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := server.Handler.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found or password is incorrect",
		}, nil
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "User not found or password is incorrect",
		}, nil
	}

	token, _ := server.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (server *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := server.Jwt.ValidToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	if result := server.Handler.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
