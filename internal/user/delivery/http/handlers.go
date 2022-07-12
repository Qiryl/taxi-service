package http

import (
	"net/http"

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

func (h *Handler) Register(c *gin.Context) {
	var req pb.RegisterRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resp, err := h.UserClient.Register(c, &req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Login(c *gin.Context) {
	var req pb.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resp, err := h.UserClient.Login(c, &req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
