package handler

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hud0shnik/osustatsapi/internal/convert"
	"github.com/hud0shnik/osustatsapi/internal/parse"
)

// mapResponse - структура респонса
type mapResponse struct {
	Artist             string              `json:"artist"`
	ArtistUnicode      string              `json:"artist_string"`
	Covers             covers              `json:"covers"`
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
	NominationsSummary nominationsSummary  `json:"nominations_summary"`
	Ranked             int                 `json:"ranked"`
	RankedDate         string              `json:"ranked_date"`
	Storyboard         bool                `json:"storyboard"`
	SubmittedDate      string              `json:"submitted_date"`
	Tags               []string            `json:"tags"`
	Beatmaps           []maps              `json:"beatmaps"`
	Converts           []maps              `json:"converts"`
	CurrentNominations []currentNomination `json:"current_nominations"`
	Description        string              `json:"description"`
	GenreId            int                 `json:"genre_id"`
	GenreName          string              `json:"genre_name"`
	LanguageId         int                 `json:"language_id"`
	LanguageName       string              `json:"language_name"`
	Ratings            []int               `json:"ratings"`
	RecentFavourites   []bmUser            `json:"recent_favourites"`
	RelatedUsers       []bmUser            `json:"related_users"`
	User               bmUser              `json:"user"`
	Comments           []comment           `json:"comments"`
	PinnedComments     []comment           `json:"pinned_comments"`
	UserFollow         bool                `json:"user_follow"`
}

// covers - заставки карты
type covers struct {
	Cover       string `json:"cover"`
	Cover2X     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2X      string `json:"card@2x"`
	List        string `json:"list"`
	List2X      string `json:"list@2x"`
	SlimCover   string `json:"slimcover"`
	SlimCover2X string `json:"slimcover@2x"`
}

// nominationsSummary - оценка номинаций карты
type nominationsSummary struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}

// maps - карта внутри сета
type maps struct {
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
	Failtimes        failtimes `json:"failtimes"`
	MaxCombo         int       `json:"max_combo"`
}

// failtimes - проигрыши карты
type failtimes struct {
	Fail []int `json:"fail"`
	Exit []int `json:"exit"`
}

// currentNomination - номинация
type currentNomination struct {
	BeatmapsetId int    `json:"beatmapset_id"`
	Rulesets     string `json:"rulesets"`
	Reset        bool   `json:"reset"`
	UserId       int    `json:"user_id"`
}

