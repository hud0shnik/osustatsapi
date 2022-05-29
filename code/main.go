package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Структура для хранения полной информации о пользователе
type UserInfo struct {
	Username                 string  `json:"username"`
	Names                    string  `json:"previous_usernames"`
	AvatarUrl                string  `json:"avatar_url"`
	UserID                   string  `json:"id"`
	CountryCode              string  `json:"country_code"`
	GlobalRank               string  `json:"global_rank"`
	CountryRank              string  `json:"country_rank"`
	PP                       string  `json:"pp"`
	PlayTime                 string  `json:"play_time"` // В секундах
	SSH                      string  `json:"ssh"`
	SS                       string  `json:"ss"`
	SH                       string  `json:"sh"`
	S                        string  `json:"s"`
	A                        string  `json:"a"`
	RankedScore              string  `json:"ranked_score"`
	Accuracy                 string  `json:"accuracy"`
	PlayCount                string  `json:"play_count"`
	TotalScore               string  `json:"total_score"`
	TotalHits                string  `json:"total_hits"`
	MaximumCombo             string  `json:"maximum_combo"`
	Replays                  string  `json:"replays"`
	Level                    string  `json:"level"`
	SupportLvl               string  `json:"support_level"`
	DefaultGroup             string  `json:"default_group"`
	IsOnline                 string  `json:"is_online"`
	IsActive                 string  `json:"is_active"`
	IsDeleted                string  `json:"is_deleted"`
	IsBot                    string  `json:"is_bot"`
	IsSupporter              string  `json:"is_supporter"`
	LastVisit                string  `json:"last_visit"`
	ProfileColor             string  `json:"profile_color"`
	RankedBeatmapsetCount    string  `json:"ranked_beatmapset_count"`
	PendingBeatmapsetCount   string  `json:"pending_beatmapset_count"`
	PmFriendsOnly            string  `json:"pm_friends_only"`
	GraveyardBeatmapsetCount string  `json:"graveyard_beatmapset_count"`
	BestBeatMap              beatMap `json:"best_beat_map"`
}

// Структура для проверки статуса пользователя
type OnlineInfo struct {
	Status string `json:"is_online"`
}

