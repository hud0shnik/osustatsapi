package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Структура респонса
type MapResponse struct {
	Error              string              `json:"error"`
	Artist             string              `json:"artist"`
	ArtistUnicode      string              `json:"artist_string"`
	Covers             Covers              `json:"covers"`
	Creator            string              `json:"creator"`
	FavoriteCount      int                 `json:"favorite_count"`
	HypeCurrent        string              `json:"hype_current"`
	HypeRequired       string              `json:"hype_required"`
	Id                 int                 `json:"id"`
	Nsfw               bool                `json:"nsfw"`
	Offset             int                 `json:"offset"`
	PlayCount          int                 `json:"play_count"`
	PreviewUrl         string              `json:"preview_url"`
	Source             string              `json:"source"`
	Spotlight          bool                `json:"spotlight"`
	Status             string              `json:"status"`
	Title              string              `json:"title"`
	TitleUnicode       string              `json:"title_unicode"`
	TrackId            int                 `json:"track_id"`
	UserId             int                 `json:"user_id"`
	Video              bool                `json:"video"`
	DownloadDisabled   bool                `json:"download_disabled"`
	Bpm                float64             `json:"bpm"`
	CanBeHyped         bool                `json:"can_be_hyped"`
	DiscussionEnabled  bool                `json:"discussion_enabled"`
	DiscussionLocked   bool                `json:"discussion_locked"`
	IsScoreable        bool                `json:"is_scoreable"`
	LastUpdated        string              `json:"last_updated"`
	LegacyThreadUrl    string              `json:"legacy_thread_url"`
	NominationsSummary NominationsSummary  `json:"nominations_summary"`
	Ranked             int                 `json:"ranked"`
	RankedDate         string              `json:"ranked_date"`
	Storyboard         bool                `json:"storyboard"`
	SubmittedDate      string              `json:"submitted_date"`
	Tags               []string            `json:"tags"`
	Beatmaps           []Maps              `json:"beatmaps"`
	Converts           []Maps              `json:"converts"`
	CurrentNominations []CurrentNomination `json:"current_nominations"`
	Description        string              `json:"description"`
	GenreId            int                 `json:"genre_id"`
	GenreName          string              `json:"genre_name"`
	LanguageId         int                 `json:"language_id"`
	LanguageName       string              `json:"language_name"`
	Ratings            []int               `json:"ratings"`
}

// Структура карты
type Maps struct {
	BeatmapSetId     int       `json:"beatmapset_id"`
	DifficultyRating float32   `json:"difficulty_rating"`
	Id               int       `json:"id"`
	Mode             string    `json:"mode"`
	Status           string    `json:"status"`
	TotalLength      int       `json:"total_length"`
	UserId           int       `json:"user_id"`
	Version          string    `json:"version"`
	Accuracy         float32   `json:"accuracy"`
	Ar               float32   `json:"ar"`
	Bpm              float32   `json:"bpm"`
	Convert          bool      `json:"convert"`
	CountCircles     int       `json:"count_circles"`
	CountSliders     int       `json:"count_sliders"`
	CountSpinners    int       `json:"count_spinners"`
	Cs               float32   `json:"cs"`
	DeletedAt        string    `json:"deleted_at"`
	Drain            float32   `json:"drain"`
	HitLength        int       `json:"hit_length"`
	IsScoreable      bool      `json:"is_scoreable"`
	LastUpdated      string    `json:"last_updated"`
	ModeInt          int       `json:"mode_int"`
	PassCount        int       `json:"pass_count"`
	PlayCount        int       `json:"play_count"`
	Ranked           int       `json:"ranked"`
	Url              string    `json:"url"`
	Checksum         string    `json:"checksum"`
	Failtimes        Failtimes `json:"failtimes"`
	MaxCombo         int       `json:"max_combo"`
}

// Структура проигрышей
type Failtimes struct {
	Fail []int `json:"fail"`
	Exit []int `json:"exit"`
}

// Структура номинации
type CurrentNomination struct {
	BeatmapsetId int    `json:"beatmapset_id"`
	Rulesets     string `json:"rulesets"`
	Reset        bool   `json:"reset"`
	UserId       int    `json:"user_id"`
}

// Структура пользователя
type BmUser struct {
	AvatarUrl     string `json:"avatar_url"`
	CountryCode   string `json:"country_code"`
	DefaultGroup  string `json:"default_group"`
	Id            int    `json:"id "`
	IsActive      bool   `json:"is_active"`
	IsBot         bool   `json:"is_bot"`
	IsDeleted     bool   `json:"is_deleted"`
	IsOnline      bool   `json:"is_online"`
	IsSupporter   bool   `json:"is_supporter"`
	LastVisit     string `json:"last_visit"`
	PmFriendsOnly bool   `json:"pm_friends_only"`
	ProfileColor  string `json:"profile_color"`
	Username      string `json:"username"`
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
