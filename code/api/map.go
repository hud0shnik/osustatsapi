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
type MapStringResponse struct {
	Error              string                    `json:"error"`
	Artist             string                    `json:"artist"`
	ArtistUnicode      string                    `json:"artist_unicode"`
	Covers             Covers                    `json:"covers"`
	Creator            string                    `json:"creator"`
	FavoriteCount      string                    `json:"favorite_count"`
	HypeCurrent        string                    `json:"hype_current"`
	HypeRequired       string                    `json:"hype_required"`
	Id                 string                    `json:"id"`
	Nsfw               string                    `json:"nsfw"`
	Offset             string                    `json:"offset"`
	PlayCount          string                    `json:"play_count"`
	PreviewUrl         string                    `json:"preview_url"`
	Source             string                    `json:"source"`
	Spotlight          string                    `json:"spotlight"`
	Status             string                    `json:"status"`
	Title              string                    `json:"title"`
	TitleUnicode       string                    `json:"title_unicode"`
	TrackId            string                    `json:"track_id"`
	UserId             string                    `json:"user_id"`
	Video              string                    `json:"video"`
	DownloadDisabled   string                    `json:"download_disabled"`
	Bpm                string                    `json:"bpm"`
	CanBeHyped         string                    `json:"can_be_hyped"`
	DiscussionEnabled  string                    `json:"discussion_enabled"`
	DiscussionLocked   string                    `json:"discussion_locked"`
	IsScoreable        string                    `json:"is_scoreable"`
	LastUpdated        string                    `json:"last_updated"`
	LegacyThreadUrl    string                    `json:"legacy_thread_url"`
	NominationsSummary NominationsSummaryString  `json:"nominations_summary"`
	Ranked             string                    `json:"ranked"`
	RankedDate         string                    `json:"ranked_date"`
	Storyboard         string                    `json:"storyboard"`
	SubmittedDate      string                    `json:"submitted_date"`
	Tags               []string                  `json:"tags"`
	Beatmaps           []MapsString              `json:"beatmaps"`
	Converts           []MapsString              `json:"converts"`
	CurrentNominations []CurrentNominationString `json:"current_nominations"`
	Description        string                    `json:"description"`
	GenreId            string                    `json:"genre_id"`
	GenreName          string                    `json:"genre_name"`
	LanguageId         string                    `json:"language_id"`
	LanguageName       string                    `json:"language_name"`
	Ratings            string                    `json:"ratings"`
	RecentFavourites   []BmUserString            `json:"recent_favourites"`
	RelatedUsers       []BmUserString            `json:"related_users"`
	User               BmUserString              `json:"user"`
	Comments           []CommentString           `json:"comments"`
	PinnedComments     []CommentString           `json:"pinned_comments"`
	UserFollow         string                    `json:"user_follow"`
}

// Структура карты
type MapsString struct {
	BeatmapSetId     string          `json:"beatmapset_id"`
	DifficultyRating string          `json:"difficulty_rating"`
	Id               string          `json:"id"`
	Mode             string          `json:"mode"`
	Status           string          `json:"status"`
	TotalLength      string          `json:"total_length"`
	UserId           string          `json:"user_id"`
	Version          string          `json:"version"`
	Accuracy         string          `json:"accuracy"`
	Ar               string          `json:"ar"`
	Bpm              string          `json:"bpm"`
	Convert          string          `json:"convert"`
	CountCircles     string          `json:"count_circles"`
	CountSliders     string          `json:"count_sliders"`
	CountSpinners    string          `json:"count_spinners"`
	Cs               string          `json:"cs"`
	DeletedAt        string          `json:"deleted_at"`
	Drain            string          `json:"drain"`
	HitLength        string          `json:"hit_length"`
	IsScoreable      string          `json:"is_scoreable"`
	LastUpdated      string          `json:"last_updated"`
	ModeInt          string          `json:"mode_int"`
	PassCount        string          `json:"pass_count"`
	PlayCount        string          `json:"play_count"`
	Ranked           string          `json:"ranked"`
	Url              string          `json:"url"`
	Checksum         string          `json:"checksum"`
	Failtimes        FailtimesString `json:"failtimes"`
	MaxCombo         string          `json:"max_combo"`
}

