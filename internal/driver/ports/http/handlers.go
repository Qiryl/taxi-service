package http

import (
	"fmt"
	"net/http"

	"github.com/Qiryl/taxi-service/internal/driver/app"
	"github.com/Qiryl/taxi-service/internal/driver/dtos"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	driver app.DriverService
	// order
}

func (h *Handler) SignUp(ctx *gin.Context) {
	var driverDto *dtos.DriverDTO
	err := ctx.BindJSON(driverDto)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.driver.SignUp(ctx, driverDto)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK,
		fmt.Sprintf("User: %s with phone number: %s created successfully", driverDto.Name, driverDto.Phone))
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var loginDto *dtos.LoginDTO
	err := ctx.BindJSON(loginDto)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	driverDto, err := h.driver.SignIn(ctx, loginDto)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, driverDto)
}
