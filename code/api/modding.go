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

// Мапы
type BeatmapsString struct {
	BeatmapSetId     string `json:"beatmapset_id"`
	DifficultyRating string `json:"difficulty_rating"`
	Id               string `json:"id"`
	Mode             string `json:"mode"`
	Status           string `json:"status"`
	TotalLength      string `json:"total_length"`
	UserId           string `json:"user_id"`
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
	PassCount        string `json:"pass_count"`
	PlayCount        string `json:"play_count"`
	Ranked           string `json:"ranked"`
	Url              string `json:"url"`
	Checksum         string `json:"checksum"`
}

// Структура события
type EventString struct {
	Id         string                  `json:"id"`
	Type       string                  `json:"type"`
	Comment    ModdingCommentString    `json:"comment"`
	CreatedAt  string                  `json:"created_at"`
	UserId     string                  `json:"user_id"`
	Beatmapset ModdingBeatmapsetString `json:"beatmapset"`
	Discussion DiscussionString        `json:"discussion"`
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
	HypeCurrent   string       `json:"hype_current"`
	HypeRequired  string       `json:"hype_required"`
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
	Id            string        `json:"id"`
	IsActive      string        `json:"is_active"`
	IsBot         string        `json:"is_bot"`
	IsDeleted     string        `json:"is_deleted"`
	IsOnline      string        `json:"is_online"`
	IsSupporter   string        `json:"is_supporter"`
	LastVisit     string        `json:"last_visit"`
	PmFriendsOnly string        `json:"pm_friends_only"`
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

// Сет мапы рекорда
type BeatmapsetString struct {
	Artist        string `json:"artist"`
	ArtistUnicode string `json:"artist_unicode"`
	Covers        Covers `json:"covers"`
	Creator       string `json:"creator"`
	FavoriteCount string `json:"favorite_count"`
	Hype          string `json:"hype"`
	Id            string `json:"id"`
	Nsfw          string `json:"nsfw"`
	Offset        string `json:"offset"`
	PlayCount     string `json:"play_count"`
	PreviewUrl    string `json:"preview_url"`
	Source        string `json:"source"`
	Spotlight     string `json:"spotlight"`
	Status        string `json:"status"`
	Title         string `json:"title"`
	TitleUnicode  string `json:"title_unicode"`
	TrackId       string `json:"track_id"`
	UserId        string `json:"userId"`
	Video         string `json:"video"`
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
func parseEvent(pageStr string, left int) (EventString, int) {

	// Структура ивента и буфер
	var ev EventString
	var buffer string

	// Конец ивента
	end := index(pageStr, "}}}", left)

	// Запись данных
	ev.Id, left = findWithIndex(pageStr, "\"id\":", ",", left, end)
	ev.Type, left = findStringWithIndex(pageStr, "\"type\":", ",", left, end)
	ev.Comment.BeatmapDiscussionId, left = findWithIndex(pageStr, "\"beatmap_discussion_id\":", ",", left, end)
	ev.Comment.BeatmapDiscussionPostId, left = findWithIndex(pageStr, "\"beatmap_discussion_post_id\":", ",", left, end)
	ev.Comment.NewVote.UserId, left = findWithIndex(pageStr, "\"user_id\":", ",", left, end)
	ev.Comment.NewVote.Score, left = findWithIndex(pageStr, "\"score\":", "}", left, end)

	buffer, left = findWithIndex(pageStr, "\"votes\":[", "created_at", left, end)

	for i := 0; index(buffer, "user_id\":", i) != -1; {

		var vt VoteString

		vt.UserId, i = findWithIndex(buffer, "\"user_id\":", ",", i, end)
		vt.Score, i = findWithIndex(buffer, "\"score\":", "}", i, end)

		ev.Comment.Votes = append(ev.Comment.Votes, vt)
	}

	ev.CreatedAt, left = findStringWithIndex(pageStr, "\"created_at\":", ",", left, end)
	ev.UserId, left = findWithIndex(pageStr, "\"user_id\":", ",", left, end)
	ev.Beatmapset.Artist, left = findStringWithIndex(pageStr, "\"artist\":", ",", left, end)
	ev.Beatmapset.ArtistUnicode, left = findStringWithIndex(pageStr, "\"artist_unicode\":", ",", left, end)

	ev.Beatmapset.Covers.Cover, left = findStringWithIndex(pageStr, "\"cover\":", ",", left, end)
	ev.Beatmapset.Covers.Cover = strings.ReplaceAll(ev.Beatmapset.Covers.Cover, "\\", "")
	ev.Beatmapset.Covers.Cover2X, left = findStringWithIndex(pageStr, "\"cover@2x\":", ",", left, end)
	ev.Beatmapset.Covers.Cover2X = strings.ReplaceAll(ev.Beatmapset.Covers.Cover2X, "\\", "")
	ev.Beatmapset.Covers.Card, left = findStringWithIndex(pageStr, "\"card\":", ",", left, end)
	ev.Beatmapset.Covers.Card = strings.ReplaceAll(ev.Beatmapset.Covers.Card, "\\", "")
	ev.Beatmapset.Covers.Card2X, left = findStringWithIndex(pageStr, "\"card@2x\":", ",", left, end)
	ev.Beatmapset.Covers.Card2X = strings.ReplaceAll(ev.Beatmapset.Covers.Card2X, "\\", "")
	ev.Beatmapset.Covers.List, left = findStringWithIndex(pageStr, "\"list\":", ",", left, end)
	ev.Beatmapset.Covers.List = strings.ReplaceAll(ev.Beatmapset.Covers.List, "\\", "")
	ev.Beatmapset.Covers.List2X, left = findStringWithIndex(pageStr, "\"list@2x\":", ",", left, end)
	ev.Beatmapset.Covers.List2X = strings.ReplaceAll(ev.Beatmapset.Covers.List2X, "\\", "")
	ev.Beatmapset.Covers.SlimCover, left = findStringWithIndex(pageStr, "\"slimcover\":", ",", left, end)
	ev.Beatmapset.Covers.SlimCover = strings.ReplaceAll(ev.Beatmapset.Covers.SlimCover, "\\", "")
	ev.Beatmapset.Covers.SlimCover2X, left = findStringWithIndex(pageStr, "\"slimcover@2x\":", "}", left, end)
	ev.Beatmapset.Covers.SlimCover2X = strings.ReplaceAll(ev.Beatmapset.Covers.SlimCover2X, "\\", "")

	ev.Beatmapset.Creator, left = findStringWithIndex(pageStr, "\"creator\":", ",", left, end)
	ev.Beatmapset.FavoriteCount, left = findWithIndex(pageStr, "\"favourite_count\":", ",", left, end)
	buffer, left = findWithIndex(pageStr, "\"hype\":", ",\"id", left, end)
	if buffer != "null" {
		ev.Beatmapset.HypeCurrent = find(buffer, "current\":", ",", 0)
		ev.Beatmapset.HypeRequired = find(buffer, "required\":", "}", 0)
	}
	ev.Beatmapset.Id, left = findWithIndex(pageStr, "\"id\":", ",", left, end)
	ev.Beatmapset.Nsfw, left = findWithIndex(pageStr, "\"nsfw\":", ",", left, end)
	ev.Beatmapset.Offset, left = findWithIndex(pageStr, "\"offset\":", ",", left, end)
	ev.Beatmapset.PlayCount, left = findWithIndex(pageStr, "\"play_count\":", ",", left, end)
	ev.Beatmapset.PreviewUrl, left = findStringWithIndex(pageStr, "\"preview_url\":\"\\/\\/", ",", left, end)
	ev.Beatmapset.PreviewUrl = strings.ReplaceAll(ev.Beatmapset.PreviewUrl, "\\", "")
	ev.Beatmapset.Source, left = findWithIndex(pageStr, "\"source\":", ",", left, end)
	ev.Beatmapset.Spotlight, left = findWithIndex(pageStr, "\"spotlight\":", ",", left, end)
	ev.Beatmapset.Status, left = findStringWithIndex(pageStr, "\"status\":", ",", left, end)
	ev.Beatmapset.Title, left = findStringWithIndex(pageStr, "\"title\":", ",", left, end)
	ev.Beatmapset.TitleUnicode, left = findStringWithIndex(pageStr, "\"title_unicode\":", ",", left, end)
	ev.Beatmapset.TrackId, left = findWithIndex(pageStr, "\"track_id\":", ",", left, end)
	ev.Beatmapset.UserId, left = findWithIndex(pageStr, "\"user_id\":", ",", left, end)
	ev.Beatmapset.Video, left = findWithIndex(pageStr, "\"video\":", ",", left, end)

	ev.Beatmapset.User.AvatarUrl, left = findStringWithIndex(pageStr, "\"avatar_url\":", ",", left, end)
	ev.Beatmapset.User.AvatarUrl = strings.ReplaceAll(ev.Beatmapset.User.AvatarUrl, "\\", "")
	ev.Beatmapset.User.CountryCode, left = findStringWithIndex(pageStr, "\"country_code\":", ",", left, end)
	ev.Beatmapset.User.DefaultGroup, left = findStringWithIndex(pageStr, "\"default_group\":", ",", left, end)
	ev.Beatmapset.User.Id, left = findWithIndex(pageStr, "\"id\":", ",", left, end)
	ev.Beatmapset.User.IsActive, left = findWithIndex(pageStr, "\"is_active\":", ",", left, end)
	ev.Beatmapset.User.IsBot, left = findWithIndex(pageStr, "\"is_bot\":", ",", left, end)
	ev.Beatmapset.User.IsDeleted, left = findWithIndex(pageStr, "\"is_deleted\":", ",", left, end)
	ev.Beatmapset.User.IsOnline, left = findWithIndex(pageStr, "\"is_online\":", ",", left, end)
	ev.Beatmapset.User.IsSupporter, left = findWithIndex(pageStr, "\"is_supporter\":", ",", left, end)
	ev.Beatmapset.User.LastVisit, left = findStringWithIndex(pageStr, "\"last_visit\":", ",", left, end)
	ev.Beatmapset.User.PmFriendsOnly, left = findWithIndex(pageStr, "\"pm_friends_only\":", ",", left, end)
	ev.Beatmapset.User.ProfileColor, left = findWithIndex(pageStr, "\"profile_colour\":", ",", left, end)
	ev.Beatmapset.User.Username, left = findStringWithIndex(pageStr, "\"username\":", "}", left, end)

	ev.Discussion.Id, left = findWithIndex(pageStr, "\"id\":", ",", left, end)
	ev.Discussion.BeatmapsetId, left = findWithIndex(pageStr, "\"beatmapset_id\":", ",", left, end)
	ev.Discussion.BeatmapId, left = findWithIndex(pageStr, "\"beatmap_id\":", ",", left, end)
	ev.Discussion.UserId, left = findWithIndex(pageStr, "\"user_id\":", ",", left, end)
	ev.Discussion.DeletedById, left = findWithIndex(pageStr, "\"deleted_by_id\":", ",", left, end)
	ev.Discussion.MessageType, left = findStringWithIndex(pageStr, "\"message_type\":", ",", left, end)
	ev.Discussion.ParentId, left = findWithIndex(pageStr, "\"parent_id\":", ",", left, end)
	ev.Discussion.Timestamp, left = findWithIndex(pageStr, "\"timestamp\":", ",", left, end)
	ev.Discussion.Resolved, left = findWithIndex(pageStr, "\"resolved\":", ",", left, end)
	ev.Discussion.CanBeResolved, left = findWithIndex(pageStr, "\"can_be_resolved\":", ",", left, end)
	ev.Discussion.CanGrantKudosu, left = findWithIndex(pageStr, "\"can_grant_kudosu\":", ",", left, end)
	ev.Discussion.CreatedAt, left = findStringWithIndex(pageStr, "\"created_at\":", ",", left, end)
	ev.Discussion.UpdatedAt, left = findStringWithIndex(pageStr, "\"updated_at\":", ",", left, end)
	ev.Discussion.DeletedAt, left = findStringWithIndex(pageStr, "\"deleted_at\":", ",", left, end)
	ev.Discussion.LastPostAt, left = findStringWithIndex(pageStr, "\"last_post_at\":", ",", left, end)
	ev.Discussion.KudosuDenied, left = findWithIndex(pageStr, "\"kudosu_denied\":", ",", left, end)

	ev.Discussion.StartingPost.BeatmapsetDiscussionId, left = findWithIndex(pageStr, "\"beatmapset_discussion_id\":", ",", left, end)
	ev.Discussion.StartingPost.CreatedAt, left = findStringWithIndex(pageStr, "\"created_at\":", ",", left, end)
	ev.Discussion.StartingPost.DeletedAt, left = findStringWithIndex(pageStr, "\"deleted_at\":", ",", left, end)
	ev.Discussion.StartingPost.DeletedById, left = findWithIndex(pageStr, "\"deleted_by_id\":", ",", left, end)
	ev.Discussion.StartingPost.Id, left = findWithIndex(pageStr, ",\"id\":", ",", left, end)
	ev.Discussion.StartingPost.LastEditorId, left = findWithIndex(pageStr, "\"last_editor_id\":", ",", left, end)
	ev.Discussion.StartingPost.Message, left = findStringWithIndex(pageStr, "\"message\":", ",\"system\":", left, end)
	ev.Discussion.StartingPost.System, left = findWithIndex(pageStr, "\"system\":", ",", left, end)
	ev.Discussion.StartingPost.UpdatedAt, left = findStringWithIndex(pageStr, "\"updated_at\":", ",", left, end)
	ev.Discussion.StartingPost.UserId, left = findWithIndex(pageStr, "\"user_id\":", "}", left, end)

	return ev, left
}

// Функция парсинга ивентов
func parseEvents(pageStr, subStr, stopChar string, left int) ([]EventString, int) {

	var result []EventString
	pageStr, end := findWithIndex(pageStr, subStr, stopChar, left, -1)

	for index(pageStr, "{\"id\":", left) != -1 {

		var ev EventString

		ev, left = parseEvent(pageStr, left)

		result = append(result, ev)

	}

	return result, end
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

	//result.Events, left = parseEvents(pageStr, "<script id=\"json-events\"", "</script>", 0)

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
	jsonResp, err := json.Marshal(map[string]string{
		"status": "Not yet available",
	})
	if err != nil {
		fmt.Print("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}
}
