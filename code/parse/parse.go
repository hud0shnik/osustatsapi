package parse

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Структура для хранения полной информации о пользователе
type UserInfo struct {
	Error                    string              `json:"error"`
	AvatarUrl                string              `json:"avatar_url"`
	CountryCode              string              `json:"country_code"`
	DefaultGroup             string              `json:"default_group"`
	UserID                   int                 `json:"id"`
	IsActive                 bool                `json:"is_active"`
	IsBot                    bool                `json:"is_bot"`
	IsDeleted                bool                `json:"is_deleted"`
	IsOnline                 bool                `json:"is_online"`
	IsSupporter              bool                `json:"is_supporter"`
	LastVisit                string              `json:"last_visit"`
	PmFriendsOnly            bool                `json:"pm_friends_only"`
	ProfileColor             string              `json:"profile_color"`
	Username                 string              `json:"username"`
	CoverUrl                 string              `json:"cover_url"`
	Discord                  string              `json:"discord"`
	HasSupported             bool                `json:"has_supported"`
	Interests                string              `json:"interests"`
	JoinDate                 string              `json:"join_date"`
	Kudosu                   int                 `json:"kudosu"`
	Location                 string              `json:"location"`
	MaxFriends               int                 `json:"max_friends"`
	MaxBLock                 int                 `json:"max_block"`
	Occupation               string              `json:"occupation"`
	Playmode                 string              `json:"playmode"`
	Playstyle                string              `json:"playstyle"`
	PostCount                int                 `json:"post_count"`
	ProfileOrder             string              `json:"profile_order"`
	Title                    string              `json:"title"`
	TitleUrl                 string              `json:"title_url"`
	Twitter                  string              `json:"twitter"`
	Website                  string              `json:"website"`
	CountyName               string              `json:"country_name"`
	UserCover                Cover               `json:"cover"`
	IsAdmin                  bool                `json:"is_admin"`
	IsBng                    bool                `json:"is_bng"`
	IsFullBan                bool                `json:"is_full_bn"`
	IsGmt                    bool                `json:"is_gmt"`
	IsLimitedBan             bool                `json:"is_limited_bn"`
	IsModerator              bool                `json:"is_moderator"`
	IsNat                    bool                `json:"is_nat"`
	IsRestricted             bool                `json:"is_restricted"`
	IsSilenced               bool                `json:"is_silenced"`
	AccountHistory           string              `json:"account_history"`
	ActiveTournamentBanner   string              `json:"active_tournament_banner"`
	Badges                   []Badge             `json:"badges"`
	CommentsCount            int                 `json:"comments_count"`
	BeatmapPlaycountsCount   int                 `json:"beatmap_playcounts_count"`
	FavoriteBeatmapsetCount  int                 `json:"favorite_beatmapset_count"`
	FollowerCount            int                 `json:"follower_count"`
	GraveyardBeatmapsetCount int                 `json:"graveyard_beatmapset_count"`
	Groups                   string              `json:"groups"`
	GuestBeatmapsetCount     int                 `json:"guest_beatmapset_count"`
	LovedBeatmapsetCount     int                 `json:"loved_beatmapset_count"`
	MappingFollowerCount     int                 `json:"mapping_follower_count"`
	MonthlyPlaycounts        []CountString       `json:"monthly_playcounts"`
	PendingBeatmapsetCount   int                 `json:"pending_beatmapset_count"`
	Names                    string              `json:"previous_usernames"`
	RankedBeatmapsetCount    int                 `json:"ranked_beatmapset_count"`
	ReplaysWatchedCount      []CountString       `json:"replays_watched_counts"`
	ScoresBestCount          int                 `json:"scores_best_count"`
	ScoresFirstCount         int                 `json:"scores_first_count"`
	ScoresPinnedCount        int                 `json:"scores_pinned_count"`
	ScoresRecentCount        int                 `json:"scores_recent_count"`
	Level                    int                 `json:"level"`
	GlobalRank               int64               `json:"global_rank"`
	PP                       float64             `json:"pp"`
	RankedScore              int                 `json:"ranked_score"`
	Accuracy                 float64             `json:"accuracy"`
	PlayCount                int                 `json:"play_count"`
	PlayTime                 string              `json:"play_time"`
	PlayTimeSeconds          int64               `json:"play_time_seconds"`
	TotalScore               int64               `json:"total_score"`
	TotalHits                int64               `json:"total_hits"`
	MaximumCombo             int                 `json:"maximum_combo"`
	Replays                  int                 `json:"replays"`
	IsRanked                 bool                `json:"is_ranked"`
	SS                       int                 `json:"ss"`
	SSH                      int                 `json:"ssh"`
	S                        int                 `json:"s"`
	SH                       int                 `json:"sh"`
	A                        int                 `json:"a"`
	CountryRank              int64               `json:"country_rank"`
	SupportLvl               int                 `json:"support_level"`
	Achievements             []AchievementString `json:"achievements"`
	RankHistory              HistoryString       `json:"rank_history"`
	RankedAndApprovedCount   int                 `json:"ranked_and_approved_beatmapset_count"`
	UnrankedBeatmapsetCount  int                 `json:"unranked_beatmapset_count"`
	ScoresBest               []ScoreString       `json:"scores_best"`
	ScoresFirst              []ScoreString       `json:"scores_first"`
	ScoresPinned             []ScoreString       `json:"scores_pinned"`
}

