package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hud0shnik/OsuStatsApi/utils"
	"github.com/sirupsen/logrus"
)

// onlineInfo - статус пользователя
type onlineInfo struct {
	Status bool `json:"is_online"`
}

// onlineInfoString - статус пользователя в формате строк
type onlineInfoString struct {
	Status string `json:"is_online"`
}

// getOnlineInfoString возвращает статус пользователя в сети в формате строк, статус код и ошибку
func getOnlineInfoString(id string) (onlineInfoString, int, error) {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return onlineInfoString{}, http.StatusInternalServerError,
			fmt.Errorf("in http.Get: %w", err)
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
		return onlineInfoString{}, resp.StatusCode,
			fmt.Errorf("in http.Get: status code is not 200: %d %s", resp.StatusCode, resp.Status)
	}

	// Запись респонса
	body, _ := io.ReadAll(resp.Body)

	// Полученная страница в формате string
	pageStr := string(body)[80000:]

	// Поиск статуса пользователя и вывод результата
	return onlineInfoString{
		Status: utils.Find(pageStr, "is_online&quot;:", ",", 0),
	}, http.StatusOK, nil

}

// getOnlineInfo возвращает статус пользователя в сети, статус код и ошибку
func getOnlineInfo(id string) (onlineInfo, int, error) {

	// Получение текстовой версии
	resultStr, statusCode, err := getOnlineInfoString(id)
	if err != nil {
		return onlineInfo{}, statusCode, err
	}

	// Перевод в классическую версию
	return onlineInfo{
		Status: utils.ToBool(resultStr.Status),
	}, http.StatusOK, nil

}

// Online - роут "/online"
func Online(w http.ResponseWriter, r *http.Request) {

	// Установка заголовков
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

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
		result, statusCode, err := getOnlineInfoString(id)
		if err != nil {
			w.WriteHeader(statusCode)
			json, _ := json.Marshal(apiError{Error: err.Error()})
			w.Write(json)
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
		result, statusCode, err := getOnlineInfo(id)
		if err != nil {
			w.WriteHeader(statusCode)
			json, _ := json.Marshal(apiError{Error: err.Error()})
			w.Write(json)
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