// Структура проигрышей
type FailtimesString struct {
	Fail string `json:"fail"`
	Exit string `json:"exit"`
}

// Структура пользователя
type BmUserString struct {
	AvatarUrl     string `json:"avatar_url"`
	CountryCode   string `json:"country_code"`
	DefaultGroup  string `json:"default_group"`
	Id            string `json:"id "`
	IsActive      string `json:"is_active"`
	IsBot         string `json:"is_bot"`
	IsDeleted     string `json:"is_deleted"`
	IsOnline      string `json:"is_online"`
	IsSupporter   string `json:"is_supporter"`
	LastVisit     string `json:"last_visit"`
	PmFriendsOnly string `json:"pm_friends_only"`
	ProfileColor  string `json:"profile_color"`
	Username      string `json:"username"`
}

// Структура комментария
type CommentString struct {
	Id              string `json:"id"`
	ParentId        string `json:"parent_id"`
	UserId          string `json:"user_id"`
	Pinned          string `json:"pinned"`
	RepliesCount    string `json:"replies_count"`
	VotesCount      string `json:"votes_count"`
	CommentableType string `json:"commentable_type"`
	CommentableId   string `json:"commentable_id"`
	LegacyName      string `json:"legacy_name"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	EditedAt        string `json:"edited_at"`
	EditedById      string `json:"edited_by_id"`
	Message         string `json:"message"`
	MessageHtml     string `json:"message_html"`
}

// Структура номинации
type CurrentNominationString struct {
	BeatmapsetId string `json:"beatmapset_id"`
	Rulesets     string `json:"rulesets"`
	Reset        string `json:"reset"`
	UserId       string `json:"user_id"`
}

// ---------------------- Классические структуры ------------------------

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
	UserFollow         bool                `json:"user_follow"`
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

// ---------------------- Функции парсинга ----------------------

// Функция парсинга текущих номинаций
func parseCurrentNominations(pageStr, subStr, stopChar string, left int) ([]CurrentNominationString, int) {

	// Индекс конца номинаций
	var end int

	// Получение рабочей части и индекса её конца
	pageStr, end = findWithIndex(pageStr, subStr, stopChar, left)

	// Проверка на наличие номинаций
	if len(pageStr) == 0 {
		return []CurrentNominationString{}, end
	}

	// Результат и индекс обработанной части
	var result []CurrentNominationString
	left = 0

	// Пока есть необработанные карты
	for index(pageStr, "beatmapset_id", left) != -1 {

		// Структура номинации
		var cn CurrentNominationString

		// Запись данных
		cn.BeatmapsetId, left = findWithIndex(pageStr, "beatmapset_id\":", ",", left)
		cn.Rulesets, left = findWithIndex(pageStr, "rulesets\":[", "]", left)
		cn.Reset, left = findWithIndex(pageStr, "reset\":", ",", left)
		cn.UserId, left = findWithIndex(pageStr, "user_id\":", "}", left)

		// Добавление карты к результату
		result = append(result, cn)
	}

	return result, end
}

// Функция парсинга карт
func parseMapsString(pageStr, subStr, stopChar string, left int) ([]MapsString, int) {

	// Индекс конца карт
	var end int

	// Получение рабочей части и индекса её конца
	pageStr, end = findWithIndex(pageStr, subStr, stopChar, left)

	// Проверка на наличие карт
	if len(pageStr) == 0 {
		return []MapsString{}, end
	}

	// Результат и индекс обработанной части
	var result []MapsString
	left = 0

	// Пока есть необработанные карты
	for index(pageStr, "\"beatmapset_id\"", left) != -1 {

		// Структура карты
		var bm MapsString

		bm, left = parseMapString(pageStr, left)

		// Добавление карты к результату
		result = append(result, bm)
	}

	return result, end
}

// Функция парсинга карты
func parseMapString(pageStr string, left int) (MapsString, int) {

	// Структура карты
	var bm MapsString

	// Запись данных
	bm.BeatmapSetId, left = findWithIndex(pageStr, "\"beatmapset_id\":", ",", left)
	bm.DifficultyRating, left = findWithIndex(pageStr, "\"difficulty_rating\":", ",", left)
	bm.Id, left = findWithIndex(pageStr, "\"id\":", ",", left)
	bm.Mode, left = findStringWithIndex(pageStr, "\"mode\":", ",", left)
	bm.Status, left = findStringWithIndex(pageStr, "\"status\":", ",", left)
	bm.TotalLength, left = findWithIndex(pageStr, "total_length\":", ",", left)
	bm.UserId, left = findWithIndex(pageStr, "\"user_id\":", ",", left)
	bm.Version, left = findWithIndex(pageStr, "\"version\":\"", "\"", left)
	bm.Accuracy, left = findWithIndex(pageStr, "\"accuracy\":", ",", left)
	bm.Ar, left = findWithIndex(pageStr, "\"ar\":", ",", left)
	bm.Bpm, left = findWithIndex(pageStr, "\"bpm\":", ",", left)
	bm.Convert, left = findWithIndex(pageStr, "\"convert\":", ",", left)
	bm.CountCircles, left = findWithIndex(pageStr, "\"count_circles\":", ",", left)
	bm.CountSliders, left = findWithIndex(pageStr, "\"count_sliders\":", ",", left)
	bm.CountSpinners, left = findWithIndex(pageStr, "\"count_spinners\":", ",", left)
	bm.Cs, left = findWithIndex(pageStr, "\"cs\":", ",", left)
	bm.DeletedAt, left = findStringWithIndex(pageStr, "\"deleted_at\":", ",", left)
	bm.Drain, left = findWithIndex(pageStr, "\"drain\":", ",", left)
	bm.HitLength, left = findWithIndex(pageStr, "\"hit_length\":", ",", left)
	bm.IsScoreable, left = findWithIndex(pageStr, "\"is_scoreable\":", ",", left)
	bm.LastUpdated, left = findStringWithIndex(pageStr, "\"last_updated\":", ",", left)
	bm.ModeInt, left = findWithIndex(pageStr, "\"mode_int\":", ",", left)
	bm.PassCount, left = findWithIndex(pageStr, "\"passcount\":", ",", left)
	bm.PlayCount, left = findWithIndex(pageStr, "\"playcount\":", ",", left)
	bm.Ranked, left = findWithIndex(pageStr, "\"ranked\":", ",", left)
	bm.Url, left = findStringWithIndex(pageStr, "\"url\":", ",", left)
	bm.Url = strings.ReplaceAll(bm.Url, "\\", "")
	bm.Checksum, left = findStringWithIndex(pageStr, "\"checksum\":", ",", left)

	bm.Failtimes.Fail, left = findWithIndex(pageStr, "\"fail\":[", "]", left)
	bm.Failtimes.Exit, left = findWithIndex(pageStr, "\"exit\":[", "]", left)

	bm.MaxCombo, left = findWithIndex(pageStr, "\"max_combo\":", "}", left)

	return bm, left
}

// Функция парсинга пользователей
func parseBmUsers(pageStr, subStr, stopChar string, left int) ([]BmUserString, int) {

	// Индекс конца пользователей
	var end int

	// Получение рабочей части и индекса её конца
	pageStr, end = findWithIndex(pageStr, subStr, stopChar, left)

	// Проверка на наличие пользователей
	if len(pageStr) == 0 {
		return []BmUserString{}, end
	}

	// Результат и индекс обработанной части
	var result []BmUserString
	left = 0

	// Пока есть необработанные пользователи
	for index(pageStr, "avatar_url", left) != -1 {

		// Структура пользователя
		var user BmUserString

		user, left = parseBmUserString(pageStr, left)

		// Добавление пользователя к результату
		result = append(result, user)
	}

	return result, end
}

// функция парсинга пользователей
func parseBmUserString(pageStr string, left int) (BmUserString, int) {

	// Структура пользователя
	var user BmUserString

	// Запись данных
	user.AvatarUrl, left = findStringWithIndex(pageStr, "avatar_url\":", ",", left)
	user.CountryCode, left = findStringWithIndex(pageStr, "country_code\":", ",", left)
	user.DefaultGroup, left = findStringWithIndex(pageStr, "default_group\":", ",", left)
	user.Id, left = findWithIndex(pageStr, "id\":", ",", left)
	user.IsActive, left = findWithIndex(pageStr, "is_active\":", ",", left)
	user.IsBot, left = findWithIndex(pageStr, "is_bot\":", ",", left)
	user.IsDeleted, left = findWithIndex(pageStr, "is_deleted\":", ",", left)
	user.IsOnline, left = findWithIndex(pageStr, "is_online\":", ",", left)
	user.IsSupporter, left = findWithIndex(pageStr, "is_supporter\":", ",", left)
	user.LastVisit, left = findStringWithIndex(pageStr, "last_visit\":", ",", left)
	user.PmFriendsOnly, left = findWithIndex(pageStr, "pm_friends_only\":", ",", left)
	user.ProfileColor, left = findStringWithIndex(pageStr, "profile_colour\":", ",", left)
	user.Username, left = findStringWithIndex(pageStr, "username\":", "}", left)

	return user, left
}

// Функция парсинга комментов
func parseCommentsString(pageStr, subStr, stopChar string, left int) ([]CommentString, int) {

	// Индекс конца комментариев
	var end int

	// Получение рабочей части и индекса её конца
	pageStr, end = findWithIndex(pageStr, subStr, stopChar, left)

	// Проверка на наличие пользователей
	if len(pageStr) == 0 {
		return []CommentString{}, end
	}

	// Результат и индекс обработанной части
	result := []CommentString{}
	left = 0

	// Пока есть необработанные пользователи
	for i := 0; index(pageStr, "id\":", left) != -1; i++ {

		// Структура комментария
		var cm CommentString
		cm.Id, left = findWithIndex(pageStr, "id\":", ",", left)
		cm.ParentId, left = findWithIndex(pageStr, "parent_id\":", ",", left)
		cm.UserId, left = findWithIndex(pageStr, "user_id\":", ",", left)
		cm.Pinned, left = findWithIndex(pageStr, "pinned\":", ",", left)
		cm.RepliesCount, left = findWithIndex(pageStr, "replies_count\":", ",", left)
		cm.VotesCount, left = findWithIndex(pageStr, "votes_count\":", ",", left)
		cm.CommentableType, left = findStringWithIndex(pageStr, "commentable_type\":", ",", left)
		cm.CommentableId, left = findWithIndex(pageStr, "commentable_id\":", ",", left)
		cm.LegacyName, left = findWithIndex(pageStr, "legacy_name\":", ",", left)
		cm.CreatedAt, left = findStringWithIndex(pageStr, "created_at\":", ",", left)
		cm.UpdatedAt, left = findStringWithIndex(pageStr, "updated_at\":", ",", left)
		cm.DeletedAt, left = findStringWithIndex(pageStr, "deleted_at\":", ",", left)
		cm.EditedAt, left = findStringWithIndex(pageStr, "edited_at\":", ",", left)
		cm.EditedById, left = findStringWithIndex(pageStr, "edited_by_id\":", ",", left)
		strings.Replace(cm.EditedById, "}", "", 1)

		if i != 0 {
			cm.Message, left = findWithIndex(pageStr, "message\":", "message_html\"", left)
			cm.MessageHtml, left = findWithIndex(pageStr, "message_html\":\"", "\"}", left)
		}

		// Добавление комментария к результату
		result = append(result, cm)
	}

	return result, end
}

// ---------------------- Функции перевода ----------------------

// Функция перевода карты
func formatBeatmap(mpss []MapsString) []Maps {

	var result []Maps

	// Обработка текстовых карт
	for _, mps := range mpss {

		mp := Maps{
			BeatmapSetId:     toInt(mps.BeatmapSetId),
			DifficultyRating: toFloat64(mps.DifficultyRating),
			Id:               toInt(mps.Id),
			Mode:             mps.Mode,
			Status:           mps.Status,
			TotalLength:      toInt(mps.TotalLength),
			UserId:           toInt(mps.UserId),
			Version:          mps.Version,
			Accuracy:         toFloat64(mps.Accuracy),
			Ar:               toFloat64(mps.Ar),
			Bpm:              toFloat64(mps.Bpm),
			Convert:          toBool(mps.Convert),
			CountCircles:     toInt(mps.CountCircles),
			CountSliders:     toInt(mps.CountSliders),
			CountSpinners:    toInt(mps.CountSpinners),
			Cs:               toFloat64(mps.Cs),
			DeletedAt:        mps.DeletedAt,
			Drain:            toFloat64(mps.Drain),
			HitLength:        toInt(mps.HitLength),
			IsScoreable:      toBool(mps.IsScoreable),
			LastUpdated:      mps.LastUpdated,
			ModeInt:          toInt(mps.ModeInt),
			PassCount:        toInt(mps.PassCount),
			PlayCount:        toInt(mps.PlayCount),
			Ranked:           toInt(mps.Ranked),
			Url:              mps.Url,
			Checksum:         mps.Checksum,
			Failtimes: Failtimes{
				Fail: toSlice(mps.Failtimes.Fail),
				Exit: toSlice(mps.Failtimes.Exit),
			},
			MaxCombo: toInt(mps.MaxCombo),
		}

		// Форматирование и добавление рекорда
		result = append(result, mp)
	}

	return result
}

// Функция перевода номинаций
func formatCurrentNominations(cns []CurrentNominationString) []CurrentNomination {

	var result []CurrentNomination

	// Обработка текстовых номинаций
	for _, cn := range cns {
		result = append(result, CurrentNomination{
			BeatmapsetId: toInt(cn.BeatmapsetId),
			Rulesets:     cn.Rulesets,
			Reset:        toBool(cn.Reset),
			UserId:       toInt(cn.UserId),
		})
	}

	return result
}

// Функция перевода пользователя
func formatBmUser(usr BmUserString) BmUser {
	return BmUser{
		AvatarUrl:     usr.AvatarUrl,
		CountryCode:   usr.CountryCode,
		DefaultGroup:  usr.DefaultGroup,
		Id:            toInt(usr.Id),
		IsActive:      toBool(usr.IsActive),
		IsBot:         toBool(usr.IsActive),
		IsDeleted:     toBool(usr.IsDeleted),
		IsOnline:      toBool(usr.IsOnline),
		IsSupporter:   toBool(usr.IsSupporter),
		LastVisit:     usr.LastVisit,
		PmFriendsOnly: toBool(usr.PmFriendsOnly),
		ProfileColor:  usr.ProfileColor,
		Username:      usr.Username,
	}
}

// Функция перевода пользователей
func formatBmUsers(usrs []BmUserString) []BmUser {
	var result []BmUser

	for _, usr := range usrs {
		result = append(result, formatBmUser(usr))
	}

	return result
}

// Функция перевода комментариев
func formatComments(cms []CommentString) []Comment {
	var result []Comment

	for _, cm := range cms {
		result = append(result, Comment{
			Id:              toInt(cm.Id),
			ParentId:        toInt(cm.ParentId),
			UserId:          toInt(cm.UserId),
			Pinned:          toBool(cm.Pinned),
			RepliesCount:    toInt(cm.RepliesCount),
			VotesCount:      toInt(cm.VotesCount),
			CommentableType: cm.CommentableType,
			CommentableId:   toInt(cm.CommentableId),
			LegacyName:      cm.LegacyName,
			CreatedAt:       cm.CreatedAt,
			UpdatedAt:       cm.UpdatedAt,
			DeletedAt:       cm.DeletedAt,
			EditedAt:        cm.EditedAt,
			EditedById:      cm.EditedById,
			Message:         cm.Message,
			MessageHtml:     cm.MessageHtml,
		})
	}

	return result
}

// ----------------- Функции получения статистики ----------------

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
		FavoriteCount:     toInt(resultStr.FavoriteCount),
		HypeCurrent:       toInt(resultStr.HypeCurrent),
		HypeRequired:      toInt(resultStr.HypeRequired),
		Id:                toInt(resultStr.Id),
		Nsfw:              toBool(resultStr.Nsfw),
		Offset:            toInt(resultStr.Offset),
		PlayCount:         toInt(resultStr.PlayCount),
		PreviewUrl:        resultStr.PreviewUrl,
		Source:            resultStr.Source,
		Spotlight:         toBool(resultStr.Spotlight),
		Status:            resultStr.Status,
		Title:             resultStr.Title,
		TitleUnicode:      resultStr.TitleUnicode,
		TrackId:           toInt(resultStr.TrackId),
		UserId:            toInt(resultStr.UserId),
		Video:             toBool(resultStr.Video),
		DownloadDisabled:  toBool(resultStr.DownloadDisabled),
		Bpm:               toFloat64(resultStr.Bpm),
		CanBeHyped:        toBool(resultStr.CanBeHyped),
		DiscussionEnabled: toBool(resultStr.DiscussionEnabled),
		DiscussionLocked:  toBool(resultStr.DiscussionLocked),
		IsScoreable:       toBool(resultStr.IsScoreable),
		LastUpdated:       resultStr.LastUpdated,
		LegacyThreadUrl:   resultStr.LegacyThreadUrl,
		NominationsSummary: NominationsSummary{
			Current:  toInt(resultStr.NominationsSummary.Current),
			Required: toInt(resultStr.NominationsSummary.Required),
		},
		Ranked:             toInt(resultStr.Ranked),
		RankedDate:         resultStr.RankedDate,
		Storyboard:         toBool(resultStr.Storyboard),
		SubmittedDate:      resultStr.SubmittedDate,
		Tags:               resultStr.Tags,
		Beatmaps:           formatBeatmap(resultStr.Beatmaps),
		Converts:           formatBeatmap(resultStr.Converts),
		CurrentNominations: formatCurrentNominations(resultStr.CurrentNominations),
		Description:        resultStr.Description,
		GenreId:            toInt(resultStr.GenreId),
		GenreName:          resultStr.GenreName,
		LanguageId:         toInt(resultStr.LanguageId),
		LanguageName:       resultStr.LanguageName,
		Ratings:            toSlice(resultStr.Ratings),
		RecentFavourites:   formatBmUsers(resultStr.RecentFavourites),
		RelatedUsers:       formatBmUsers(resultStr.RelatedUsers),
		User:               formatBmUser(resultStr.User),
		Comments:           formatComments(resultStr.Comments),
		PinnedComments:     formatComments(resultStr.PinnedComments),
		UserFollow:         toBool(resultStr.UserFollow),
	}

	return result
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

	// Проверка на ошибки
	if resp.StatusCode != 200 {
		return MapStringResponse{
			Error: resp.Status,
		}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// HTML полученной страницы в формате string
	pageStr := string(body)
	pageStr = pageStr[index(pageStr, "<script id=\"json-beatmapset\" type=\"application/json", 80000)+61:]

	// Сохранение html"ки в файл sample.html (для тестов)
	/*if err := os.WriteFile("sampleVotes.html", []byte(pageStr), 0666); err != nil {
		log.Fatal(err)
	}*/

	result := MapStringResponse{}
	left := 0

	result.Artist, left = findWithIndex(pageStr, "\"artist\":\"", "\",", left)
	result.ArtistUnicode, left = findWithIndex(pageStr, "\"artist_unicode\":\"", "\",", left)

	result.Covers.Cover, left = findWithIndex(pageStr, "\"cover\":\"", "\"", left)
	result.Covers.Cover = strings.ReplaceAll(result.Covers.Cover, "\\", "")
	result.Covers.Cover2X, left = findWithIndex(pageStr, "\"cover@2x\":\"", "\"", left)
	result.Covers.Cover2X = strings.ReplaceAll(result.Covers.Cover2X, "\\", "")
	result.Covers.Card, left = findWithIndex(pageStr, "\"card\":\"", "\"", left)
	result.Covers.Card = strings.ReplaceAll(result.Covers.Card, "\\", "")
	result.Covers.Card2X, left = findWithIndex(pageStr, "\"card@2x\":\"", "\"", left)
	result.Covers.Card2X = strings.ReplaceAll(result.Covers.Card2X, "\\", "")
	result.Covers.List, left = findWithIndex(pageStr, "\"list\":\"", "\"", left)
	result.Covers.List = strings.ReplaceAll(result.Covers.List, "\\", "")
	result.Covers.List2X, left = findWithIndex(pageStr, "\"list@2x\":\"", "\"", left)
	result.Covers.List2X = strings.ReplaceAll(result.Covers.List2X, "\\", "")
	result.Covers.SlimCover, left = findWithIndex(pageStr, "\"slimcover\":\"", "\"", left)
	result.Covers.SlimCover = strings.ReplaceAll(result.Covers.SlimCover, "\\", "")
	result.Covers.SlimCover2X, left = findWithIndex(pageStr, "\"slimcover@2x\":\"", "\"", left)
	result.Covers.SlimCover2X = strings.ReplaceAll(result.Covers.SlimCover2X, "\\", "")

	result.Creator, left = findWithIndex(pageStr, "\"creator\":\"", "\",", left)
	result.FavoriteCount, left = findWithIndex(pageStr, "\"favourite_count\":", ",", left)

	if !contains(pageStr, "\"hype\":null", left) {
		result.HypeCurrent, left = findWithIndex(pageStr, "hype\":{\"current\":", ",", left)
		result.HypeRequired, left = findWithIndex(pageStr, "required\":", "}", left)
	}

	result.Id, left = findWithIndex(pageStr, "\"id\":", ",", left)
	result.Nsfw, left = findWithIndex(pageStr, "\"nsfw\":", ",", left)
	result.Offset, left = findWithIndex(pageStr, "\"offset\":", ",", left)
	result.PlayCount, left = findWithIndex(pageStr, "\"play_count\":", ",", left)
	result.PreviewUrl, left = findWithIndex(pageStr, "\"preview_url\":\"\\/\\/", "\",", left)
	result.PreviewUrl = strings.ReplaceAll(result.PreviewUrl, "\\", "")
	result.Source, left = findWithIndex(pageStr, "\"source\":\"", "\",", left)
	result.Spotlight, left = findWithIndex(pageStr, "\"spotlight\":", ",", left)
	result.Status, left = findWithIndex(pageStr, "\"status\":\"", "\",", left)
	result.Title, left = findWithIndex(pageStr, "\"title\":\"", "\",", left)
	result.TitleUnicode, left = findWithIndex(pageStr, "\"title_unicode\":\"", "\",", left)
	result.TrackId, left = findWithIndex(pageStr, "\"track_id\":", ",", left)
	result.UserId, left = findWithIndex(pageStr, "\"user_id\":", ",", left)
	result.Video, left = findWithIndex(pageStr, "\"video\":", ",", left)
	result.DownloadDisabled, left = findWithIndex(pageStr, "\"download_disabled\":", ",", left)
	result.Bpm, left = findWithIndex(pageStr, "\"bpm\":", ",", left)
	result.CanBeHyped, left = findWithIndex(pageStr, "\"can_be_hyped\":", ",", left)
	result.DiscussionEnabled, left = findWithIndex(pageStr, "\"discussion_enabled\":", ",", left)
	result.DiscussionLocked, left = findWithIndex(pageStr, "\"discussion_locked\":", ",", left)
	result.IsScoreable, left = findWithIndex(pageStr, "\"is_scoreable\":", ",", left)
	result.LastUpdated, left = findStringWithIndex(pageStr, "\"last_updated\":", ",", left)
	result.LegacyThreadUrl, left = findStringWithIndex(pageStr, "\"legacy_thread_url\":", ",", left)
	result.LegacyThreadUrl = strings.ReplaceAll(result.LegacyThreadUrl, "\\", "")

	result.NominationsSummary.Current, left = findWithIndex(pageStr, "\"current\":", ",", left)
	result.NominationsSummary.Required, left = findWithIndex(pageStr, "\"required\":", "}", left)

	result.Ranked, left = findWithIndex(pageStr, "\"ranked\":", ",", left)
	result.RankedDate, left = findStringWithIndex(pageStr, "\"ranked_date\":", ",", left)
	result.Storyboard, left = findWithIndex(pageStr, "\"storyboard\":", ",", left)
	result.SubmittedDate, left = findStringWithIndex(pageStr, "\"submitted_date\":", ",", left)

	result.Tags = strings.Split(find(pageStr, "\"tags\":\"", "\",", left), " ")
	if result.Tags[0] == "" {
		result.Tags = nil
	}

	result.Beatmaps, left = parseMapsString(pageStr, "\"beatmaps\":[", "],\"converts\":[", left)
	result.Converts, left = parseMapsString(pageStr, "\"converts\":[", "\"current_nominations\":[", left)
	result.CurrentNominations, left = parseCurrentNominations(pageStr, "current_nominations\":[", "\"description\":{\"description\":", left)

	result.Description, left = findWithIndex(pageStr, "\"description\":{\"description\":\"", "},\"genre\":", left)
	result.GenreId, left = findWithIndex(pageStr, "\"genre\":{\"id\":", ",", left)
	result.GenreName, left = findWithIndex(pageStr, "\"name\":\"", "\"}", left)
	result.LanguageId, left = findWithIndex(pageStr, "\"language\":{\"id\":", ",", left)
	result.LanguageName, left = findWithIndex(pageStr, "\"name\":\"", "\"}", left)
	result.Ratings, left = findWithIndex(pageStr, "ratings\":[", "]", left)

	result.RecentFavourites, left = parseBmUsers(pageStr, "recent_favourites\":[", "related_users\":[", left)
	result.RelatedUsers, left = parseBmUsers(pageStr, "related_users\":[", "user\":{", left)
	result.User, left = parseBmUserString(pageStr, left)

	result.Comments, left = parseCommentsString(pageStr, "comments\":[", "has_more\":", left)
	result.PinnedComments, left = parseCommentsString(pageStr, "\"pinned_comments\":[", "\"user_votes\":[", left)

	result.UserFollow, _ = findWithIndex(pageStr, "\"user_follow\":", ",", left)

	return result
}

// Роут "/map" для vercel
func Map(w http.ResponseWriter, r *http.Request) {

	// Получение параметра id из реквеста
	id := r.URL.Query().Get("id")
	beatmapset := r.URL.Query().Get("beatmapset")

	// Если параметра нет, отправка ошибки
	if id == "" || beatmapset == "" {
		http.NotFound(w, r)
		return
	}

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Проверка на тип, получение статистики, форматирование и отправка
	if r.URL.Query().Get("type") == "string" {
		jsonResp, err := json.Marshal(GetMapInfoString(beatmapset, id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Print("Error: ", err)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
		}
	} else {
		jsonResp, err := json.Marshal(GetMapInfo(beatmapset, id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Print("Error: ", err)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
		}
	}
}
