package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"osustatsapi/utils"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// ---------------------- Классические структуры ------------------------

// Информация о пользователе
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

// Ковёр пользователя
type cover struct {
	CustomUrl string `json:"custom_url"`
	Url       string `json:"url"`
	Id        int    `json:"id"`
}

// Значок профиля
type badge struct {
	AwardedAt   string `json:"awarded_at"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

// Достижение
type achievement struct {
	AchievedAt    string `json:"achieved_at"`
	AchievementId int    `json:"achievement_id"`
}

// История рейтинга
type history struct {
	Mode string `json:"mode"`
	Data []int  `json:"data"`
}

// ---------------------- Структуры для парсинга ------------------------

// Информация о пользователе
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

// Ковёр пользователя
type coverString struct {
	CustomUrl string `json:"custom_url"`
	Url       string `json:"url"`
	Id        string `json:"id"`
}

// Достижение
type achievementString struct {
	AchievedAt    string `json:"achieved_at"`
	AchievementId string `json:"achievement_id"`
}

// История рейтинга
type historyString struct {
	Mode string   `json:"mode"`
	Data []string `json:"data"`
}

// Структура ошибки
type apiError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// ----------------- Функции получения статистики ----------------

// Функция получения текстовой информации о пользователе
func getUserInfoString(id string) (userInfoString, error) {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return userInfoString{}, fmt.Errorf("in http.Get: %w", err)
	}
	defer resp.Body.Close()

	// Проверка статускода
	if resp.StatusCode != 200 {
		return userInfoString{}, fmt.Errorf("response status: %s", resp.Status)
	}

	// Запись респонса
	body, _ := ioutil.ReadAll(resp.Body)

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

	result.AvatarUrl, left = utils.FindWithIndex(pageStr, "avatar_url : ", " ", left, -1)
	result.AvatarUrl = strings.ReplaceAll(result.AvatarUrl, "\\", "")
	result.CountryCode, left = utils.FindWithIndex(pageStr, "country_code : ", " ", left, -1)
	result.DefaultGroup, left = utils.FindWithIndex(pageStr, "default_group : ", " ", left, -1)
	result.UserID, left = utils.FindWithIndex(pageStr, " id :", ",", left, -1)
	result.IsActive, left = utils.FindWithIndex(pageStr, "is_active :", ",", left, -1)
	result.IsBot, left = utils.FindWithIndex(pageStr, "is_bot :", ",", left, -1)
	result.IsDeleted, left = utils.FindWithIndex(pageStr, "is_deleted :", ",", left, -1)
	result.IsOnline, left = utils.FindWithIndex(pageStr, "is_online :", ",", left, -1)
	result.IsSupporter, left = utils.FindWithIndex(pageStr, "is_supporter :", ",", left, -1)
	result.LastVisit, left = utils.FindWithIndex(pageStr, "last_visit : ", " ", left, -1)
	result.PmFriendsOnly, left = utils.FindWithIndex(pageStr, "pm_friends_only :", ",", left, -1)
	result.ProfileColor, left = utils.FindWithIndex(pageStr, "profile_colour : ", " ,", left, -1)
	result.Username, left = utils.FindWithIndex(pageStr, "username : ", " ", left, -1)
	result.CoverUrl, left = utils.FindWithIndex(pageStr, "cover_url : ", " ", left, -1)
	result.CoverUrl = strings.ReplaceAll(result.CoverUrl, "\\", "")
	result.Discord, left = utils.FindWithIndex(pageStr, "discord : ", " ,", left, -1)
	result.HasSupported, left = utils.FindWithIndex(pageStr, "has_supported :", ",", left, -1)
	result.Interests, left = utils.FindWithIndex(pageStr, "interests : ", " , join_date", left, -1)
	result.JoinDate, left = utils.FindWithIndex(pageStr, "join_date : ", " ,", left, -1)
	result.Kudosu, left = utils.FindWithIndex(pageStr, "kudosu :{ total :", ",", left, -1)
	result.Location, left = utils.FindWithIndex(pageStr, "location : ", " ,", left, -1)
	result.MaxBLock, left = utils.FindWithIndex(pageStr, "max_blocks :", ",", left, -1)
	result.MaxFriends, left = utils.FindWithIndex(pageStr, "max_friends :", ",", left, -1)
	result.Occupation, left = utils.FindWithIndex(pageStr, "occupation : ", " ,", left, -1)
	result.Playmode, left = utils.FindWithIndex(pageStr, "playmode : ", " ,", left, -1)
	result.Playstyle = strings.Split(utils.Find(pageStr, "playstyle :[ ", " ],", left), " , ")
	if result.Playstyle[0] == "" {
		result.Playstyle = nil
	}
	result.PostCount, left = utils.FindWithIndex(pageStr, "post_count :", ",", left, -1)
	result.ProfileOrder = strings.Split(utils.Find(pageStr, "profile_order :[ ", " ]", left), " , ")
	if result.ProfileOrder[0] == "" {
		result.ProfileOrder = nil
	}
	result.Title, left = utils.FindWithIndex(pageStr, "title :", ",", left, -1)
	result.TitleUrl, left = utils.FindWithIndex(pageStr, "title_url : ", " ,", left, -1)
	result.Twitter, left = utils.FindWithIndex(pageStr, "twitter : ", " ,", left, -1)
	result.Website, left = utils.FindWithIndex(pageStr, "website : ", " ,", left, -1)
	result.Website = strings.ReplaceAll(result.Website, "\\", "")
	result.CountyName, left = utils.FindWithIndex(pageStr, " name : ", " }", left, -1)

	result.UserCover.CustomUrl, left = utils.FindWithIndex(pageStr, "custom_url : ", " ,", left, -1)
	result.UserCover.CustomUrl = strings.ReplaceAll(result.UserCover.CustomUrl, "\\", "")
	result.UserCover.Url, left = utils.FindWithIndex(pageStr, "url : ", " ,", left, -1)
	result.UserCover.Url = strings.ReplaceAll(result.UserCover.Url, "\\", "")
	result.UserCover.Id, left = utils.FindWithIndex(pageStr, " , id : ", " }", left, -1)

	result.IsAdmin, left = utils.FindWithIndex(pageStr, "is_admin :", ",", left, -1)
	result.IsBng, left = utils.FindWithIndex(pageStr, "is_bng :", ",", left, -1)
	result.IsFullBan, left = utils.FindWithIndex(pageStr, "is_full_bn :", ",", left, -1)
	result.IsGmt, left = utils.FindWithIndex(pageStr, "is_gmt :", ",", left, -1)
	result.IsLimitedBan, left = utils.FindWithIndex(pageStr, "is_limited_bn :", ",", left, -1)
	result.IsModerator, left = utils.FindWithIndex(pageStr, "is_moderator :", ",", left, -1)
	result.IsNat, left = utils.FindWithIndex(pageStr, "is_nat :", ",", left, -1)
	result.IsRestricted, left = utils.FindWithIndex(pageStr, "is_restricted :", ",", left, -1)
	result.IsSilenced, left = utils.FindWithIndex(pageStr, "is_silenced :", ",", left, -1)
	result.ActiveTournamentBanner, left = utils.FindWithIndex(pageStr, "active_tournament_banner :", ", badges", left, -1)
	result.ActiveTournamentBanner = strings.ReplaceAll(result.ActiveTournamentBanner, "\\", "")

	// Значки
	for c := utils.Index(pageStr, "badges :[", left, -1); pageStr[c] != ']'; c++ {
		if pageStr[c:c+13] == "awarded_at : " {
			result.Badges = append(result.Badges, badge{
				AwardedAt:   utils.Find(pageStr[c:], "awarded_at : ", " ", 0),
				Description: utils.Find(pageStr[c:], "description : ", " ,", 0),
				ImageUrl:    strings.ReplaceAll(utils.Find(pageStr[c:], "image_url : ", " ", 0), "\\", ""),
			})
		}
	}

	result.CommentsCount, left = utils.FindWithIndex(pageStr, "comments_count :", ",", left, -1)
	result.FollowerCount, left = utils.FindWithIndex(pageStr, "follower_count :", ",", left, -1)

	// Принадлежность к группам
	for c := utils.Index(pageStr, "groups :[", left, -1); pageStr[c] != ']'; c++ {
		if pageStr[c] == '{' {
			result.Groups += utils.Find(pageStr[c:], "name : ", " ,", 0) + ", "
		}
	}
	if result.Groups != "" {
		result.Groups = result.Groups[:len(result.Groups)-2]
	}

	result.MappingFollowerCount, left = utils.FindWithIndex(pageStr, "mapping_follower_count :", ",", left, -1)
	result.PendingBeatmapsetCount, left = utils.FindWithIndex(pageStr, "pending_beatmapset_count :", ",", left, -1)
	result.Names = strings.Split(utils.Find(pageStr, "previous_usernames :[ ", " ],", left), " , ")
	if result.Names[0] == "" {
		result.Names = nil
	}
	result.RankHighest, left = utils.FindWithIndex(pageStr, "rank_highest :{ rank :", ",", left, -1)
	result.Count100, left = utils.FindWithIndex(pageStr, "count_100 :", ",", left, -1)
	result.Count300, left = utils.FindWithIndex(pageStr, "count_300 :", ",", left, -1)
	result.Count50, left = utils.FindWithIndex(pageStr, "count_50 :", ",", left, -1)
	result.CountMiss, left = utils.FindWithIndex(pageStr, "count_miss :", ",", left, -1)
	result.Level, left = utils.FindWithIndex(pageStr, "level :{ current :", ",", left, -1)
	result.GlobalRank, left = utils.FindWithIndex(pageStr, "global_rank :", ",", left, -1)
	result.PP, left = utils.FindWithIndex(pageStr, "pp :", ",", left, -1)
	result.PPExp, left = utils.FindWithIndex(pageStr, "pp_exp :", ",", left, -1)
	result.RankedScore, left = utils.FindWithIndex(pageStr, "ranked_score :", ",", left, -1)
	result.Accuracy, left = utils.FindWithIndex(pageStr, "hit_accuracy :", ",", left, -1)
	result.PlayCount, left = utils.FindWithIndex(pageStr, "play_count :", ",", left, -1)
	result.PlayTimeSeconds, left = utils.FindWithIndex(pageStr, "play_time :", ",", left, -1)
	duration, _ := time.ParseDuration(result.PlayTimeSeconds + "s")
	result.PlayTime = duration.String()
	result.TotalScore, left = utils.FindWithIndex(pageStr, "total_score :", ",", left, -1)
	result.TotalHits, left = utils.FindWithIndex(pageStr, "total_hits :", ",", left, -1)
	result.MaximumCombo, left = utils.FindWithIndex(pageStr, "maximum_combo :", ",", left, -1)
	result.Replays, left = utils.FindWithIndex(pageStr, "replays_watched_by_others :", ",", left, -1)
	result.IsRanked, left = utils.FindWithIndex(pageStr, "is_ranked :", ",", left, -1)
	result.SS, left = utils.FindWithIndex(pageStr, "grade_counts :{ ss :", ",", left, -1)
	result.SSH, left = utils.FindWithIndex(pageStr, "ssh :", ",", left, -1)
	result.S, left = utils.FindWithIndex(pageStr, "s :", ",", left, -1)
	result.SH, left = utils.FindWithIndex(pageStr, "sh :", ",", left, -1)
	result.A, left = utils.FindWithIndex(pageStr, "a :", "}", left, -1)
	result.CountryRank, left = utils.FindWithIndex(pageStr, "country_rank :", ",", left, -1)
	result.SupportLvl, left = utils.FindWithIndex(pageStr, "support_level :", ",", left, -1)

	// Проверка на наличие достижений
	if !utils.Contains(pageStr, "user_achievements :[]", left) {

		// Конец блока достижений
		end := utils.Index(pageStr, "]", left, -1) - 10
		medals := 0

		// Цикл обработки достижений
		for ; left < end; medals++ {

			// Инициализация достижения
			var achieve achievementString

			// Генерация достижения
			achieve.AchievedAt, left = utils.FindWithIndex(pageStr, "achieved_at : ", " ,", left, -1)
			achieve.AchievementId, left = utils.FindWithIndex(pageStr, "achievement_id :", "}", left, -1)

			// Добавление достижения
			result.Achievements = append(result.Achievements, achieve)

		}

		// Запись количества медалей
		result.Medals = strconv.Itoa(medals)

	}

	// Проверка на наличие статистики
	if !utils.Contains(pageStr, " rank_history :null", left) {
		result.RankHistory.Mode, left = utils.FindWithIndex(pageStr, "mode : ", " ,", left, -1)
		result.RankHistory.Data = strings.Split(utils.Find(pageStr, "data :[", "]", left), ",")
	}

	result.UnrankedBeatmapsetCount, _ = utils.FindWithIndex(pageStr, "unranked_beatmapset_count :", "}", left, -1)

	return result, nil

}

// Функция получения информации о пользователе
func getUserInfo(id string) (userInfo, error) {

	// Получение текстовой версии статистики
	resultStr, err := getUserInfoString(id)
	if err != nil {
		return userInfo{}, err
	}

	// Перевод в классическую версию
	result := userInfo{
		AvatarUrl:     resultStr.AvatarUrl,
		CountryCode:   resultStr.CountryCode,
		DefaultGroup:  resultStr.DefaultGroup,
		UserID:        utils.ToInt(resultStr.UserID),
		IsActive:      utils.ToBool(resultStr.IsActive),
		IsBot:         utils.ToBool(resultStr.IsBot),
		IsDeleted:     utils.ToBool(resultStr.IsDeleted),
		IsOnline:      utils.ToBool(resultStr.IsOnline),
		IsSupporter:   utils.ToBool(resultStr.IsSupporter),
		LastVisit:     resultStr.LastVisit,
		PmFriendsOnly: utils.ToBool(resultStr.PmFriendsOnly),
		ProfileColor:  resultStr.ProfileColor,
		Username:      resultStr.Username,
		CoverUrl:      resultStr.CoverUrl,
		Discord:       resultStr.Discord,
		HasSupported:  utils.ToBool(resultStr.HasSupported),
		Interests:     resultStr.Interests,
		JoinDate:      resultStr.JoinDate,
		Kudosu:        utils.ToInt(resultStr.Kudosu),
		Location:      resultStr.Location,
		MaxFriends:    utils.ToInt(resultStr.MaxFriends),
		MaxBLock:      utils.ToInt(resultStr.MaxBLock),
		Occupation:    resultStr.Occupation,
		Playmode:      resultStr.Playmode,
		Playstyle:     resultStr.Playstyle,
		PostCount:     utils.ToInt(resultStr.PostCount),
		ProfileOrder:  resultStr.ProfileOrder,
		Title:         resultStr.Title,
		TitleUrl:      resultStr.TitleUrl,
		Twitter:       resultStr.Twitter,
		Website:       resultStr.Website,
		CountyName:    resultStr.CountyName,
		UserCover: cover{
			CustomUrl: resultStr.UserCover.CustomUrl,
			Url:       resultStr.UserCover.Url,
			Id:        utils.ToInt(resultStr.UserCover.Id),
		},
		IsAdmin:                 utils.ToBool(resultStr.IsAdmin),
		IsBng:                   utils.ToBool(resultStr.IsBng),
		IsFullBan:               utils.ToBool(resultStr.IsFullBan),
		IsGmt:                   utils.ToBool(resultStr.IsGmt),
		IsLimitedBan:            utils.ToBool(resultStr.IsLimitedBan),
		IsModerator:             utils.ToBool(resultStr.IsModerator),
		IsNat:                   utils.ToBool(resultStr.IsNat),
		IsRestricted:            utils.ToBool(resultStr.IsRestricted),
		IsSilenced:              utils.ToBool(resultStr.IsSilenced),
		AccountHistory:          resultStr.AccountHistory,
		ActiveTournamentBanner:  resultStr.ActiveTournamentBanner,
		Badges:                  resultStr.Badges,
		CommentsCount:           utils.ToInt(resultStr.CommentsCount),
		FollowerCount:           utils.ToInt(resultStr.FollowerCount),
		Groups:                  resultStr.Groups,
		MappingFollowerCount:    utils.ToInt(resultStr.MappingFollowerCount),
		PendingBeatmapsetCount:  utils.ToInt(resultStr.PendingBeatmapsetCount),
		Names:                   resultStr.Names,
		RankHighest:             utils.ToInt(resultStr.RankHighest),
		Count100:                utils.ToInt(resultStr.Count100),
		Count300:                utils.ToInt(resultStr.Count300),
		Count50:                 utils.ToInt(resultStr.Count50),
		CountMiss:               utils.ToInt(resultStr.CountMiss),
		Level:                   utils.ToInt(resultStr.Level),
		GlobalRank:              utils.ToInt64(resultStr.GlobalRank),
		PP:                      utils.ToFloat64(resultStr.PP),
		PPExp:                   utils.ToInt(resultStr.PPExp),
		RankedScore:             utils.ToInt(resultStr.RankedScore),
		Accuracy:                utils.ToFloat64(resultStr.Accuracy),
		PlayCount:               utils.ToInt(resultStr.PlayCount),
		PlayTime:                resultStr.PlayTime,
		PlayTimeSeconds:         utils.ToInt64(resultStr.PlayTimeSeconds),
		TotalScore:              utils.ToInt64(resultStr.TotalScore),
		TotalHits:               utils.ToInt64(resultStr.TotalHits),
		MaximumCombo:            utils.ToInt(resultStr.MaximumCombo),
		Replays:                 utils.ToInt(resultStr.Replays),
		IsRanked:                utils.ToBool(resultStr.IsRanked),
		SS:                      utils.ToInt(resultStr.SS),
		SSH:                     utils.ToInt(resultStr.SSH),
		S:                       utils.ToInt(resultStr.S),
		SH:                      utils.ToInt(resultStr.SH),
		A:                       utils.ToInt(resultStr.A),
		CountryRank:             utils.ToInt(resultStr.CountryRank),
		SupportLvl:              utils.ToInt(resultStr.SupportLvl),
		UnrankedBeatmapsetCount: utils.ToInt(resultStr.UnrankedBeatmapsetCount),
	}

	// Перевод достижений
	for _, c := range resultStr.Achievements {
		result.Achievements = append(result.Achievements, achievement{
			AchievedAt:    c.AchievedAt,
			AchievementId: utils.ToInt(c.AchievementId),
		})
	}

	// Перевод количества медалей и истории рейтинга
	result.Medals = utils.ToInt(resultStr.Medals)
	result.RankHistory.Mode = resultStr.RankHistory.Mode

	for _, d := range resultStr.RankHistory.Data {
		result.RankHistory.Data = append(result.RankHistory.Data, utils.ToInt(d))
	}

	return result, nil

}

// Роут "/user"
func User(w http.ResponseWriter, r *http.Request) {

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Получение параметра id из реквеста
	id := r.URL.Query().Get("id")

	// Проверка на наличие параметра
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json, _ := json.Marshal(apiError{Error: "please insert user id"})
		w.Write(json)
		return
	}

	// Проверка на тип
	if r.URL.Query().Get("type") == "string" {

		// Получение статистики
		result, err := getUserInfoString(id)
		if err != nil {
			if err.Error() == "response status: 404 Not Found" {
				w.WriteHeader(http.StatusNotFound)
				json, _ := json.Marshal(apiError{Error: "not found"})
				w.Write(json)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: "internal server error"})
			w.Write(json)
			logrus.Printf("getUserInfo err: %s", err)
			return
		}

		// Перевод в json
		jsonResp, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: "internal server error"})
			w.Write(json)
			logrus.Printf("json.Marshal err: %s", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)

	} else {

		// Получение статистики
		result, err := getUserInfo(id)
		if err != nil {
			if err.Error() == "response status: 404 Not Found" {
				w.WriteHeader(http.StatusNotFound)
				json, _ := json.Marshal(apiError{Error: "not found"})
				w.Write(json)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: "internal server error"})
			w.Write(json)
			logrus.Printf("getUserInfo err: %s", err)
			return
		}

		// Перевод в json
		jsonResp, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: "internal server error"})
			w.Write(json)
			logrus.Printf("json.Marshal err: %s", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)

	}

}
