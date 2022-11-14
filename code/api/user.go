package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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

// ---------------------- Функции перевода ------------------------

func ToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("parsing error: \t", s)
		return 0
	}

	return i
}

func ToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		fmt.Println("parsing error: \t", s)
		return 0
	}

	return i
}

func ToBool(s string) bool {
	f, err := strconv.ParseBool(s)

	if err != nil {
		fmt.Println("parsing error: \t", s)
		return false
	}

	return f
}

func ToFloat64(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)

	if err != nil {
		fmt.Println("parsing error: \t", s)
		return 0
	}

	return i
}

// ---------------------- Функции поиска ------------------------

// Функция поиска. Возвращает искомое значение и индекс последнего символа
func findWithIndex(str, subStr, stopChar string, start int) (string, int) {

	// Обрезка левой границы поиска
	str = str[start:]

	// Проверка на существование нужной строки
	if strings.Contains(str, subStr) {

		// Поиск индекса начала нужной строки
		left := strings.Index(str, subStr) + len(subStr)

		// Поиск правой границы
		right := left + strings.Index(str[left:], stopChar)

		// Обрезка и вывод результата
		return str[left:right], right + start
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

// Функция перевода карты
func FormatBeatmaps(bms []BeatmapString) []Beatmap {

	var result []Beatmap

	for _, bm := range bms {

		result = append(result, Beatmap{
			Artist:            bm.Artist,
			ArtistUnicode:     bm.ArtistUnicode,
			Covers:            bm.Covers,
			Creator:           bm.Creator,
			FavoriteCount:     ToInt(bm.FavoriteCount),
			Hype:              bm.Hype,
			Id:                ToInt(bm.Id),
			Nsfw:              ToBool(bm.Nsfw),
			Offset:            ToInt(bm.Offset),
			PlayCount:         ToInt(bm.PlayCount),
			PreviewUrl:        bm.PreviewUrl,
			Source:            bm.Source,
			Spotlight:         ToBool(bm.Spotlight),
			Status:            bm.Status,
			Title:             bm.Title,
			TitleUnicode:      bm.TitleUnicode,
			TrackId:           bm.TrackId,
			UserId:            ToInt(bm.UserId),
			Video:             ToBool(bm.Video),
			DownloadDisabled:  ToBool(bm.DownloadDisabled),
			Bpm:               ToFloat64(bm.Bpm),
			CanBeHyped:        ToBool(bm.CanBeHyped),
			DiscussionEnabled: ToBool(bm.DiscussionEnabled),
			DiscussionLocked:  ToBool(bm.DiscussionLocked),
			IsScoreable:       ToBool(bm.IsScoreable),
			LastUpdated:       bm.LastUpdated,
			LegacyThreadUrl:   bm.LegacyThreadUrl,
			Nominations: NominationsSummary{
				Current:  ToInt(bm.Nominations.Current),
				Required: ToInt(bm.Nominations.Required),
			},
			Ranked:        ToInt(bm.Ranked),
			RankedDate:    bm.RankedDate,
			Storyboard:    ToBool(bm.Storyboard),
			SubmittedDate: bm.SubmittedDate,
			Tags:          bm.Tags,
			Beatmap: Beatmaps{
				BeatmapSetId:     ToInt(bm.Beatmap.BeatmapSetId),
				DifficultyRating: ToFloat64(bm.Beatmap.DifficultyRating),
				Id:               ToInt(bm.Beatmap.Id),
				Mode:             bm.Beatmap.Mode,
				Status:           bm.Beatmap.Status,
				TotalLength:      ToInt(bm.Beatmap.TotalLength),
				UserId:           ToInt(bm.Beatmap.UserId),
				Version:          bm.Beatmap.Version,
				Accuracy:         ToFloat64(bm.Beatmap.Accuracy),
				Ar:               ToFloat64(bm.Beatmap.Ar),
				Bpm:              ToFloat64(bm.Beatmap.Bpm),
				Convert:          ToBool(bm.Beatmap.Convert),
				CountCircles:     ToInt(bm.Beatmap.CountCircles),
				CountSliders:     ToInt(bm.Beatmap.CountSliders),
				CountSpinners:    ToInt(bm.Beatmap.CountSpinners),
				Cs:               ToFloat64(bm.Beatmap.Cs),
				DeletedAt:        bm.Beatmap.DeletedAt,
				Drain:            ToFloat64(bm.Beatmap.Drain),
				HitLength:        ToInt(bm.Beatmap.HitLength),
				IsScoreable:      ToBool(bm.Beatmap.IsScoreable),
				LastUpdated:      bm.Beatmap.LastUpdated,
				ModeInt:          ToInt(bm.Beatmap.ModeInt),
				PassCount:        ToInt(bm.Beatmap.PassCount),
				PlayCount:        ToInt(bm.Beatmap.PlayCount),
				Ranked:           ToInt(bm.Beatmap.Ranked),
				Url:              bm.Beatmap.Url,
				Checksum:         bm.Beatmap.Checksum,
			},
		})

	}

	return result
}

// Функция перевода рекорда
func FormatScores(scs []ScoreString) []Score {

	var result []Score

	for _, sc := range scs {

		result = append(result, Score{
			Accuracy:         ToFloat64(sc.Accuracy),
			BeatmapId:        ToInt(sc.BeatmapId),
			BuildId:          sc.BuildId,
			EndedAt:          sc.EndedAt,
			LegacyScoreId:    sc.LegacyScoreId,
			LegacyTotalScore: sc.LegacyTotalScore,
			MaximumCombo:     ToInt(sc.MaximumCombo),
			MaximumStatistics: Statistics{
				Good:    ToInt(sc.MaximumStatistics.Good),
				Great:   ToInt(sc.MaximumStatistics.Great),
				Meh:     ToInt(sc.MaximumStatistics.Meh),
				Miss:    ToInt(sc.MaximumStatistics.Miss),
				Ok:      ToInt(sc.MaximumStatistics.Ok),
				Perfect: ToInt(sc.MaximumStatistics.Perfect),
			},
			Mods:                  sc.Mods,
			Passed:                ToBool(sc.Passed),
			Rank:                  sc.Rank,
			RulesetId:             ToInt(sc.RulesetId),
			StartedAt:             sc.StartedAt,
			TotalScore:            ToInt(sc.TotalScore),
			UserId:                ToInt(sc.UserId),
			BestId:                ToInt(sc.BestId),
			Id:                    ToInt(sc.Id),
			LegacyPerfect:         ToBool(sc.LegacyPerfect),
			PP:                    ToFloat64(sc.PP),
			Replay:                ToBool(sc.Replay),
			Type:                  sc.Type,
			CurrentUserAttributes: sc.CurrentUserAttributes,
			Beatmap: Beatmaps{
				BeatmapSetId:     ToInt(sc.Beatmap.BeatmapSetId),
				DifficultyRating: ToFloat64(sc.Beatmap.DifficultyRating),
				Id:               ToInt(sc.Beatmap.Id),
				Mode:             sc.Beatmap.Mode,
				Status:           sc.Beatmap.Status,
				TotalLength:      ToInt(sc.Beatmap.TotalLength),
				UserId:           ToInt(sc.Beatmap.UserId),
				Version:          sc.Beatmap.Version,
				Accuracy:         ToFloat64(sc.Beatmap.Accuracy),
				Ar:               ToFloat64(sc.Beatmap.Ar),
				Bpm:              ToFloat64(sc.Beatmap.Bpm),
				Convert:          ToBool(sc.Beatmap.Convert),
				CountCircles:     ToInt(sc.Beatmap.CountCircles),
				CountSliders:     ToInt(sc.Beatmap.CountSliders),
				CountSpinners:    ToInt(sc.Beatmap.CountSpinners),
				Cs:               ToFloat64(sc.Beatmap.Cs),
				DeletedAt:        sc.Beatmap.DeletedAt,
				Drain:            ToFloat64(sc.Beatmap.Drain),
				HitLength:        ToInt(sc.Beatmap.HitLength),
				IsScoreable:      ToBool(sc.Beatmap.IsScoreable),
				LastUpdated:      sc.Beatmap.LastUpdated,
				ModeInt:          ToInt(sc.Beatmap.ModeInt),
				PassCount:        ToInt(sc.Beatmap.PassCount),
				PlayCount:        ToInt(sc.Beatmap.PlayCount),
				Ranked:           ToInt(sc.Beatmap.Ranked),
				Url:              sc.Beatmap.Url,
				Checksum:         sc.Beatmap.Checksum,
			},
			Beatmapset: Beatmapset{
				Artist:        sc.Beatmapset.Artist,
				ArtistUnicode: sc.Beatmapset.ArtistUnicode,
				Covers:        sc.Beatmapset.Covers,
				Creator:       sc.Beatmapset.Creator,
				FavoriteCount: ToInt(sc.Beatmapset.FavoriteCount),
				Hype:          sc.Beatmapset.Hype,
				Id:            ToInt(sc.Beatmapset.Id),
				Nsfw:          ToBool(sc.Beatmapset.Nsfw),
				Offset:        ToInt(sc.Beatmapset.Offset),
				PlayCount:     ToInt(sc.Beatmapset.PlayCount),
				PreviewUrl:    sc.Beatmapset.PreviewUrl,
				Source:        sc.Beatmapset.Source,
				Spotlight:     ToBool(sc.Beatmapset.Spotlight),
				Status:        sc.Beatmapset.Status,
				Title:         sc.Beatmapset.Title,
				TitleUnicode:  sc.Beatmapset.TitleUnicode,
				TrackId:       sc.Beatmapset.TrackId,
				UserId:        ToInt(sc.Beatmapset.UserId),
				Video:         ToBool(sc.Beatmapset.Video),
			},
			Weight: Weight{
				Percentage: ToFloat64(sc.Weight.Percentage),
				PP:         ToFloat64(sc.Weight.PP),
			},
		})

	}

	return result
}

// Функция получения информации о пользователе
func GetUserInfo(id string) UserInfo {

	resultStr := GetUserInfoString(id)

	result := UserInfo{
		Error:         resultStr.Error,
		AvatarUrl:     resultStr.AvatarUrl,
		CountryCode:   resultStr.CountryCode,
		DefaultGroup:  resultStr.DefaultGroup,
		UserID:        ToInt(resultStr.UserID),
		IsActive:      ToBool(resultStr.IsActive),
		IsBot:         ToBool(resultStr.IsBot),
		IsDeleted:     ToBool(resultStr.IsDeleted),
		IsOnline:      ToBool(resultStr.IsOnline),
		IsSupporter:   ToBool(resultStr.IsSupporter),
		LastVisit:     resultStr.LastVisit,
		PmFriendsOnly: ToBool(resultStr.PmFriendsOnly),
		ProfileColor:  resultStr.ProfileColor,
		Username:      resultStr.Username,
		CoverUrl:      resultStr.CoverUrl,
		Discord:       resultStr.Discord,
		HasSupported:  ToBool(resultStr.HasSupported),
		Interests:     resultStr.Interests,
		JoinDate:      resultStr.JoinDate,
		Kudosu:        ToInt(resultStr.Kudosu),
		Location:      resultStr.Location,
		MaxFriends:    ToInt(resultStr.MaxFriends),
		MaxBLock:      ToInt(resultStr.MaxBLock),
		Occupation:    resultStr.Occupation,
		Playmode:      resultStr.Playmode,
		Playstyle:     resultStr.Playstyle,
		PostCount:     ToInt(resultStr.PostCount),
		ProfileOrder:  resultStr.ProfileOrder,
		Title:         resultStr.Title,
		TitleUrl:      resultStr.TitleUrl,
		Twitter:       resultStr.Twitter,
		Website:       resultStr.Website,
		CountyName:    resultStr.CountyName,
		UserCover: Cover{
			CustomUrl: resultStr.UserCover.CustomUrl,
			Url:       resultStr.UserCover.Url,
			Id:        ToInt(resultStr.UserCover.Id),
		},
		IsAdmin:                 ToBool(resultStr.IsAdmin),
		IsBng:                   ToBool(resultStr.IsBng),
		IsFullBan:               ToBool(resultStr.IsFullBan),
		IsGmt:                   ToBool(resultStr.IsGmt),
		IsLimitedBan:            ToBool(resultStr.IsLimitedBan),
		IsModerator:             ToBool(resultStr.IsModerator),
		IsNat:                   ToBool(resultStr.IsNat),
		IsRestricted:            ToBool(resultStr.IsRestricted),
		IsSilenced:              ToBool(resultStr.IsSilenced),
		AccountHistory:          resultStr.AccountHistory,
		ActiveTournamentBanner:  resultStr.ActiveTournamentBanner,
		Badges:                  resultStr.Badges,
		CommentsCount:           ToInt(resultStr.CommentsCount),
		FollowerCount:           ToInt(resultStr.FollowerCount),
		Groups:                  resultStr.Groups,
		MappingFollowerCount:    ToInt(resultStr.MappingFollowerCount),
		PendingBeatmapsetCount:  ToInt(resultStr.PendingBeatmapsetCount),
		Names:                   resultStr.Names,
		Level:                   ToInt(resultStr.Level),
		GlobalRank:              ToInt64(resultStr.GlobalRank),
		PP:                      ToFloat64(resultStr.PP),
		RankedScore:             ToInt(resultStr.RankedScore),
		Accuracy:                ToFloat64(resultStr.Accuracy),
		PlayCount:               ToInt(resultStr.PlayCount),
		PlayTime:                resultStr.PlayTime,
		PlayTimeSeconds:         ToInt64(resultStr.PlayTimeSeconds),
		TotalScore:              ToInt64(resultStr.TotalScore),
		TotalHits:               ToInt64(resultStr.TotalHits),
		MaximumCombo:            ToInt(resultStr.MaximumCombo),
		Replays:                 ToInt(resultStr.Replays),
		IsRanked:                ToBool(resultStr.IsRanked),
		SS:                      ToInt(resultStr.SS),
		SSH:                     ToInt(resultStr.SSH),
		S:                       ToInt(resultStr.S),
		SH:                      ToInt(resultStr.SH),
		A:                       ToInt(resultStr.A),
		CountryRank:             ToInt(resultStr.CountryRank),
		SupportLvl:              ToInt(resultStr.SupportLvl),
		UnrankedBeatmapsetCount: ToInt(resultStr.UnrankedBeatmapsetCount),
	}

	for _, c := range resultStr.Achievements {
		result.Achievements = append(result.Achievements, Achievement{
			AchievedAt:    c.AchievedAt,
			AchievementId: ToInt(c.AchievementId),
		})
	}

	result.RankHistory.Mode = resultStr.RankHistory.Mode

	for _, d := range resultStr.RankHistory.Data {
		result.RankHistory.Data = append(result.RankHistory.Data, ToInt(d))
	}

	result.FavoriteBeatmaps = FormatBeatmaps(resultStr.FavoriteBeatmaps)
	result.GraveyardBeatmaps = FormatBeatmaps(resultStr.GraveyardBeatmaps)
	result.GuestBeatmaps = FormatBeatmaps(resultStr.GuestBeatmaps)
	result.LovedBeatmaps = FormatBeatmaps(resultStr.LovedBeatmaps)
	result.RankedBeatmaps = FormatBeatmaps(resultStr.RankedBeatmaps)
	result.PendingBeatmaps = FormatBeatmaps(resultStr.PendingBeatmaps)

	for _, k := range resultStr.KudosuItems {
		result.KudosuItems = append(result.KudosuItems, Kudosu{
			Id:        ToInt(k.Id),
			Action:    k.Action,
			Amount:    ToInt(k.Amount),
			Model:     k.Model,
			CreatedAt: k.CreatedAt,
			Giver:     k.Giver,
			Post:      k.Post,
			Details:   k.Details,
		})
	}

	for _, a := range resultStr.RecentActivity {
		result.RecentActivity = append(result.RecentActivity, Activity{
			CreatedAt:    a.CreatedAt,
			Id:           ToInt(a.Id),
			Type:         a.Type,
			ScoreRank:    a.ScoreRank,
			Rank:         ToInt(a.Rank),
			Mode:         a.Mode,
			BeatmapTitle: a.BeatmapTitle,
			BeatmapUrl:   a.BeatmapUrl,
		})
	}

	result.Best = FormatScores(resultStr.Best)
	result.Firsts = FormatScores(resultStr.Firsts)
	result.Pinned = FormatScores(resultStr.Pinned)

	for _, pc := range resultStr.BeatmapPlaycounts {
		result.BeatmapPlaycounts = append(result.BeatmapPlaycounts, PlayCount{
			BeatmapId: ToInt(pc.BeatmapId),
			Count:     ToInt(pc.Count),
			Beatmap: PlayCountBeatmap{
				BeatmapsetId:     ToInt(pc.Beatmap.BeatmapsetId),
				DifficultyRating: ToFloat64(pc.Beatmap.DifficultyRating),
				Id:               ToInt(pc.Beatmap.Id),
				Status:           pc.Beatmap.Status,
				TotalLength:      ToInt(pc.Beatmap.TotalLength),
				UserId:           ToInt(pc.Beatmap.UserId),
				Version:          pc.Beatmap.Version,
			},
			Beatmapset: Beatmapset{
				Artist:        pc.Beatmapset.Artist,
				ArtistUnicode: pc.Beatmapset.ArtistUnicode,
				Covers:        pc.Beatmapset.Covers,
				Creator:       pc.Beatmapset.Creator,
				FavoriteCount: ToInt(pc.Beatmapset.FavoriteCount),
				Hype:          pc.Beatmapset.Hype,
				Id:            ToInt(pc.Beatmapset.Id),
				Nsfw:          ToBool(pc.Beatmapset.Nsfw),
				Offset:        ToInt(pc.Beatmapset.Offset),
				PlayCount:     ToInt(pc.Beatmapset.PlayCount),
				PreviewUrl:    pc.Beatmapset.PreviewUrl,
				Source:        pc.Beatmapset.Source,
				Spotlight:     ToBool(pc.Beatmapset.Spotlight),
				Status:        pc.Beatmapset.Status,
				Title:         pc.Beatmapset.Title,
				TitleUnicode:  pc.Beatmapset.TitleUnicode,
				TrackId:       pc.Beatmapset.TrackId,
				UserId:        ToInt(pc.Beatmapset.UserId),
				Video:         ToBool(pc.Beatmapset.Video),
			},
		})
	}

	for _, c := range resultStr.MonthlyPlaycounts {
		result.MonthlyPlaycounts = append(result.MonthlyPlaycounts, Count{
			StartDate: c.StartDate,
			Count:     ToInt(c.Count),
		})
	}

	for _, c := range resultStr.ReplaysWatchedCount {
		result.ReplaysWatchedCount = append(result.ReplaysWatchedCount, Count{
			StartDate: c.StartDate,
			Count:     ToInt(c.Count),
		})
	}

	return result
}

// Роут "/user"
func User(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	resp := GetUserInfo(id)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Print("Error: ", err)
	} else {
		w.Write(jsonResp)
	}
}