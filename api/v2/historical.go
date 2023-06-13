package api2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type response struct {
	BeatmapPlaycounts beatmapPlaycounts `json:"beatmap_playcounts"`
	Recent            recent            `json:"recent"`
}

type beatmapPlaycounts struct {
	Items []playcount `json:"items"`
}

type playcount struct {
	BeatmapsetID     int     `json:"beatmapset_id"`
	DifficultyRating float64 `json:"difficulty_rating"`
	ID               int     `json:"id"`
	Mode             string  `json:"mode"`
	Status           string  `json:"status"`
	TotalLength      int     `json:"total_length"`
	UserID           int     `json:"user_id"`
	Version          string  `json:"version"`
}

type beatmapSmall struct {
	BeatmapsetID     int     `json:"beatmapset_id"`
	DifficultyRating float64 `json:"difficulty_rating"`
	ID               int     `json:"id"`
	Mode             string  `json:"mode"`
	Status           string  `json:"status"`
	TotalLength      int     `json:"total_length"`
	UserID           int     `json:"user_id"`
	Version          string  `json:"version"`
}

type recent struct {
	Items []recentScores `json:"items"`
}

type recentScores struct {
	Accuracy  float64 `json:"accuracy"`
	BeatmapID int     `json:"beatmap_id"`
	EndedAt   string  `json:"ended_at"`
	MaxCombo  int     `json:"max_combo"`
	Mods      []struct {
		Acronym string `json:"acronym"`
	} `json:"mods"`
	Passed        bool       `json:"passed"`
	Rank          string     `json:"rank"`
	RulesetID     int        `json:"ruleset_id"`
	Statistics    statistics `json:"statistics"`
	TotalScore    int        `json:"total_score"`
	UserID        int        `json:"user_id"`
	BestID        int64      `json:"best_id"`
	ID            int64      `json:"id"`
	LegacyPerfect bool       `json:"legacy_perfect"`
	Pp            float64    `json:"pp"`
	Replay        bool       `json:"replay"`
	Type          string     `json:"type"`
	Beatmap       Beatmap    `json:"beatmap"`
	Beatmapset    beatmapset `json:"beatmapset"`
}

type statistics struct {
	Great int `json:"great"`
	Miss  int `json:"miss"`
	Ok    int `json:"ok"`
}

type Beatmap struct {
	BeatmapsetID     int     `json:"beatmapset_id"`
	DifficultyRating float64 `json:"difficulty_rating"`
	ID               int     `json:"id"`
	Mode             string  `json:"mode"`
	Status           string  `json:"status"`
	TotalLength      int     `json:"total_length"`
	UserID           int     `json:"user_id"`
	Version          string  `json:"version"`
	Accuracy         float64 `json:"accuracy"`
	Ar               float64 `json:"ar"`
	Bpm              int     `json:"bpm"`
	Convert          bool    `json:"convert"`
	CountCircles     int     `json:"count_circles"`
	CountSliders     int     `json:"count_sliders"`
	CountSpinners    int     `json:"count_spinners"`
	Cs               int     `json:"cs"`
	DeletedAt        string  `json:"deleted_at"`
	Drain            int     `json:"drain"`
	HitLength        int     `json:"hit_length"`
	IsScoreable      bool    `json:"is_scoreable"`
	LastUpdated      string  `json:"last_updated"`
	ModeInt          int     `json:"mode_int"`
	Passcount        int     `json:"passcount"`
	Playcount        int     `json:"playcount"`
	Ranked           int     `json:"ranked"`
	URL              string  `json:"url"`
	Checksum         string  `json:"checksum"`
}

type beatmapset struct {
	Artist         string `json:"artist"`
	ArtistUnicode  string `json:"artist_unicode"`
	Covers         covers `json:"covers"`
	Creator        string `json:"creator"`
	FavouriteCount int    `json:"favourite_count"`
	ID             int    `json:"id"`
	Nsfw           bool   `json:"nsfw"`
	Offset         int    `json:"offset"`
	PlayCount      int    `json:"play_count"`
	PreviewURL     string `json:"preview_url"`
	Source         string `json:"source"`
	Spotlight      bool   `json:"spotlight"`
	Status         string `json:"status"`
	Title          string `json:"title"`
	TitleUnicode   string `json:"title_unicode"`
	TrackID        int    `json:"track_id"`
	UserID         int    `json:"user_id"`
	Video          bool   `json:"video"`
}

// Функция получения текстовой информации о пользователе
func getUserHistorical(id string) (response, error) {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id + "/extra-pages/historical?mode=osu")
	if err != nil {
		return response{}, fmt.Errorf("in http.Get: %w", err)
	}
	defer resp.Body.Close()

	// Проверка статускода
	if resp.StatusCode != 200 {
		return response{}, fmt.Errorf("response status: %s", resp.Status)
	}

	// Запись респонса
	body, _ := ioutil.ReadAll(resp.Body)

	var result response

	// Запись статистики в структуру
	err = json.Unmarshal(body, &result)
	if err != nil {
		return response{}, err
	}

	return result, nil

}

// Роут "/historical"
func Historical(w http.ResponseWriter, r *http.Request) {

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Получение параметра id из реквеста
	id := r.URL.Query().Get("id")

	// Проверка на наличие параметра
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json, _ := json.Marshal(apiError{Error: "please insert user id"})
		w.Write(json)
		return
	}

	// Получение статистики
	result, err := getUserHistorical(id)
	if err != nil {
		if err.Error() == "response status: 404 Not Found" {
			w.WriteHeader(http.StatusNotFound)
			json, _ := json.Marshal(apiError{Error: "not found"})
			w.Write(json)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(apiError{Error: "internal server error"})
		w.Write(json)
		logrus.Printf("getUserHistorical err: %s", err)
		return
	}

	// Перевод в json
	jsonResp, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(apiError{Error: "internal server error"})
		w.Write(json)
		logrus.Printf("json.Marshal err: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

}
