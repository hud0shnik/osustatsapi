package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Структура респонса
type MapResponse struct {
	Error              string             `json:"error"`
	Artist             string             `json:"artist"`
	ArtistUnicode      string             `json:"artist_string"`
	Covers             Covers             `json:"covers"`
	Creator            string             `json:"creator"`
	FavoriteCount      int                `json:"favorite_count"`
	HypeCurrent        string             `json:"hype_current"`
	HypeRequired       string             `json:"hype_required"`
	Id                 int                `json:"id"`
	Nsfw               bool               `json:"nsfw"`
	Offset             int                `json:"offset"`
	PlayCount          int                `json:"play_count"`
	PreviewUrl         string             `json:"preview_url"`
	Source             string             `json:"source"`
	Spotlight          bool               `json:"spotlight"`
	Status             string             `json:"status"`
	Title              string             `json:"title"`
	TitleUnicode       string             `json:"title_unicode"`
	TrackId            int                `json:"track_id"`
	UserId             int                `json:"user_id"`
	Video              bool               `json:"video"`
	DownloadDisabled   bool               `json:"download_disabled"`
	Bpm                float64            `json:"bpm"`
	CanBeHyped         bool               `json:"can_be_hyped"`
	DiscussionEnabled  bool               `json:"discussion_enabled"`
	DiscussionLocked   bool               `json:"discussion_locked"`
	IsScoreable        bool               `json:"is_scoreable"`
	LastUpdated        string             `json:"last_updated"`
	LegacyThreadUrl    string             `json:"legacy_thread_url"`
	NominationsSummary NominationsSummary `json:"nominations_summary"`
	Ranked             int                `json:"ranked"`
	RankedDate         string             `json:"ranked_date"`
	Storyboard         bool               `json:"storyboard"`
	SubmittedDate      string             `json:"submitted_date"`
	Tags               []string           `json:"tags"`
}

// Структура карты
type Maps struct {
	BeatmapSetId     int     `json:"beatmapset_id"`
	DifficultyRating float32 `json:"difficulty_rating"`
	Id               int     `json:"id"`
	Mode             string  `json:"mode"`
	Status           string  `json:"status"`
	TotalLength      int     `json:"total_length"`
	UserId           int     `json:"user_id"`
	Version          string  `json:"version"`
	Accuracy         float32 `json:"accuracy"`
	Ar               float32 `json:"ar"`
	Bpm              float32 `json:"bpm"`
	Convert          bool    `json:"convert"`
	CountCircles     int     `json:"count_circles"`
	CountSliders     int     `json:"count_sliders"`
	CountSpinners    int     `json:"count_spinners"`
}

// Роут "/map" для vercel
func Map(w http.ResponseWriter, r *http.Request) {

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
	jsonResp, err := json.Marshal(GetMapInfo(beatmapset, id))
	if err != nil {
		fmt.Print("Error: ", err)
	} else {
		w.Write(jsonResp)
	}
}

// Функция получения статистики карты
func GetMapInfo(beatmapset, id string) MapResponse {
	return MapResponse{
		Error: "",
	}
}
