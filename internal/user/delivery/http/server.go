package http

import (
	"github.com/Qiryl/taxi-service/internal/user/repo/psql"
	"github.com/Qiryl/taxi-service/internal/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) *gin.Engine {
	userRepo := psql.NewPsqlUserRepo(db)
	userUc := usecase.NewUserUsecase(userRepo)

	router := gin.Default()
	handler := NewHandler(*userUc)
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	return router
}
