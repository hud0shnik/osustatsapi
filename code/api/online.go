package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Статуса пользователя
type OnlineInfo struct {
	Error  string `json:"error"`
	Status string `json:"is_online"`
}

// Функция получения информации о пользователе
func GetOnlineInfo(id string) OnlineInfo {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return OnlineInfo{
			Error: "http get error",
		}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// HTML полученной страницы в формате string
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
			Error: "user not found",
		}
	}

	// Поиск статуса пользователя и вывод результата
	return OnlineInfo{
		Status: find(pageStr, "is_online&quot;:", ",", 0),
	}
}

// Роут "/online"  для vercel
func Online(w http.ResponseWriter, r *http.Request) {

	// Получение параметра id из реквеста
	id := r.URL.Query().Get("id")

	// Если параметра нет, отправка ошибки
	if id == "" {
		http.NotFound(w, r)
		return
	}

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Получение статистики, форматирование и отправка
	jsonResp, err := json.Marshal(GetOnlineInfo(id))
	if err != nil {
		fmt.Print("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}
}
