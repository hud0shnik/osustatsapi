package handler2

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
	Status  bool   `json:"is_online"`
}

// Структура статуса пользователя для парсинга
type OnlineInfoString struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Status  string `json:"is_online"`
}

// Функция парсинга информации о пользователе
func GetOnlineInfoString(id string) OnlineInfoString {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return OnlineInfoString{
			Success: false,
			Error:   "can't reach osu.ppy.sh",
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
		return OnlineInfoString{
			Success: false,
			Error:   "not found",
		}
	}

	// Поиск статуса пользователя и вывод результата
	return OnlineInfoString{
		Success: true,
		Status:  find(pageStr, "is_online&quot;:", ",", 0),
	}
}

// Функция получения информации о пользователе
func GetOnlineInfo(id string) OnlineInfo {
	resultStr := GetOnlineInfoString(id)
	return OnlineInfo{
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

	// Проверка на наличие параметра
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json, _ := json.Marshal(ApiError{Error: "please insert user id"})
		w.Write(json)
		return
	}

	// Проверка на тип
	if r.URL.Query().Get("type") == "string" {

		// Получение статистики и перевод в json
		result := GetOnlineInfoString(id)
		jsonResp, err := json.Marshal(result)

		// Обработчик ошибок
		switch {
		case err != nil:
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(ApiError{Error: "internal server error"})
			w.Write(json)
			log.Printf("json.Marshal error: %s", err)
		case result.Error == "not found":
			w.WriteHeader(http.StatusNotFound)
			json, _ := json.Marshal(ApiError{Error: "not found"})
			w.Write(json)
		case !result.Success:
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(ApiError{Error: result.Error})
			w.Write(json)
		default:
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
		}

	} else {

		// Получение статистики и перевод в json
		result := GetOnlineInfo(id)
		jsonResp, err := json.Marshal(result)

		// Обработчик ошибок
		switch {
		case err != nil:
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(ApiError{Error: "internal server error"})
			w.Write(json)
			log.Printf("json.Marshal error: %s", err)
		case result.Error == "not found":
			w.WriteHeader(http.StatusNotFound)
			json, _ := json.Marshal(ApiError{Error: "not found"})
			w.Write(json)
		case !result.Success:
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(ApiError{Error: result.Error})
			w.Write(json)
		default:
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
		}

	}

}
