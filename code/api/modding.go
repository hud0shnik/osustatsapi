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
	Success    string                  `json:"success"`
	Error      string                  `json:"error"`
	Id         string                  `json:"id"`
	Type       string                  `json:"type"`
	Comment    ModdingCommentString    `json:"comment"`
	CreatedAt  string                  `json:"created_at"`
	UserId     string                  `json:"user_id"`
	Beatmapset ModdingBeatmapsetString `json:"beatmapset"`
}

// Структура комментария
type ModdingCommentString struct {
	BeatmapDiscussionId     string       `json:"beatmap_discussion_id"`
	BeatmapDiscussionPostId string       `json:"beatmap_discussion_post_id"`
	NewVote                 VoteString   `json:"new_vote"`
	Votes                   []VoteString `json:"votes"`
}

// Структура голоса
type VoteString struct {
	UserId string `json:"user_id"`
	Score  string `json:"score"`
}

// Структура карты
type ModdingBeatmapsetString struct {
	Artist        string       `json:"artist"`
	ArtistUnicode string       `json:"artist_unicode"`
	Covers        Covers       `json:"covers"`
	Creator       string       `json:"creator"`
	FavoriteCount string       `json:"favorite_count"`
	Hype          string       `json:"hype"`
	Id            string       `json:"id"`
	Nsfw          string       `json:"nsfw"`
	Offset        string       `json:"offset"`
	PlayCount     string       `json:"play_count"`
	PreviewUrl    string       `json:"preview_url"`
	Source        string       `json:"source"`
	Spotlight     string       `json:"spotlight"`
	Status        string       `json:"status"`
	Title         string       `json:"title"`
	TitleUnicode  string       `json:"title_unicode"`
	TrackId       string       `json:"track_id"`
	UserId        string       `json:"userId"`
	Video         string       `json:"video"`
	User          BmUserString `json:"user"`
}

// Структура дискуссии
type DiscussionString struct {
	Id             string       `json:"id"`
	BeatmapsetId   string       `json:"beatmapset_id"`
	BeatmapId      string       `json:"beatmap_id"`
	UserId         string       `json:"user_id"`
	DeletedById    string       `json:"deleted_by_id"`
	MessageType    string       `json:"message_type"`
	ParentId       string       `json:"parent_id"`
	Timestamp      string       `json:"timestamp"`
	Resolved       string       `json:"resolved"`
	CanBeResolved  string       `json:"can_be_resolved"`
	CanGrantKudosu string       `json:"can_grant_kudosu"`
	CreatedAt      string       `json:"created_at"`
	UpdatedAt      string       `json:"updated_at"`
	DeletedAt      string       `json:"deleted_at"`
	LastPostAt     string       `json:"last_post_at"`
	KudosuDenied   string       `json:"kudosu_denied"`
	StartingPost   StartingPost `json:"starting_post"`
}

// Структура первого поста
type StartingPost struct {
	BeatmapsetDiscussionId string `json:"beatmapset_discussion_id"`
	CreatedAt              string `json:"created_at"`
	DeletedAt              string `json:"deleted_at"`
	DeletedById            string `json:"deleted_by_id"`
	Id                     string `json:"id"`
	LastEditorId           string `json:"last_editor_id"`
	Message                string `json:"message"`
	System                 string `json:"system"`
	UpdatedAt              string `json:"updated_at"`
	UserId                 string `json:"user_id"`
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
