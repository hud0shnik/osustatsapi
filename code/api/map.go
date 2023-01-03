package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Структура респонса
type MapResponse struct {
	Error              string              `json:"error"`
	Artist             string              `json:"artist"`
	ArtistUnicode      string              `json:"artist_string"`
	Covers             Covers              `json:"covers"`
	Creator            string              `json:"creator"`
	FavoriteCount      int                 `json:"favorite_count"`
	HypeCurrent        int                 `json:"hype_current"`
	HypeRequired       int                 `json:"hype_required"`
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
	RecentFavourites   []BmUser            `json:"recent_favourites"`
	RelatedUsers       []BmUser            `json:"related_users"`
	User               BmUser              `json:"user"`
	Comments           []Comment           `json:"comments"`
	PinnedComments     []Comment           `json:"pinned_comments"`
	UserFollow         string              `json:"user_follow"`
}

// Структура карты
type Maps struct {
	BeatmapSetId     int       `json:"beatmapset_id"`
	DifficultyRating float64   `json:"difficulty_rating"`
	Id               int       `json:"id"`
	Mode             string    `json:"mode"`
	Status           string    `json:"status"`
	TotalLength      int       `json:"total_length"`
	UserId           int       `json:"user_id"`
	Version          string    `json:"version"`
	Accuracy         float64   `json:"accuracy"`
	Ar               float64   `json:"ar"`
	Bpm              float64   `json:"bpm"`
	Convert          bool      `json:"convert"`
	CountCircles     int       `json:"count_circles"`
	CountSliders     int       `json:"count_sliders"`
	CountSpinners    int       `json:"count_spinners"`
	Cs               float64   `json:"cs"`
	DeletedAt        string    `json:"deleted_at"`
	Drain            float64   `json:"drain"`
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

// Структура комментария
type Comment struct {
	Id              int    `json:"id"`
	ParentId        int    `json:"parent_id"`
	UserId          int    `json:"user_id"`
	Pinned          bool   `json:"pinned"`
	RepliesCount    int    `json:"replies_count"`
	VotesCount      int    `json:"votes_count"`
	CommentableType string `json:"commentable_type"`
	CommentableId   int    `json:"commentable_id"`
	LegacyName      string `json:"legacy_name"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	EditedAt        string `json:"edited_at"`
	EditedById      string `json:"edited_by_id"`
	Message         string `json:"message"`
	MessageHtml     string `json:"message_html"`
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

// Функция перевода карты
func FormatBeatmap(mpss []MapsString) []Maps {

	var result []Maps

	// Обработка текстовых рекордов
	for _, mps := range mpss {

		mp := Maps{
			BeatmapSetId:     ToInt(mps.BeatmapSetId),
			DifficultyRating: ToFloat64(mps.DifficultyRating),
			Id:               ToInt(mps.Id),
			Mode:             mps.Mode,
			Status:           mps.Status,
			TotalLength:      ToInt(mps.TotalLength),
			UserId:           ToInt(mps.UserId),
			Version:          mps.Version,
			Accuracy:         ToFloat64(mps.Version),
			Ar:               ToFloat64(mps.Ar),
			Bpm:              ToFloat64(mps.Bpm),
			Convert:          ToBool(mps.Convert),
			CountCircles:     ToInt(mps.CountCircles),
			CountSliders:     ToInt(mps.CountSliders),
			CountSpinners:    ToInt(mps.CountSpinners),
			Cs:               ToFloat64(mps.Cs),
			DeletedAt:        mps.DeletedAt,
			Drain:            ToFloat64(mps.Drain),
			HitLength:        ToInt(mps.HitLength),
			IsScoreable:      ToBool(mps.IsScoreable),
			LastUpdated:      mps.LastUpdated,
			ModeInt:          ToInt(mps.ModeInt),
			PassCount:        ToInt(mps.PassCount),
			PlayCount:        ToInt(mps.PlayCount),
			Ranked:           ToInt(mps.Ranked),
			Url:              mps.Url,
			Checksum:         mps.Checksum,
			MaxCombo:         ToInt(mps.MaxCombo),
		}

		for _, f := range strings.Split(mps.Failtimes.Fail, ",") {
			mp.Failtimes.Fail = append(mp.Failtimes.Fail, ToInt(f))
		}

		for _, e := range strings.Split(mps.Failtimes.Exit, ",") {
			mp.Failtimes.Exit = append(mp.Failtimes.Exit, ToInt(e))
		}

		// Форматирование и добавление рекорда
		result = append(result, mp)
	}

	return result
}

// Функция получения статистики карты
func GetMapInfo(beatmapset, id string) MapResponse {

	// Получение текстовой версии статистики
	resultStr := GetMapInfoString(beatmapset, id)

	// Проверка на ошибки при парсинге
	if resultStr.Error != "" {
		return MapResponse{
			Error: "",
		}
	}

	// Перевод в классическую версию
	result := MapResponse{
		Error:             resultStr.Error,
		Artist:            resultStr.Artist,
		ArtistUnicode:     resultStr.ArtistUnicode,
		Covers:            resultStr.Covers,
		Creator:           resultStr.Creator,
		FavoriteCount:     ToInt(resultStr.FavoriteCount),
		HypeCurrent:       ToInt(resultStr.HypeCurrent),
		HypeRequired:      ToInt(resultStr.HypeRequired),
		Id:                ToInt(resultStr.Id),
		Nsfw:              ToBool(resultStr.Nsfw),
		Offset:            ToInt(resultStr.Offset),
		PlayCount:         ToInt(resultStr.PlayCount),
		PreviewUrl:        resultStr.PreviewUrl,
		Source:            resultStr.Source,
		Spotlight:         ToBool(resultStr.Spotlight),
		Status:            resultStr.Status,
		Title:             resultStr.Title,
		TitleUnicode:      resultStr.TitleUnicode,
		TrackId:           ToInt(resultStr.TrackId),
		UserId:            ToInt(resultStr.UserId),
		Video:             ToBool(resultStr.Video),
		DownloadDisabled:  ToBool(resultStr.DownloadDisabled),
		Bpm:               ToFloat64(resultStr.Bpm),
		CanBeHyped:        ToBool(resultStr.CanBeHyped),
		DiscussionEnabled: ToBool(resultStr.DiscussionEnabled),
		DiscussionLocked:  ToBool(resultStr.DiscussionLocked),
		IsScoreable:       ToBool(resultStr.IsScoreable),
		LastUpdated:       resultStr.LastUpdated,
		LegacyThreadUrl:   resultStr.LegacyThreadUrl,
		NominationsSummary: NominationsSummary{
			Current:  ToInt(resultStr.NominationsSummary.Current),
			Required: ToInt(resultStr.NominationsSummary.Required),
		},
		Ranked: ToInt(resultStr.Ranked),
	}

	return result
}
