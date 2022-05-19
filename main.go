package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Структура для хранения полной информации о пользователе
type UserInfo struct {
	UserID       string `json:"id"`
	Username     string `json:"username"`
	GlobalRank   string `json:"globalRank"`
	CountryRank  string `json:"countryRank"`
	Accuracy     string `json:"accuracy"`
	PlayCount    string `json:"playCount"`
	PlayTime     string `json:"playTime"` // В секундах
	TotalScore   string `json:"totalScore"`
	TotalHits    string `json:"totalHits"`
	MaximumCombo string `json:"maximumCombo"`
	Replays      string `json:"replays"`
	X            string `json:"ss"`
	XH           string `json:"xh"`
}

// Функция поиска. Возвращает искомое значение и индекс
func find(str string, subStr string, char byte) (string, int) {

	// Поиск индекса начала нужной строки
	left := strings.Index(str, subStr) + len(subStr)

	// Проверка на существование нужной строки
	if left > len(subStr)-1 {

		// Крайняя часть искомой строки
		right := left

		for ; str[right] != char; right++ {
			// Доводит str[right] до символа char
		}
		return str[left:right], right
	}

	return "", 0
}

// Функция получения информации о пользователе
func getUserInfo(id string) UserInfo {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return UserInfo{}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// HTML полученной страницы в формате string
	pageStr := string(body)

	// Сохранение html'ки в файл sample.html
	/*if err := os.WriteFile("sample.html", body, 0666); err != nil {
		log.Fatal(err)
	}*/

	// Структура, которую будет возвращать функция
	result := UserInfo{
		UserID: id,
	}

	// Обрезка юзелесс части html'ки
	pageStr = pageStr[strings.Index(pageStr, "js-react--profile-page osu-layout osu-layout--full\""):]

	i := 0

	/* -----------------------------------------------------------
	# Далее происходит заполнение полей функцией find			 #
	# после каждого поиска тело сайта обрезается для оптимизации #
	------------------------------------------------------------ */

	// Юзернейм
	result.Username, i = find(pageStr, "username&quot;:&quot;", '&')
	pageStr = pageStr[i:]

	// Глобальный рейтинг
	result.GlobalRank, i = find(pageStr, "global_rank&quot;:", ',')
	pageStr = pageStr[i:]

	// Точность попаданий
	result.Accuracy, i = find(pageStr, "hit_accuracy&quot;:", ',')
	pageStr = pageStr[i:]

	// Количество игр
	result.PlayCount, i = find(pageStr, "play_count&quot;:", ',')
	pageStr = pageStr[i:]

	// Времени в игре
	result.PlayTime, i = find(pageStr, "play_time&quot;:", ',')
	pageStr = pageStr[i:]

	// Рейтинговые очки
	result.TotalScore, i = find(pageStr, "total_score&quot;:", ',')
	pageStr = pageStr[i:]

	// Всего попаданий
	result.TotalHits, i = find(pageStr, "total_hits&quot;:", ',')
	pageStr = pageStr[i:]

	// Максимальное комбо
	result.MaximumCombo, i = find(pageStr, "maximum_combo&quot;:", ',')
	pageStr = pageStr[i:]

	// Реплеев просмотрено другими
	result.Replays, i = find(pageStr, "replays_watched_by_others&quot;:", ',')
	pageStr = pageStr[i:]

	// X'ки
	result.X, i = find(pageStr, "grade_counts&quot;:{&quot;ss&quot;:", ',')
	pageStr = pageStr[i:]

	// XH'ки
	result.XH, i = find(pageStr, "ssh&quot;:", ',')
	pageStr = pageStr[i:]

	// Рейтинг в стране
	result.CountryRank, _ = find(pageStr, "country_rank&quot;:", ',')

	return result
}

// Функция отправки информации о пользователе
func sendUserInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(getUserInfo(mux.Vars(request)["id"]))
}

func main() {

	// Вывод времени начала работы
	fmt.Println("API Start: " + string(time.Now().Format("2006-01-02 15:04:05")))

	// Роутер
	router := mux.NewRouter()

	// Маршруты

	router.HandleFunc("/user/{id}", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/", sendUserInfo).Methods("GET")

	// Запуск API

	// Для Heroku
	//log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

	// Для локалхоста (127.0.0.1:8080/)
	log.Fatal(http.ListenAndServe(":8080", router))
}
