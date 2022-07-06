package grpc

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "github.com/Qiryl/taxi-service/proto/user"
	"github.com/jmoiron/sqlx"
)

type UserServer struct {
	pb.UnimplementedUserServer
	db *sqlx.DB
}

func NewUserServer(db *sqlx.DB) *UserServer {
	return &UserServer{db: db}
}

func (s *UserServer) connect(ctx context.Context) (*sqlx.Conn, error) {
	c, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	c, err := s.connect(ctx)
	if err != nil {
		fmt.Println("Login conn err", err)
		return nil, err
	}
	defer c.Close()

	var pass string
	err = s.db.SelectContext(ctx, pass, "SELECT user_password FROM users WHERE user_phone = $1", req.Phone)
	if err != nil {
		fmt.Println("Login db query err:", err)
		return nil, err
	}

	if pass == req.Password {
		return &pb.AuthResponse{Token: ""}, nil
	} else {
		return nil, errors.New("invalid password")
	}
}

func (s *UserServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	c, err := s.connect(ctx)
	if err != nil {
		log.Fatalln("REGISTER CONNECTION ERROR:", err)
		return nil, err
	}
	defer c.Close()

	query := "INSERT INTO users (user_name, user_phone, user_email, user_password) VALUES ($1, $2, $3, $4)"
	_, err = s.db.ExecContext(ctx, query, req.Username, req.Phone, req.Email, req.Password)
	if err != nil {
		log.Println("REGISTER DB QUERY ERR", err)
		return nil, err
	}

	return &pb.AuthResponse{Token: ""}, nil
}
