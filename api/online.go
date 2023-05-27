package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// Структура статуса пользователя
type onlineInfo struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Status  bool   `json:"is_online"`
}

// Структура статуса пользователя для парсинга
type onlineInfoString struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Status  string `json:"is_online"`
}

// Функция парсинга информации о пользователе
func getOnlineInfoString(id string) onlineInfoString {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return onlineInfoString{
			Success: false,
			Error:   "http get error",
		}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// Полученная страница в формате string
	pageStr := string(body)[80000:]

	// Сохранение html"ки в файл sample.html (для тестов)
	/*
		if err := os.WriteFile("sample.html", []byte(pageStr), 0666); err != nil {
			logrus.Fatal(err)
		}
	*/

	// Проверка на страницу пользователя
	if !strings.Contains(pageStr, "js-react--profile") {
		return onlineInfoString{
			Success: false,
			Error:   "user not found",
		}
	}

	// Поиск статуса пользователя и вывод результата
	return onlineInfoString{
		Success: true,
		Status:  find(pageStr, "is_online&quot;:", ",", 0),
	}
}

// Функция получения информации о пользователе
func getOnlineInfo(id string) onlineInfo {
	resultStr := getOnlineInfoString(id)
	return onlineInfo{
		Success: resultStr.Success,
		Error:   resultStr.Error,
		Status:  toBool(resultStr.Status),
	}
}

// Роут "/online"
func Online(w http.ResponseWriter, r *http.Request) {

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Получение параметра id из реквеста
	id := r.URL.Query().Get("id")

	// Если параметра нет, отправка ошибки
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json, _ := json.Marshal(map[string]string{"Error": "Please insert user id"})
		w.Write(json)
		return
	}

	// Проверка на тип, получение статистики, форматирование и отправка
	if r.URL.Query().Get("type") == "string" {
		jsonResp, err := json.Marshal(getOnlineInfoString(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(map[string]string{"Error": "Internal Server Error"})
			w.Write(json)
			logrus.Printf("json.Marshal error: %s", err)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
		}
	} else {
		jsonResp, err := json.Marshal(getOnlineInfo(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(map[string]string{"Error": "Internal Server Error"})
			w.Write(json)
			logrus.Printf("json.Marshal error: %s", err)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
		}
	}

}
