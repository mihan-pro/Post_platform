package main

import (
	Router "go-postgres/router"
	"net/http"
	"os"

	TelegramBot "go-postgres/TelegramBot"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	log := logger.Sugar()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := Router.Router()

	//Запуск бота телеги
	if os.Getenv("MOD") != "develop" {
		go TelegramBot.InitTBot()
	}

	ip := os.Getenv("APP_IP")
	port := os.Getenv("APP_PORT")

	log.Info("App starts on the port: ", ip+":"+port)

	log.Fatal(http.ListenAndServe(ip+":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))

	//log.Error(http.ListenAndServe(ip+":"+port, r))

}
