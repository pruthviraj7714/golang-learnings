package main

import (
	"context"
	"grpc/internal/config"
	"grpc/internal/database"
	"grpc/internal/models"
	"grpc/pb"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type userserver struct {
	pb.UnimplementedUserServiceServer
	DB *gorm.DB
}

func (s *userserver) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		},
	}, nil
}

func (s *userserver) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var user models.User

	err := s.DB.First(&user, req.Id).Error

	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		},
	}, nil
}

func main() {
	r := gin.Default()

	cfg := config.LoadConfig()

	db := database.Connect(cfg.DBURL)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userService := &userserver{
		DB: db,
	}

	pb.RegisterUserServiceServer(grpcServer, userService)

	log.Printf("gRPC server starting on port :50051...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	r.Run(cfg.PORT)
}
