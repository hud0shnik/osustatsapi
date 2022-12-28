package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	api "osustatsapi/api"
	"time"

	"github.com/gorilla/mux"
)

// Функция отправки информации о пользователе в формате строк
func sendUserInfoString(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(api.GetUserInfoString(request.URL.Query().Get("id")))
}

// Функция отправки информации о пользователе
func sendUserInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(api.GetUserInfo(request.URL.Query().Get("id")))
}

// Функция отправки информации о статусе пользователя
func sendOnlineInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(api.GetOnlineInfo(request.URL.Query().Get("id")))
}

// Функция отправки информации о карте в формате строк
func sendMapInfoString(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(api.GetMapInfoString(request.URL.Query().Get("beatmapset"), request.URL.Query().Get("id")))
}

// Функция отправки информации о карте
func sendMapInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(api.GetMapInfo(request.URL.Query().Get("beatmapset"), request.URL.Query().Get("id")))
}

func main() {

	// Вывод времени начала работы
	fmt.Println("API Start: " + string(time.Now().Format("2006-01-02 15:04:05")))
	fmt.Println("Port:\t" + os.Getenv("PORT"))

	// Роутер
	router := mux.NewRouter()

	// Маршруты

	router.HandleFunc("/api/user", sendUserInfo).Methods("GET")

	router.HandleFunc("/api/userstring", sendUserInfoString).Methods("GET")

	router.HandleFunc("/api/online", sendOnlineInfo).Methods("GET")

	router.HandleFunc("/api/mapstring", sendMapInfoString).Methods("GET")

	router.HandleFunc("/api/map", sendMapInfo).Methods("GET")

	// Запуск API
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
