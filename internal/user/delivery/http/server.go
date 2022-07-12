package http

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Qiryl/taxi-service/internal/user/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ListenAndServe(httpCfg *config.HttpConfig, grpcCfg *config.GrpcConfig) {
	grpcClientConn, err := grpc.Dial(":"+grpcCfg.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	defer grpcClientConn.Close()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	router := gin.Default()
	handler := NewHandler(grpcClientConn)
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	if err := router.Run(":" + httpCfg.Port); err != nil {
		log.Fatalln(err)
	}
}
