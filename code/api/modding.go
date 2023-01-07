package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ---------------------- Структуры для парсинга ------------------------

// Структура респонса
type ModdingResponseString struct {
	Success string `json:"success"`
	Error   string `json:"error"`
}

// Функция получения текстовой информации
func GetModdingInfoString(id string) ModdingResponseString {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id + "/modding")
	if err != nil {
		return ModdingResponseString{
			Success: "false",
			Error:   "http get error",
		}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// Полученная страница в формате string
	pageStr := string(body)[80000:]

	// Проверка на страницу пользователя
	if strings.Contains(pageStr, "<h1>User not found! ;_;</h1>") {
		return ModdingResponseString{
			Success: "false",
			Error:   "user not found",
		}
	}

	// Обрезка юзелесс части html"ки
	pageStr = pageStr[strings.Index(pageStr, "<script id=\"json-events\" type=\"application/json\">"):]

	// Сохранение html"ки в файл sample.html (для тестов)

	/*if err := os.WriteFile("sample2.html", []byte(pageStr), 0666); err != nil {
		log.Fatal(err)
	}*/

	// Структура, которую будет возвращать функция
	result := ModdingResponseString{}

	// Крайняя левая граница поиска
	//left := 0

	return result
}

// Роут "/modding"  для vercel
func Modding(w http.ResponseWriter, r *http.Request) {

	// Получение параметра id из реквеста
	id := r.URL.Query().Get("id")

	// Если параметра нет, отправка ошибки
	if id == "" {
		http.NotFound(w, r)
		return
	}

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Формирование и отправка статистики
	jsonResp, err := json.Marshal(GetModdingInfoString(id))
	if err != nil {
		fmt.Print("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}
}
