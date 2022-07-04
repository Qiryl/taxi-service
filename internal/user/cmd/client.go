package main

import (
	"fmt"
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
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("GRPC client started")

	// Starting http server
	h := rest.NewHandler(conn)

	router := gin.Default()
	router.POST("/register", h.RegisterHandler)
	router.POST("/login", h.LoginHandler)
	router.Run("localhost:8080")
}
