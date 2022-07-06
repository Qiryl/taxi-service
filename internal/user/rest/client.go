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

type RegisterReq struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) RegisterHandler(c *gin.Context) {
	var req pb.RegisterRequest

	err := c.BindJSON(&req)
	if err != nil {
		log.Fatalln("BindJSON error:", err)
	}

	fmt.Println("IN REGISTER HANDLER")
	resp, err := h.UserClient.Register(c, &req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Register auth req", &resp)

	// c.JSON(200, resp)
}

func (h *Handler) LoginHandler(c *gin.Context) {
	var req pb.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err.Error())
	}
	log.Println("Login req", &req)

	resp, err := h.UserClient.Login(c, &req)
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println("Login auth req", &resp)

	c.JSON(200, resp)
}
