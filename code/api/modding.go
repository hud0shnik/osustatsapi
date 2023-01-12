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
	Success        string                     `json:"success"`
	Error          string                     `json:"error"`
	Events         []EventString              `json:"events"`
	Users          []ModdingUserString        `json:"users"`
	Beatmaps       []BeatmapsString           `json:"beatmaps"`
	Beatmapsets    []ModdingBeatmapsetsString `json:"beatmapsets"`
	Discussions    []DiscussionString         `json:"discussions"`
	Posts          []PostString               `json:"posts"`
	ReceivedKudosu ReceivedKudosuString       `json:"recently_received_kudosu"`
}

// Структура события
type EventString struct {
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

// Структура пользователя
type ModdingUserString struct {
	AvatarUrl     string        `json:"avatar_url"`
	CountryCode   string        `json:"country_code"`
	DefaultGroup  string        `json:"default_group"`
	Id            int           `json:"id "`
	IsActive      bool          `json:"is_active"`
	IsBot         bool          `json:"is_bot"`
	IsDeleted     bool          `json:"is_deleted"`
	IsOnline      bool          `json:"is_online"`
	IsSupporter   bool          `json:"is_supporter"`
	LastVisit     string        `json:"last_visit"`
	PmFriendsOnly bool          `json:"pm_friends_only"`
	ProfileColor  string        `json:"profile_color"`
	Username      string        `json:"username"`
	Groups        []GroupString `json:"groups"`
}

// Структура группы
type GroupString struct {
	Colour         string   `json:"colour"`
	HasListing     string   `json:"has_listing"`
	HasPlaymodes   string   `json:"has_playmodes"`
	Id             string   `json:"id"`
	Identifier     string   `json:"identifier"`
	IsProbationary string   `json:"is_probationary"`
	Name           string   `json:"name"`
	ShortName      string   `json:"short_name"`
	Playmodes      []string `json:"playmodes"`
}

// Структура сета
type ModdingBeatmapsetsString struct {
	Artist            string                   `json:"artist"`
	ArtistUnicode     string                   `json:"artist_unicode"`
	Covers            Covers                   `json:"covers"`
	Creator           string                   `json:"creator"`
	FavoriteCount     string                   `json:"favorite_count"`
	Hype              string                   `json:"hype"`
	Id                string                   `json:"id"`
	Nsfw              string                   `json:"nsfw"`
	Offset            string                   `json:"offset"`
	PlayCount         string                   `json:"play_count"`
	PreviewUrl        string                   `json:"preview_url"`
	Source            string                   `json:"source"`
	Spotlight         string                   `json:"spotlight"`
	Status            string                   `json:"status"`
	Title             string                   `json:"title"`
	TitleUnicode      string                   `json:"title_unicode"`
	TrackId           string                   `json:"track_id"`
	UserId            string                   `json:"userId"`
	Video             string                   `json:"video"`
	DownloadDisabled  string                   `json:"download_disabled"`
	Bpm               string                   `json:"bpm"`
	CanBeHyped        string                   `json:"can_be_hyped"`
	DiscussionEnabled string                   `json:"discussion_enabled"`
	DiscussionLocked  string                   `json:"discussion_locked"`
	IsScoreable       string                   `json:"is_scoreable"`
	LastUpdated       string                   `json:"last_updated"`
	LegacyThreadUrl   string                   `json:"legacy_thread_url"`
	Nominations       NominationsSummaryString `json:"nominations_summary"`
	Ranked            string                   `json:"ranked"`
	RankedDate        string                   `json:"ranked_date"`
	Storyboard        string                   `json:"storyboard"`
	SubmittedDate     string                   `json:"submitted_date"`
	Tags              []string                 `json:"tags"`
}

// Структура поста
type PostString struct {
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

// Структура дискуссии поста
type DiscussionPostString struct {
	Id             string           `json:"id"`
	BeatmapsetId   string           `json:"beatmapset_id"`
	BeatmapId      string           `json:"beatmap_id"`
	UserId         string           `json:"user_id"`
	DeletedById    string           `json:"deleted_by_id"`
	MessageType    string           `json:"message_type"`
	ParentId       string           `json:"parent_id"`
	Timestamp      string           `json:"timestamp"`
	Resolved       string           `json:"resolved"`
	CanBeResolved  string           `json:"can_be_resolved"`
	CanGrantKudosu string           `json:"can_grant_kudosu"`
	CreatedAt      string           `json:"created_at"`
	UpdatedAt      string           `json:"updated_at"`
	DeletedAt      string           `json:"deleted_at"`
	LastPostAt     string           `json:"last_post_at"`
	KudosuDenied   string           `json:"kudosu_denied"`
	StartingPost   StartingPost     `json:"starting_post"`
	Beatmapset     BeatmapsetString `json:"beatmapset"`
}

// Структура голосов
type VotesString struct {
	Given    []VoteString `json:"given"`
	Received []VoteString `json:"received"`
}

// Структура кудосу
type ReceivedKudosuString struct {
	Id        string `json:"id"`
	Action    string `json:"action"`
	Amount    string `json:"amount"`
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Giver     string `json:"giver"`
	PostUrl   string `json:"post_url"`
	PostTitle string `json:"post_title"`
	Details   string `json:"details"`
}

// Структура пользователя
type UserString struct {
	Success                string      `json:"success"`
	Error                  string      `json:"error"`
	AvatarUrl              string      `json:"avatar_url"`
	CountryCode            string      `json:"country_code"`
	DefaultGroup           string      `json:"default_group"`
	Id                     string      `json:"id"`
	IsActive               string      `json:"is_active"`
	IsBot                  string      `json:"is_bot"`
	IsDeleted              string      `json:"is_deleted"`
	IsOnline               string      `json:"is_online"`
	IsSupporter            string      `json:"is_supporter"`
	LastVisit              string      `json:"last_visit"`
	PmFriendsOnly          string      `json:"pm_friends_only"`
	ProfileColor           string      `json:"profile_color"`
	Username               string      `json:"username"`
	CoverUrl               string      `json:"cover_url"`
	Discord                string      `json:"discord"`
	HasSupported           string      `json:"has_supported"`
	Interests              string      `json:"interests"`
	JoinDate               string      `json:"join_date"`
	Kudosu                 string      `json:"kudosu"`
	Location               string      `json:"location"`
	MaxFriends             string      `json:"max_friends"`
	MaxBLock               string      `json:"max_block"`
	Occupation             string      `json:"occupation"`
	Playmode               string      `json:"playmode"`
	Playstyle              []string    `json:"playstyle"`
	PostCount              string      `json:"post_count"`
	ProfileOrder           []string    `json:"profile_order"`
	Title                  string      `json:"title"`
	TitleUrl               string      `json:"title_url"`
	Twitter                string      `json:"twitter"`
	Website                string      `json:"website"`
	CountyName             string      `json:"country_name"`
	UserCover              CoverString `json:"cover"`
	IsAdmin                string      `json:"is_admin"`
	IsBng                  string      `json:"is_bng"`
	IsFullBan              string      `json:"is_full_bn"`
	IsGmt                  string      `json:"is_gmt"`
	IsLimitedBan           string      `json:"is_limited_bn"`
	IsModerator            string      `json:"is_moderator"`
	IsNat                  string      `json:"is_nat"`
	IsRestricted           string      `json:"is_restricted"`
	IsSilenced             string      `json:"is_silenced"`
	ActiveTournamentBanner string      `json:"active_tournament_banner"`
	Badges                 []Badge     `json:"badges"`
	CommentsCount          string      `json:"comments_count"`
	FollowerCount          string      `json:"follower_count"`
	Groups                 string      `json:"groups"`
	MappingFollowerCount   string      `json:"mapping_follower_count"`
	PendingBeatmapsetCount string      `json:"pending_beatmapset_count"`
	Names                  []string    `json:"previous_usernames"`
	RankedBeatmapsetCount  string      `json:"ranked_beatmapset_count"`
	Level                  string      `json:"level"`
	GlobalRank             string      `json:"global_rank"`
	PP                     string      `json:"pp"`
	RankedScore            string      `json:"ranked_score"`
	Accuracy               string      `json:"accuracy"`
	PlayCount              string      `json:"play_count"`
	PlayTime               string      `json:"play_time"`
	PlayTimeSeconds        string      `json:"play_time_seconds"`
	TotalScore             string      `json:"total_score"`
	TotalHits              string      `json:"total_hits"`
	MaximumCombo           string      `json:"maximum_combo"`
	Replays                string      `json:"replays"`
	IsRanked               string      `json:"is_ranked"`
	SS                     string      `json:"ss"`
	SSH                    string      `json:"ssh"`
	S                      string      `json:"s"`
	SH                     string      `json:"sh"`
	A                      string      `json:"a"`
	CountryRank            string      `json:"country_rank"`
	SupportLvl             string      `json:"support_level"`
}

// ---------------------- Функции парсинга ----------------------

// Функция парсинга ивента
func parseEvent(pageStr string, left int) EventString {

	// Структура ивента
	var ev EventString

	return ev
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
