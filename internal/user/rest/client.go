package rest

import (
	"fmt"
	"log"

	pb "github.com/Qiryl/taxi-service/proto/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Handler struct {
	pb.UserClient
}

func NewHandler(userConn *grpc.ClientConn) Handler {
	return Handler{pb.NewUserClient(userConn)}

}

func (h *Handler) RegisterHandler(c *gin.Context) {
	var req pb.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Println("BindJSON error:", err)
	}
	fmt.Println("Register req", &req)

	resp, err := h.UserClient.Register(c, &req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Register auth req", &resp)

	c.JSON(200, resp)
}

func (h *Handler) LoginHandler(c *gin.Context) {
	var req pb.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Register req", &req)

	resp, err := h.UserClient.Login(c, &req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Register auth req", &resp)

	c.JSON(200, resp)
}
