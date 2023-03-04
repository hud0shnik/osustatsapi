package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Структура статуса пользователя
type OnlineInfo struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Status  string `json:"is_online"`
}

// Функция получения информации о пользователе
func GetOnlineInfo(id string) OnlineInfo {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return OnlineInfo{
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
			log.Fatal(err)
		}
	*/

	// Проверка на страницу пользователя
	if !strings.Contains(pageStr, "js-react--profile") {
		return OnlineInfo{
			Success: false,
			Error:   "user not found",
		}
	}

	// Поиск статуса пользователя и вывод результата
	return OnlineInfo{
		Success: true,
		Status:  find(pageStr, "is_online&quot;:", ",", 0),
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

	// Получение статистики, форматирование и отправка
	jsonResp, err := json.Marshal(GetOnlineInfo(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(map[string]string{"Error": "Internal Server Error"})
		w.Write(json)
		log.Printf("json.Marshal error: %s", err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}
}
