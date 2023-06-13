package api2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// ---------------------- Классические структуры ------------------------

type response struct {
	Recent recent `json:"recent"`
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
	PP            float64    `json:"PP"`
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
	Bpm              float64 `json:"bpm"`
	Convert          bool    `json:"convert"`
	CountCircles     int     `json:"count_circles"`
	CountSliders     int     `json:"count_sliders"`
	CountSpinners    int     `json:"count_spinners"`
	Cs               float64 `json:"cs"`
	DeletedAt        string  `json:"deleted_at"`
	Drain            float64 `json:"drain"`
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

// ---------------------- Текстовые структуры ------------------------

type responseString struct {
	Recent recentString `json:"recent"`
}

type recentString struct {
	Items []recentScoresString `json:"items"`
}

type recentScoresString struct {
	Accuracy  string `json:"accuracy"`
	BeatmapID string `json:"beatmap_id"`
	EndedAt   string `json:"ended_at"`
	MaxCombo  string `json:"max_combo"`
	Mods      []struct {
		Acronym string `json:"acronym"`
	} `json:"mods"`
	Passed        string           `json:"passed"`
	Rank          string           `json:"rank"`
	RulesetID     string           `json:"ruleset_id"`
	Statistics    statisticsString `json:"statistics"`
	TotalScore    string           `json:"total_score"`
	UserID        string           `json:"user_id"`
	BestID        string           `json:"best_id"`
	ID            string           `json:"id"`
	LegacyPerfect string           `json:"legacy_perfect"`
	PP            string           `json:"PP"`
	Replay        string           `json:"replay"`
	Type          string           `json:"type"`
	Beatmap       BeatmapString    `json:"beatmap"`
	Beatmapset    beatmapsetString `json:"beatmapset"`
}

type statisticsString struct {
	Great string `json:"great"`
	Miss  string `json:"miss"`
	Ok    string `json:"ok"`
}

type BeatmapString struct {
	BeatmapsetID     string `json:"beatmapset_id"`
	DifficultyRating string `json:"difficulty_rating"`
	ID               string `json:"id"`
	Mode             string `json:"mode"`
	Status           string `json:"status"`
	TotalLength      string `json:"total_length"`
	UserID           string `json:"user_id"`
	Version          string `json:"version"`
	Accuracy         string `json:"accuracy"`
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
	Passcount        string `json:"passcount"`
	Playcount        string `json:"playcount"`
	Ranked           string `json:"ranked"`
	URL              string `json:"url"`
	Checksum         string `json:"checksum"`
}

type beatmapsetString struct {
	Artist         string `json:"artist"`
	ArtistUnicode  string `json:"artist_unicode"`
	Covers         covers `json:"covers"`
	Creator        string `json:"creator"`
	FavouriteCount string `json:"favourite_count"`
	ID             string `json:"id"`
	Nsfw           string `json:"nsfw"`
	Offset         string `json:"offset"`
	PlayCount      string `json:"play_count"`
	PreviewURL     string `json:"preview_url"`
	Source         string `json:"source"`
	Spotlight      string `json:"spotlight"`
	Status         string `json:"status"`
	Title          string `json:"title"`
	TitleUnicode   string `json:"title_unicode"`
	TrackID        string `json:"track_id"`
	UserID         string `json:"user_id"`
	Video          string `json:"video"`
}

// Функция получения недавних карт
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

// Функция конвертации недавних карт
func getUserHistoricalString(id string) (responseString, error) {

	// Получение классической версии
	classic, err := getUserHistorical(id)
	if err != nil {
		return responseString{}, err
	}

	var result responseString

	// Конвертация
	for _, c := range classic.Recent.Items {
		recent := recentScoresString{
			Accuracy:  fmt.Sprint(c.Accuracy),
			BeatmapID: fmt.Sprint(c.BeatmapID),
			EndedAt:   c.EndedAt,
			MaxCombo:  fmt.Sprint(c.MaxCombo),
			Mods:      c.Mods,
			Passed:    fmt.Sprint(c.Passed),
			Rank:      c.Rank,
			RulesetID: fmt.Sprint(c.RulesetID),
			Statistics: statisticsString{
				Great: fmt.Sprint(c.Statistics.Great),
				Miss:  fmt.Sprint(c.Statistics.Miss),
				Ok:    fmt.Sprint(c.Statistics.Ok),
			},
			TotalScore:    fmt.Sprint(c.TotalScore),
			UserID:        fmt.Sprint(c.UserID),
			BestID:        fmt.Sprint(c.BestID),
			ID:            fmt.Sprint(c.ID),
			LegacyPerfect: fmt.Sprint(c.LegacyPerfect),
			PP:            fmt.Sprint(c.PP),
			Replay:        fmt.Sprint(c.Replay),
			Type:          c.Type,
			Beatmap: BeatmapString{
				BeatmapsetID:     fmt.Sprint(c.BeatmapID),
				DifficultyRating: fmt.Sprint(c.Beatmap.DifficultyRating),
				ID:               fmt.Sprint(c.Beatmap.ID),
				Mode:             c.Beatmap.Mode,
				Status:           c.Beatmap.Status,
				TotalLength:      fmt.Sprint(c.Beatmap.TotalLength),
				UserID:           fmt.Sprint(c.Beatmap.UserID),
				Version:          c.Beatmap.Version,
				Accuracy:         fmt.Sprint(c.Beatmap.Accuracy),
				Ar:               fmt.Sprint(c.Beatmap.Ar),
				Bpm:              fmt.Sprint(c.Beatmap.Bpm),
				Convert:          fmt.Sprint(c.Beatmap.Convert),
				CountCircles:     fmt.Sprint(c.Beatmap.CountCircles),
				CountSliders:     fmt.Sprint(c.Beatmap.CountSliders),
				CountSpinners:    fmt.Sprint(c.Beatmap.CountSpinners),
				Cs:               fmt.Sprint(c.Beatmap.Cs),
				DeletedAt:        c.Beatmap.DeletedAt,
				Drain:            fmt.Sprint(c.Beatmap.Drain),
				HitLength:        fmt.Sprint(c.Beatmap.HitLength),
				IsScoreable:      fmt.Sprint(c.Beatmap.IsScoreable),
				LastUpdated:      c.Beatmap.LastUpdated,
				ModeInt:          fmt.Sprint(c.Beatmap.ModeInt),
				Passcount:        fmt.Sprint(c.Beatmap.Passcount),
				Playcount:        fmt.Sprint(c.Beatmap.Playcount),
				Ranked:           fmt.Sprint(c.Beatmap.Ranked),
				URL:              c.Beatmap.URL,
				Checksum:         c.Beatmap.Checksum,
			},
			Beatmapset: beatmapsetString{
				Artist:         c.Beatmapset.Artist,
				ArtistUnicode:  c.Beatmapset.ArtistUnicode,
				Covers:         c.Beatmapset.Covers,
				Creator:        c.Beatmapset.Creator,
				FavouriteCount: fmt.Sprint(c.Beatmapset.FavouriteCount),
				ID:             fmt.Sprint(c.Beatmapset.ID),
				Nsfw:           fmt.Sprint(c.Beatmapset.Nsfw),
				Offset:         fmt.Sprint(c.Beatmapset.Offset),
				PlayCount:      fmt.Sprint(c.Beatmapset.PlayCount),
				PreviewURL:     c.Beatmapset.PreviewURL,
				Source:         c.Beatmapset.Source,
				Spotlight:      fmt.Sprint(c.Beatmapset.Spotlight),
				Status:         c.Beatmapset.Status,
				Title:          c.Beatmapset.Title,
				TitleUnicode:   c.Beatmapset.TitleUnicode,
				TrackID:        fmt.Sprint(c.Beatmapset.TrackID),
				UserID:         fmt.Sprint(c.Beatmapset.UserID),
				Video:          fmt.Sprint(c.Beatmapset.Video),
			},
		}
		result.Recent.Items = append(result.Recent.Items, recent)

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

	// Проверка на тип
	if r.URL.Query().Get("type") == "string" {

		// Получение статистики
		result, err := getUserHistoricalString(id)
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

	} else {

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

}