// bmUser - пользователь карты
type bmUser struct {
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

// comment - комментарий к карте
type comment struct {
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

// mapStringResponse - респонс в формате строк
type mapStringResponse struct {
	Artist             string                    `json:"artist"`
	ArtistUnicode      string                    `json:"artist_unicode"`
	Covers             covers                    `json:"covers"`
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
	NominationsSummary nominationsSummaryString  `json:"nominations_summary"`
	Ranked             string                    `json:"ranked"`
	RankedDate         string                    `json:"ranked_date"`
	Storyboard         string                    `json:"storyboard"`
	SubmittedDate      string                    `json:"submitted_date"`
	Tags               []string                  `json:"tags"`
	Beatmaps           []mapsString              `json:"beatmaps"`
	Converts           []mapsString              `json:"converts"`
	CurrentNominations []currentNominationString `json:"current_nominations"`
	Description        string                    `json:"description"`
	GenreId            string                    `json:"genre_id"`
	GenreName          string                    `json:"genre_name"`
	LanguageId         string                    `json:"language_id"`
	LanguageName       string                    `json:"language_name"`
	Ratings            string                    `json:"ratings"`
	RecentFavourites   []bmUserString            `json:"recent_favourites"`
	RelatedUsers       []bmUserString            `json:"related_users"`
	User               bmUserString              `json:"user"`
	Comments           []commentString           `json:"comments"`
	PinnedComments     []commentString           `json:"pinned_comments"`
	UserFollow         string                    `json:"user_follow"`
}

// nominationsSummaryString - оценка номинаций в формате строк
type nominationsSummaryString struct {
	Current  string `json:"current"`
	Required string `json:"required"`
}

// mapsString - карта в формате строк
type mapsString struct {
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
	Failtimes        failtimesString `json:"failtimes"`
	MaxCombo         string          `json:"max_combo"`
}

// failtimesString - проигрыши в формате строк
type failtimesString struct {
	Fail string `json:"fail"`
	Exit string `json:"exit"`
}

// currentNominationString - номинация в формате строк
type currentNominationString struct {
	BeatmapsetId string `json:"beatmapset_id"`
	Rulesets     string `json:"rulesets"`
	Reset        string `json:"reset"`
	UserId       string `json:"user_id"`
}

// bmUserString - пользователь в формате строк
type bmUserString struct {
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

// commentString - комментарий в формате строк
type commentString struct {
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

// parseCurrentNominations парсит номинации в pageStr[left+index(subStr):] и возвращает их в формате строк вместе с индексом конца
func parseCurrentNominations(pageStr, subStr, stopChar string, left int) ([]currentNominationString, int) {

	// Индекс конца номинаций
	var end int

	// Получение рабочей части и индекса её конца
	pageStr, end = parse.FindWithIndex(pageStr, subStr, stopChar, left, -1)

	// Проверка на наличие номинаций
	if len(pageStr) == 0 {
		return []currentNominationString{}, end
	}

	// Результат и индекс обработанной части
	var result []currentNominationString
	left = 0

	// Пока есть необработанные карты
	for parse.Index(pageStr, "beatmapset_id", left, -1) != -1 {

		// Структура номинации
		var cn currentNominationString

		// Запись данных
		cn.BeatmapsetId, left = parse.FindWithIndex(pageStr, "beatmapset_id\":", ",", left, -1)
		cn.Rulesets, left = parse.FindWithIndex(pageStr, "rulesets\":[", "]", left, -1)
		cn.Reset, left = parse.FindWithIndex(pageStr, "reset\":", ",", left, -1)
		cn.UserId, left = parse.FindWithIndex(pageStr, "user_id\":", "}", left, -1)

		// Добавление карты к результату
		result = append(result, cn)
	}

	return result, end

}

// parseMapString парсит карту в pageStr[left:] и возвращает их в формате строк с индексом конца
func parseMapString(pageStr string, left int) (mapsString, int) {

	// Структура карты
	var bm mapsString

	// Запись данных
	bm.BeatmapSetId, left = parse.FindWithIndex(pageStr, "\"beatmapset_id\":", ",", left, -1)
	bm.DifficultyRating, left = parse.FindWithIndex(pageStr, "\"difficulty_rating\":", ",", left, -1)
	bm.Id, left = parse.FindWithIndex(pageStr, "\"id\":", ",", left, -1)
	bm.Mode, left = parse.FindStringWithIndex(pageStr, "\"mode\":", ",", left, -1)
	bm.Status, left = parse.FindStringWithIndex(pageStr, "\"status\":", ",", left, -1)
	bm.TotalLength, left = parse.FindWithIndex(pageStr, "total_length\":", ",", left, -1)
	bm.UserId, left = parse.FindWithIndex(pageStr, "\"user_id\":", ",", left, -1)
	bm.Version, left = parse.FindWithIndex(pageStr, "\"version\":\"", "\"", left, -1)
	bm.Accuracy, left = parse.FindWithIndex(pageStr, "\"accuracy\":", ",", left, -1)
	bm.Ar, left = parse.FindWithIndex(pageStr, "\"ar\":", ",", left, -1)
	bm.Bpm, left = parse.FindWithIndex(pageStr, "\"bpm\":", ",", left, -1)
	bm.Convert, left = parse.FindWithIndex(pageStr, "\"convert\":", ",", left, -1)
	bm.CountCircles, left = parse.FindWithIndex(pageStr, "\"count_circles\":", ",", left, -1)
	bm.CountSliders, left = parse.FindWithIndex(pageStr, "\"count_sliders\":", ",", left, -1)
	bm.CountSpinners, left = parse.FindWithIndex(pageStr, "\"count_spinners\":", ",", left, -1)
	bm.Cs, left = parse.FindWithIndex(pageStr, "\"cs\":", ",", left, -1)
	bm.DeletedAt, left = parse.FindStringWithIndex(pageStr, "\"deleted_at\":", ",", left, -1)
	bm.Drain, left = parse.FindWithIndex(pageStr, "\"drain\":", ",", left, -1)
	bm.HitLength, left = parse.FindWithIndex(pageStr, "\"hit_length\":", ",", left, -1)
	bm.IsScoreable, left = parse.FindWithIndex(pageStr, "\"is_scoreable\":", ",", left, -1)
	bm.LastUpdated, left = parse.FindStringWithIndex(pageStr, "\"last_updated\":", ",", left, -1)
	bm.ModeInt, left = parse.FindWithIndex(pageStr, "\"mode_int\":", ",", left, -1)
	bm.PassCount, left = parse.FindWithIndex(pageStr, "\"passcount\":", ",", left, -1)
	bm.PlayCount, left = parse.FindWithIndex(pageStr, "\"playcount\":", ",", left, -1)
	bm.Ranked, left = parse.FindWithIndex(pageStr, "\"ranked\":", ",", left, -1)
	bm.Url, left = parse.FindStringWithIndex(pageStr, "\"url\":", ",", left, -1)
	bm.Url = strings.ReplaceAll(bm.Url, "\\", "")
	bm.Checksum, left = parse.FindStringWithIndex(pageStr, "\"checksum\":", ",", left, -1)

	bm.Failtimes.Fail, left = parse.FindWithIndex(pageStr, "\"fail\":[", "]", left, -1)
	bm.Failtimes.Exit, left = parse.FindWithIndex(pageStr, "\"exit\":[", "]", left, -1)

	bm.MaxCombo, left = parse.FindWithIndex(pageStr, "\"max_combo\":", "}", left, -1)

	return bm, left

}

// parseMapsString парсит карты сета в pageStr[left+index(subStr):] и возвращает их в формате строк с индексом конца
func parseMapsString(pageStr, subStr, stopChar string, left int) ([]mapsString, int) {

	// Индекс конца карт
	var end int

	// Получение рабочей части и индекса её конца
	pageStr, end = parse.FindWithIndex(pageStr, subStr, stopChar, left, -1)

	// Проверка на наличие карт
	if len(pageStr) == 0 {
		return []mapsString{}, end
	}

	// Результат и индекс обработанной части
	var result []mapsString
	left = 0

	// Пока есть необработанные карты
	for parse.Index(pageStr, "\"beatmapset_id\"", left, -1) != -1 {

		// Структура карты
		var bm mapsString

		bm, left = parseMapString(pageStr, left)

		// Добавление карты к результату
		result = append(result, bm)

	}

	return result, end

}

// parseBmUsers парсит список игроков карты в pageStr[left+index(subStr):] и возвращает их в формате строк с индексом конца
func parseBmUsers(pageStr, subStr, stopChar string, left int) ([]bmUserString, int) {

	// Индекс конца пользователей
	var end int

	// Получение рабочей части и индекса её конца
	pageStr, end = parse.FindWithIndex(pageStr, subStr, stopChar, left, -1)

	// Проверка на наличие пользователей
	if len(pageStr) == 0 {
		return []bmUserString{}, end
	}

	// Результат и индекс обработанной части
	var result []bmUserString
	left = 0

	// Пока есть необработанные пользователи
	for parse.Index(pageStr, "avatar_url", left, -1) != -1 {

		// Структура пользователя
		var user bmUserString

		user, left = parseBmUserString(pageStr, left)

		// Добавление пользователя к результату
		result = append(result, user)

	}

	return result, end

}

// parseBmUserString парсит игрока карты в pageStr[left+index(subStr):] и возвращает его в формате строк с индексом конца
func parseBmUserString(pageStr string, left int) (bmUserString, int) {

	// Структура пользователя
	var user bmUserString

	// Запись данных
	user.AvatarUrl, left = parse.FindStringWithIndex(pageStr, "avatar_url\":", ",", left, -1)
	user.CountryCode, left = parse.FindStringWithIndex(pageStr, "country_code\":", ",", left, -1)
	user.DefaultGroup, left = parse.FindStringWithIndex(pageStr, "default_group\":", ",", left, -1)
	user.Id, left = parse.FindWithIndex(pageStr, "id\":", ",", left, -1)
	user.IsActive, left = parse.FindWithIndex(pageStr, "is_active\":", ",", left, -1)
	user.IsBot, left = parse.FindWithIndex(pageStr, "is_bot\":", ",", left, -1)
	user.IsDeleted, left = parse.FindWithIndex(pageStr, "is_deleted\":", ",", left, -1)
	user.IsOnline, left = parse.FindWithIndex(pageStr, "is_online\":", ",", left, -1)
	user.IsSupporter, left = parse.FindWithIndex(pageStr, "is_supporter\":", ",", left, -1)
	user.LastVisit, left = parse.FindStringWithIndex(pageStr, "last_visit\":", ",", left, -1)
	user.PmFriendsOnly, left = parse.FindWithIndex(pageStr, "pm_friends_only\":", ",", left, -1)
	user.ProfileColor, left = parse.FindStringWithIndex(pageStr, "profile_colour\":", ",", left, -1)
	user.Username, left = parse.FindStringWithIndex(pageStr, "username\":", "}", left, -1)

	return user, left

}

// parseCommentsString парсит комментарии в pageStr[left+index(subStr):index(stopChar)] и возвращает их в формате строк с индексом конца
func parseCommentsString(pageStr, subStr, stopChar string, left int) ([]commentString, int) {

	// Индекс конца комментариев
	var end int

	// Получение рабочей части и индекса её конца
	pageStr, end = parse.FindWithIndex(pageStr, subStr, stopChar, left, -1)

	// Проверка на наличие пользователей
	if len(pageStr) == 0 {
		return []commentString{}, end
	}

	// Результат и индекс обработанной части
	result := []commentString{}
	left = 0

	// Пока есть необработанные пользователи
	for i := 0; parse.Index(pageStr, "id\":", left, -1) != -1; i++ {

		// Структура комментария
		var cm commentString
		cm.Id, left = parse.FindWithIndex(pageStr, "id\":", ",", left, -1)
		cm.ParentId, left = parse.FindWithIndex(pageStr, "parent_id\":", ",", left, -1)
		cm.UserId, left = parse.FindWithIndex(pageStr, "user_id\":", ",", left, -1)
		cm.Pinned, left = parse.FindWithIndex(pageStr, "pinned\":", ",", left, -1)
		cm.RepliesCount, left = parse.FindWithIndex(pageStr, "replies_count\":", ",", left, -1)
		cm.VotesCount, left = parse.FindWithIndex(pageStr, "votes_count\":", ",", left, -1)
		cm.CommentableType, left = parse.FindStringWithIndex(pageStr, "commentable_type\":", ",", left, -1)
		cm.CommentableId, left = parse.FindWithIndex(pageStr, "commentable_id\":", ",", left, -1)
		cm.LegacyName, left = parse.FindWithIndex(pageStr, "legacy_name\":", ",", left, -1)
		cm.CreatedAt, left = parse.FindStringWithIndex(pageStr, "created_at\":", ",", left, -1)
		cm.UpdatedAt, left = parse.FindStringWithIndex(pageStr, "updated_at\":", ",", left, -1)
		cm.DeletedAt, left = parse.FindStringWithIndex(pageStr, "deleted_at\":", ",", left, -1)
		cm.EditedAt, left = parse.FindStringWithIndex(pageStr, "edited_at\":", ",", left, -1)
		cm.EditedById, left = parse.FindStringWithIndex(pageStr, "edited_by_id\":", ",", left, -1)
		strings.Replace(cm.EditedById, "}", "", 1)

		if i != 0 {
			cm.Message, left = parse.FindWithIndex(pageStr, "message\":", "message_html\"", left, -1)
			cm.MessageHtml, left = parse.FindWithIndex(pageStr, "message_html\":\"", "\"}", left, -1)
		}

		// Добавление комментария к результату
		result = append(result, cm)

	}

	return result, end

}

// formatBeatmap переводит карту из string формата в стандартный
func formatBeatmap(mpss []mapsString) []maps {

	var result []maps

	// Обработка текстовых карт
	for _, mps := range mpss {

		mp := maps{
			BeatmapSetId:     convert.ToInt(mps.BeatmapSetId),
			DifficultyRating: convert.ToFloat64(mps.DifficultyRating),
			Id:               convert.ToInt(mps.Id),
			Mode:             mps.Mode,
			Status:           mps.Status,
			TotalLength:      convert.ToInt(mps.TotalLength),
			UserId:           convert.ToInt(mps.UserId),
			Version:          mps.Version,
			Accuracy:         convert.ToFloat64(mps.Accuracy),
			Ar:               convert.ToFloat64(mps.Ar),
			Bpm:              convert.ToFloat64(mps.Bpm),
			Convert:          convert.ToBool(mps.Convert),
			CountCircles:     convert.ToInt(mps.CountCircles),
			CountSliders:     convert.ToInt(mps.CountSliders),
			CountSpinners:    convert.ToInt(mps.CountSpinners),
			Cs:               convert.ToFloat64(mps.Cs),
			DeletedAt:        mps.DeletedAt,
			Drain:            convert.ToFloat64(mps.Drain),
			HitLength:        convert.ToInt(mps.HitLength),
			IsScoreable:      convert.ToBool(mps.IsScoreable),
			LastUpdated:      mps.LastUpdated,
			ModeInt:          convert.ToInt(mps.ModeInt),
			PassCount:        convert.ToInt(mps.PassCount),
			PlayCount:        convert.ToInt(mps.PlayCount),
			Ranked:           convert.ToInt(mps.Ranked),
			Url:              mps.Url,
			Checksum:         mps.Checksum,
			Failtimes: failtimes{
				Fail: convert.ToSlice(mps.Failtimes.Fail),
				Exit: convert.ToSlice(mps.Failtimes.Exit),
			},
			MaxCombo: convert.ToInt(mps.MaxCombo),
		}

		// Форматирование и добавление рекорда
		result = append(result, mp)
	}

	return result

}

// formatCurrentNominations переводит номинации из string формата в стандартный
func formatCurrentNominations(cns []currentNominationString) []currentNomination {

	var result []currentNomination

	// Обработка текстовых номинаций
	for _, cn := range cns {
		result = append(result, currentNomination{
			BeatmapsetId: convert.ToInt(cn.BeatmapsetId),
			Rulesets:     cn.Rulesets,
			Reset:        convert.ToBool(cn.Reset),
			UserId:       convert.ToInt(cn.UserId),
		})
	}

	return result

}

// formatBmUser переводит пользователя из string формата в стандартный
func formatBmUser(usr bmUserString) bmUser {

	return bmUser{
		AvatarUrl:     usr.AvatarUrl,
		CountryCode:   usr.CountryCode,
		DefaultGroup:  usr.DefaultGroup,
		Id:            convert.ToInt(usr.Id),
		IsActive:      convert.ToBool(usr.IsActive),
		IsBot:         convert.ToBool(usr.IsActive),
		IsDeleted:     convert.ToBool(usr.IsDeleted),
		IsOnline:      convert.ToBool(usr.IsOnline),
		IsSupporter:   convert.ToBool(usr.IsSupporter),
		LastVisit:     usr.LastVisit,
		PmFriendsOnly: convert.ToBool(usr.PmFriendsOnly),
		ProfileColor:  usr.ProfileColor,
		Username:      usr.Username,
	}

}

// formatBmUsers переводит пользователей из string формата в стандартный
func formatBmUsers(usrs []bmUserString) []bmUser {

	var result []bmUser

	for _, usr := range usrs {
		result = append(result, formatBmUser(usr))
	}

	return result

}

// formatComments переводит комментарии из string формата в стандартный
func formatComments(cms []commentString) []comment {

	var result []comment

	for _, cm := range cms {
		result = append(result, comment{
			Id:              convert.ToInt(cm.Id),
			ParentId:        convert.ToInt(cm.ParentId),
			UserId:          convert.ToInt(cm.UserId),
			Pinned:          convert.ToBool(cm.Pinned),
			RepliesCount:    convert.ToInt(cm.RepliesCount),
			VotesCount:      convert.ToInt(cm.VotesCount),
			CommentableType: cm.CommentableType,
			CommentableId:   convert.ToInt(cm.CommentableId),
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

// getMapInfoString возвращает статистику карты в формате строк, статус код и ошибку
func getMapInfoString(id string) (mapStringResponse, int, error) {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/b/" + id + "?m=0")
	if err != nil {
		return mapStringResponse{}, http.StatusInternalServerError,
			fmt.Errorf("in http.Get: %w", err)
	}
	defer resp.Body.Close()

	// Проверка статускода
	if resp.StatusCode != 200 {
		return mapStringResponse{}, resp.StatusCode,
			fmt.Errorf("in http.Get: %s", resp.Status)
	}

	// Запись респонса
	body, _ := io.ReadAll(resp.Body)

	// Полученная страница в формате string
	pageStr := string(body)
	pageStr = pageStr[parse.Index(pageStr, "<script id=\"json-beatmapset\" type=\"application/json", 80000, -1)+61:]

	// Сохранение html"ки в файл sample.html (для тестов)
	/*if err := os.WriteFile("sampleMap.html", []byte(pageStr), 0666); err != nil {
		logrus.Fatal(err)
	}*/

	// Структура результата и указатель отработанной части
	result := mapStringResponse{}
	left := 0

	result.Artist, left = parse.FindWithIndex(pageStr, "\"artist\":\"", "\",", left, -1)
	result.ArtistUnicode, left = parse.FindWithIndex(pageStr, "\"artist_unicode\":\"", "\",", left, -1)

	result.Covers.Cover, left = parse.FindWithIndex(pageStr, "\"cover\":\"", "\"", left, -1)
	result.Covers.Cover = strings.ReplaceAll(result.Covers.Cover, "\\", "")
	result.Covers.Cover2X, left = parse.FindWithIndex(pageStr, "\"cover@2x\":\"", "\"", left, -1)
	result.Covers.Cover2X = strings.ReplaceAll(result.Covers.Cover2X, "\\", "")
	result.Covers.Card, left = parse.FindWithIndex(pageStr, "\"card\":\"", "\"", left, -1)
	result.Covers.Card = strings.ReplaceAll(result.Covers.Card, "\\", "")
	result.Covers.Card2X, left = parse.FindWithIndex(pageStr, "\"card@2x\":\"", "\"", left, -1)
	result.Covers.Card2X = strings.ReplaceAll(result.Covers.Card2X, "\\", "")
	result.Covers.List, left = parse.FindWithIndex(pageStr, "\"list\":\"", "\"", left, -1)
	result.Covers.List = strings.ReplaceAll(result.Covers.List, "\\", "")
	result.Covers.List2X, left = parse.FindWithIndex(pageStr, "\"list@2x\":\"", "\"", left, -1)
	result.Covers.List2X = strings.ReplaceAll(result.Covers.List2X, "\\", "")
	result.Covers.SlimCover, left = parse.FindWithIndex(pageStr, "\"slimcover\":\"", "\"", left, -1)
	result.Covers.SlimCover = strings.ReplaceAll(result.Covers.SlimCover, "\\", "")
	result.Covers.SlimCover2X, left = parse.FindWithIndex(pageStr, "\"slimcover@2x\":\"", "\"", left, -1)
	result.Covers.SlimCover2X = strings.ReplaceAll(result.Covers.SlimCover2X, "\\", "")

	result.Creator, left = parse.FindWithIndex(pageStr, "\"creator\":\"", "\",", left, -1)
	result.FavoriteCount, left = parse.FindWithIndex(pageStr, "\"favourite_count\":", ",", left, -1)

	if !parse.Contains(pageStr, "\"hype\":null", left) {
		result.HypeCurrent, left = parse.FindWithIndex(pageStr, "hype\":{\"current\":", ",", left, -1)
		result.HypeRequired, left = parse.FindWithIndex(pageStr, "required\":", "}", left, -1)
	}

	result.Id, left = parse.FindWithIndex(pageStr, "\"id\":", ",", left, -1)
	result.Nsfw, left = parse.FindWithIndex(pageStr, "\"nsfw\":", ",", left, -1)
	result.Offset, left = parse.FindWithIndex(pageStr, "\"offset\":", ",", left, -1)
	result.PlayCount, left = parse.FindWithIndex(pageStr, "\"play_count\":", ",", left, -1)
	result.PreviewUrl, left = parse.FindWithIndex(pageStr, "\"preview_url\":\"\\/\\/", "\",", left, -1)
	result.PreviewUrl = strings.ReplaceAll(result.PreviewUrl, "\\", "")
	result.Source, left = parse.FindWithIndex(pageStr, "\"source\":\"", "\",", left, -1)
	result.Spotlight, left = parse.FindWithIndex(pageStr, "\"spotlight\":", ",", left, -1)
	result.Status, left = parse.FindWithIndex(pageStr, "\"status\":\"", "\",", left, -1)
	result.Title, left = parse.FindWithIndex(pageStr, "\"title\":\"", "\",", left, -1)
	result.TitleUnicode, left = parse.FindWithIndex(pageStr, "\"title_unicode\":\"", "\",", left, -1)
	result.TrackId, left = parse.FindWithIndex(pageStr, "\"track_id\":", ",", left, -1)
	result.UserId, left = parse.FindWithIndex(pageStr, "\"user_id\":", ",", left, -1)
	result.Video, left = parse.FindWithIndex(pageStr, "\"video\":", ",", left, -1)
	result.Bpm, left = parse.FindWithIndex(pageStr, "\"bpm\":", ",", left, -1)
	result.CanBeHyped, left = parse.FindWithIndex(pageStr, "\"can_be_hyped\":", ",", left, -1)
	result.DiscussionEnabled, left = parse.FindWithIndex(pageStr, "\"discussion_enabled\":", ",", left, -1)
	result.DiscussionLocked, left = parse.FindWithIndex(pageStr, "\"discussion_locked\":", ",", left, -1)
	result.IsScoreable, left = parse.FindWithIndex(pageStr, "\"is_scoreable\":", ",", left, -1)
	result.LastUpdated, left = parse.FindStringWithIndex(pageStr, "\"last_updated\":", ",", left, -1)
	result.LegacyThreadUrl, left = parse.FindStringWithIndex(pageStr, "\"legacy_thread_url\":", ",", left, -1)
	result.LegacyThreadUrl = strings.ReplaceAll(result.LegacyThreadUrl, "\\", "")
	result.NominationsSummary.Current, left = parse.FindWithIndex(pageStr, "\"current\":", ",", left, -1)
	result.NominationsSummary.Required, left = parse.FindWithIndex(pageStr, "\"required\":", "}", left, -1)

	result.Ranked, left = parse.FindWithIndex(pageStr, "\"ranked\":", ",", left, -1)
	result.RankedDate, left = parse.FindStringWithIndex(pageStr, "\"ranked_date\":", ",", left, -1)
	result.Storyboard, left = parse.FindWithIndex(pageStr, "\"storyboard\":", ",", left, -1)
	result.SubmittedDate, left = parse.FindStringWithIndex(pageStr, "\"submitted_date\":", ",", left, -1)

	result.Tags = strings.Split(parse.Find(pageStr, "\"tags\":\"", "\",", left), " ")
	if result.Tags[0] == "" {
		result.Tags = nil
	}

	result.DownloadDisabled, left = parse.FindWithIndex(pageStr, "\"download_disabled\":", ",", left, -1)
	result.Beatmaps, left = parseMapsString(pageStr, "\"beatmaps\":[", "],\"converts\":[", left)
	result.Converts, left = parseMapsString(pageStr, "\"converts\":[", "\"current_nominations\":[", left)
	result.CurrentNominations, left = parseCurrentNominations(pageStr, "current_nominations\":[", "\"description\":{\"description\":", left)

	result.Description, left = parse.FindWithIndex(pageStr, "\"description\":{\"description\":\"", "},\"genre\":", left, -1)
	result.GenreId, left = parse.FindWithIndex(pageStr, "\"genre\":{\"id\":", ",", left, -1)
	result.GenreName, left = parse.FindWithIndex(pageStr, "\"name\":\"", "\"}", left, -1)
	result.LanguageId, left = parse.FindWithIndex(pageStr, "\"language\":{\"id\":", ",", left, -1)
	result.LanguageName, left = parse.FindWithIndex(pageStr, "\"name\":\"", "\"}", left, -1)
	result.Ratings, left = parse.FindWithIndex(pageStr, "ratings\":[", "]", left, -1)

	result.RecentFavourites, left = parseBmUsers(pageStr, "recent_favourites\":[", "related_users\":[", left)
	result.RelatedUsers, left = parseBmUsers(pageStr, "related_users\":[", "user\":{", left)
	result.User, left = parseBmUserString(pageStr, left)

	result.Comments, left = parseCommentsString(pageStr, "comments\":[", "has_more\":", left)
	result.PinnedComments, left = parseCommentsString(pageStr, "\"pinned_comments\":[", "\"user_votes\":[", left)

	result.UserFollow, _ = parse.FindWithIndex(pageStr, "\"user_follow\":", ",", left, -1)

	return result, http.StatusOK, nil

}

// getMapInfo возвращает статистику карты, статус код и ошибку
func getMapInfo(id string) (mapResponse, int, error) {

	// Получение текстовой версии статистики
	resultStr, statusCode, err := getMapInfoString(id)
	if err != nil {
		return mapResponse{}, statusCode, err
	}

	// Перевод в классическую версию
	result := mapResponse{
		Artist:            resultStr.Artist,
		ArtistUnicode:     resultStr.ArtistUnicode,
		Covers:            resultStr.Covers,
		Creator:           resultStr.Creator,
		FavoriteCount:     convert.ToInt(resultStr.FavoriteCount),
		HypeCurrent:       convert.ToInt(resultStr.HypeCurrent),
		HypeRequired:      convert.ToInt(resultStr.HypeRequired),
		Id:                convert.ToInt(resultStr.Id),
		Nsfw:              convert.ToBool(resultStr.Nsfw),
		Offset:            convert.ToInt(resultStr.Offset),
		PlayCount:         convert.ToInt(resultStr.PlayCount),
		PreviewUrl:        resultStr.PreviewUrl,
		Source:            resultStr.Source,
		Spotlight:         convert.ToBool(resultStr.Spotlight),
		Status:            resultStr.Status,
		Title:             resultStr.Title,
		TitleUnicode:      resultStr.TitleUnicode,
		TrackId:           convert.ToInt(resultStr.TrackId),
		UserId:            convert.ToInt(resultStr.UserId),
		Video:             convert.ToBool(resultStr.Video),
		DownloadDisabled:  convert.ToBool(resultStr.DownloadDisabled),
		Bpm:               convert.ToFloat64(resultStr.Bpm),
		CanBeHyped:        convert.ToBool(resultStr.CanBeHyped),
		DiscussionEnabled: convert.ToBool(resultStr.DiscussionEnabled),
		DiscussionLocked:  convert.ToBool(resultStr.DiscussionLocked),
		IsScoreable:       convert.ToBool(resultStr.IsScoreable),
		LastUpdated:       resultStr.LastUpdated,
		LegacyThreadUrl:   resultStr.LegacyThreadUrl,
		NominationsSummary: nominationsSummary{
			Current:  convert.ToInt(resultStr.NominationsSummary.Current),
			Required: convert.ToInt(resultStr.NominationsSummary.Required),
		},
		Ranked:             convert.ToInt(resultStr.Ranked),
		RankedDate:         resultStr.RankedDate,
		Storyboard:         convert.ToBool(resultStr.Storyboard),
		SubmittedDate:      resultStr.SubmittedDate,
		Tags:               resultStr.Tags,
		Beatmaps:           formatBeatmap(resultStr.Beatmaps),
		Converts:           formatBeatmap(resultStr.Converts),
		CurrentNominations: formatCurrentNominations(resultStr.CurrentNominations),
		Description:        resultStr.Description,
		GenreId:            convert.ToInt(resultStr.GenreId),
		GenreName:          resultStr.GenreName,
		LanguageId:         convert.ToInt(resultStr.LanguageId),
		LanguageName:       resultStr.LanguageName,
		Ratings:            convert.ToSlice(resultStr.Ratings),
		RecentFavourites:   formatBmUsers(resultStr.RecentFavourites),
		RelatedUsers:       formatBmUsers(resultStr.RelatedUsers),
		User:               formatBmUser(resultStr.User),
		Comments:           formatComments(resultStr.Comments),
		PinnedComments:     formatComments(resultStr.PinnedComments),
		UserFollow:         convert.ToBool(resultStr.UserFollow),
	}

	return result, http.StatusOK, nil

}

// Map - роут "/map"
func Map(w http.ResponseWriter, r *http.Request) {

	// Получение параметров из реквеста
	id := r.URL.Query().Get("id")

	// Проверка на наличие параметров
	if id == "" {
		response(w, http.StatusBadRequest, apiError{Error: "please insert map id"})
		return
	}

	// Проверка на тип
	if r.URL.Query().Get("type") == "string" {

		// Получение статистики
		result, statusCode, err := getMapInfoString(id)
		if err != nil {
			response(w, statusCode, apiError{Error: err.Error()})
			return
		}

		response(w, statusCode, result)

	} else {

		// Получение статистики
		result, statusCode, err := getMapInfo(id)
		if err != nil {
			response(w, statusCode, apiError{Error: err.Error()})
			return
		}

		response(w, statusCode, result)

	}

}
