package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"osustatsapi/api"
	api2 "osustatsapi/api/v2"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {

	// Настройка логгера
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Вывод времени начала работы
	fmt.Println("API Start: " + string(time.Now().Format("2006-01-02 15:04:05")))
	fmt.Println("Port:\t" + os.Getenv("PORT"))

	// Роутер
	router := mux.NewRouter()

	// Маршруты

	router.HandleFunc("/api/user", api.User).Methods("GET")
	router.HandleFunc("/api/v2/user", api2.User).Methods("GET")

	router.HandleFunc("/api/online", api.Online).Methods("GET")
	router.HandleFunc("/api/v2/online", api2.Online).Methods("GET")

	router.HandleFunc("/api/map", api.Map).Methods("GET")
	router.HandleFunc("/api/v2/map", api2.Map).Methods("GET")

	router.HandleFunc("/api/modding", api.Modding).Methods("GET")
	router.HandleFunc("/api/v2/historical", api2.Historical).Methods("GET")

	// Запуск API
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
