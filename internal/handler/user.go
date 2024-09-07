package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hud0shnik/osustatsapi/internal/convert"
	"github.com/hud0shnik/osustatsapi/internal/parse"
)

// apiError - структура ошибки
type apiError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// userInfo - информация о пользователе
type userInfo struct {
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
	UserCover               cover         `json:"cover"`
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
	Badges                  []badge       `json:"badges"`
	CommentsCount           int           `json:"comments_count"`
	FollowerCount           int           `json:"follower_count"`
	Groups                  string        `json:"groups"`
	MappingFollowerCount    int           `json:"mapping_follower_count"`
	PendingBeatmapsetCount  int           `json:"pending_beatmapset_count"`
	Names                   []string      `json:"previous_usernames"`
	RankHighest             int           `json:"rank_highest"`
	Count100                int           `json:"count_100"`
	Count300                int           `json:"count_300"`
	Count50                 int           `json:"count_50"`
	CountMiss               int           `json:"count_miss"`
	Level                   int           `json:"level"`
	GlobalRank              int64         `json:"global_rank"`
	PP                      float64       `json:"pp"`
	PPExp                   int           `json:"pp_exp"`
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
	Achievements            []achievement `json:"achievements"`
	Medals                  int           `json:"medals"`
	RankHistory             history       `json:"rank_history"`
	UnrankedBeatmapsetCount int           `json:"unranked_beatmapset_count"`
}

// cover - ковёр пользователя
type cover struct {
	CustomUrl string `json:"custom_url"`
	Url       string `json:"url"`
	Id        int    `json:"id"`
}

