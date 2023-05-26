package api2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Структура статуса пользователя
type onlineInfo struct {
	Status bool `json:"is_online"`
}

// Структура статуса пользователя для парсинга
type onlineInfoString struct {
	Status string `json:"is_online"`
}

// Функция парсинга информации о пользователе
func getOnlineInfoString(id string) (onlineInfoString, error) {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return onlineInfoString{}, fmt.Errorf("in http.Get: %w", err)
	}
	defer resp.Body.Close()

	// Сохранение html"ки в файл sample.html (для тестов)
	/*
		if err := os.WriteFile("sample.html", []byte(pageStr), 0666); err != nil {
			logrus.Fatal(err)
		}
	*/

	// Проверка статускода
	if resp.StatusCode != 200 {
		return onlineInfoString{}, fmt.Errorf("response status: %s", resp.Status)
	}

	// Запись респонса
	body, _ := ioutil.ReadAll(resp.Body)

	// Полученная страница в формате string
	pageStr := string(body)[80000:]

	// Поиск статуса пользователя и вывод результата
	return onlineInfoString{
		Status: find(pageStr, "is_online&quot;:", ",", 0),
	}, nil

}

// Функция получения информации о пользователе
func getOnlineInfo(id string) (onlineInfo, error) {

	// Получение текстовой версии
	resultStr, err := getOnlineInfoString(id)
	if err != nil {
		return onlineInfo{}, err
	}

	// Перевод в классическую версию
	return onlineInfo{
		Status: toBool(resultStr.Status),
	}, nil

}

// Роут "/online"
func Online(w http.ResponseWriter, r *http.Request) {

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Получение параметра id из реквеста
	id := r.URL.Query().Get("id")

	// Проверка на наличие параметра
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json, _ := json.Marshal(apiError{Error: "please insert user id"})
		w.Write(json)
		return
	}

	// Проверка на тип
	if r.URL.Query().Get("type") == "string" {

		// Получение статистики
		result, err := getOnlineInfoString(id)
		if err != nil {
			if err.Error() == "response status: 404 Not Found" {
				w.WriteHeader(http.StatusNotFound)
				json, _ := json.Marshal(apiError{Error: "not found"})
				w.Write(json)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: "internal server error"})
			w.Write(json)
			logrus.Printf("getOnlineInfo err: %s", err)
			return
		}

		// Перевод в json
		jsonResp, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: "internal server error"})
			w.Write(json)
			logrus.Printf("json.Marshal err: %s", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)

	} else {

		// Получение статистики
		result, err := getOnlineInfo(id)
		if err != nil {
			if err.Error() == "response status: 404 Not Found" {
				w.WriteHeader(http.StatusNotFound)
				json, _ := json.Marshal(apiError{Error: "not found"})
				w.Write(json)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: "internal server error"})
			w.Write(json)
			logrus.Printf("getOnlineInfo err: %s", err)
			return
		}

		// Перевод в json
		jsonResp, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: "internal server error"})
			w.Write(json)
			logrus.Printf("json.Marshal err: %s", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)

	}

}
