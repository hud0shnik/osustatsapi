package main

import (
	"log"
	"net/http"
	"os"
	"osustatsapi/api"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {

	// Настройка логгера
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Вывод времени начала работы
	logrus.Info("API Start: " + string(time.Now().Format("2006-01-02 15:04:05")))
	logrus.Info("Port:\t" + os.Getenv("PORT"))

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
