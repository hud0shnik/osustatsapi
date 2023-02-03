package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	api "osustatsapi/api"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	// Вывод времени начала работы
	fmt.Println("API Start: " + string(time.Now().Format("2006-01-02 15:04:05")))
	fmt.Println("Port:\t" + os.Getenv("PORT"))

	// Роутер
	router := mux.NewRouter()

	// Маршруты

	router.HandleFunc("/api/user", api.User).Methods("GET")

	router.HandleFunc("/api/online", api.Online).Methods("GET")

	router.HandleFunc("/api/map", api.Map).Methods("GET")

	router.HandleFunc("/api/modding", api.Modding).Methods("GET")

	// Запуск API
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
