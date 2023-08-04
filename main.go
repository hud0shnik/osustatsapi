package main

import (
	"log"
	"net/http"

	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hud0shnik/OsuStatsApi/api"
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
	router := chi.NewRouter()

	// Маршруты

	router.Get("/api/user", api.User)
	router.Get("/api/online", api.Online)
	router.Get("/api/map", api.Map)
	router.Get("/api/historical", api.Historical)

	// Запуск API
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
