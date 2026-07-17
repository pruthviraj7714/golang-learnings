package main

import (
	"context"
	"grpc/internal/config"
	"grpc/internal/database"
	"grpc/internal/models"
	"grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	DB *gorm.DB
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

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

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
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

func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var user models.User

	err := s.DB.Find(&user, req.Id).Error

	if err != nil {
		return nil, err
	}

	user.Email = req.Email
	user.Name = req.Name
	user.Password = req.Password

	err = s.DB.Save(&user).Error

	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		User: &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		},
	}, nil

}

func (s *UserServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	var user models.User

	err := s.DB.Delete(&user, req.Id).Error

	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}

func main() {
	cfg := config.LoadConfig()

	db := database.Connect(cfg.DBURL)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userService := &UserServer{
		DB: db,
	}

	pb.RegisterUserServiceServer(grpcServer, userService)
	reflection.Register(grpcServer)

	log.Printf("gRPC server starting on port :50051...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
