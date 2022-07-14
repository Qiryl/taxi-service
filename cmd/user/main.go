package main

import (
	"log"

	"github.com/Qiryl/taxi-service/internal/user/config"
	"github.com/Qiryl/taxi-service/internal/user/delivery/http"
	"github.com/Qiryl/taxi-service/internal/user/repo/psql"
	"github.com/Qiryl/taxi-service/internal/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	dbCfg, err := config.GetEnvPostgresConfig()
	if err != nil {
		log.Fatalln(err)
	}

	httpCfg, err := config.GetEnvHttpConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sqlx.Open(dbCfg.Driver, dbCfg.Url)
	if err != nil {
		log.Println(err)
	}

	userRepo := psql.NewPsqlUserRepo(db)
	userUc := usecase.NewUserUsecase(userRepo)

	router := gin.Default()
	handler := http.NewHandler(*userUc)
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	if err := router.Run(":" + httpCfg.Port); err != nil {
		log.Fatalln(err)
	}
}
