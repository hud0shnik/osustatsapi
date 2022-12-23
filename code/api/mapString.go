package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// ---------------------- Структуры для парсинга ------------------------

type MapStringResponse struct {
	Error              string                   `json:"error"`
	Artist             string                   `json:"artist"`
	ArtistUnicode      string                   `json:"artist_unicode"`
	Covers             Covers                   `json:"covers"`
	Creator            string                   `json:"creator"`
	FavoriteCount      string                   `json:"favorite_count"`
	HypeCurrent        string                   `json:"hype_current"`
	HypeRequired       string                   `json:"hype_required"`
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
	UserId             string                   `json:"user_id"`
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
	Converts           []MapsString             `json:"converts"`
	Description      string    `json:"description"`
	GenreId          string    `json:"genre_id"`
	GenreName        string    `json:"genre_name"`
	LanguageId       string    `json:"language_id"`
	LanguageName     string    `json:"language_name"`
	Ratings          string    `json:"ratings"`
	RecentFavourites []BmUser  `json:"recent_favourites"`
	RelatedUsers     []BmUser  `json:"related_users"`
	User             BmUser    `json:"user"`
	Comments         []Comment `json:"comments"`
	PinnedComments   []Comment `json:"pinned_comments"`
	UserFollow       string    `json:"user_follow"`
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
	Fail string `json:"fail"`
	Exit string `json:"exit"`
}

type BmUser struct {
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

type Comment struct {
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
	if err := os.WriteFile("sampleVotes.html", []byte(pageStr), 0666); err != nil {
		log.Fatal(err)
	}

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
	result.HypeCurrent, left = findWithIndex(pageStr, "hype\":{\"current\":", ",", left)
	result.HypeRequired, left = findWithIndex(pageStr, "required\":", "}", left)
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

	result.Beatmaps, left = parseMapsString(pageStr, left, "beatmaps")
	result.Converts, left = parseMapsString(pageStr, left, "convert")

	result.Description, left = findWithIndex(pageStr, "\"description\":{\"description\":\"", "},\"genre\":", left)
	result.GenreId, left = findWithIndex(pageStr, "\"genre\":{\"id\":", ",", left)
	result.GenreName, left = findWithIndex(pageStr, "\"name\":\"", "\"}", left)
	result.LanguageId, left = findWithIndex(pageStr, "\"language\":{\"id\":", ",", left)
	result.LanguageName, left = findWithIndex(pageStr, "\"name\":\"", "\"}", left)
	result.Ratings, left = findWithIndex(pageStr, "ratings\":[", "]", left)

	result.RecentFavourites, left = parseBmUsers(pageStr, "recent_favourites\":[", "related_users\":[", left)
	result.RelatedUsers, left = parseBmUsers(pageStr, "related_users\":[", "user\":{", left)
	result.User, left = parseBmUser(pageStr, left)

	result.Comments, left = parseComments(pageStr, "comments\":[", "has_more\":", left)
	result.PinnedComments, left = parseComments(pageStr, "\"pinned_comments\":[", "\"user_votes\":[", left)

	result.UserFollow, _ = findWithIndex(pageStr, "\"user_follow\":", ",", left)

	return result
}

// Функция парсинга карт
func parseMapsString(pageStr string, left int, mapType string) ([]MapsString, int) {

	// Индекс конца карт
	var end int

	// Получение рабочей части и индекса её конца в зависимости от типа карт
	if mapType == "beatmaps" {
		pageStr, end = findWithIndex(pageStr, "\"beatmaps\":[", "],\"converts\":[", left)
	} else {
		pageStr, end = findWithIndex(pageStr, "\"converts\":[", "],\"current_nominations\"", left)
	}

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

		// Добавление карты к результату
		result = append(result, bm)
	}

	return result, end

}

// Функция парсинга пользователей
func parseBmUsers(pageStr, subStr, stopChar string, left int) ([]BmUser, int) {

	// Индекс конца пользователей
	var end int

	// Получение рабочей части и индекса её конца
	pageStr, end = findWithIndex(pageStr, subStr, stopChar, left)

	// Проверка на наличие пользователей
	if len(pageStr) == 0 {
		return []BmUser{}, end
	}

	// Результат и индекс обработанной части
	var result []BmUser
	left = 0

	// Пока есть необработанные пользователи
	for index(pageStr, "avatar_url", left) != -1 {

		// Структура пользователя
		var user BmUser

		user, left = parseBmUser(pageStr, left)

		// Добавление пользователя к результату
		result = append(result, user)
	}

	return result, end

}

// функция парсинга пользователей
func parseBmUser(pageStr string, left int) (BmUser, int) {

	// Структура пользователя
	var user BmUser

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
func parseComments(pageStr, subStr, stopChar string, left int) ([]Comment, int) {

	// Индекс конца комментариев
	var end int

	// Получение рабочей части и индекса её конца
	pageStr, end = findWithIndex(pageStr, subStr, stopChar, left)

	// Проверка на наличие пользователей
	if len(pageStr) == 0 {
		return []Comment{}, end
	}

	// Результат и индекс обработанной части
	result := []Comment{}
	left = 0

	// Пока есть необработанные пользователи
	for i := 0; index(pageStr, "id\":", left) != -1; i++ {

		// Структура комментария
		var cm Comment
		cm.Id, left = findWithIndex(pageStr, "id\":", ",", left)
		cm.ParentId, left = findWithIndex(pageStr, "parent_id\":", ",", left)
		cm.UserId, left = findWithIndex(pageStr, "user_id\":", ",", left)
		cm.Pinned, left = findWithIndex(pageStr, "pinned\":", ",", left)
		cm.RepliesCount, left = findWithIndex(pageStr, "replies_count\":", ",", left)
		cm.VotesCount, left = findWithIndex(pageStr, "votes_count\":", ",", left)
		cm.CommentableType, left = findStringWithIndex(pageStr, "commentable_type:", ",", left)
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