// badge - значок профиля
type badge struct {
	AwardedAt   string `json:"awarded_at"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

// achievement - достижение
type achievement struct {
	AchievedAt    string `json:"achieved_at"`
	AchievementId int    `json:"achievement_id"`
}

// history - история рейтинга
type history struct {
	Mode string `json:"mode"`
	Data []int  `json:"data"`
}

// userInfoString - информация о пользователе в формате строк
type userInfoString struct {
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
	UserCover               coverString         `json:"cover"`
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
	Badges                  []badge             `json:"badges"`
	CommentsCount           string              `json:"comments_count"`
	FollowerCount           string              `json:"follower_count"`
	Groups                  string              `json:"groups"`
	MappingFollowerCount    string              `json:"mapping_follower_count"`
	PendingBeatmapsetCount  string              `json:"pending_beatmapset_count"`
	Names                   []string            `json:"previous_usernames"`
	RankHighest             string              `json:"rank_highest"`
	Count100                string              `json:"count_100"`
	Count300                string              `json:"count_300"`
	Count50                 string              `json:"count_50"`
	CountMiss               string              `json:"count_miss"`
	Level                   string              `json:"level"`
	GlobalRank              string              `json:"global_rank"`
	PP                      string              `json:"pp"`
	PPExp                   string              `json:"pp_exp"`
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
	Achievements            []achievementString `json:"achievements"`
	Medals                  string              `json:"medals"`
	RankHistory             historyString       `json:"rank_history"`
	UnrankedBeatmapsetCount string              `json:"unranked_beatmapset_count"`
}

// coverString - ковёр пользователя в формате строк
type coverString struct {
	CustomUrl string `json:"custom_url"`
	Url       string `json:"url"`
	Id        string `json:"id"`
}

// achievementString - достижение в формате строк
type achievementString struct {
	AchievedAt    string `json:"achieved_at"`
	AchievementId string `json:"achievement_id"`
}

// historyString - история рейтинга в формате строк
type historyString struct {
	Mode string   `json:"mode"`
	Data []string `json:"data"`
}

// getUserInfoString возвращает структуру с информацией пользователя в формате строк, статус код и ошибку
func getUserInfoString(id string) (userInfoString, int, error) {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return userInfoString{}, http.StatusInternalServerError,
			fmt.Errorf("in http.Get: %w", err)
	}
	defer resp.Body.Close()

	// Проверка статускода
	if resp.StatusCode != 200 {
		return userInfoString{}, resp.StatusCode,
			fmt.Errorf("in http.Get: %s", resp.Status)
	}

	// Запись респонса
	body, _ := io.ReadAll(resp.Body)

	// HTML полученной страницы в формате string
	pageStr := string(body)[80000:]

	// Обрезка юзелесс части html"ки
	pageStr = strings.ReplaceAll(pageStr[strings.Index(pageStr, "current_mode"):], "&quot;", " ")

	// Сохранение html"ки в файл sample.html (для тестов)
	/*
		if err := os.WriteFile("sample.html", []byte(pageStr), 0666); err != nil {
			logrus.Fatal(err)
		}
	*/

	// Структура, которую будет возвращать функция
	result := userInfoString{}

	// Крайняя левая граница поиска
	left := 0

	//--------------------------- Статистика игрока ------------------------------

	result.AvatarUrl, left = parse.FindWithIndex(pageStr, "avatar_url : ", " ", left, -1)
	result.AvatarUrl = strings.ReplaceAll(result.AvatarUrl, "\\", "")
	result.CountryCode, left = parse.FindWithIndex(pageStr, "country_code : ", " ", left, -1)
	result.DefaultGroup, left = parse.FindWithIndex(pageStr, "default_group : ", " ", left, -1)
	result.UserID, left = parse.FindWithIndex(pageStr, " id :", ",", left, -1)
	result.IsActive, left = parse.FindWithIndex(pageStr, "is_active :", ",", left, -1)
	result.IsBot, left = parse.FindWithIndex(pageStr, "is_bot :", ",", left, -1)
	result.IsDeleted, left = parse.FindWithIndex(pageStr, "is_deleted :", ",", left, -1)
	result.IsOnline, left = parse.FindWithIndex(pageStr, "is_online :", ",", left, -1)
	result.IsSupporter, left = parse.FindWithIndex(pageStr, "is_supporter :", ",", left, -1)
	result.LastVisit, left = parse.FindWithIndex(pageStr, "last_visit : ", " ", left, -1)
	result.PmFriendsOnly, left = parse.FindWithIndex(pageStr, "pm_friends_only :", ",", left, -1)
	result.ProfileColor, left = parse.FindWithIndex(pageStr, "profile_colour : ", " ,", left, -1)
	result.Username, left = parse.FindWithIndex(pageStr, "username : ", " ", left, -1)
	result.CoverUrl, left = parse.FindWithIndex(pageStr, "cover_url : ", " ", left, -1)
	result.CoverUrl = strings.ReplaceAll(result.CoverUrl, "\\", "")
	result.Discord, left = parse.FindWithIndex(pageStr, "discord : ", " ,", left, -1)
	result.HasSupported, left = parse.FindWithIndex(pageStr, "has_supported :", ",", left, -1)
	result.Interests, left = parse.FindWithIndex(pageStr, "interests : ", " , join_date", left, -1)
	result.JoinDate, left = parse.FindWithIndex(pageStr, "join_date : ", " ,", left, -1)
	result.Kudosu, left = parse.FindWithIndex(pageStr, "kudosu :{ total :", ",", left, -1)
	result.Location, left = parse.FindWithIndex(pageStr, "location : ", " ,", left, -1)
	result.MaxBLock, left = parse.FindWithIndex(pageStr, "max_blocks :", ",", left, -1)
	result.MaxFriends, left = parse.FindWithIndex(pageStr, "max_friends :", ",", left, -1)
	result.Occupation, left = parse.FindWithIndex(pageStr, "occupation : ", " ,", left, -1)
	result.Playmode, left = parse.FindWithIndex(pageStr, "playmode : ", " ,", left, -1)
	result.Playstyle = strings.Split(parse.Find(pageStr, "playstyle :[ ", " ],", left), " , ")
	if result.Playstyle[0] == "" {
		result.Playstyle = nil
	}
	result.PostCount, left = parse.FindWithIndex(pageStr, "post_count :", ",", left, -1)
	result.ProfileOrder = strings.Split(parse.Find(pageStr, "profile_order :[ ", " ]", left), " , ")
	if result.ProfileOrder[0] == "" {
		result.ProfileOrder = nil
	}
	result.Title, left = parse.FindWithIndex(pageStr, "title :", ",", left, -1)
	result.TitleUrl, left = parse.FindWithIndex(pageStr, "title_url : ", " ,", left, -1)
	result.Twitter, left = parse.FindWithIndex(pageStr, "twitter : ", " ,", left, -1)
	result.Website, left = parse.FindWithIndex(pageStr, "website : ", " ,", left, -1)
	result.Website = strings.ReplaceAll(result.Website, "\\", "")
	result.CountyName, left = parse.FindWithIndex(pageStr, " name : ", " }", left, -1)

	result.UserCover.CustomUrl, left = parse.FindWithIndex(pageStr, "custom_url : ", " ,", left, -1)
	result.UserCover.CustomUrl = strings.ReplaceAll(result.UserCover.CustomUrl, "\\", "")
	result.UserCover.Url, left = parse.FindWithIndex(pageStr, "url : ", " ,", left, -1)
	result.UserCover.Url = strings.ReplaceAll(result.UserCover.Url, "\\", "")
	result.UserCover.Id, left = parse.FindWithIndex(pageStr, " , id : ", " }", left, -1)

	result.IsAdmin, left = parse.FindWithIndex(pageStr, "is_admin :", ",", left, -1)
	result.IsBng, left = parse.FindWithIndex(pageStr, "is_bng :", ",", left, -1)
	result.IsFullBan, left = parse.FindWithIndex(pageStr, "is_full_bn :", ",", left, -1)
	result.IsGmt, left = parse.FindWithIndex(pageStr, "is_gmt :", ",", left, -1)
	result.IsLimitedBan, left = parse.FindWithIndex(pageStr, "is_limited_bn :", ",", left, -1)
	result.IsModerator, left = parse.FindWithIndex(pageStr, "is_moderator :", ",", left, -1)
	result.IsNat, left = parse.FindWithIndex(pageStr, "is_nat :", ",", left, -1)
	result.IsRestricted, left = parse.FindWithIndex(pageStr, "is_restricted :", ",", left, -1)
	result.IsSilenced, left = parse.FindWithIndex(pageStr, "is_silenced :", ",", left, -1)
	result.ActiveTournamentBanner, left = parse.FindWithIndex(pageStr, "active_tournament_banner :", ", badges", left, -1)
	result.ActiveTournamentBanner = strings.ReplaceAll(result.ActiveTournamentBanner, "\\", "")

	// Значки
	for c := parse.Index(pageStr, "badges :[", left, -1); pageStr[c] != ']'; c++ {
		if pageStr[c:c+13] == "awarded_at : " {
			result.Badges = append(result.Badges, badge{
				AwardedAt:   parse.Find(pageStr[c:], "awarded_at : ", " ", 0),
				Description: parse.Find(pageStr[c:], "description : ", " ,", 0),
				ImageUrl:    strings.ReplaceAll(parse.Find(pageStr[c:], "image_url : ", " ", 0), "\\", ""),
			})
		}
	}

	result.CommentsCount, left = parse.FindWithIndex(pageStr, "comments_count :", ",", left, -1)
	result.FollowerCount, left = parse.FindWithIndex(pageStr, "follower_count :", ",", left, -1)

	// Принадлежность к группам
	for c := parse.Index(pageStr, "groups :[", left, -1); pageStr[c] != ']'; c++ {
		if pageStr[c] == '{' {
			result.Groups += parse.Find(pageStr[c:], "name : ", " ,", 0) + ", "
		}
	}
	if result.Groups != "" {
		result.Groups = result.Groups[:len(result.Groups)-2]
	}

	result.MappingFollowerCount, left = parse.FindWithIndex(pageStr, "mapping_follower_count :", ",", left, -1)
	result.PendingBeatmapsetCount, left = parse.FindWithIndex(pageStr, "pending_beatmapset_count :", ",", left, -1)
	result.Names = strings.Split(parse.Find(pageStr, "previous_usernames :[ ", " ],", left), " , ")
	if result.Names[0] == "" {
		result.Names = nil
	}
	result.RankHighest, left = parse.FindWithIndex(pageStr, "rank_highest :{ rank :", ",", left, -1)
	result.Count100, left = parse.FindWithIndex(pageStr, "count_100 :", ",", left, -1)
	result.Count300, left = parse.FindWithIndex(pageStr, "count_300 :", ",", left, -1)
	result.Count50, left = parse.FindWithIndex(pageStr, "count_50 :", ",", left, -1)
	result.CountMiss, left = parse.FindWithIndex(pageStr, "count_miss :", ",", left, -1)
	result.Level, left = parse.FindWithIndex(pageStr, "level :{ current :", ",", left, -1)
	result.GlobalRank, left = parse.FindWithIndex(pageStr, "global_rank :", ",", left, -1)
	result.PP, left = parse.FindWithIndex(pageStr, "pp :", ",", left, -1)
	result.PPExp, left = parse.FindWithIndex(pageStr, "pp_exp :", ",", left, -1)
	result.RankedScore, left = parse.FindWithIndex(pageStr, "ranked_score :", ",", left, -1)
	result.Accuracy, left = parse.FindWithIndex(pageStr, "hit_accuracy :", ",", left, -1)
	result.PlayCount, left = parse.FindWithIndex(pageStr, "play_count :", ",", left, -1)
	result.PlayTimeSeconds, left = parse.FindWithIndex(pageStr, "play_time :", ",", left, -1)
	duration, _ := time.ParseDuration(result.PlayTimeSeconds + "s")
	result.PlayTime = duration.String()
	result.TotalScore, left = parse.FindWithIndex(pageStr, "total_score :", ",", left, -1)
	result.TotalHits, left = parse.FindWithIndex(pageStr, "total_hits :", ",", left, -1)
	result.MaximumCombo, left = parse.FindWithIndex(pageStr, "maximum_combo :", ",", left, -1)
	result.Replays, left = parse.FindWithIndex(pageStr, "replays_watched_by_others :", ",", left, -1)
	result.IsRanked, left = parse.FindWithIndex(pageStr, "is_ranked :", ",", left, -1)
	result.SS, left = parse.FindWithIndex(pageStr, "grade_counts :{ ss :", ",", left, -1)
	result.SSH, left = parse.FindWithIndex(pageStr, "ssh :", ",", left, -1)
	result.S, left = parse.FindWithIndex(pageStr, "s :", ",", left, -1)
	result.SH, left = parse.FindWithIndex(pageStr, "sh :", ",", left, -1)
	result.A, left = parse.FindWithIndex(pageStr, "a :", "}", left, -1)
	result.CountryRank, left = parse.FindWithIndex(pageStr, "country_rank :", ",", left, -1)
	result.SupportLvl, left = parse.FindWithIndex(pageStr, "support_level :", ",", left, -1)

	// Проверка на наличие достижений
	if !parse.Contains(pageStr, "user_achievements :[]", left) {

		// Конец блока достижений
		end := parse.Index(pageStr, "]", left, -1) - 10
		medals := 0

		// Цикл обработки достижений
		for ; left < end; medals++ {

			// Инициализация достижения
			var achieve achievementString

			// Генерация достижения
			achieve.AchievedAt, left = parse.FindWithIndex(pageStr, "achieved_at : ", " ,", left, -1)
			achieve.AchievementId, left = parse.FindWithIndex(pageStr, "achievement_id :", "}", left, -1)

			// Добавление достижения
			result.Achievements = append(result.Achievements, achieve)

		}

		// Запись количества медалей
		result.Medals = strconv.Itoa(medals)

	}

	// Проверка на наличие статистики
	if !parse.Contains(pageStr, " rank_history :null", left) {
		result.RankHistory.Mode, left = parse.FindWithIndex(pageStr, "mode : ", " ,", left, -1)
		result.RankHistory.Data = strings.Split(parse.Find(pageStr, "data :[", "]", left), ",")
	}

	result.UnrankedBeatmapsetCount, _ = parse.FindWithIndex(pageStr, "unranked_beatmapset_count :", "}", left, -1)

	return result, http.StatusOK, nil

}

// getUserInfo возвращает структуру с информацией пользователя, статус код и ошибку
func getUserInfo(id string) (userInfo, int, error) {

	// Получение текстовой версии статистики
	resultStr, statusCode, err := getUserInfoString(id)
	if err != nil {
		return userInfo{}, statusCode, err
	}

	// Перевод в классическую версию
	result := userInfo{
		AvatarUrl:     resultStr.AvatarUrl,
		CountryCode:   resultStr.CountryCode,
		DefaultGroup:  resultStr.DefaultGroup,
		UserID:        convert.ToInt(resultStr.UserID),
		IsActive:      convert.ToBool(resultStr.IsActive),
		IsBot:         convert.ToBool(resultStr.IsBot),
		IsDeleted:     convert.ToBool(resultStr.IsDeleted),
		IsOnline:      convert.ToBool(resultStr.IsOnline),
		IsSupporter:   convert.ToBool(resultStr.IsSupporter),
		LastVisit:     resultStr.LastVisit,
		PmFriendsOnly: convert.ToBool(resultStr.PmFriendsOnly),
		ProfileColor:  resultStr.ProfileColor,
		Username:      resultStr.Username,
		CoverUrl:      resultStr.CoverUrl,
		Discord:       resultStr.Discord,
		HasSupported:  convert.ToBool(resultStr.HasSupported),
		Interests:     resultStr.Interests,
		JoinDate:      resultStr.JoinDate,
		Kudosu:        convert.ToInt(resultStr.Kudosu),
		Location:      resultStr.Location,
		MaxFriends:    convert.ToInt(resultStr.MaxFriends),
		MaxBLock:      convert.ToInt(resultStr.MaxBLock),
		Occupation:    resultStr.Occupation,
		Playmode:      resultStr.Playmode,
		Playstyle:     resultStr.Playstyle,
		PostCount:     convert.ToInt(resultStr.PostCount),
		ProfileOrder:  resultStr.ProfileOrder,
		Title:         resultStr.Title,
		TitleUrl:      resultStr.TitleUrl,
		Twitter:       resultStr.Twitter,
		Website:       resultStr.Website,
		CountyName:    resultStr.CountyName,
		UserCover: cover{
			CustomUrl: resultStr.UserCover.CustomUrl,
			Url:       resultStr.UserCover.Url,
			Id:        convert.ToInt(resultStr.UserCover.Id),
		},
		IsAdmin:                 convert.ToBool(resultStr.IsAdmin),
		IsBng:                   convert.ToBool(resultStr.IsBng),
		IsFullBan:               convert.ToBool(resultStr.IsFullBan),
		IsGmt:                   convert.ToBool(resultStr.IsGmt),
		IsLimitedBan:            convert.ToBool(resultStr.IsLimitedBan),
		IsModerator:             convert.ToBool(resultStr.IsModerator),
		IsNat:                   convert.ToBool(resultStr.IsNat),
		IsRestricted:            convert.ToBool(resultStr.IsRestricted),
		IsSilenced:              convert.ToBool(resultStr.IsSilenced),
		AccountHistory:          resultStr.AccountHistory,
		ActiveTournamentBanner:  resultStr.ActiveTournamentBanner,
		Badges:                  resultStr.Badges,
		CommentsCount:           convert.ToInt(resultStr.CommentsCount),
		FollowerCount:           convert.ToInt(resultStr.FollowerCount),
		Groups:                  resultStr.Groups,
		MappingFollowerCount:    convert.ToInt(resultStr.MappingFollowerCount),
		PendingBeatmapsetCount:  convert.ToInt(resultStr.PendingBeatmapsetCount),
		Names:                   resultStr.Names,
		RankHighest:             convert.ToInt(resultStr.RankHighest),
		Count100:                convert.ToInt(resultStr.Count100),
		Count300:                convert.ToInt(resultStr.Count300),
		Count50:                 convert.ToInt(resultStr.Count50),
		CountMiss:               convert.ToInt(resultStr.CountMiss),
		Level:                   convert.ToInt(resultStr.Level),
		GlobalRank:              convert.ToInt64(resultStr.GlobalRank),
		PP:                      convert.ToFloat64(resultStr.PP),
		PPExp:                   convert.ToInt(resultStr.PPExp),
		RankedScore:             convert.ToInt(resultStr.RankedScore),
		Accuracy:                convert.ToFloat64(resultStr.Accuracy),
		PlayCount:               convert.ToInt(resultStr.PlayCount),
		PlayTime:                resultStr.PlayTime,
		PlayTimeSeconds:         convert.ToInt64(resultStr.PlayTimeSeconds),
		TotalScore:              convert.ToInt64(resultStr.TotalScore),
		TotalHits:               convert.ToInt64(resultStr.TotalHits),
		MaximumCombo:            convert.ToInt(resultStr.MaximumCombo),
		Replays:                 convert.ToInt(resultStr.Replays),
		IsRanked:                convert.ToBool(resultStr.IsRanked),
		SS:                      convert.ToInt(resultStr.SS),
		SSH:                     convert.ToInt(resultStr.SSH),
		S:                       convert.ToInt(resultStr.S),
		SH:                      convert.ToInt(resultStr.SH),
		A:                       convert.ToInt(resultStr.A),
		CountryRank:             convert.ToInt(resultStr.CountryRank),
		SupportLvl:              convert.ToInt(resultStr.SupportLvl),
		UnrankedBeatmapsetCount: convert.ToInt(resultStr.UnrankedBeatmapsetCount),
	}

	// Перевод достижений
	for _, c := range resultStr.Achievements {
		result.Achievements = append(result.Achievements, achievement{
			AchievedAt:    c.AchievedAt,
			AchievementId: convert.ToInt(c.AchievementId),
		})
	}

	// Перевод количества медалей и истории рейтинга
	result.Medals = convert.ToInt(resultStr.Medals)
	result.RankHistory.Mode = resultStr.RankHistory.Mode

	for _, d := range resultStr.RankHistory.Data {
		result.RankHistory.Data = append(result.RankHistory.Data, convert.ToInt(d))
	}

	return result, http.StatusOK, nil

}

// Response отправляет ответ на реквест
func response(w http.ResponseWriter, statusCode int, body any) {

	// Установка заголовков
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Установка статускода и запись тела респонса
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)

}

// User - роут "/user"
func User(w http.ResponseWriter, r *http.Request) {

	// Получение параметра id из реквеста
	id := r.URL.Query().Get("id")

	// Проверка на наличие параметра
	if id == "" {
		response(w, http.StatusBadRequest, apiError{Error: "please insert user id"})
		return
	}

	// Проверка на тип
	if r.URL.Query().Get("type") == "string" {

		// Получение статистики
		result, statusCode, err := getUserInfoString(id)
		if err != nil {
			response(w, statusCode, apiError{Error: err.Error()})
			return
		}

		response(w, statusCode, result)

	} else {

		// Получение статистики
		result, statusCode, err := getUserInfo(id)
		if err != nil {
			response(w, statusCode, apiError{Error: err.Error()})
			return
		}

		response(w, statusCode, result)

	}

}
