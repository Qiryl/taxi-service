package grpc

import (
	"github.com/Qiryl/taxi-service/internal/user/config"
	_ "github.com/lib/pq"
)

func ListenAndServe(grpcCfg *config.GrpcConfig, dbCfg *config.PostgresConfig) {

	// userRepo := psql.NewPsqlUserRepo(db)
	// userUc := usecase.NewUserUsecase(userRepo)
	// s := NewGrpcUserServer(userUc)

	// lis, err := net.Listen("tcp", ":"+grpcCfg.Port)
	// if err != nil {
	// 	log.Println(err)
	// }

	// grpcServer := grpc.NewServer()
	// userPb.RegisterUserServer(grpcServer, s)

	// go func() {
	// 	if err = grpcServer.Serve(lis); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }()
}
