package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/Qiryl/taxi-service/internal/user/domain"
	pb "github.com/Qiryl/taxi-service/proto/user"
)

type GrpcUserServer struct {
	pb.UnimplementedUserServer
	userUc domain.UserUsecase
}

func NewGrpcUserServer(uc domain.UserUsecase) *GrpcUserServer {
	return &GrpcUserServer{
		userUc: uc,
	}
}

func (s *GrpcUserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	loginReq := &domain.LoginRequest{
		Phone:    req.Phone,
		Password: req.Password,
	}
	if err := s.userUc.Login(ctx, loginReq); err != nil {
		return nil, fmt.Errorf("Grpc Login: %w", err)
	}

	return &pb.AuthResponse{Token: ""}, nil
}

func (s *GrpcUserServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	user := &domain.User{
		ID:           [16]byte{},
		Name:         req.Username,
		Phone:        req.Phone,
		Email:        req.Email,
		Password:     req.Password,
		RegisteredAt: time.Time{},
	}
	if err := s.userUc.Register(ctx, user); err != nil {
		return nil, fmt.Errorf("Grpc Register: %w", err)
	}

	return &pb.AuthResponse{Token: ""}, nil
}
