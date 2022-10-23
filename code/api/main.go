package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"osustatsapi/parse"
	"time"

	"github.com/gorilla/mux"
)

// Функция отправки информации о пользователе в формате строк
func sendUserInfoString(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(parse.GetUserInfoString(mux.Vars(request)["id"], mux.Vars(request)["mode"]))
}

// Функция отправки информации о пользователе
func sendUserInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(parse.GetUserInfo(mux.Vars(request)["id"], mux.Vars(request)["mode"]))
}

// Функция отправки информации о статусе пользователя
func sendOnlineInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(parse.GetOnlineInfo(mux.Vars(request)["id"]))
}

func main() {

	// Вывод времени начала работы
	fmt.Println("API Start: " + string(time.Now().Format("2006-01-02 15:04:05")))
	fmt.Println("Port:\t" + os.Getenv("PORT"))

	// Роутер
	router := mux.NewRouter()

	// Маршруты

	router.HandleFunc("/user", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/", sendUserInfo).Methods("GET")

	router.HandleFunc("/user/{id}", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/{mode}", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/{mode}/", sendUserInfo).Methods("GET")

	router.HandleFunc("/userString", sendUserInfoString).Methods("GET")
	router.HandleFunc("/userString/", sendUserInfoString).Methods("GET")

	router.HandleFunc("/userString/{id}", sendUserInfoString).Methods("GET")
	router.HandleFunc("/userString/{id}/", sendUserInfoString).Methods("GET")
	router.HandleFunc("/userString/{id}/{mode}", sendUserInfoString).Methods("GET")
	router.HandleFunc("/userString/{id}/{mode}/", sendUserInfoString).Methods("GET")

	router.HandleFunc("/online", sendOnlineInfo).Methods("GET")
	router.HandleFunc("/online/", sendOnlineInfo).Methods("GET")

	router.HandleFunc("/online/{id}", sendOnlineInfo).Methods("GET")
	router.HandleFunc("/online/{id}/", sendOnlineInfo).Methods("GET")

	// Запуск API
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
