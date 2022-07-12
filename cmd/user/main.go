package main

import (
	"log"

	"github.com/Qiryl/taxi-service/internal/user/config"
	"github.com/Qiryl/taxi-service/internal/user/delivery/grpc"
	"github.com/Qiryl/taxi-service/internal/user/delivery/http"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	grpcCfg, err := config.GetEnvGrpcConfig()
	if err != nil {
		log.Fatalln(err)
	}

	dbCfg, err := config.GetEnvPostgresConfig()
	if err != nil {
		log.Fatalln(err)
	}

	httpCfg, err := config.GetEnvHttpConfig()
	if err != nil {
		log.Fatalln(err)
	}

	grpc.ListenAndServe(grpcCfg, dbCfg)
	http.ListenAndServe(httpCfg, grpcCfg)
}