// Структура для подсчёта
type Count struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}

// Достижение
type Achievement struct {
	AchievedAt    string `json:"achieved_at"`
	AchievementId int    `json:"achievement_id"`
}

// Структура для истории рейтинга
type History struct {
	Mode string `json:"mode"`
	Data []int  `json:"data"`
}

// Рекорд
type Score struct {
	Accuracy              float64          `json:"accuracy"`
	BeatMapId             int              `json:"beatmap_id"`
	BuildId               string           `json:"build_id"`
	EndedAt               string           `json:"ended_at"`
	MaximumCombo          int              `json:"maximum_combo"`
	Mods                  []string         `json:"mods"`
	Passed                bool             `json:"passed"`
	Rank                  string           `json:"rank"`
	RulesetId             int              `json:"ruleset_id"`
	StartedAt             string           `json:"started_at"`
	Statistics            string           `json:"statistics"`
	TotalScore            int              `json:"total_score"`
	UserId                int              `json:"user_id"`
	BestId                int              `json:"best_id"`
	Id                    int              `json:"id"`
	LegacyPerfect         bool             `json:"legacy_perfect"`
	PP                    float64          `json:"pp"`
	Replay                bool             `json:"replay"`
	Type                  string           `json:"type"`
	CurrentUserAttributes string           `json:"current_user_attributes"`
	BeatMap               BeatMapString    `json:"beatmap"`
	BeatMapSet            BeatMapSetString `json:"beatmapset"`
	Weight                WeightString     `json:"weight"`
}

