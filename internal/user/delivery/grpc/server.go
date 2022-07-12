package grpc

import (
	"log"
	"net"

	"github.com/Qiryl/taxi-service/internal/user/config"
	"github.com/Qiryl/taxi-service/internal/user/repo/psql"
	"github.com/Qiryl/taxi-service/internal/user/usecase"
	userPb "github.com/Qiryl/taxi-service/proto/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func ListenAndServe(grpcCfg *config.GrpcConfig, dbCfg *config.PostgresConfig) {
	db, err := sqlx.Open(dbCfg.Driver, dbCfg.Url)
	if err != nil {
		log.Println(err)
	}

	userRepo := psql.NewPsqlUserRepo(db)
	userUc := usecase.NewUserUsecase(userRepo)
	s := NewGrpcUserServer(userUc)

	lis, err := net.Listen("tcp", ":"+grpcCfg.Port)
	if err != nil {
		log.Println(err)
	}

	grpcServer := grpc.NewServer()
	userPb.RegisterUserServer(grpcServer, s)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()
}
