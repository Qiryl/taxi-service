package http

import (
	"net/http"

	"github.com/Qiryl/taxi-service/internal/user/domain"
	"github.com/Qiryl/taxi-service/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.UserUsecase
}

func NewHandler(uc usecase.UserUsecase) Handler {
	return Handler{uc: uc}
}

func (h *Handler) Register(c *gin.Context) {
	var inp registerInp

	err := c.BindJSON(&inp)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.uc.Register(c, &domain.User{
		Name:     inp.Name,
		Phone:    inp.Phone,
		Email:    inp.Email,
		Password: inp.Password,
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// TODO: Return success message
	c.JSON(http.StatusOK, "")
}

func (h *Handler) Login(c *gin.Context) {
	var inp loginInp
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user, err := h.uc.Login(c, &domain.LoginRequest{
		Phone:    inp.Phone,
		Password: inp.Password,
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &loginOut{
		Name:     user.Name,
		Phone:    user.Phone,
		Email:    user.Email,
		Password: user.Password,
		Token:    "",
	})
}