// Мапа
type BeatMap struct {
	BeatMapSetId     int     `json:"beatmapset_id"`
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

// Мап сет
type BeatMapSet struct {
	Artist        string `json:"artist"`
	ArtistUnicode string `json:"artist_unicode"`
	Covers        Covers `json:"covers"`
	Creator       string `json:"creator"`
	FavoriteCount int    `json:"favorite_count"`
	Hype          string `json:"hype"`
	Id            int    `json:"id"`
	Nsfw          bool   `json:"nsfw"`
	Offset        int    `json:"offset"`
	PlayCount     int64  `json:"play_count"`
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

// Статистика
type Weight struct {
	Percentage int     `json:"percentage"`
	PP         float64 `json:"pp"`
}

// Статуса пользователя
type OnlineInfo struct {
	Error  string `json:"error"`
	Status string `json:"is_online"`
}

// Функция получения информации о пользователе
func GetOnlineInfo(id string) OnlineInfo {

	// Если пользователь не ввёл id, по умолчанию ставит мой id
	if id == "" {
		id = "29829158"
	}

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return OnlineInfo{
			Error: "http.Get error",
		}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// HTML полученной страницы в формате string
	pageStr := string(body)[90000:]

	// Сохранение html"ки в файл sample.html (для тестов)
	/*
		if err := os.WriteFile("sample.html", []byte(pageStr), 0666); err != nil {
			log.Fatal(err)
		}
	*/

	// Проверка на страницу пользователя
	if strings.Contains(pageStr, "js-react--profile") {
		return OnlineInfo{
			Status: find(pageStr, "is_online&quot;:", ","),
		}
	}

	return OnlineInfo{
		Error: "User not found",
	}

}

// Функция получения информации о пользователе
func GetUserInfo(id, mode string) UserInfo {

	resultStr := GetUserInfoString(id, mode)

	result := UserInfo{
		Error:                    resultStr.Error,
		AvatarUrl:                resultStr.AvatarUrl,
		CountryCode:              resultStr.CountryCode,
		DefaultGroup:             resultStr.DefaultGroup,
		UserID:                   ToInt(resultStr.UserID),
		IsActive:                 ToBool(resultStr.IsActive),
		IsBot:                    ToBool(resultStr.IsBot),
		IsDeleted:                ToBool(resultStr.IsDeleted),
		IsOnline:                 ToBool(resultStr.IsOnline),
		IsSupporter:              ToBool(resultStr.IsSupporter),
		LastVisit:                resultStr.LastVisit,
		PmFriendsOnly:            ToBool(resultStr.PmFriendsOnly),
		ProfileColor:             resultStr.ProfileColor,
		Username:                 resultStr.Username,
		CoverUrl:                 resultStr.CoverUrl,
		Discord:                  resultStr.Discord,
		HasSupported:             ToBool(resultStr.HasSupported),
		Interests:                resultStr.Interests,
		JoinDate:                 resultStr.JoinDate,
		Kudosu:                   ToInt(resultStr.Kudosu),
		Location:                 resultStr.Location,
		MaxFriends:               ToInt(resultStr.MaxFriends),
		MaxBLock:                 ToInt(resultStr.MaxBLock),
		Occupation:               resultStr.Occupation,
		Playmode:                 resultStr.Playmode,
		Playstyle:                resultStr.Playstyle,
		PostCount:                ToInt(resultStr.PostCount),
		ProfileOrder:             resultStr.ProfileOrder,
		Title:                    resultStr.Title,
		TitleUrl:                 resultStr.TitleUrl,
		Twitter:                  resultStr.Twitter,
		Website:                  resultStr.Website,
		CountyName:               resultStr.CountyName,
		UserCover:                resultStr.UserCover,
		IsAdmin:                  ToBool(resultStr.IsAdmin),
		IsBng:                    ToBool(resultStr.IsBng),
		IsFullBan:                ToBool(resultStr.IsFullBan),
		IsGmt:                    ToBool(resultStr.IsGmt),
		IsLimitedBan:             ToBool(resultStr.IsLimitedBan),
		IsModerator:              ToBool(resultStr.IsModerator),
		IsNat:                    ToBool(resultStr.IsNat),
		IsRestricted:             ToBool(resultStr.IsRestricted),
		IsSilenced:               ToBool(resultStr.IsSilenced),
		AccountHistory:           resultStr.AccountHistory,
		ActiveTournamentBanner:   resultStr.ActiveTournamentBanner,
		Badges:                   resultStr.Badges,
		CommentsCount:            ToInt(resultStr.CommentsCount),
		BeatmapPlaycountsCount:   ToInt(resultStr.BeatmapPlaycountsCount),
		FavoriteBeatmapsetCount:  ToInt(resultStr.FavoriteBeatmapsetCount),
		GraveyardBeatmapsetCount: ToInt(resultStr.GraveyardBeatmapsetCount),
		Groups:                   resultStr.Groups,
		GuestBeatmapsetCount:     ToInt(resultStr.GuestBeatmapsetCount),
		LovedBeatmapsetCount:     ToInt(resultStr.LovedBeatmapsetCount),
		MappingFollowerCount:     ToInt(resultStr.MappingFollowerCount),
		// monthly_playcounts
		PendingBeatmapsetCount: ToInt(resultStr.PendingBeatmapsetCount),
		Names:                  resultStr.Names,
		RankedBeatmapsetCount:  ToInt(resultStr.RankedBeatmapsetCount),
		// replays_watched_counts
		ScoresBestCount:   ToInt(resultStr.ScoresBestCount),
		ScoresFirstCount:  ToInt(resultStr.ScoresFirstCount),
		ScoresPinnedCount: ToInt(resultStr.ScoresPinnedCount),
		ScoresRecentCount: ToInt(resultStr.ScoresRecentCount),
		Level:             ToInt(resultStr.Level),
		GlobalRank:        ToInt64(resultStr.GlobalRank),
		PP:                ToFloat64(resultStr.PP),
		RankedScore:       ToInt(resultStr.RankedScore),
		Accuracy:          ToFloat64(resultStr.Accuracy),
		PlayCount:         ToInt(resultStr.PlayCount),
		PlayTime:          resultStr.PlayTime,
		PlayTimeSeconds:   ToInt64(resultStr.PlayTimeSeconds),
		TotalScore:        ToInt64(resultStr.TotalScore),
		TotalHits:         ToInt64(resultStr.TotalHits),
		MaximumCombo:      ToInt(resultStr.MaximumCombo),
		Replays:           ToInt(resultStr.Replays),
		IsRanked:          ToBool(resultStr.IsRanked),
		SS:                ToInt(resultStr.SS),
		SSH:               ToInt(resultStr.SSH),
		S:                 ToInt(resultStr.S),
		SH:                ToInt(resultStr.SH),
		A:                 ToInt(resultStr.A),
		// achievements
		// rank history
		RankedAndApprovedCount:  ToInt(resultStr.RankedAndApprovedCount),
		UnrankedBeatmapsetCount: ToInt(resultStr.UnrankedBeatmapsetCount),
		// scores_best
		// scores_first
		// scores_pinned
	}

	return result
}
