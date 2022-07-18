package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Qiryl/taxi-service/internal/user/config"
	"github.com/Qiryl/taxi-service/internal/user/delivery/http"
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

	router := http.NewRouter(db)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if err := router.Run(":" + httpCfg.Port); err != nil {
		log.Fatalln(err)
	}

	<-done
}