// Структура для хранения информации о мапе
type beatMap struct {
	Id               string `json:"id"`
	Rank             string `json:"rank"`
	EndedAt          string `json:"ended_at"`
	Accuracy         string `json:"accuracy"`
	MaximumCombo     string `json:"maximum_combo"`
	PP               string `json:"pp"`
	Passed           string `json:"passed"`
	TotalScore       string `json:"total_score"`
	LegacyPerfect    string `json:"legacy_perfect"`
	Replay           string `json:"replay"`
	DifficultyRating string `json:"difficulty_rating"`
	Mode             string `json:"mode"`
	Status           string `json:"status"`
	TotalLength      string `json:"total_length"`
	Ar               string `json:"ar"`
	Bpm              string `json:"bpm"`
	Convert          string `json:"convert"`
	CountCircles     string `json:"count_circles"`
	CountSliders     string `json:"count_sliders"`
	CountSpinners    string `json:"count_spinners"`
	Cs               string `json:"cs"`
	DeletedAt        string `json:"deleted_at"`
	Drain            string `json:"drain"`
	HitLength        string `json:"hit_length"`
	IsScoreable      string `json:"is_scoreable"`
	LastUpdated      string `json:"last_updated"`
	ModeInt          string `json:"mode_int"`
	PassCount        string `json:"pass_count"`
	PlayCount        string `json:"play_count"`
	Ranked           string `json:"ranked"`
	Url              string `json:"url"`
	Checksum         string `json:"checksum"`
	Creator          string `json:"creator"`
	FavoriteCount    string `json:"favorite_count"`
	Hype             string `json:"hype"`
	Nsfw             string `json:"nsfw"`
	Offset           string `json:"offset"`
	Spotlight        string `json:"spotlight"`
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
func getUserInfo(id, mode string) UserInfo {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id + "/" + mode)
	if err != nil {
		return UserInfo{}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// HTML полученной страницы в формате string
	pageStr := string(body)

	// Проверка на страницу пользователя
	if !strings.Contains(pageStr, "js-react--profile") {
		return UserInfo{}
	}

	// Обрезка юзелесс части html'ки
	pageStr = pageStr[strings.Index(pageStr, "current_mode"):]
	pageStr = strings.ReplaceAll(pageStr, "&quot;", " ")

	// Сохранение html'ки в файл sample.html

	if err := os.WriteFile("sample.html", []byte(pageStr), 0666); err != nil {
		log.Fatal(err)
	}

	// Структура, которую будет возвращать функция
	result := UserInfo{
		UserID: id,
	}

	i := 0

	/* -----------------------------------------------------------
	# Далее происходит заполнение полей функцией find			 #
	# после каждого поиска тело сайта обрезается для оптимизации #
	------------------------------------------------------------ */

	// Лучшая мапа
	result.BestBeatMap.Accuracy, _ = find(pageStr, "accuracy :", ',')
	result.BestBeatMap.Id, _ = find(pageStr, "beatmap_id :", ',')
	result.BestBeatMap.EndedAt, _ = find(pageStr, "ended_at : ", ' ')
	result.BestBeatMap.MaximumCombo, _ = find(pageStr, "max_combo :", ',')
	result.BestBeatMap.Passed, _ = find(pageStr, "passed :", ',')
	result.BestBeatMap.Rank, _ = find(pageStr, "rank : ", ' ')
	result.BestBeatMap.TotalScore, _ = find(pageStr, "total_score :", ',')
	result.BestBeatMap.LegacyPerfect, _ = find(pageStr, "legacy_perfect :", ',')
	result.BestBeatMap.PP, _ = find(pageStr, "pp :", ',')
	result.BestBeatMap.Replay, _ = find(pageStr, "replay :", ',')
	result.BestBeatMap.DifficultyRating, _ = find(pageStr, "difficulty_rating :", ',')
	result.BestBeatMap.Mode, _ = find(pageStr, "mode : ", ' ')
	result.BestBeatMap.Status, _ = find(pageStr, "status : ", ' ')
	result.BestBeatMap.TotalLength, _ = find(pageStr, "total_length :", ',')
	result.BestBeatMap.Ar, _ = find(pageStr, "ar :", ',')
	result.BestBeatMap.Bpm, _ = find(pageStr, "bpm :", ',')
	result.BestBeatMap.Convert, _ = find(pageStr, "convert :", ',')
	result.BestBeatMap.CountCircles, _ = find(pageStr, "count_circles :", ',')
	result.BestBeatMap.CountSliders, _ = find(pageStr, "count_sliders :", ',')
	result.BestBeatMap.CountSpinners, _ = find(pageStr, "count_spinners :", ',')
	result.BestBeatMap.Cs, _ = find(pageStr, " cs :", ',')
	result.BestBeatMap.DeletedAt, _ = find(pageStr, "deleted_at :", ',')
	result.BestBeatMap.Drain, _ = find(pageStr, "drain :", ',')
	result.BestBeatMap.HitLength, _ = find(pageStr, "hit_length :", ',')
	result.BestBeatMap.IsScoreable, _ = find(pageStr, "is_scoreable :", ',')
	result.BestBeatMap.LastUpdated, _ = find(pageStr, "last_updated : ", ' ')
	result.BestBeatMap.ModeInt, _ = find(pageStr, "mode_int :", ',')
	result.BestBeatMap.PassCount, _ = find(pageStr, "passcount :", ',')
	result.BestBeatMap.PlayCount, _ = find(pageStr, "playcount :", ',')
	result.BestBeatMap.Ranked, _ = find(pageStr, "ranked :", ',')
	result.BestBeatMap.Url, _ = find(pageStr, "url : ", ' ')
	result.BestBeatMap.Url = strings.ReplaceAll(result.BestBeatMap.Url, "\\", "")
	result.BestBeatMap.Checksum, i = find(pageStr, "checksum : ", ' ')
	pageStr = pageStr[i:]
	result.BestBeatMap.Creator, _ = find(pageStr, "creator : ", ' ')
	result.BestBeatMap.FavoriteCount, _ = find(pageStr, "favourite_count :", ',')
	result.BestBeatMap.Hype, _ = find(pageStr, "hype :", ',')
	result.BestBeatMap.Nsfw, _ = find(pageStr, "nsfw :", ',')
	result.BestBeatMap.Offset, _ = find(pageStr, "offset :", ',')
	result.BestBeatMap.Spotlight, _ = find(pageStr, "spotlight :", ',')

	pageStr = pageStr[i:]

	// В последний раз был в сети
	result.LastVisit, _ = find(pageStr, "last_visit :", ',')

	// Сообщения только от друзей
	result.PmFriendsOnly, _ = find(pageStr, "pm_friends_only :", ',')

	// Ссылка на аватар
	result.AvatarUrl, i = find(pageStr, "avatar_url : ", ' ')
	result.AvatarUrl = strings.ReplaceAll(result.AvatarUrl, "\\", "")
	pageStr = pageStr[i:]

	// Код страны
	result.CountryCode, _ = find(pageStr, "country_code : ", ' ')

	// Группа
	result.DefaultGroup, _ = find(pageStr, "default_group : ", ' ')

	// Активность
	result.IsActive, _ = find(pageStr, "is_active :", ',')

	// Бот
	result.IsBot, _ = find(pageStr, "is_bot :", ',')

	// Удалённый профиль
	result.IsDeleted, _ = find(pageStr, "is_deleted :", ',')

	// Статус в сети
	result.IsOnline, _ = find(pageStr, "is_online :", ',')

	// Подписка
	result.IsSupporter, _ = find(pageStr, "is_supporter :", ',')

	// Цвет профиля
	result.ProfileColor, _ = find(pageStr, "profile_colour :", ',')

	// Юзернейм
	result.Username, i = find(pageStr, "username : ", ' ')
	pageStr = pageStr[i:]

	// Заброшенные карты
	result.GraveyardBeatmapsetCount, _ = find(pageStr, "graveyard_beatmapset_count", ',')

	// Карты на рассмотрении
	result.PendingBeatmapsetCount, _ = find(pageStr, "pending_beatmapset_count :", ',')

	// Юзернеймы
	result.Names, _ = find(pageStr, "previous_usernames :[ ", ']')

	// Рейтинговые и одобренные карты
	result.RankedBeatmapsetCount, _ = find(pageStr, "ranked_beatmapset_count :", ',')

	// Уровень
	result.Level, i = find(pageStr, "level :{ current :", ',')
	pageStr = pageStr[i:]

	// Глобальный рейтинг
	result.GlobalRank, i = find(pageStr, "global_rank :", ',')
	pageStr = pageStr[i:]

	// PP-хи
	result.PP, i = find(pageStr, "pp :", ',')
	pageStr = pageStr[i:]

	// Всего очков
	result.RankedScore, i = find(pageStr, "ranked_score :", ',')
	pageStr = pageStr[i:]

	// Точность попаданий
	result.Accuracy, i = find(pageStr, "hit_accuracy :", ',')
	pageStr = pageStr[i:]

	// Количество игр
	result.PlayCount, i = find(pageStr, "play_count :", ',')
	pageStr = pageStr[i:]

	// Времени в игре
	result.PlayTime, i = find(pageStr, "play_time :", ',')
	pageStr = pageStr[i:]

	// Рейтинговые очки
	result.TotalScore, i = find(pageStr, "total_score :", ',')
	pageStr = pageStr[i:]

	// Всего попаданий
	result.TotalHits, i = find(pageStr, "total_hits :", ',')
	pageStr = pageStr[i:]

	// Максимальное комбо
	result.MaximumCombo, i = find(pageStr, "maximum_combo :", ',')
	pageStr = pageStr[i:]

	// Реплеев просмотрено другими
	result.Replays, i = find(pageStr, "replays_watched_by_others :", ',')
	pageStr = pageStr[i:]

	// SS-ки
	result.SS, i = find(pageStr, "grade_counts :{ ss :", ',')
	pageStr = pageStr[i:]

	// SSH-ки
	result.SSH, i = find(pageStr, "ssh :", ',')
	pageStr = pageStr[i:]

	// S-ки
	result.S, i = find(pageStr, "s :", ',')
	pageStr = pageStr[i:]

	// SH-ки
	result.SH, i = find(pageStr, "sh :", ',')
	pageStr = pageStr[i:]

	// A-хи
	result.A, i = find(pageStr, "a :", '}')
	pageStr = pageStr[i:]

	// Рейтинг в стране
	result.CountryRank, i = find(pageStr, "country_rank :", ',')
	pageStr = pageStr[i:]

	// Уровень подписки
	result.SupportLvl, _ = find(pageStr, "support_level :", ',')

	return result
}

// Функция получения информации о пользователе
func getOnlineInfo(id string) OnlineInfo {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return OnlineInfo{}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// Проверка на страницу пользователя
	if !strings.Contains(string(body), "js-react--profile") {
		return OnlineInfo{}
	}

	// Структура, которую будет возвращать функция
	result := OnlineInfo{}

	// Статус в сети
	result.Status, _ = find(string(body), "is_online&quot;:", ',')

	return result
}

// Функция отправки информации о пользователе
func sendUserInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(getUserInfo(mux.Vars(request)["id"], mux.Vars(request)["mode"]))
}

// Функция отправки информации о статусе пользователя
func sendOnlineInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(getOnlineInfo(mux.Vars(request)["id"]))
}

func main() {

	// Вывод времени начала работы
	fmt.Println("API Start: " + string(time.Now().Format("2006-01-02 15:04:05")))

	// Роутер
	router := mux.NewRouter()

	// Маршруты

	router.HandleFunc("/user/{id}", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/{mode}", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/{mode}/", sendUserInfo).Methods("GET")

	router.HandleFunc("/online/{id}", sendOnlineInfo).Methods("GET")
	router.HandleFunc("/online/{id}/", sendOnlineInfo).Methods("GET")

	// Запуск API

	// Для Heroku
	//log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

	// Для локалхоста (127.0.0.1:8080/)
	log.Fatal(http.ListenAndServe(":8080", router))
}
