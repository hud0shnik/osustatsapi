package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MapStringResponse struct {
	Error              string                   `json:"error"`
	Artist             string                   `json:"artist"`
	ArtistUnicode      string                   `json:"artist_unicode"`
	Covers             CoverString              `json:"covers"`
	Creator            string                   `json:"creator"`
	FavoriteCount      string                   `json:"favorite_count"`
	Hype               string                   `json:"hype"`
	Id                 string                   `json:"id"`
	Nsfw               string                   `json:"nsfw"`
	Offset             string                   `json:"offset"`
	PlayCount          string                   `json:"play_count"`
	PreviewUrl         string                   `json:"preview_url"`
	Source             string                   `json:"source"`
	Spotlight          string                   `json:"spotlight"`
	Status             string                   `json:"status"`
	Title              string                   `json:"title"`
	TitleUnicode       string                   `json:"title_unicode"`
	TrackId            string                   `json:"track_id"`
	Video              string                   `json:"video"`
	DownloadDisabled   string                   `json:"download_disabled"`
	Bpm                string                   `json:"bpm"`
	CanBeHyped         string                   `json:"can_be_hyped"`
	DiscussionEnabled  string                   `json:"discussion_enabled"`
	DiscussionLocked   string                   `json:"discussion_locked"`
	IsScoreable        string                   `json:"is_scoreable"`
	LastUpdated        string                   `json:"last_updated"`
	LegacyThreadUrl    string                   `json:"legacy_thread_url"`
	NominationsSummary NominationsSummaryString `json:"nominations_summary"`
	Ranked             string                   `json:"ranked"`
	RankedDate         string                   `json:"ranked_date"`
	Storyboard         string                   `json:"storyboard"`
	SubmittedDate      string                   `json:"submitted_date"`
	Tags               []string                 `json:"tags"`
	Beatmaps           []MapsString             `json:"beatmaps"`
}

type MapsString struct {
	BeatmapSetId     string    `json:"beatmapset_id"`
	DifficultyRating string    `json:"difficulty_rating"`
	Id               string    `json:"id"`
	Mode             string    `json:"mode"`
	Status           string    `json:"status"`
	TotalLength      string    `json:"total_length"`
	UserId           string    `json:"user_id"`
	Version          string    `json:"version"`
	Accuracy         string    `json:"accuracy"`
	Ar               string    `json:"ar"`
	Bpm              string    `json:"bpm"`
	Convert          string    `json:"convert"`
	CountCircles     string    `json:"count_circles"`
	CountSliders     string    `json:"count_sliders"`
	CountSpinners    string    `json:"count_spinners"`
	Cs               string    `json:"cs"`
	DeletedAt        string    `json:"deleted_at"`
	Drain            string    `json:"drain"`
	HitLength        string    `json:"hit_length"`
	IsScoreable      string    `json:"is_scoreable"`
	LastUpdated      string    `json:"last_updated"`
	ModeInt          string    `json:"mode_int"`
	PassCount        string    `json:"pass_count"`
	PlayCount        string    `json:"play_count"`
	Ranked           string    `json:"ranked"`
	Url              string    `json:"url"`
	Checksum         string    `json:"checksum"`
	Failtimes        Failtimes `json:"failtimes"`
	MaxCombo         string    `json:"max_combo"`
}

type Failtimes struct {
	Fail []string `json:"fail"`
	Exit []string `json:"exit"`
}

// Роут "/mapstring" для vercel
func MapString(w http.ResponseWriter, r *http.Request) {

	// Формирование заголовка респонса по статускоду
	w.WriteHeader(http.StatusCreated)

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Получение параметра id из реквеста
	id := r.URL.Query().Get("id")
	beatmapset := r.URL.Query().Get("beatmapset")

	// Если параметра нет, отправка ошибки
	if id == "" || beatmapset == "" {
		http.NotFound(w, r)
		return
	}

	// Получение статистики, форматирование и отправка
	jsonResp, err := json.Marshal(GetMapInfoString(beatmapset, id))
	if err != nil {
		fmt.Print("Error: ", err)
	} else {
		w.Write(jsonResp)
	}
}

// Функция получения статистики карты
func GetMapInfoString(beatmapset, id string) MapStringResponse {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/beatmapsets/" + beatmapset + "#osu/" + id)
	if err != nil {
		return MapStringResponse{
			Error: "http get error",
		}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// HTML полученной страницы в формате string
	pageStr := string(body)
	pageStr = pageStr[index(pageStr, "<script id=\"json-beatmapset\" type=\"application/json", 80000)+61:]

	// Сохранение html"ки в файл sample.html (для тестов)

	/*if err := os.WriteFile("sample.html", []byte(pageStr), 0666); err != nil {
		log.Fatal(err)
	}*/

	return MapStringResponse{}
}
