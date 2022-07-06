package main

import (
	"log"
	"net"

	userGrpc "github.com/Qiryl/taxi-service/internal/user/grpc"
	userPb "github.com/Qiryl/taxi-service/proto/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	// Starting grpc server
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Open connection to db
	db, err := sqlx.Open("postgres", "postgres://pguser:pass@pg_db:5432/userdb?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	s := userGrpc.NewUserServer(db)
	grpcServer := grpc.NewServer()
	userPb.RegisterUserServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
