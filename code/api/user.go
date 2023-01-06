package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ---------------------- Структуры для парсинга ------------------------

// Информация о пользователе
type UserInfoString struct {
	Error                   string              `json:"error"`
	AvatarUrl               string              `json:"avatar_url"`
	CountryCode             string              `json:"country_code"`
	DefaultGroup            string              `json:"default_group"`
	UserID                  string              `json:"id"`
	IsActive                string              `json:"is_active"`
	IsBot                   string              `json:"is_bot"`
	IsDeleted               string              `json:"is_deleted"`
	IsOnline                string              `json:"is_online"`
	IsSupporter             string              `json:"is_supporter"`
	LastVisit               string              `json:"last_visit"`
	PmFriendsOnly           string              `json:"pm_friends_only"`
	ProfileColor            string              `json:"profile_color"`
	Username                string              `json:"username"`
	CoverUrl                string              `json:"cover_url"`
	Discord                 string              `json:"discord"`
	HasSupported            string              `json:"has_supported"`
	Interests               string              `json:"interests"`
	JoinDate                string              `json:"join_date"`
	Kudosu                  string              `json:"kudosu"`
	Location                string              `json:"location"`
	MaxFriends              string              `json:"max_friends"`
	MaxBLock                string              `json:"max_block"`
	Occupation              string              `json:"occupation"`
	Playmode                string              `json:"playmode"`
	Playstyle               []string            `json:"playstyle"`
	PostCount               string              `json:"post_count"`
	ProfileOrder            []string            `json:"profile_order"`
	Title                   string              `json:"title"`
	TitleUrl                string              `json:"title_url"`
	Twitter                 string              `json:"twitter"`
	Website                 string              `json:"website"`
	CountyName              string              `json:"country_name"`
	UserCover               CoverString         `json:"cover"`
	IsAdmin                 string              `json:"is_admin"`
	IsBng                   string              `json:"is_bng"`
	IsFullBan               string              `json:"is_full_bn"`
	IsGmt                   string              `json:"is_gmt"`
	IsLimitedBan            string              `json:"is_limited_bn"`
	IsModerator             string              `json:"is_moderator"`
	IsNat                   string              `json:"is_nat"`
	IsRestricted            string              `json:"is_restricted"`
	IsSilenced              string              `json:"is_silenced"`
	AccountHistory          string              `json:"account_history"`
	ActiveTournamentBanner  string              `json:"active_tournament_banner"`
	Badges                  []Badge             `json:"badges"`
	CommentsCount           string              `json:"comments_count"`
	FollowerCount           string              `json:"follower_count"`
	Groups                  string              `json:"groups"`
	MappingFollowerCount    string              `json:"mapping_follower_count"`
	PendingBeatmapsetCount  string              `json:"pending_beatmapset_count"`
	Names                   []string            `json:"previous_usernames"`
	Level                   string              `json:"level"`
	GlobalRank              string              `json:"global_rank"`
	PP                      string              `json:"pp"`
	RankedScore             string              `json:"ranked_score"`
	Accuracy                string              `json:"accuracy"`
	PlayCount               string              `json:"play_count"`
	PlayTime                string              `json:"play_time"`
	PlayTimeSeconds         string              `json:"play_time_seconds"`
	TotalScore              string              `json:"total_score"`
	TotalHits               string              `json:"total_hits"`
	MaximumCombo            string              `json:"maximum_combo"`
	Replays                 string              `json:"replays"`
	IsRanked                string              `json:"is_ranked"`
	SS                      string              `json:"ss"`
	SSH                     string              `json:"ssh"`
	S                       string              `json:"s"`
	SH                      string              `json:"sh"`
	A                       string              `json:"a"`
	CountryRank             string              `json:"country_rank"`
	SupportLvl              string              `json:"support_level"`
	Achievements            []AchievementString `json:"achievements"`
	Medals                  string              `json:"medals"`
	RankHistory             HistoryString       `json:"rank_history"`
	UnrankedBeatmapsetCount string              `json:"unranked_beatmapset_count"`
	FavoriteBeatmaps        []BeatmapString     `json:"favorite_beatmaps"`
	GraveyardBeatmaps       []BeatmapString     `json:"graveyard_beatmaps"`
	GuestBeatmaps           []BeatmapString     `json:"guest_beatmaps"`
	LovedBeatmaps           []BeatmapString     `json:"loved_beatmaps"`
	RankedBeatmaps          []BeatmapString     `json:"ranked_beatmaps"`
	PendingBeatmaps         []BeatmapString     `json:"pending_beatmaps"`
	KudosuItems             []KudosuString      `json:"kudosu_items"`
	RecentActivity          []ActivityString    `json:"recent_activity"`
	Best                    []ScoreString       `json:"best"`
	Firsts                  []ScoreString       `json:"firsts"`
	Pinned                  []ScoreString       `json:"pinned"`
	BeatmapPlaycounts       []PlayCountString   `json:"beatmap_playcounts"`
	MonthlyPlaycounts       []CountString       `json:"monthly_playcounts"`
	ReplaysWatchedCount     []CountString       `json:"replays_watched_counts"`
}

// Ковёр пользователя
type CoverString struct {
	CustomUrl string `json:"custom_url"`
	Url       string `json:"url"`
	Id        string `json:"id"`
}

