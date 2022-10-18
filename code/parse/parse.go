package parse

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Структура для хранения полной информации о пользователе
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
	Names                   string              `json:"previous_usernames"`
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

// Структура значка профиля
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

// Структура для истории рейтинга
type HistoryString struct {
	Mode string `json:"mode"`
	Data string `json:"data"`
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
	Tags              string                   `json:"tags"`
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
	Great string `json:"great"`
	Meh   string `json:"meh"`
	Miss  string `json:"miss"`
	Ok    string `json:"ok"`
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

// Статуса пользователя
type OnlineInfo struct {
	Error  string `json:"error"`
	Status string `json:"is_online"`
}

// Функция для парсинга карты
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
		bm.Tags, left = findWithIndex(pageStr, "tags : ", " ", left)

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

// Функция для парсинга рекорда
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
			Great: find(statisticsString, "great :", ",", 0),
			Meh:   find(statisticsString, "meh :", ",", 0),
			Ok:    find(statisticsString, "ok :", ",", 0),
			Miss:  find(statisticsString, "miss :", ",", 0),
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

// Функция получения информации о пользователе
func GetUserInfoString(id, mode string) UserInfoString {

	// Если пользователь не ввёл id, по умолчанию ставит мой id
	if id == "" {
		id = "29829158"
	}

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id + "/" + mode)
	if err != nil {
		return UserInfoString{
			Error: "http.Get error",
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
			Error: "User not found",
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
	result.Names, left = findWithIndex(pageStr, "previous_usernames :[ ", " ],", left)
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

		// Цикл обработки достижений
		for left < end {

			// Инициализация достижения
			var achieve AchievementString

			// Генерация достижения
			achieve.AchievedAt, left = findWithIndex(pageStr, "achieved_at : ", " ,", left)
			achieve.AchievementId, left = findWithIndex(pageStr, "achievement_id :", "}", left)

			// Добавление достижения
			result.Achievements = append(result.Achievements, achieve)

		}

	}

	// Проверка на наличие статистики
	if !contains(pageStr, " rank_history :null", left) {
		result.RankHistory.Mode, left = findWithIndex(pageStr, "mode : ", " ,", left)
		result.RankHistory.Data, left = findWithIndex(pageStr, "data :[", "]", left)
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
			pc.Beatmapset.Hype, left = findWithIndex(pageStr, "hype :", ",", left)
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
	pageStr := string(body)[80000:]

	// Сохранение html"ки в файл sample.html (для тестов)
	/*
		if err := os.WriteFile("sample.html", []byte(pageStr), 0666); err != nil {
			log.Fatal(err)
		}
	*/

	// Проверка на страницу пользователя
	if strings.Contains(pageStr, "js-react--profile") {
		return OnlineInfo{
			Status: find(pageStr, "is_online&quot;:", ",", 0),
		}
	}

	return OnlineInfo{
		Error: "User not found",
	}

}
