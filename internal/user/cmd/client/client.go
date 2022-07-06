package main

import (
	"log"

	"github.com/Qiryl/taxi-service/internal/user/rest"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Starting client connection
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	// Starting http server
	h := rest.NewHandler(conn)

	router := gin.Default()
	router.POST("/register", h.RegisterHandler)
	router.POST("/login", h.LoginHandler)
	router.Run(":8080")
}
