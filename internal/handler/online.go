package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/hud0shnik/osustatsapi/utils"
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
			fmt.Errorf("in http.Get: %s", resp.Status)
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

	// Получение параметра id из реквеста
	id := r.URL.Query().Get("id")

	// Проверка на наличие параметра
	if id == "" {
		response(w, http.StatusBadRequest, apiError{Error: "please insert user id"})
		return
	}

	// Проверка на тип
	if r.URL.Query().Get("type") == "string" {

		// Получение статистики
		result, statusCode, err := getOnlineInfoString(id)
		if err != nil {
			response(w, statusCode, apiError{Error: err.Error()})
			return
		}

		response(w, statusCode, result)

	} else {

		// Получение статистики
		result, statusCode, err := getOnlineInfo(id)
		if err != nil {
			response(w, statusCode, apiError{Error: err.Error()})
			return
		}

		response(w, statusCode, result)

	}

}
