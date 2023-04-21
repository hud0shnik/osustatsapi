package handler2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ---------------------- Классические структуры ------------------------

// Информация о пользователе
type userInfo struct {
	Success                 bool          `json:"success"`
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
	Success                 bool                `json:"success"`
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

// ---------------------- Функции поиска ------------------------

// Функция поиска. Возвращает искомое значение и индекс последнего символа
func findWithIndex(str, subStr, stopChar string, start, end int) (string, int) {

	// Обрезка левой границы поиска
	str = str[start:]

	// Поиск индекса начала нужной строки
	left := strings.Index(str, subStr) + len(subStr)

	// Проверка на существование нужной строки и попадание в диапазон
	if left != len(subStr)-1 && ((end == -1) || (left+start < end)) {

		// Поиск и проверка правой границы
		right := strings.Index(str[left:], stopChar)
		if right == -1 {
			return "", start
		}

		// Обрезка и вывод результата
		return str[left : left+right], right + left + start
	}

	// Вывод ненайденных значений для тестов
	// fmt.Println("error foundn't \t", subStr, "-")

	return "", start
}

// Функция поиска. Возвращает искомое значение без кавычек и индекс последнего символа
func findStringWithIndex(str, subStr, stopChar string, start, end int) (string, int) {

	// Обрезка левой границы поиска
	str = str[start:]

	// Поиск индекса начала нужной строки
	left := strings.Index(str, subStr) + len(subStr)

	// Проверка на существование нужной строки и попадание в диапазон
	if left != len(subStr)-1 && ((end == -1) || (left+start < end)) {

		// Поиск и проверка правой границы
		right := strings.Index(str[left:], stopChar)
		if right == -1 {
			return "", start
		}

		// Обрезка и вывод результата
		return strings.ReplaceAll(str[left:left+right], "\"", ""), right + left + start
	}

	// Вывод ненайденных значений для тестов
	// fmt.Println("error foundn't \t", subStr, "-")

	return "", start
}

// Облегчённая функция поиска. Возвращает только искомое значение
func find(str, subStr, stopChar string, start int) string {

	// Обрезка левой границы поиска
	str = str[start:]

	// Поиск индекса начала нужной строки
	left := strings.Index(str, subStr)

	// Проверка на существование нужной строки
	if left != -1 {

		// Обрезка левой части
		str = str[left+len(subStr):]

		// Поиск и проверка правой границы
		right := strings.Index(str, stopChar)
		if right == -1 {
			return ""
		}

		// Обрезка правой части и вывод результата
		return str[:right]
	}

	return ""
}

// Функция поиска индекса
func index(str, subStr string, start, end int) int {

	res := strings.Index(str[start:], subStr)

	// Проверка на существование нужной строки в диапазоне
	if res != -1 && ((end == -1) || (res+start < end)) {

		//fmt.Println(res+start, " - ", subStr)
		return res + start
	}

	//fmt.Println("index error: \t", subStr)
	return -1
}

// Функция проверки наличия подстроки
func contains(str, subStr string, left int) bool {

	return strings.Contains(str[left:], subStr)
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

// ----------------- Функции получения статистики ----------------

// Функция получения текстовой информации о пользователе
func getUserInfoString(id string) userInfoString {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return userInfoString{
			Success: false,
			Error:   "can't reach osu.ppy.sh",
		}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// HTML полученной страницы в формате string
	pageStr := string(body)[80000:]

	// Проверка на страницу пользователя
	if strings.Contains(pageStr, "<h1>User not found! ;_;</h1>") {
		return userInfoString{
			Success: false,
			Error:   "not found",
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
	result := userInfoString{
		Success: true,
	}

	// Крайняя левая граница поиска
	left := 0

	//--------------------------- Статистика игрока ------------------------------

	result.AvatarUrl, left = findWithIndex(pageStr, "avatar_url : ", " ", left, -1)
	result.AvatarUrl = strings.ReplaceAll(result.AvatarUrl, "\\", "")
	result.CountryCode, left = findWithIndex(pageStr, "country_code : ", " ", left, -1)
	result.DefaultGroup, left = findWithIndex(pageStr, "default_group : ", " ", left, -1)
	result.UserID, left = findWithIndex(pageStr, " id :", ",", left, -1)
	result.IsActive, left = findWithIndex(pageStr, "is_active :", ",", left, -1)
	result.IsBot, left = findWithIndex(pageStr, "is_bot :", ",", left, -1)
	result.IsDeleted, left = findWithIndex(pageStr, "is_deleted :", ",", left, -1)
	result.IsOnline, left = findWithIndex(pageStr, "is_online :", ",", left, -1)
	result.IsSupporter, left = findWithIndex(pageStr, "is_supporter :", ",", left, -1)
	result.LastVisit, left = findWithIndex(pageStr, "last_visit : ", " ", left, -1)
	result.PmFriendsOnly, left = findWithIndex(pageStr, "pm_friends_only :", ",", left, -1)
	result.ProfileColor, left = findWithIndex(pageStr, "profile_colour : ", " ,", left, -1)
	result.Username, left = findWithIndex(pageStr, "username : ", " ", left, -1)
	result.CoverUrl, left = findWithIndex(pageStr, "cover_url : ", " ", left, -1)
	result.CoverUrl = strings.ReplaceAll(result.CoverUrl, "\\", "")
	result.Discord, left = findWithIndex(pageStr, "discord : ", " ,", left, -1)
	result.HasSupported, left = findWithIndex(pageStr, "has_supported :", ",", left, -1)
	result.Interests, left = findWithIndex(pageStr, "interests : ", " , join_date", left, -1)
	result.JoinDate, left = findWithIndex(pageStr, "join_date : ", " ,", left, -1)
	result.Kudosu, left = findWithIndex(pageStr, "kudosu :{ total :", ",", left, -1)
	result.Location, left = findWithIndex(pageStr, "location : ", " ,", left, -1)
	result.MaxBLock, left = findWithIndex(pageStr, "max_blocks :", ",", left, -1)
	result.MaxFriends, left = findWithIndex(pageStr, "max_friends :", ",", left, -1)
	result.Occupation, left = findWithIndex(pageStr, "occupation : ", " ,", left, -1)
	result.Playmode, left = findWithIndex(pageStr, "playmode : ", " ,", left, -1)
	result.Playstyle = strings.Split(find(pageStr, "playstyle :[ ", " ],", left), " , ")
	if result.Playstyle[0] == "" {
		result.Playstyle = nil
	}
	result.PostCount, left = findWithIndex(pageStr, "post_count :", ",", left, -1)
	result.ProfileOrder = strings.Split(find(pageStr, "profile_order :[ ", " ]", left), " , ")
	if result.ProfileOrder[0] == "" {
		result.ProfileOrder = nil
	}
	result.Title, left = findWithIndex(pageStr, "title :", ",", left, -1)
	result.TitleUrl, left = findWithIndex(pageStr, "title_url : ", " ,", left, -1)
	result.Twitter, left = findWithIndex(pageStr, "twitter : ", " ,", left, -1)
	result.Website, left = findWithIndex(pageStr, "website : ", " ,", left, -1)
	result.Website = strings.ReplaceAll(result.Website, "\\", "")
	result.CountyName, left = findWithIndex(pageStr, " name : ", " }", left, -1)

	result.UserCover.CustomUrl, left = findWithIndex(pageStr, "custom_url : ", " ,", left, -1)
	result.UserCover.CustomUrl = strings.ReplaceAll(result.UserCover.CustomUrl, "\\", "")
	result.UserCover.Url, left = findWithIndex(pageStr, "url : ", " ,", left, -1)
	result.UserCover.Url = strings.ReplaceAll(result.UserCover.Url, "\\", "")
	result.UserCover.Id, left = findWithIndex(pageStr, " , id : ", " }", left, -1)

	result.IsAdmin, left = findWithIndex(pageStr, "is_admin :", ",", left, -1)
	result.IsBng, left = findWithIndex(pageStr, "is_bng :", ",", left, -1)
	result.IsFullBan, left = findWithIndex(pageStr, "is_full_bn :", ",", left, -1)
	result.IsGmt, left = findWithIndex(pageStr, "is_gmt :", ",", left, -1)
	result.IsLimitedBan, left = findWithIndex(pageStr, "is_limited_bn :", ",", left, -1)
	result.IsModerator, left = findWithIndex(pageStr, "is_moderator :", ",", left, -1)
	result.IsNat, left = findWithIndex(pageStr, "is_nat :", ",", left, -1)
	result.IsRestricted, left = findWithIndex(pageStr, "is_restricted :", ",", left, -1)
	result.IsSilenced, left = findWithIndex(pageStr, "is_silenced :", ",", left, -1)
	result.ActiveTournamentBanner, left = findWithIndex(pageStr, "active_tournament_banner :", ", badges", left, -1)
	result.ActiveTournamentBanner = strings.ReplaceAll(result.ActiveTournamentBanner, "\\", "")

	// Значки
	for c := index(pageStr, "badges :[", left, -1); pageStr[c] != ']'; c++ {
		if pageStr[c:c+13] == "awarded_at : " {
			result.Badges = append(result.Badges, badge{
				AwardedAt:   find(pageStr[c:], "awarded_at : ", " ", 0),
				Description: find(pageStr[c:], "description : ", " ,", 0),
				ImageUrl:    strings.ReplaceAll(find(pageStr[c:], "image_url : ", " ", 0), "\\", ""),
			})
		}
	}

	result.CommentsCount, left = findWithIndex(pageStr, "comments_count :", ",", left, -1)
	result.FollowerCount, left = findWithIndex(pageStr, "follower_count :", ",", left, -1)

	// Принадлежность к группам
	for c := index(pageStr, "groups :[", left, -1); pageStr[c] != ']'; c++ {
		if pageStr[c] == '{' {
			result.Groups += find(pageStr[c:], "name : ", " ,", 0) + ", "
		}
	}
	if result.Groups != "" {
		result.Groups = result.Groups[:len(result.Groups)-2]
	}

	result.MappingFollowerCount, left = findWithIndex(pageStr, "mapping_follower_count :", ",", left, -1)
	result.PendingBeatmapsetCount, left = findWithIndex(pageStr, "pending_beatmapset_count :", ",", left, -1)
	result.Names = strings.Split(find(pageStr, "previous_usernames :[ ", " ],", left), " , ")
	if result.Names[0] == "" {
		result.Names = nil
	}
	result.Level, left = findWithIndex(pageStr, "level :{ current :", ",", left, -1)
	result.GlobalRank, left = findWithIndex(pageStr, "global_rank :", ",", left, -1)
	result.PP, left = findWithIndex(pageStr, "pp :", ",", left, -1)
	result.RankedScore, left = findWithIndex(pageStr, "ranked_score :", ",", left, -1)
	result.Accuracy, left = findWithIndex(pageStr, "hit_accuracy :", ",", left, -1)
	result.PlayCount, left = findWithIndex(pageStr, "play_count :", ",", left, -1)
	result.PlayTimeSeconds, left = findWithIndex(pageStr, "play_time :", ",", left, -1)
	duration, _ := time.ParseDuration(result.PlayTimeSeconds + "s")
	result.PlayTime = duration.String()
	result.TotalScore, left = findWithIndex(pageStr, "total_score :", ",", left, -1)
	result.TotalHits, left = findWithIndex(pageStr, "total_hits :", ",", left, -1)
	result.MaximumCombo, left = findWithIndex(pageStr, "maximum_combo :", ",", left, -1)
	result.Replays, left = findWithIndex(pageStr, "replays_watched_by_others :", ",", left, -1)
	result.IsRanked, left = findWithIndex(pageStr, "is_ranked :", ",", left, -1)
	result.SS, left = findWithIndex(pageStr, "grade_counts :{ ss :", ",", left, -1)
	result.SSH, left = findWithIndex(pageStr, "ssh :", ",", left, -1)
	result.S, left = findWithIndex(pageStr, "s :", ",", left, -1)
	result.SH, left = findWithIndex(pageStr, "sh :", ",", left, -1)
	result.A, left = findWithIndex(pageStr, "a :", "}", left, -1)
	result.CountryRank, left = findWithIndex(pageStr, "country_rank :", ",", left, -1)
	result.SupportLvl, left = findWithIndex(pageStr, "support_level :", ",", left, -1)

	// Проверка на наличие достижений
	if !contains(pageStr, "user_achievements :[]", left) {

		// Конец блока достижений
		end := index(pageStr, "]", left, -1) - 10
		medals := 0

		// Цикл обработки достижений
		for ; left < end; medals++ {

			// Инициализация достижения
			var achieve achievementString

			// Генерация достижения
			achieve.AchievedAt, left = findWithIndex(pageStr, "achieved_at : ", " ,", left, -1)
			achieve.AchievementId, left = findWithIndex(pageStr, "achievement_id :", "}", left, -1)

			// Добавление достижения
			result.Achievements = append(result.Achievements, achieve)

		}

		// Запись количества медалей
		result.Medals = strconv.Itoa(medals)

	}

	// Проверка на наличие статистики
	if !contains(pageStr, " rank_history :null", left) {
		result.RankHistory.Mode, left = findWithIndex(pageStr, "mode : ", " ,", left, -1)
		result.RankHistory.Data = strings.Split(find(pageStr, "data :[", "]", left), ",")
	}

	result.UnrankedBeatmapsetCount, _ = findWithIndex(pageStr, "unranked_beatmapset_count :", "}", left, -1)

	return result
}

// Функция получения информации о пользователе
func getUserInfo(id string) userInfo {

	// Получение текстовой версии статистики
	resultStr := getUserInfoString(id)

	// Проверка на ошибки при парсинге
	if !resultStr.Success {
		return userInfo{
			Success: false,
			Error:   resultStr.Error,
		}
	}

	// Перевод в классическую версию
	result := userInfo{
		Success:       true,
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
		UserCover: cover{
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
		result.Achievements = append(result.Achievements, achievement{
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

	return result
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

		// Получение статистики и перевод в json
		result := getUserInfoString(id)
		jsonResp, err := json.Marshal(result)

		// Обработчик ошибок
		switch {
		case err != nil:
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: "internal server error"})
			w.Write(json)
			log.Printf("json.Marshal error: %s", err)
		case result.Error == "not found":
			w.WriteHeader(http.StatusNotFound)
			json, _ := json.Marshal(apiError{Error: "not found"})
			w.Write(json)
		case !result.Success:
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: result.Error})
			w.Write(json)
		default:
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
		}

	} else {

		// Получение статистики и перевод в json
		result := getUserInfo(id)
		jsonResp, err := json.Marshal(result)

		// Обработчик ошибок
		switch {
		case err != nil:
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: "internal server error"})
			w.Write(json)
			log.Printf("json.Marshal error: %s", err)
		case result.Error == "not found":
			w.WriteHeader(http.StatusNotFound)
			json, _ := json.Marshal(apiError{Error: "not found"})
			w.Write(json)
		case !result.Success:
			w.WriteHeader(http.StatusInternalServerError)
			json, _ := json.Marshal(apiError{Error: result.Error})
			w.Write(json)
		default:
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
		}

	}

}