// Значок профиля
type Badge struct {
	AwardedAt   string `json:"awarded_at"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

// Достижение
type AchievementString struct {
	AchievedAt    string `json:"achieved_at"`
	AchievementId string `json:"achievement_id"`
}

// История рейтинга
type HistoryString struct {
	Mode string   `json:"mode"`
	Data []string `json:"data"`
}

// Структура карты
type BeatmapString struct {
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
	Beatmap           BeatmapsString           `json:"beatmap"`
}

// Картинки
type Covers struct {
	Cover       string `json:"cover"`
	Cover2X     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2X      string `json:"card@2x"`
	List        string `json:"list"`
	List2X      string `json:"list@2x"`
	SlimCover   string `json:"slimcover"`
	SlimCover2X string `json:"slimcover@2x"`
}

// Оценка номинаций
type NominationsSummaryString struct {
	Current  string `json:"current"`
	Required string `json:"required"`
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

// Кудосу
type KudosuString struct {
	Id        string      `json:"id"`
	Action    string      `json:"action"`
	Amount    string      `json:"amount"`
	Model     string      `json:"model"`
	CreatedAt string      `json:"created_at"`
	Giver     KudosuGiver `json:"giver"`
	Post      KudosuPost  `json:"post"`
	Details   string      `json:"details"`
}

// Источник кудосу
type KudosuGiver struct {
	Url      string `json:"url"`
	Username string `json:"username"`
}

// Пост кудосу
type KudosuPost struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

// Активность
type ActivityString struct {
	CreatedAt    string `json:"created_at"`
	Id           string `json:"id"`
	Type         string `json:"type"`
	ScoreRank    string `json:"score_rank"`
	Rank         string `json:"rank"`
	Mode         string `json:"mode"`
	BeatmapTitle string `json:"beatmap_title"`
	BeatmapUrl   string `json:"beatmap_url"`
}

// Структура для подсчёта
type CountString struct {
	StartDate string `json:"start_date"`
	Count     string `json:"count"`
}

// Рекорд
type ScoreString struct {
	Accuracy              string           `json:"accuracy"`
	BeatmapId             string           `json:"beatmap_id"`
	BuildId               string           `json:"build_id"`
	EndedAt               string           `json:"ended_at"`
	LegacyScoreId         string           `json:"legacy_score_id"`
	LegacyTotalScore      string           `json:"legacy_total_score"`
	MaximumCombo          string           `json:"max_combo"`
	MaximumStatistics     StatisticsString `json:"maximum_statistics"`
	Mods                  []string         `json:"mods"`
	Passed                string           `json:"passed"`
	Rank                  string           `json:"rank"`
	RulesetId             string           `json:"ruleset_id"`
	StartedAt             string           `json:"started_at"`
	Statistics            StatisticsString `json:"statistics"`
	TotalScore            string           `json:"total_score"`
	UserId                string           `json:"user_id"`
	BestId                string           `json:"best_id"`
	Id                    string           `json:"id"`
	LegacyPerfect         string           `json:"legacy_perfect"`
	PP                    string           `json:"pp"`
	Replay                string           `json:"replay"`
	Type                  string           `json:"type"`
	CurrentUserAttributes string           `json:"current_user_attributes"`
	Beatmap               BeatmapsString   `json:"beatmap"`
	Beatmapset            BeatmapsetString `json:"beatmapset"`
	Weight                WeightString     `json:"weight"`
}

// Статистика рекорда
type StatisticsString struct {
	Good    string `json:"good"`
	Great   string `json:"great"`
	Meh     string `json:"meh"`
	Miss    string `json:"miss"`
	Ok      string `json:"ok"`
	Perfect string `json:"perfect"`
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

// Статистика рекорда
type WeightString struct {
	Percentage string `json:"percentage"`
	PP         string `json:"pp"`
}

// Структура для подсчёта количества игр
type PlayCountString struct {
	BeatmapId  string                 `json:"beatmap_id"`
	Count      string                 `json:"count"`
	Beatmap    PlayCountBeatmapString `json:"beatmap"`
	Beatmapset BeatmapsetString       `json:"beatmapset"`
}

// Сет мапы подсчёта
type PlayCountBeatmapString struct {
	BeatmapsetId     string `json:"beatmapset_id"`
	DifficultyRating string `json:"difficulty_rating"`
	Id               string `json:"id"`
	Status           string `json:"status"`
	TotalLength      string `json:"total_length"`
	UserId           string `json:"user_id"`
	Version          string `json:"version"`
}

// ---------------------- Классические структуры ------------------------

// Информация о пользователе
type UserInfo struct {
	Error                   string        `json:"error"`
	AvatarUrl               string        `json:"avatar_url"`
	CountryCode             string        `json:"country_code"`
	DefaultGroup            string        `json:"default_group"`
	UserID                  int           `json:"id"`
	IsActive                bool          `json:"is_active"`
	IsBot                   bool          `json:"is_bot"`
	IsDeleted               bool          `json:"is_deleted"`
	IsOnline                bool          `json:"is_online"`
	IsSupporter             bool          `json:"is_supporter"`
	LastVisit               string        `json:"last_visit"`
	PmFriendsOnly           bool          `json:"pm_friends_only"`
	ProfileColor            string        `json:"profile_color"`
	Username                string        `json:"username"`
	CoverUrl                string        `json:"cover_url"`
	Discord                 string        `json:"discord"`
	HasSupported            bool          `json:"has_supported"`
	Interests               string        `json:"interests"`
	JoinDate                string        `json:"join_date"`
	Kudosu                  int           `json:"kudosu"`
	Location                string        `json:"location"`
	MaxFriends              int           `json:"max_friends"`
	MaxBLock                int           `json:"max_block"`
	Occupation              string        `json:"occupation"`
	Playmode                string        `json:"playmode"`
	Playstyle               []string      `json:"playstyle"`
	PostCount               int           `json:"post_count"`
	ProfileOrder            []string      `json:"profile_order"`
	Title                   string        `json:"title"`
	TitleUrl                string        `json:"title_url"`
	Twitter                 string        `json:"twitter"`
	Website                 string        `json:"website"`
	CountyName              string        `json:"country_name"`
	UserCover               Cover         `json:"cover"`
	IsAdmin                 bool          `json:"is_admin"`
	IsBng                   bool          `json:"is_bng"`
	IsFullBan               bool          `json:"is_full_bn"`
	IsGmt                   bool          `json:"is_gmt"`
	IsLimitedBan            bool          `json:"is_limited_bn"`
	IsModerator             bool          `json:"is_moderator"`
	IsNat                   bool          `json:"is_nat"`
	IsRestricted            bool          `json:"is_restricted"`
	IsSilenced              bool          `json:"is_silenced"`
	AccountHistory          string        `json:"account_history"`
	ActiveTournamentBanner  string        `json:"active_tournament_banner"`
	Badges                  []Badge       `json:"badges"`
	CommentsCount           int           `json:"comments_count"`
	FollowerCount           int           `json:"follower_count"`
	Groups                  string        `json:"groups"`
	MappingFollowerCount    int           `json:"mapping_follower_count"`
	PendingBeatmapsetCount  int           `json:"pending_beatmapset_count"`
	Names                   []string      `json:"previous_usernames"`
	Level                   int           `json:"level"`
	GlobalRank              int64         `json:"global_rank"`
	PP                      float64       `json:"pp"`
	RankedScore             int           `json:"ranked_score"`
	Accuracy                float64       `json:"accuracy"`
	PlayCount               int           `json:"play_count"`
	PlayTime                string        `json:"play_time"`
	PlayTimeSeconds         int64         `json:"play_time_seconds"`
	TotalScore              int64         `json:"total_score"`
	TotalHits               int64         `json:"total_hits"`
	MaximumCombo            int           `json:"maximum_combo"`
	Replays                 int           `json:"replays"`
	IsRanked                bool          `json:"is_ranked"`
	SS                      int           `json:"ss"`
	SSH                     int           `json:"ssh"`
	S                       int           `json:"s"`
	SH                      int           `json:"sh"`
	A                       int           `json:"a"`
	CountryRank             int           `json:"country_rank"`
	SupportLvl              int           `json:"support_level"`
	Achievements            []Achievement `json:"achievements"`
	Medals                  int           `json:"medals"`
	RankHistory             History       `json:"rank_history"`
	UnrankedBeatmapsetCount int           `json:"unranked_beatmapset_count"`
	FavoriteBeatmaps        []Beatmap     `json:"favorite_beatmaps"`
	GraveyardBeatmaps       []Beatmap     `json:"graveyard_beatmaps"`
	GuestBeatmaps           []Beatmap     `json:"guest_beatmaps"`
	LovedBeatmaps           []Beatmap     `json:"loved_beatmaps"`
	RankedBeatmaps          []Beatmap     `json:"ranked_beatmaps"`
	PendingBeatmaps         []Beatmap     `json:"pending_beatmaps"`
	KudosuItems             []Kudosu      `json:"kudosu_items"`
	RecentActivity          []Activity    `json:"recent_activity"`
	Best                    []Score       `json:"best"`
	Firsts                  []Score       `json:"firsts"`
	Pinned                  []Score       `json:"pinned"`
	BeatmapPlaycounts       []PlayCount   `json:"beatmap_playcounts"`
	MonthlyPlaycounts       []Count       `json:"monthly_playcounts"`
	ReplaysWatchedCount     []Count       `json:"replays_watched_counts"`
}

// Ковёр пользователя
type Cover struct {
	CustomUrl string `json:"custom_url"`
	Url       string `json:"url"`
	Id        int    `json:"id"`
}

// Достижение
type Achievement struct {
	AchievedAt    string `json:"achieved_at"`
	AchievementId int    `json:"achievement_id"`
}

// История рейтинга
type History struct {
	Mode string `json:"mode"`
	Data []int  `json:"data"`
}

// Структура карты
type Beatmap struct {
	Artist            string             `json:"artist"`
	ArtistUnicode     string             `json:"artist_unicode"`
	Covers            Covers             `json:"covers"`
	Creator           string             `json:"creator"`
	FavoriteCount     int                `json:"favorite_count"`
	Hype              string             `json:"hype"`
	Id                int                `json:"id"`
	Nsfw              bool               `json:"nsfw"`
	Offset            int                `json:"offset"`
	PlayCount         int                `json:"play_count"`
	PreviewUrl        string             `json:"preview_url"`
	Source            string             `json:"source"`
	Spotlight         bool               `json:"spotlight"`
	Status            string             `json:"status"`
	Title             string             `json:"title"`
	TitleUnicode      string             `json:"title_unicode"`
	TrackId           string             `json:"track_id"`
	UserId            int                `json:"userId"`
	Video             bool               `json:"video"`
	DownloadDisabled  bool               `json:"download_disabled"`
	Bpm               float64            `json:"bpm"`
	CanBeHyped        bool               `json:"can_be_hyped"`
	DiscussionEnabled bool               `json:"discussion_enabled"`
	DiscussionLocked  bool               `json:"discussion_locked"`
	IsScoreable       bool               `json:"is_scoreable"`
	LastUpdated       string             `json:"last_updated"`
	LegacyThreadUrl   string             `json:"legacy_thread_url"`
	Nominations       NominationsSummary `json:"nominations_summary"`
	Ranked            int                `json:"ranked"`
	RankedDate        string             `json:"ranked_date"`
	Storyboard        bool               `json:"storyboard"`
	SubmittedDate     string             `json:"submitted_date"`
	Tags              []string           `json:"tags"`
	Beatmap           Beatmaps           `json:"beatmap"`
}

// Оценка номинаций
type NominationsSummary struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}

// Мапы
type Beatmaps struct {
	BeatmapSetId     int     `json:"beatmapset_id"`
	DifficultyRating float64 `json:"difficulty_rating"`
	Id               int     `json:"id"`
	Mode             string  `json:"mode"`
	Status           string  `json:"status"`
	TotalLength      int     `json:"total_length"`
	UserId           int     `json:"user_id"`
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
	PassCount        int     `json:"pass_count"`
	PlayCount        int     `json:"play_count"`
	Ranked           int     `json:"ranked"`
	Url              string  `json:"url"`
	Checksum         string  `json:"checksum"`
}

// Кудосу
type Kudosu struct {
	Id        int         `json:"id"`
	Action    string      `json:"action"`
	Amount    int         `json:"amount"`
	Model     string      `json:"model"`
	CreatedAt string      `json:"created_at"`
	Giver     KudosuGiver `json:"giver"`
	Post      KudosuPost  `json:"post"`
	Details   string      `json:"details"`
}

// Активность
type Activity struct {
	CreatedAt    string `json:"created_at"`
	Id           int    `json:"id"`
	Type         string `json:"type"`
	ScoreRank    string `json:"score_rank"`
	Rank         int    `json:"rank"`
	Mode         string `json:"mode"`
	BeatmapTitle string `json:"beatmap_title"`
	BeatmapUrl   string `json:"beatmap_url"`
}

// Структура для подсчёта
type Count struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}

// Рекорд
type Score struct {
	Accuracy              float64    `json:"accuracy"`
	BeatmapId             int        `json:"beatmap_id"`
	BuildId               string     `json:"build_id"`
	EndedAt               string     `json:"ended_at"`
	LegacyScoreId         string     `json:"legacy_score_id"`
	LegacyTotalScore      string     `json:"legacy_total_score"`
	MaximumCombo          int        `json:"max_combo"`
	MaximumStatistics     Statistics `json:"maximum_statistics"`
	Mods                  []string   `json:"mods"`
	Passed                bool       `json:"passed"`
	Rank                  string     `json:"rank"`
	RulesetId             int        `json:"ruleset_id"`
	StartedAt             string     `json:"started_at"`
	Statistics            Statistics `json:"statistics"`
	TotalScore            int        `json:"total_score"`
	UserId                int        `json:"user_id"`
	BestId                int        `json:"best_id"`
	Id                    int        `json:"id"`
	LegacyPerfect         bool       `json:"legacy_perfect"`
	PP                    float64    `json:"pp"`
	Replay                bool       `json:"replay"`
	Type                  string     `json:"type"`
	CurrentUserAttributes string     `json:"current_user_attributes"`
	Beatmap               Beatmaps   `json:"beatmap"`
	Beatmapset            Beatmapset `json:"beatmapset"`
	Weight                Weight     `json:"weight"`
}

// Статистика рекорда
type Statistics struct {
	Good    int `json:"good"`
	Great   int `json:"great"`
	Meh     int `json:"meh"`
	Miss    int `json:"miss"`
	Ok      int `json:"ok"`
	Perfect int `json:"perfect"`
}

// Сет мапы рекорда
type Beatmapset struct {
	Artist        string `json:"artist"`
	ArtistUnicode string `json:"artist_unicode"`
	Covers        Covers `json:"covers"`
	Creator       string `json:"creator"`
	FavoriteCount int    `json:"favorite_count"`
	Hype          string `json:"hype"`
	Id            int    `json:"id"`
	Nsfw          bool   `json:"nsfw"`
	Offset        int    `json:"offset"`
	PlayCount     int    `json:"play_count"`
	PreviewUrl    string `json:"preview_url"`
	Source        string `json:"source"`
	Spotlight     bool   `json:"spotlight"`
	Status        string `json:"status"`
	Title         string `json:"title"`
	TitleUnicode  string `json:"title_unicode"`
	TrackId       string `json:"track_id"`
	UserId        int    `json:"userId"`
	Video         bool   `json:"video"`
}

// Статистика рекорда
type Weight struct {
	Percentage float64 `json:"percentage"`
	PP         float64 `json:"pp"`
}

// Структура для подсчёта количества игр
type PlayCount struct {
	BeatmapId  int              `json:"beatmap_id"`
	Count      int              `json:"count"`
	Beatmap    PlayCountBeatmap `json:"beatmap"`
	Beatmapset Beatmapset       `json:"beatmapset"`
}

// Сет мапы подсчёта
type PlayCountBeatmap struct {
	BeatmapsetId     int     `json:"beatmapset_id"`
	DifficultyRating float64 `json:"difficulty_rating"`
	Id               int     `json:"id"`
	Status           string  `json:"status"`
	TotalLength      int     `json:"total_length"`
	UserId           int     `json:"user_id"`
	Version          string  `json:"version"`
}

// ---------------------- Функции поиска ------------------------

// Функция поиска. Возвращает искомое значение и индекс последнего символа
func findWithIndex(str, subStr, stopChar string, start int) (string, int) {

	// Обрезка левой границы поиска
	str = str[start:]

	// Поиск индекса начала нужной строки
	left := strings.Index(str, subStr) + len(subStr)

	// Проверка на существование нужной строки
	if left != len(subStr)-1 {

		// Поиск правой границы
		right := left + strings.Index(str[left:], stopChar)

		// Обрезка и вывод результата
		return str[left:right], right + start
	}

	// Вывод ненайденных значений для тестов
	// fmt.Println("error foundn't \t", subStr, "-")

	return "", start
}

// Функция поиска. Возвращает искомое значение без кавычек и индекс последнего символа
func findStringWithIndex(str, subStr, stopChar string, start int) (string, int) {

	// Обрезка левой границы поиска
	str = str[start:]

	// Поиск индекса начала нужной строки
	left := strings.Index(str, subStr) + len(subStr)

	// Проверка на существование нужной строки
	if left != len(subStr)-1 {

		// Поиск правой границы
		right := left + strings.Index(str[left:], stopChar)

		// Обрезка и вывод результата
		return strings.ReplaceAll(str[left:right], "\"", ""), right + start
	}

	// Вывод ненайденных значений для тестов
	// fmt.Println("error foundn't \t", subStr, "-")

	return "", start
}

// Облегчённая функция поиска. Возвращает только искомое значение
func find(str, subStr, stopChar string, start int) string {

	str = str[start:]
	left := strings.Index(str, subStr)

	// Проверка на существование нужной строки
	if left != -1 {

		// Обрезка левой части
		str = str[left+len(subStr):]

		// Обрезка правой части и вывод результата
		return str[:strings.Index(str, stopChar)]
	}

	return ""
}

// Функция поиска индекса
func index(str, subStr string, start int) int {

	res := strings.Index(str[start:], subStr)

	// Проверка на существование нужной строки
	if res == -1 {

		//fmt.Println("index error: \t", subStr)

		return -1
	}

	//fmt.Println(res+start, " - ", subStr)
	return res + start
}

// Функция проверки наличия подстроки
func contains(str, subStr string, left int) bool {

	return strings.Contains(str[left:], subStr)
}

// ---------------------- Функции парсинга ----------------------

// Функция парсинга карты
func parseBeatmapsString(pageStr string, left int) ([]BeatmapString, int) {

	// Получение рабочей части и индекса её конца
	pageStr, end := findWithIndex(pageStr, "items :[", "], pagination", left)

	// Проверка на наличие карт
	if len(pageStr) == 0 {
		return nil, end
	}

	// Результат и индекс обработанной части
	var result []BeatmapString
	left = 0

	// Пока есть необработанные карты
	for index(pageStr, "{ artist", left) != -1 {

		// Инициализация карты
		var bm BeatmapString

		// Запись данных
		bm.Artist, left = findWithIndex(pageStr, "artist : ", " , artist_", left)
		bm.ArtistUnicode, left = findWithIndex(pageStr, "artist_unicode : ", " ,", left)

		bm.Covers.Cover, left = findWithIndex(pageStr, "cover : ", " , cover", left)
		bm.Covers.Cover = strings.ReplaceAll(bm.Covers.Cover, "\\", "")
		bm.Covers.Cover2X, left = findWithIndex(pageStr, "cover@2x : ", " ,", left)
		bm.Covers.Cover2X = strings.ReplaceAll(bm.Covers.Cover2X, "\\", "")
		bm.Covers.Card, left = findWithIndex(pageStr, "card : ", " , card@2x", left)
		bm.Covers.Card = strings.ReplaceAll(bm.Covers.Card, "\\", "")
		bm.Covers.Card2X, left = findWithIndex(pageStr, "card@2x : ", " ,", left)
		bm.Covers.Card2X = strings.ReplaceAll(bm.Covers.Card2X, "\\", "")
		bm.Covers.List, left = findWithIndex(pageStr, "list : ", " ,", left)
		bm.Covers.List = strings.ReplaceAll(bm.Covers.List, "\\", "")
		bm.Covers.List2X, left = findWithIndex(pageStr, "list@2x : ", " ,", left)
		bm.Covers.List2X = strings.ReplaceAll(bm.Covers.List2X, "\\", "")
		bm.Covers.SlimCover, left = findWithIndex(pageStr, "slimcover : ", " , slimcover", left)
		bm.Covers.SlimCover = strings.ReplaceAll(bm.Covers.SlimCover, "\\", "")
		bm.Covers.SlimCover2X, left = findWithIndex(pageStr, "slimcover@2x : ", " }", left)
		bm.Covers.SlimCover2X = strings.ReplaceAll(bm.Covers.SlimCover2X, "\\", "")

		bm.Creator, left = findWithIndex(pageStr, "creator : ", " ", left)
		bm.FavoriteCount, left = findWithIndex(pageStr, "favourite_count :", ",", left)
		bm.Hype, left = findWithIndex(pageStr, "hype :", ",", left)
		bm.Id, left = findWithIndex(pageStr, "id :", ",", left)
		bm.Nsfw, left = findWithIndex(pageStr, "nsfw :", ",", left)
		bm.Offset, left = findWithIndex(pageStr, "offset :", ",", left)
		bm.PlayCount, left = findWithIndex(pageStr, "play_count :", ",", left)
		bm.PreviewUrl, left = findWithIndex(pageStr, "preview_url : \\/\\/", " , source", left)
		bm.PreviewUrl = strings.ReplaceAll(bm.PreviewUrl, "\\", "")
		bm.Source, left = findWithIndex(pageStr, "source :", " ", left)
		bm.Spotlight, left = findWithIndex(pageStr, "spotlight :", ",", left)
		bm.Status, left = findWithIndex(pageStr, "status : ", " ,", left)
		bm.Title, left = findWithIndex(pageStr, "title : ", " , title_unicode", left)
		bm.TitleUnicode, left = findWithIndex(pageStr, "title_unicode : ", " ,", left)
		bm.TrackId, left = findWithIndex(pageStr, "track_id :", ",", left)
		bm.UserId, left = findWithIndex(pageStr, "user_id :", ",", left)
		bm.Video, left = findWithIndex(pageStr, "video :", ",", left)
		bm.DownloadDisabled, left = findWithIndex(pageStr, "download_disabled :", ",", left)

		bm.Bpm, left = findWithIndex(pageStr, "bpm :", ",", left)
		bm.CanBeHyped, left = findWithIndex(pageStr, "can_be_hyped :", ",", left)
		bm.DiscussionEnabled, left = findWithIndex(pageStr, "discussion_enabled :", ",", left)
		bm.DiscussionLocked, left = findWithIndex(pageStr, "discussion_locked :", ",", left)
		bm.IsScoreable, left = findWithIndex(pageStr, "is_scoreable :", ",", left)
		bm.LastUpdated, left = findWithIndex(pageStr, "last_updated : ", " ", left)
		bm.LegacyThreadUrl, left = findWithIndex(pageStr, "legacy_thread_url : ", " ", left)
		bm.LegacyThreadUrl = strings.ReplaceAll(bm.LegacyThreadUrl, "\\", "")

		bm.Nominations.Current, left = findWithIndex(pageStr, "current :", ",", left)
		bm.Nominations.Required, left = findWithIndex(pageStr, "required :", "}", left)

		bm.Ranked, left = findWithIndex(pageStr, "ranked :", ",", left)
		bm.RankedDate, left = findWithIndex(pageStr, "ranked_date : ", " ", left)
		bm.Storyboard, left = findWithIndex(pageStr, "storyboard :", ",", left)
		bm.SubmittedDate, left = findWithIndex(pageStr, "submitted_date : ", " ", left)
		bm.Tags = strings.Split(find(pageStr, "tags : ", " ,", left), " ")
		if bm.Tags[0] == "" {
			bm.Tags = nil
		}

		bm.Beatmap.BeatmapSetId, left = findWithIndex(pageStr, "beatmapset_id :", ",", left)
		bm.Beatmap.DifficultyRating, left = findWithIndex(pageStr, "difficulty_rating :", ",", left)
		bm.Beatmap.Id, left = findWithIndex(pageStr, "id :", ",", left)
		bm.Beatmap.Mode, left = findWithIndex(pageStr, "mode : ", " ", left)
		bm.Beatmap.Status, left = findWithIndex(pageStr, "status : ", " ", left)
		bm.Beatmap.TotalLength, left = findWithIndex(pageStr, "total_length :", ",", left)
		bm.Beatmap.UserId, left = findWithIndex(pageStr, " user_id :", ",", left)
		bm.Beatmap.Version, left = findWithIndex(pageStr, "version : ", " , accuracy", left)
		bm.Beatmap.Accuracy, left = findWithIndex(pageStr, "accuracy :", ",", left)
		bm.Beatmap.Ar, left = findWithIndex(pageStr, "ar :", ",", left)
		bm.Beatmap.Convert, left = findWithIndex(pageStr, "convert :", ",", left)
		bm.Beatmap.CountCircles, left = findWithIndex(pageStr, "count_circles :", ",", left)
		bm.Beatmap.CountSliders, left = findWithIndex(pageStr, "count_sliders :", ",", left)
		bm.Beatmap.CountSpinners, left = findWithIndex(pageStr, "count_spinners :", ",", left)
		bm.Beatmap.Cs, left = findWithIndex(pageStr, " cs :", ",", left)
		bm.Beatmap.DeletedAt, left = findWithIndex(pageStr, "deleted_at :", ",", left)
		bm.Beatmap.Drain, left = findWithIndex(pageStr, "drain :", ",", left)
		bm.Beatmap.HitLength, left = findWithIndex(pageStr, "hit_length :", ",", left)
		bm.Beatmap.IsScoreable, left = findWithIndex(pageStr, "is_scoreable :", ",", left)
		bm.Beatmap.LastUpdated, left = findWithIndex(pageStr, "last_updated : ", " ", left)
		bm.Beatmap.ModeInt, left = findWithIndex(pageStr, "mode_int :", ",", left)
		bm.Beatmap.PassCount, left = findWithIndex(pageStr, "passcount :", ",", left)
		bm.Beatmap.PlayCount, left = findWithIndex(pageStr, "playcount :", ",", left)
		bm.Beatmap.Ranked, left = findWithIndex(pageStr, "ranked :", ",", left)
		bm.Beatmap.Url, left = findWithIndex(pageStr, "url : ", " ", left)
		bm.Beatmap.Url = strings.ReplaceAll(bm.Beatmap.Url, "\\", "")
		bm.Beatmap.Checksum, left = findWithIndex(pageStr, "checksum : ", " ", left)

		// Добавление карты к результату
		result = append(result, bm)

	}

	return result, end
}

// Функция парсинга рекорда
func parseScoresString(pageStr, scoreType string, left int) ([]ScoreString, int) {

	// Получение рабочей части и индекса её конца
	pageStr, end := findWithIndex(pageStr, "items :[", "], pagination", left)

	// Проверка на наличие рекордов
	if len(pageStr) == 0 {
		return nil, end
	}

	// Результат и индекс обработанной части
	var result []ScoreString
	left = 0

	// Пока есть необработанные карты
	for index(pageStr, "{ accuracy :", left) != -1 {

		// Инициализация карты
		var sc ScoreString

		// Запись данных
		sc.Accuracy, left = findWithIndex(pageStr, "accuracy :", ",", left)
		sc.BeatmapId, left = findWithIndex(pageStr, "beatmap_id :", ",", left)
		sc.BuildId, left = findWithIndex(pageStr, "build_id :", ",", left)
		sc.EndedAt, left = findWithIndex(pageStr, "ended_at : ", " ,", left)
		sc.LegacyScoreId, left = findWithIndex(pageStr, "legacy_score_id :", ",", left)
		sc.LegacyTotalScore, left = findWithIndex(pageStr, "legacy_total_score :", ",", left)
		sc.MaximumCombo, left = findWithIndex(pageStr, "max_combo :", ",", left)

		// Обработка максимальной статистики рекорда
		statisticsString, left := findWithIndex(pageStr, "maximum_statistics :{", "}", left)
		statisticsString += ","

		sc.MaximumStatistics = StatisticsString{
			Good:    find(statisticsString, "good :", ",", 0),
			Great:   find(statisticsString, "great :", ",", 0),
			Meh:     find(statisticsString, "meh :", ",", 0),
			Ok:      find(statisticsString, "ok :", ",", 0),
			Miss:    find(statisticsString, "miss :", ",", 0),
			Perfect: find(statisticsString, "perfect :", ",", 0),
		}

		// Цикл для обработки модов
		for c := 0; pageStr[c] != ']'; c++ {
			if pageStr[c:c+10] == "acronym : " {
				sc.Mods = append(sc.Mods, pageStr[c+10:c+12])
			}
		}

		sc.Passed, left = findWithIndex(pageStr, "passed :", ",", left)
		sc.Rank, left = findWithIndex(pageStr, "rank : ", " ", left)
		sc.RulesetId, left = findWithIndex(pageStr, "ruleset_id :", ",", left)
		sc.StartedAt, left = findWithIndex(pageStr, "started_at :", ",", left)

		// Обработка статистики рекорда
		statisticsString, left = findWithIndex(pageStr, "statistics :{", "}", left)
		statisticsString += ","

		sc.Statistics = StatisticsString{
			Great: find(statisticsString, "great :", ",", 0),
			Meh:   find(statisticsString, "meh :", ",", 0),
			Ok:    find(statisticsString, "ok :", ",", 0),
			Miss:  find(statisticsString, "miss :", ",", 0),
		}

		sc.TotalScore, left = findWithIndex(pageStr, "total_score :", ",", left)
		sc.UserId, left = findWithIndex(pageStr, " user_id :", ",", left)
		sc.BestId, left = findWithIndex(pageStr, " best_id :", ",", left)
		sc.Id, left = findWithIndex(pageStr, " id :", ",", left)
		sc.LegacyPerfect, left = findWithIndex(pageStr, "legacy_perfect :", ",", left)
		sc.PP, left = findWithIndex(pageStr, "pp :", ",", left)
		sc.Replay, left = findWithIndex(pageStr, "replay :", ",", left)
		sc.Type, left = findWithIndex(pageStr, "type : ", " ", left)
		sc.CurrentUserAttributes, left = findWithIndex(pageStr, "current_user_attributes :{ ", "},", left)

		sc.Beatmap.BeatmapSetId, left = findWithIndex(pageStr, "beatmapset_id :", ",", left)
		sc.Beatmap.DifficultyRating, left = findWithIndex(pageStr, "difficulty_rating :", ",", left)
		sc.Beatmap.Id, left = findWithIndex(pageStr, "id :", ",", left)
		sc.Beatmap.Mode, left = findWithIndex(pageStr, "mode : ", " ", left)
		sc.Beatmap.Status, left = findWithIndex(pageStr, "status : ", " ", left)
		sc.Beatmap.TotalLength, left = findWithIndex(pageStr, "total_length :", ",", left)
		sc.Beatmap.UserId, left = findWithIndex(pageStr, " user_id :", ",", left)
		sc.Beatmap.Version, left = findWithIndex(pageStr, "version : ", " , accuracy", left)
		sc.Beatmap.Accuracy, left = findWithIndex(pageStr, "accuracy :", ",", left)
		sc.Beatmap.Ar, left = findWithIndex(pageStr, "ar :", ",", left)
		sc.Beatmap.Bpm, left = findWithIndex(pageStr, "bpm :", ",", left)
		sc.Beatmap.Convert, left = findWithIndex(pageStr, "convert :", ",", left)
		sc.Beatmap.CountCircles, left = findWithIndex(pageStr, "count_circles :", ",", left)
		sc.Beatmap.CountSliders, left = findWithIndex(pageStr, "count_sliders :", ",", left)
		sc.Beatmap.CountSpinners, left = findWithIndex(pageStr, "count_spinners :", ",", left)
		sc.Beatmap.Cs, left = findWithIndex(pageStr, " cs :", ",", left)
		sc.Beatmap.DeletedAt, left = findWithIndex(pageStr, "deleted_at :", ",", left)
		sc.Beatmap.Drain, left = findWithIndex(pageStr, "drain :", ",", left)
		sc.Beatmap.HitLength, left = findWithIndex(pageStr, "hit_length :", ",", left)
		sc.Beatmap.IsScoreable, left = findWithIndex(pageStr, "is_scoreable :", ",", left)
		sc.Beatmap.LastUpdated, left = findWithIndex(pageStr, "last_updated : ", " ", left)
		sc.Beatmap.ModeInt, left = findWithIndex(pageStr, "mode_int :", ",", left)
		sc.Beatmap.PassCount, left = findWithIndex(pageStr, "passcount :", ",", left)
		sc.Beatmap.PlayCount, left = findWithIndex(pageStr, "playcount :", ",", left)
		sc.Beatmap.Ranked, left = findWithIndex(pageStr, "ranked :", ",", left)
		sc.Beatmap.Url, left = findWithIndex(pageStr, "url : ", " ", left)
		sc.Beatmap.Url = strings.ReplaceAll(sc.Beatmap.Url, "\\", "")
		sc.Beatmap.Checksum, left = findWithIndex(pageStr, "checksum : ", " ", left)

		sc.Beatmapset.Artist, left = findWithIndex(pageStr, "artist : ", " , artist_", left)
		sc.Beatmapset.ArtistUnicode, left = findWithIndex(pageStr, "artist_unicode : ", " ,", left)

		sc.Beatmapset.Covers.Cover, left = findWithIndex(pageStr, "cover : ", " , cover", left)
		sc.Beatmapset.Covers.Cover = strings.ReplaceAll(sc.Beatmapset.Covers.Cover, "\\", "")
		sc.Beatmapset.Covers.Cover2X, left = findWithIndex(pageStr, "cover@2x : ", " ,", left)
		sc.Beatmapset.Covers.Cover2X = strings.ReplaceAll(sc.Beatmapset.Covers.Cover2X, "\\", "")
		sc.Beatmapset.Covers.Card, left = findWithIndex(pageStr, "card : ", " , card@2x", left)
		sc.Beatmapset.Covers.Card = strings.ReplaceAll(sc.Beatmapset.Covers.Card, "\\", "")
		sc.Beatmapset.Covers.Card2X, left = findWithIndex(pageStr, "card@2x : ", " ,", left)
		sc.Beatmapset.Covers.Card2X = strings.ReplaceAll(sc.Beatmapset.Covers.Card2X, "\\", "")
		sc.Beatmapset.Covers.List, left = findWithIndex(pageStr, "list : ", " ,", left)
		sc.Beatmapset.Covers.List = strings.ReplaceAll(sc.Beatmapset.Covers.List, "\\", "")
		sc.Beatmapset.Covers.List2X, left = findWithIndex(pageStr, "list@2x : ", " ,", left)
		sc.Beatmapset.Covers.List2X = strings.ReplaceAll(sc.Beatmapset.Covers.List2X, "\\", "")
		sc.Beatmapset.Covers.SlimCover, left = findWithIndex(pageStr, "slimcover : ", " , slimcover", left)
		sc.Beatmapset.Covers.SlimCover = strings.ReplaceAll(sc.Beatmapset.Covers.SlimCover, "\\", "")
		sc.Beatmapset.Covers.SlimCover2X, left = findWithIndex(pageStr, "slimcover@2x : ", " }", left)
		sc.Beatmapset.Covers.SlimCover2X = strings.ReplaceAll(sc.Beatmapset.Covers.SlimCover2X, "\\", "")

		sc.Beatmapset.Creator, left = findWithIndex(pageStr, "creator : ", " ", left)
		sc.Beatmapset.FavoriteCount, left = findWithIndex(pageStr, "favourite_count :", ",", left)
		sc.Beatmapset.Hype, left = findWithIndex(pageStr, "hype :", ",", left)
		sc.Beatmapset.Id, left = findWithIndex(pageStr, "id :", ",", left)
		sc.Beatmapset.Nsfw, left = findWithIndex(pageStr, "nsfw :", ",", left)
		sc.Beatmapset.Offset, left = findWithIndex(pageStr, "offset :", ",", left)
		sc.Beatmapset.PlayCount, left = findWithIndex(pageStr, "play_count :", ",", left)
		sc.Beatmapset.PreviewUrl, left = findWithIndex(pageStr, "preview_url : \\/\\/", " , source", left)
		sc.Beatmapset.PreviewUrl = strings.ReplaceAll(sc.Beatmapset.PreviewUrl, "\\", "")
		sc.Beatmapset.Source, left = findWithIndex(pageStr, "source :", " ", left)
		sc.Beatmapset.Spotlight, left = findWithIndex(pageStr, "spotlight :", ",", left)
		sc.Beatmapset.Status, left = findWithIndex(pageStr, "status : ", " ,", left)
		sc.Beatmapset.Title, left = findWithIndex(pageStr, "title : ", " , title_unicode", left)
		sc.Beatmapset.TitleUnicode, left = findWithIndex(pageStr, "title_unicode : ", " ,", left)
		sc.Beatmapset.TrackId, left = findWithIndex(pageStr, "track_id :", ",", left)
		sc.Beatmapset.UserId, left = findWithIndex(pageStr, "user_id :", ",", left)
		sc.Beatmapset.Video, left = findWithIndex(pageStr, "video :", "}", left)

		if scoreType == "best" {
			sc.Weight.Percentage, left = findWithIndex(pageStr, "percentage :", ",", left)
			sc.Weight.PP, _ = findWithIndex(pageStr, "pp :", "}", left)
		}

		// Добавление карты к результату
		result = append(result, sc)

	}

	return result, end
}

// ---------------------- Функции перевода ----------------------

func toInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("parsing error: \t", s)
		return 0
	}

	return i
}

func toInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		fmt.Println("parsing error: \t", s)
		return 0
	}

	return i
}

func toBool(s string) bool {
	f, err := strconv.ParseBool(s)

	if err != nil {
		fmt.Println("parsing error: \t", s)
		return false
	}

	return f
}

func toFloat64(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)

	if err != nil {
		fmt.Println("parsing error: \t", s)
		return 0
	}

	return i
}

func toSlice(s string) []int {

	var result []int
	sliceStr := strings.Split(s, ",")

	if len(sliceStr) == 1 && sliceStr[0] == "" {
		return nil
	}

	for _, digit := range sliceStr {
		result = append(result, toInt(digit))
	}

	return result
}

// Функция перевода карты
func formatBeatmaps(bms []BeatmapString) []Beatmap {

	var result []Beatmap

	// Обработка текстовых карт
	for _, bm := range bms {

		// Форматирование и добавление карты
		result = append(result, Beatmap{
			Artist:            bm.Artist,
			ArtistUnicode:     bm.ArtistUnicode,
			Covers:            bm.Covers,
			Creator:           bm.Creator,
			FavoriteCount:     toInt(bm.FavoriteCount),
			Hype:              bm.Hype,
			Id:                toInt(bm.Id),
			Nsfw:              toBool(bm.Nsfw),
			Offset:            toInt(bm.Offset),
			PlayCount:         toInt(bm.PlayCount),
			PreviewUrl:        bm.PreviewUrl,
			Source:            bm.Source,
			Spotlight:         toBool(bm.Spotlight),
			Status:            bm.Status,
			Title:             bm.Title,
			TitleUnicode:      bm.TitleUnicode,
			TrackId:           bm.TrackId,
			UserId:            toInt(bm.UserId),
			Video:             toBool(bm.Video),
			DownloadDisabled:  toBool(bm.DownloadDisabled),
			Bpm:               toFloat64(bm.Bpm),
			CanBeHyped:        toBool(bm.CanBeHyped),
			DiscussionEnabled: toBool(bm.DiscussionEnabled),
			DiscussionLocked:  toBool(bm.DiscussionLocked),
			IsScoreable:       toBool(bm.IsScoreable),
			LastUpdated:       bm.LastUpdated,
			LegacyThreadUrl:   bm.LegacyThreadUrl,
			Nominations: NominationsSummary{
				Current:  toInt(bm.Nominations.Current),
				Required: toInt(bm.Nominations.Required),
			},
			Ranked:        toInt(bm.Ranked),
			RankedDate:    bm.RankedDate,
			Storyboard:    toBool(bm.Storyboard),
			SubmittedDate: bm.SubmittedDate,
			Tags:          bm.Tags,
			Beatmap: Beatmaps{
				BeatmapSetId:     toInt(bm.Beatmap.BeatmapSetId),
				DifficultyRating: toFloat64(bm.Beatmap.DifficultyRating),
				Id:               toInt(bm.Beatmap.Id),
				Mode:             bm.Beatmap.Mode,
				Status:           bm.Beatmap.Status,
				TotalLength:      toInt(bm.Beatmap.TotalLength),
				UserId:           toInt(bm.Beatmap.UserId),
				Version:          bm.Beatmap.Version,
				Accuracy:         toFloat64(bm.Beatmap.Accuracy),
				Ar:               toFloat64(bm.Beatmap.Ar),
				Bpm:              toFloat64(bm.Beatmap.Bpm),
				Convert:          toBool(bm.Beatmap.Convert),
				CountCircles:     toInt(bm.Beatmap.CountCircles),
				CountSliders:     toInt(bm.Beatmap.CountSliders),
				CountSpinners:    toInt(bm.Beatmap.CountSpinners),
				Cs:               toFloat64(bm.Beatmap.Cs),
				DeletedAt:        bm.Beatmap.DeletedAt,
				Drain:            toFloat64(bm.Beatmap.Drain),
				HitLength:        toInt(bm.Beatmap.HitLength),
				IsScoreable:      toBool(bm.Beatmap.IsScoreable),
				LastUpdated:      bm.Beatmap.LastUpdated,
				ModeInt:          toInt(bm.Beatmap.ModeInt),
				PassCount:        toInt(bm.Beatmap.PassCount),
				PlayCount:        toInt(bm.Beatmap.PlayCount),
				Ranked:           toInt(bm.Beatmap.Ranked),
				Url:              bm.Beatmap.Url,
				Checksum:         bm.Beatmap.Checksum,
			},
		})

	}

	return result
}

// Функция перевода рекорда
func formatScores(scs []ScoreString) []Score {

	var result []Score

	// Обработка текстовых рекордов
	for _, sc := range scs {

		// Форматирование и добавление рекорда
		result = append(result, Score{
			Accuracy:         toFloat64(sc.Accuracy),
			BeatmapId:        toInt(sc.BeatmapId),
			BuildId:          sc.BuildId,
			EndedAt:          sc.EndedAt,
			LegacyScoreId:    sc.LegacyScoreId,
			LegacyTotalScore: sc.LegacyTotalScore,
			MaximumCombo:     toInt(sc.MaximumCombo),
			MaximumStatistics: Statistics{
				Good:    toInt(sc.MaximumStatistics.Good),
				Great:   toInt(sc.MaximumStatistics.Great),
				Meh:     toInt(sc.MaximumStatistics.Meh),
				Miss:    toInt(sc.MaximumStatistics.Miss),
				Ok:      toInt(sc.MaximumStatistics.Ok),
				Perfect: toInt(sc.MaximumStatistics.Perfect),
			},
			Mods:                  sc.Mods,
			Passed:                toBool(sc.Passed),
			Rank:                  sc.Rank,
			RulesetId:             toInt(sc.RulesetId),
			StartedAt:             sc.StartedAt,
			TotalScore:            toInt(sc.TotalScore),
			UserId:                toInt(sc.UserId),
			BestId:                toInt(sc.BestId),
			Id:                    toInt(sc.Id),
			LegacyPerfect:         toBool(sc.LegacyPerfect),
			PP:                    toFloat64(sc.PP),
			Replay:                toBool(sc.Replay),
			Type:                  sc.Type,
			CurrentUserAttributes: sc.CurrentUserAttributes,
			Beatmap: Beatmaps{
				BeatmapSetId:     toInt(sc.Beatmap.BeatmapSetId),
				DifficultyRating: toFloat64(sc.Beatmap.DifficultyRating),
				Id:               toInt(sc.Beatmap.Id),
				Mode:             sc.Beatmap.Mode,
				Status:           sc.Beatmap.Status,
				TotalLength:      toInt(sc.Beatmap.TotalLength),
				UserId:           toInt(sc.Beatmap.UserId),
				Version:          sc.Beatmap.Version,
				Accuracy:         toFloat64(sc.Beatmap.Accuracy),
				Ar:               toFloat64(sc.Beatmap.Ar),
				Bpm:              toFloat64(sc.Beatmap.Bpm),
				Convert:          toBool(sc.Beatmap.Convert),
				CountCircles:     toInt(sc.Beatmap.CountCircles),
				CountSliders:     toInt(sc.Beatmap.CountSliders),
				CountSpinners:    toInt(sc.Beatmap.CountSpinners),
				Cs:               toFloat64(sc.Beatmap.Cs),
				DeletedAt:        sc.Beatmap.DeletedAt,
				Drain:            toFloat64(sc.Beatmap.Drain),
				HitLength:        toInt(sc.Beatmap.HitLength),
				IsScoreable:      toBool(sc.Beatmap.IsScoreable),
				LastUpdated:      sc.Beatmap.LastUpdated,
				ModeInt:          toInt(sc.Beatmap.ModeInt),
				PassCount:        toInt(sc.Beatmap.PassCount),
				PlayCount:        toInt(sc.Beatmap.PlayCount),
				Ranked:           toInt(sc.Beatmap.Ranked),
				Url:              sc.Beatmap.Url,
				Checksum:         sc.Beatmap.Checksum,
			},
			Beatmapset: Beatmapset{
				Artist:        sc.Beatmapset.Artist,
				ArtistUnicode: sc.Beatmapset.ArtistUnicode,
				Covers:        sc.Beatmapset.Covers,
				Creator:       sc.Beatmapset.Creator,
				FavoriteCount: toInt(sc.Beatmapset.FavoriteCount),
				Hype:          sc.Beatmapset.Hype,
				Id:            toInt(sc.Beatmapset.Id),
				Nsfw:          toBool(sc.Beatmapset.Nsfw),
				Offset:        toInt(sc.Beatmapset.Offset),
				PlayCount:     toInt(sc.Beatmapset.PlayCount),
				PreviewUrl:    sc.Beatmapset.PreviewUrl,
				Source:        sc.Beatmapset.Source,
				Spotlight:     toBool(sc.Beatmapset.Spotlight),
				Status:        sc.Beatmapset.Status,
				Title:         sc.Beatmapset.Title,
				TitleUnicode:  sc.Beatmapset.TitleUnicode,
				TrackId:       sc.Beatmapset.TrackId,
				UserId:        toInt(sc.Beatmapset.UserId),
				Video:         toBool(sc.Beatmapset.Video),
			},
			Weight: Weight{
				Percentage: toFloat64(sc.Weight.Percentage),
				PP:         toFloat64(sc.Weight.PP),
			},
		})

	}

	return result
}

// ----------------- Функции получения статистики ----------------

// Функция получения информации о пользователе
func GetUserInfo(id string) UserInfo {

	// Получение текстовой версии статистики
	resultStr := GetUserInfoString(id)

	// Проверка на ошибки при парсинге
	if resultStr.Error != "" {
		return UserInfo{
			Error: resultStr.Error,
		}
	}

	// Перевод в классическую версию
	result := UserInfo{
		Error:         resultStr.Error,
		AvatarUrl:     resultStr.AvatarUrl,
		CountryCode:   resultStr.CountryCode,
		DefaultGroup:  resultStr.DefaultGroup,
		UserID:        toInt(resultStr.UserID),
		IsActive:      toBool(resultStr.IsActive),
		IsBot:         toBool(resultStr.IsBot),
		IsDeleted:     toBool(resultStr.IsDeleted),
		IsOnline:      toBool(resultStr.IsOnline),
		IsSupporter:   toBool(resultStr.IsSupporter),
		LastVisit:     resultStr.LastVisit,
		PmFriendsOnly: toBool(resultStr.PmFriendsOnly),
		ProfileColor:  resultStr.ProfileColor,
		Username:      resultStr.Username,
		CoverUrl:      resultStr.CoverUrl,
		Discord:       resultStr.Discord,
		HasSupported:  toBool(resultStr.HasSupported),
		Interests:     resultStr.Interests,
		JoinDate:      resultStr.JoinDate,
		Kudosu:        toInt(resultStr.Kudosu),
		Location:      resultStr.Location,
		MaxFriends:    toInt(resultStr.MaxFriends),
		MaxBLock:      toInt(resultStr.MaxBLock),
		Occupation:    resultStr.Occupation,
		Playmode:      resultStr.Playmode,
		Playstyle:     resultStr.Playstyle,
		PostCount:     toInt(resultStr.PostCount),
		ProfileOrder:  resultStr.ProfileOrder,
		Title:         resultStr.Title,
		TitleUrl:      resultStr.TitleUrl,
		Twitter:       resultStr.Twitter,
		Website:       resultStr.Website,
		CountyName:    resultStr.CountyName,
		UserCover: Cover{
			CustomUrl: resultStr.UserCover.CustomUrl,
			Url:       resultStr.UserCover.Url,
			Id:        toInt(resultStr.UserCover.Id),
		},
		IsAdmin:                 toBool(resultStr.IsAdmin),
		IsBng:                   toBool(resultStr.IsBng),
		IsFullBan:               toBool(resultStr.IsFullBan),
		IsGmt:                   toBool(resultStr.IsGmt),
		IsLimitedBan:            toBool(resultStr.IsLimitedBan),
		IsModerator:             toBool(resultStr.IsModerator),
		IsNat:                   toBool(resultStr.IsNat),
		IsRestricted:            toBool(resultStr.IsRestricted),
		IsSilenced:              toBool(resultStr.IsSilenced),
		AccountHistory:          resultStr.AccountHistory,
		ActiveTournamentBanner:  resultStr.ActiveTournamentBanner,
		Badges:                  resultStr.Badges,
		CommentsCount:           toInt(resultStr.CommentsCount),
		FollowerCount:           toInt(resultStr.FollowerCount),
		Groups:                  resultStr.Groups,
		MappingFollowerCount:    toInt(resultStr.MappingFollowerCount),
		PendingBeatmapsetCount:  toInt(resultStr.PendingBeatmapsetCount),
		Names:                   resultStr.Names,
		Level:                   toInt(resultStr.Level),
		GlobalRank:              toInt64(resultStr.GlobalRank),
		PP:                      toFloat64(resultStr.PP),
		RankedScore:             toInt(resultStr.RankedScore),
		Accuracy:                toFloat64(resultStr.Accuracy),
		PlayCount:               toInt(resultStr.PlayCount),
		PlayTime:                resultStr.PlayTime,
		PlayTimeSeconds:         toInt64(resultStr.PlayTimeSeconds),
		TotalScore:              toInt64(resultStr.TotalScore),
		TotalHits:               toInt64(resultStr.TotalHits),
		MaximumCombo:            toInt(resultStr.MaximumCombo),
		Replays:                 toInt(resultStr.Replays),
		IsRanked:                toBool(resultStr.IsRanked),
		SS:                      toInt(resultStr.SS),
		SSH:                     toInt(resultStr.SSH),
		S:                       toInt(resultStr.S),
		SH:                      toInt(resultStr.SH),
		A:                       toInt(resultStr.A),
		CountryRank:             toInt(resultStr.CountryRank),
		SupportLvl:              toInt(resultStr.SupportLvl),
		UnrankedBeatmapsetCount: toInt(resultStr.UnrankedBeatmapsetCount),
	}

	// Перевод достижений
	for _, c := range resultStr.Achievements {
		result.Achievements = append(result.Achievements, Achievement{
			AchievedAt:    c.AchievedAt,
			AchievementId: toInt(c.AchievementId),
		})
	}

	// Перевод количества медалей и истории рейтинга
	result.Medals = toInt(resultStr.Medals)
	result.RankHistory.Mode = resultStr.RankHistory.Mode

	for _, d := range resultStr.RankHistory.Data {
		result.RankHistory.Data = append(result.RankHistory.Data, toInt(d))
	}

	// Перевод карт
	result.FavoriteBeatmaps = formatBeatmaps(resultStr.FavoriteBeatmaps)
	result.GraveyardBeatmaps = formatBeatmaps(resultStr.GraveyardBeatmaps)
	result.GuestBeatmaps = formatBeatmaps(resultStr.GuestBeatmaps)
	result.LovedBeatmaps = formatBeatmaps(resultStr.LovedBeatmaps)
	result.RankedBeatmaps = formatBeatmaps(resultStr.RankedBeatmaps)
	result.PendingBeatmaps = formatBeatmaps(resultStr.PendingBeatmaps)

	// Перевод кудосу
	for _, k := range resultStr.KudosuItems {
		result.KudosuItems = append(result.KudosuItems, Kudosu{
			Id:        toInt(k.Id),
			Action:    k.Action,
			Amount:    toInt(k.Amount),
			Model:     k.Model,
			CreatedAt: k.CreatedAt,
			Giver:     k.Giver,
			Post:      k.Post,
			Details:   k.Details,
		})
	}

	// Перевод последней активности
	for _, a := range resultStr.RecentActivity {
		result.RecentActivity = append(result.RecentActivity, Activity{
			CreatedAt:    a.CreatedAt,
			Id:           toInt(a.Id),
			Type:         a.Type,
			ScoreRank:    a.ScoreRank,
			Rank:         toInt(a.Rank),
			Mode:         a.Mode,
			BeatmapTitle: a.BeatmapTitle,
			BeatmapUrl:   a.BeatmapUrl,
		})
	}

	// Перевод рекордов
	result.Best = formatScores(resultStr.Best)
	result.Firsts = formatScores(resultStr.Firsts)
	result.Pinned = formatScores(resultStr.Pinned)

	// Перевод карт с количеством игр
	for _, pc := range resultStr.BeatmapPlaycounts {
		result.BeatmapPlaycounts = append(result.BeatmapPlaycounts, PlayCount{
			BeatmapId: toInt(pc.BeatmapId),
			Count:     toInt(pc.Count),
			Beatmap: PlayCountBeatmap{
				BeatmapsetId:     toInt(pc.Beatmap.BeatmapsetId),
				DifficultyRating: toFloat64(pc.Beatmap.DifficultyRating),
				Id:               toInt(pc.Beatmap.Id),
				Status:           pc.Beatmap.Status,
				TotalLength:      toInt(pc.Beatmap.TotalLength),
				UserId:           toInt(pc.Beatmap.UserId),
				Version:          pc.Beatmap.Version,
			},
			Beatmapset: Beatmapset{
				Artist:        pc.Beatmapset.Artist,
				ArtistUnicode: pc.Beatmapset.ArtistUnicode,
				Covers:        pc.Beatmapset.Covers,
				Creator:       pc.Beatmapset.Creator,
				FavoriteCount: toInt(pc.Beatmapset.FavoriteCount),
				Hype:          pc.Beatmapset.Hype,
				Id:            toInt(pc.Beatmapset.Id),
				Nsfw:          toBool(pc.Beatmapset.Nsfw),
				Offset:        toInt(pc.Beatmapset.Offset),
				PlayCount:     toInt(pc.Beatmapset.PlayCount),
				PreviewUrl:    pc.Beatmapset.PreviewUrl,
				Source:        pc.Beatmapset.Source,
				Spotlight:     toBool(pc.Beatmapset.Spotlight),
				Status:        pc.Beatmapset.Status,
				Title:         pc.Beatmapset.Title,
				TitleUnicode:  pc.Beatmapset.TitleUnicode,
				TrackId:       pc.Beatmapset.TrackId,
				UserId:        toInt(pc.Beatmapset.UserId),
				Video:         toBool(pc.Beatmapset.Video),
			},
		})
	}

	for _, c := range resultStr.MonthlyPlaycounts {
		result.MonthlyPlaycounts = append(result.MonthlyPlaycounts, Count{
			StartDate: c.StartDate,
			Count:     toInt(c.Count),
		})
	}

	// Перевод статистики просмотров повторов
	for _, c := range resultStr.ReplaysWatchedCount {
		result.ReplaysWatchedCount = append(result.ReplaysWatchedCount, Count{
			StartDate: c.StartDate,
			Count:     toInt(c.Count),
		})
	}

	return result
}

// Функция получения текстовой информации о пользователе
func GetUserInfoString(id string) UserInfoString {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return UserInfoString{
			Error: "http get error",
		}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// HTML полученной страницы в формате string
	pageStr := string(body)[80000:]

	// Проверка на страницу пользователя
	if !strings.Contains(pageStr, "js-react--profile") {
		return UserInfoString{
			Error: "user not found",
		}
	}

	// Обрезка юзелесс части html"ки
	pageStr = strings.ReplaceAll(pageStr[strings.Index(pageStr, "current_mode"):], "&quot;", " ")

	// Сохранение html"ки в файл sample.html (для тестов)
	/*
		if err := os.WriteFile("sample.html", []byte(pageStr), 0666); err != nil {
			log.Fatal(err)
		}
	*/

	// Структура, которую будет возвращать функция
	result := UserInfoString{}

	// Крайняя левая граница поиска
	left := 0

	//--------------------------- Статистика игрока ------------------------------

	result.AvatarUrl, left = findWithIndex(pageStr, "avatar_url : ", " ", left)
	result.AvatarUrl = strings.ReplaceAll(result.AvatarUrl, "\\", "")
	result.CountryCode, left = findWithIndex(pageStr, "country_code : ", " ", left)
	result.DefaultGroup, left = findWithIndex(pageStr, "default_group : ", " ", left)
	result.UserID, left = findWithIndex(pageStr, " id :", ",", left)
	result.IsActive, left = findWithIndex(pageStr, "is_active :", ",", left)
	result.IsBot, left = findWithIndex(pageStr, "is_bot :", ",", left)
	result.IsDeleted, left = findWithIndex(pageStr, "is_deleted :", ",", left)
	result.IsOnline, left = findWithIndex(pageStr, "is_online :", ",", left)
	result.IsSupporter, left = findWithIndex(pageStr, "is_supporter :", ",", left)
	result.LastVisit, left = findWithIndex(pageStr, "last_visit : ", " ", left)
	result.PmFriendsOnly, left = findWithIndex(pageStr, "pm_friends_only :", ",", left)
	result.ProfileColor, left = findWithIndex(pageStr, "profile_colour : ", " ,", left)
	result.Username, left = findWithIndex(pageStr, "username : ", " ", left)
	result.CoverUrl, left = findWithIndex(pageStr, "cover_url : ", " ", left)
	result.CoverUrl = strings.ReplaceAll(result.CoverUrl, "\\", "")
	result.Discord, left = findWithIndex(pageStr, "discord : ", " ,", left)
	result.HasSupported, left = findWithIndex(pageStr, "has_supported :", ",", left)
	result.Interests, left = findWithIndex(pageStr, "interests : ", " , join_date", left)
	result.JoinDate, left = findWithIndex(pageStr, "join_date : ", " ,", left)
	result.Kudosu, left = findWithIndex(pageStr, "kudosu :{ total :", ",", left)
	result.Location, left = findWithIndex(pageStr, "location : ", " ,", left)
	result.MaxBLock, left = findWithIndex(pageStr, "max_blocks :", ",", left)
	result.MaxFriends, left = findWithIndex(pageStr, "max_friends :", ",", left)
	result.Occupation, left = findWithIndex(pageStr, "occupation : ", " ,", left)
	result.Playmode, left = findWithIndex(pageStr, "playmode : ", " ,", left)
	result.Playstyle = strings.Split(find(pageStr, "playstyle :[ ", " ],", left), " , ")
	if result.Playstyle[0] == "" {
		result.Playstyle = nil
	}
	result.PostCount, left = findWithIndex(pageStr, "post_count :", ",", left)
	result.ProfileOrder = strings.Split(find(pageStr, "profile_order :[ ", " ]", left), " , ")
	if result.ProfileOrder[0] == "" {
		result.ProfileOrder = nil
	}
	result.Title, left = findWithIndex(pageStr, "title :", ",", left)
	result.TitleUrl, left = findWithIndex(pageStr, "title_url : ", " ,", left)
	result.Twitter, left = findWithIndex(pageStr, "twitter : ", " ,", left)
	result.Website, left = findWithIndex(pageStr, "website : ", " ,", left)
	result.Website = strings.ReplaceAll(result.Website, "\\", "")
	result.CountyName, left = findWithIndex(pageStr, " name : ", " }", left)

	result.UserCover.CustomUrl, left = findWithIndex(pageStr, "custom_url : ", " ,", left)
	result.UserCover.CustomUrl = strings.ReplaceAll(result.UserCover.CustomUrl, "\\", "")
	result.UserCover.Url, left = findWithIndex(pageStr, "url : ", " ,", left)
	result.UserCover.Url = strings.ReplaceAll(result.UserCover.Url, "\\", "")
	result.UserCover.Id, left = findWithIndex(pageStr, " , id : ", " }", left)

	result.IsAdmin, left = findWithIndex(pageStr, "is_admin :", ",", left)
	result.IsBng, left = findWithIndex(pageStr, "is_bng :", ",", left)
	result.IsFullBan, left = findWithIndex(pageStr, "is_full_bn :", ",", left)
	result.IsGmt, left = findWithIndex(pageStr, "is_gmt :", ",", left)
	result.IsLimitedBan, left = findWithIndex(pageStr, "is_limited_bn :", ",", left)
	result.IsModerator, left = findWithIndex(pageStr, "is_moderator :", ",", left)
	result.IsNat, left = findWithIndex(pageStr, "is_nat :", ",", left)
	result.IsRestricted, left = findWithIndex(pageStr, "is_restricted :", ",", left)
	result.IsSilenced, left = findWithIndex(pageStr, "is_silenced :", ",", left)
	result.ActiveTournamentBanner, left = findWithIndex(pageStr, "active_tournament_banner :", ", badges", left)
	result.ActiveTournamentBanner = strings.ReplaceAll(result.ActiveTournamentBanner, "\\", "")

	// Значки
	for c := index(pageStr, "badges :[", left); pageStr[c] != ']'; c++ {
		if pageStr[c:c+13] == "awarded_at : " {
			result.Badges = append(result.Badges, Badge{
				AwardedAt:   find(pageStr[c:], "awarded_at : ", " ", 0),
				Description: find(pageStr[c:], "description : ", " ,", 0),
				ImageUrl:    strings.ReplaceAll(find(pageStr[c:], "image_url : ", " ", 0), "\\", ""),
			})
		}
	}

	result.CommentsCount, left = findWithIndex(pageStr, "comments_count :", ",", left)
	result.FollowerCount, left = findWithIndex(pageStr, "follower_count :", ",", left)

	// Принадлежность к группам
	for c := index(pageStr, "groups :[", left); pageStr[c] != ']'; c++ {
		if pageStr[c] == '{' {
			result.Groups += find(pageStr[c:], "name : ", " ,", 0) + ", "
		}
	}
	if result.Groups != "" {
		result.Groups = result.Groups[:len(result.Groups)-2]
	}

	result.MappingFollowerCount, left = findWithIndex(pageStr, "mapping_follower_count :", ",", left)
	result.PendingBeatmapsetCount, left = findWithIndex(pageStr, "pending_beatmapset_count :", ",", left)
	result.Names = strings.Split(find(pageStr, "previous_usernames :[ ", " ],", left), " , ")
	result.Level, left = findWithIndex(pageStr, "level :{ current :", ",", left)
	result.GlobalRank, left = findWithIndex(pageStr, "global_rank :", ",", left)
	result.PP, left = findWithIndex(pageStr, "pp :", ",", left)
	result.RankedScore, left = findWithIndex(pageStr, "ranked_score :", ",", left)
	result.Accuracy, left = findWithIndex(pageStr, "hit_accuracy :", ",", left)
	result.PlayCount, left = findWithIndex(pageStr, "play_count :", ",", left)
	result.PlayTimeSeconds, left = findWithIndex(pageStr, "play_time :", ",", left)
	duration, _ := time.ParseDuration(result.PlayTimeSeconds + "s")
	result.PlayTime = duration.String()
	result.TotalScore, left = findWithIndex(pageStr, "total_score :", ",", left)
	result.TotalHits, left = findWithIndex(pageStr, "total_hits :", ",", left)
	result.MaximumCombo, left = findWithIndex(pageStr, "maximum_combo :", ",", left)
	result.Replays, left = findWithIndex(pageStr, "replays_watched_by_others :", ",", left)
	result.IsRanked, left = findWithIndex(pageStr, "is_ranked :", ",", left)
	result.SS, left = findWithIndex(pageStr, "grade_counts :{ ss :", ",", left)
	result.SSH, left = findWithIndex(pageStr, "ssh :", ",", left)
	result.S, left = findWithIndex(pageStr, "s :", ",", left)
	result.SH, left = findWithIndex(pageStr, "sh :", ",", left)
	result.A, left = findWithIndex(pageStr, "a :", "}", left)
	result.CountryRank, left = findWithIndex(pageStr, "country_rank :", ",", left)
	result.SupportLvl, left = findWithIndex(pageStr, "support_level :", ",", left)

	// Проверка на наличие достижений
	if !contains(pageStr, "user_achievements :[]", left) {

		// Конец блока достижений
		end := index(pageStr, "]", left) - 10
		medals := 0

		// Цикл обработки достижений
		for ; left < end; medals++ {

			// Инициализация достижения
			var achieve AchievementString

			// Генерация достижения
			achieve.AchievedAt, left = findWithIndex(pageStr, "achieved_at : ", " ,", left)
			achieve.AchievementId, left = findWithIndex(pageStr, "achievement_id :", "}", left)

			// Добавление достижения
			result.Achievements = append(result.Achievements, achieve)

		}

		// Запись количества медалей
		result.Medals = strconv.Itoa(medals)

	}

	// Проверка на наличие статистики
	if !contains(pageStr, " rank_history :null", left) {
		result.RankHistory.Mode, left = findWithIndex(pageStr, "mode : ", " ,", left)
		result.RankHistory.Data = strings.Split(find(pageStr, "data :[", "]", left), ",")
	}

	result.UnrankedBeatmapsetCount, left = findWithIndex(pageStr, "unranked_beatmapset_count :", "}", left)

	// Обрезка левой части и обнуление левого индекса
	pageStr = pageStr[left:]
	left = 0

	// Карты
	result.FavoriteBeatmaps, left = parseBeatmapsString(pageStr, left)
	result.GraveyardBeatmaps, left = parseBeatmapsString(pageStr, left)
	result.GuestBeatmaps, left = parseBeatmapsString(pageStr, left)
	result.LovedBeatmaps, left = parseBeatmapsString(pageStr, left)
	result.RankedBeatmaps, left = parseBeatmapsString(pageStr, left)
	result.PendingBeatmaps, left = parseBeatmapsString(pageStr, left)

	// Проверка на наличие статистики
	if !contains(pageStr, "kudosu :{ items :[]", left) {

		// Пока есть необработанные данные
		for index(pageStr, "giver :{", left) != -1 {

			// Инициализация кудосу
			var kudosu KudosuString

			// Запись данных
			kudosu.Id, left = findWithIndex(pageStr, "id :", ",", left)
			kudosu.Action, left = findWithIndex(pageStr, "action : ", " ,", left)
			kudosu.Amount, left = findWithIndex(pageStr, "amount :", ",", left)
			kudosu.Model, left = findWithIndex(pageStr, "model : ", " ", left)
			kudosu.CreatedAt, left = findWithIndex(pageStr, "created_at : ", " ", left)

			kudosu.Giver.Url, left = findWithIndex(pageStr, "url : ", " ", left)
			kudosu.Giver.Url = strings.ReplaceAll(kudosu.Giver.Url, "\\", "")
			kudosu.Giver.Username, left = findWithIndex(pageStr, "username : ", " },", left)

			kudosu.Post.Url, left = findWithIndex(pageStr, "url : ", " ", left)
			kudosu.Post.Url = strings.ReplaceAll(kudosu.Post.Url, "\\", "")
			kudosu.Post.Title, left = findWithIndex(pageStr, "title : ", " },", left)

			kudosu.Details, left = findWithIndex(pageStr, "details :", "}", left)

			// Добавление данных к результату
			result.KudosuItems = append(result.KudosuItems, kudosu)

		}

	} else {

		// Смещение указателя на конец
		left = index(pageStr, "items :[]", left) + 1

	}

	// Проверка на наличие активности
	if !contains(pageStr, "recent_activity :{ items :[]", left) {

		// Пока есть необработанная активность
		for index(pageStr, "scoreRank", left) != -1 {

			// Инициализация активности
			var act ActivityString

			// Запись данных
			act.CreatedAt, left = findWithIndex(pageStr, "created_at : ", " ", left)
			act.Id, left = findWithIndex(pageStr, "id :", ",", left)
			act.Type, left = findWithIndex(pageStr, "type : ", " ", left)
			act.ScoreRank, left = findWithIndex(pageStr, "scoreRank : ", " ", left)
			act.Rank, left = findWithIndex(pageStr, "rank :", ",", left)
			act.Mode, left = findWithIndex(pageStr, "mode : ", " ", left)
			act.BeatmapTitle, left = findWithIndex(pageStr, "title : ", " , url", left)
			act.BeatmapUrl, left = findWithIndex(pageStr, "url : ", " }", left)

			// Добавление статистики
			result.RecentActivity = append(result.RecentActivity, act)

		}

	} else {

		// Смещение указателя на конец
		left = index(pageStr, "items :[]", left) + 1

	}

	// Обрезка левой части и обнуление левого индекса
	pageStr = pageStr[left:]
	left = 0

	result.Best, left = parseScoresString(pageStr, "best", left)
	result.Firsts, left = parseScoresString(pageStr, "first", left)
	result.Pinned, left = parseScoresString(pageStr, "pinned", left)

	// Проверка на наличие статистики
	if !contains(pageStr, "beatmap_playcounts :{ items :[]", left) {

		// Пока есть необработанная активность
		for index(pageStr, "{ beatmap_id", left) != -1 {

			// Инициализация активности
			var pc PlayCountString

			// Запись данных
			pc.BeatmapId, left = findWithIndex(pageStr, "beatmap_id :", ",", left)
			pc.Count, left = findWithIndex(pageStr, "count :", ",", left)

			pc.Beatmap.BeatmapsetId, left = findWithIndex(pageStr, "beatmapset_id :", ",", left)
			pc.Beatmap.DifficultyRating, left = findWithIndex(pageStr, "difficulty_rating :", ",", left)
			pc.Beatmap.Id, left = findWithIndex(pageStr, "id :", ",", left)
			pc.Beatmap.Status, left = findWithIndex(pageStr, "status : ", " ,", left)
			pc.Beatmap.TotalLength, left = findWithIndex(pageStr, "total_length :", ",", left)
			pc.Beatmap.UserId, left = findWithIndex(pageStr, "user_id :", ",", left)
			pc.Beatmap.Version, left = findWithIndex(pageStr, "version : ", " },", left)

			pc.Beatmapset.Artist, left = findWithIndex(pageStr, "artist : ", " , artist_", left)
			pc.Beatmapset.ArtistUnicode, left = findWithIndex(pageStr, "artist_unicode : ", " , ", left)

			pc.Beatmapset.Covers.Cover, left = findWithIndex(pageStr, "cover : ", " ", left)
			pc.Beatmapset.Covers.Cover = strings.ReplaceAll(pc.Beatmapset.Covers.Cover, "\\", "")
			pc.Beatmapset.Covers.Cover2X, left = findWithIndex(pageStr, "cover@2x : ", " ", left)
			pc.Beatmapset.Covers.Cover2X = strings.ReplaceAll(pc.Beatmapset.Covers.Cover2X, "\\", "")
			pc.Beatmapset.Covers.Card, left = findWithIndex(pageStr, "card : ", " ", left)
			pc.Beatmapset.Covers.Card = strings.ReplaceAll(pc.Beatmapset.Covers.Card, "\\", "")
			pc.Beatmapset.Covers.Card2X, left = findWithIndex(pageStr, "card@2x : ", " ", left)
			pc.Beatmapset.Covers.Card2X = strings.ReplaceAll(pc.Beatmapset.Covers.Card2X, "\\", "")
			pc.Beatmapset.Covers.List, left = findWithIndex(pageStr, "list : ", " ", left)
			pc.Beatmapset.Covers.List = strings.ReplaceAll(pc.Beatmapset.Covers.List, "\\", "")
			pc.Beatmapset.Covers.List2X, left = findWithIndex(pageStr, "list@2x : ", " ", left)
			pc.Beatmapset.Covers.List2X = strings.ReplaceAll(pc.Beatmapset.Covers.List2X, "\\", "")
			pc.Beatmapset.Covers.SlimCover, left = findWithIndex(pageStr, "slimcover : ", " ", left)
			pc.Beatmapset.Covers.SlimCover = strings.ReplaceAll(pc.Beatmapset.Covers.SlimCover, "\\", "")
			pc.Beatmapset.Covers.SlimCover2X, left = findWithIndex(pageStr, "slimcover@2x : ", " ", left)
			pc.Beatmapset.Covers.SlimCover2X = strings.ReplaceAll(pc.Beatmapset.Covers.SlimCover2X, "\\", "")

			pc.Beatmapset.Creator, left = findWithIndex(pageStr, "creator : ", " ,", left)
			pc.Beatmapset.FavoriteCount, left = findWithIndex(pageStr, "favourite_count :", ",", left)
			pc.Beatmapset.Hype, left = findWithIndex(pageStr, "hype :", ", id", left)
			pc.Beatmapset.Id, left = findWithIndex(pageStr, "id :", ",", left)
			pc.Beatmapset.Nsfw, left = findWithIndex(pageStr, "nsfw :", ",", left)
			pc.Beatmapset.Offset, left = findWithIndex(pageStr, "offset :", ",", left)
			pc.Beatmapset.PlayCount, left = findWithIndex(pageStr, "play_count :", ",", left)
			pc.Beatmapset.PreviewUrl, left = findWithIndex(pageStr, "preview_url : ", " ,", left)
			pc.Beatmapset.PreviewUrl = strings.ReplaceAll(pc.Beatmapset.PreviewUrl, "\\", "")
			pc.Beatmapset.Source, left = findWithIndex(pageStr, "source : ", " , spotlight", left)
			pc.Beatmapset.Spotlight, left = findWithIndex(pageStr, "spotlight :", ",", left)
			pc.Beatmapset.Status, left = findWithIndex(pageStr, "status : ", " ,", left)
			pc.Beatmapset.Title, left = findWithIndex(pageStr, "title : ", " , title_", left)
			pc.Beatmapset.TitleUnicode, left = findWithIndex(pageStr, "title_unicode : ", " , track_", left)
			pc.Beatmapset.TrackId, left = findWithIndex(pageStr, "track_id :", ",", left)
			pc.Beatmapset.UserId, left = findWithIndex(pageStr, "user_id :", ",", left)
			pc.Beatmapset.Video, left = findWithIndex(pageStr, "video :", "}}", left)

			// Добавление статистики
			result.BeatmapPlaycounts = append(result.BeatmapPlaycounts, pc)

		}

	} else {

		// Смещение указателя на конец
		left = index(pageStr, "items :[]", left) + 1

	}

	// Проверка на наличие статистики
	if !contains(pageStr, "replays_watched_counts :[]", left) {

		// Конец части со статистикой
		end := index(pageStr, "]}}", left) - 10

		// Цикл обработки статистики
		for left < end {

			// Инициализация структуры подсчета
			var count CountString

			// Генерация подсчета
			count.StartDate, left = findWithIndex(pageStr, "start_date : ", " ", left)
			count.Count, left = findWithIndex(pageStr, "count :", "}", left)

			// Добавление статистики
			result.ReplaysWatchedCount = append(result.ReplaysWatchedCount, count)

		}
	}

	return result
}

// Роут "/user"  для vercel
func User(w http.ResponseWriter, r *http.Request) {

	// Получение параметра id из реквеста
	id := r.URL.Query().Get("id")

	// Если параметра нет, отправка ошибки
	if id == "" {
		http.NotFound(w, r)
		return
	}

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Проверка на тип, получение статистики, форматирование и отправка
	if r.URL.Query().Get("type") == "string" {
		jsonResp, err := json.Marshal(GetUserInfoString(id))
		if err != nil {
			fmt.Print("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
		}
	} else {
		jsonResp, err := json.Marshal(GetUserInfo(id))
		if err != nil {
			fmt.Print("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
		}
	}
}
