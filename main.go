package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hud0shnik/OsuStatsApi/api"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {

	// Настройка логгера
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})

	// Вывод времени начала работы
	logrus.Info("API Start")
	logrus.Info("Port: " + os.Getenv("PORT"))

	// Роутер
	router := mux.NewRouter()

	// Маршруты

	router.HandleFunc("/api/user", api.User).Methods("GET")
	router.HandleFunc("/api/online", api.Online).Methods("GET")
	router.HandleFunc("/api/map", api.Map).Methods("GET")
	router.HandleFunc("/api/historical", api.Historical).Methods("GET")

	// Запуск API
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
