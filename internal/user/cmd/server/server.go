package main

import (
	"log"
	"net"

	userGrpc "github.com/Qiryl/taxi-service/internal/user/delivery/grpc"
	"github.com/Qiryl/taxi-service/internal/user/repo/psql"
	"github.com/Qiryl/taxi-service/internal/user/usecase"
	userPb "github.com/Qiryl/taxi-service/proto/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	db, err := sqlx.Open("postgres", "postgres://pguser:pass@pg_db:5432/userdb?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	userRepo := psql.NewPsqlUserRepo(db)
	userUc := usecase.NewUserUsecase(userRepo)
	s := userGrpc.NewGrpcUserServer(userUc)

	// Starting grpc server
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("FAILED TO LISTEN: %v", err)
	}

	grpcServer := grpc.NewServer()
	userPb.RegisterUserServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
