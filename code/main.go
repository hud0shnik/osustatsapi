package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Структура для хранения полной информации о пользователе
type UserInfo struct {
	Error                    string        `json:"error"`
	AvatarUrl                string        `json:"avatar_url"`
	CountryCode              string        `json:"country_code"`
	DefaultGroup             string        `json:"default_group"`
	UserID                   string        `json:"id"`
	IsActive                 string        `json:"is_active"`
	IsBot                    string        `json:"is_bot"`
	IsDeleted                string        `json:"is_deleted"`
	IsOnline                 string        `json:"is_online"`
	IsSupporter              string        `json:"is_supporter"`
	LastVisit                string        `json:"last_visit"`
	PmFriendsOnly            string        `json:"pm_friends_only"`
	ProfileColor             string        `json:"profile_color"`
	Username                 string        `json:"username"`
	CoverUrl                 string        `json:"cover_url"`
	Discord                  string        `json:"discord"`
	HasSupported             string        `json:"has_supported"`
	Interests                string        `json:"interests"`
	JoinDate                 string        `json:"join_date"`
	Kudosu                   string        `json:"kudosu"`
	Location                 string        `json:"location"`
	MaxFriends               string        `json:"max_friends"`
	MaxBLock                 string        `json:"max_block"`
	Occupation               string        `json:"occupation"`
	Playmode                 string        `json:"playmode"`
	Playstyle                string        `json:"playstyle"`
	PostCount                string        `json:"post_count"`
	ProfileOrder             string        `json:"profile_order"`
	Title                    string        `json:"title"`
	TitleUrl                 string        `json:"title_url"`
	Twitter                  string        `json:"twitter"`
	Website                  string        `json:"website"`
	CountyName               string        `json:"country_name"`
	UserCover                Cover         `json:"cover"`
	IsAdmin                  string        `json:"is_admin"`
	IsBng                    string        `json:"is_bng"`
	IsFullBan                string        `json:"is_full_bn"`
	IsGmt                    string        `json:"is_gmt"`
	IsLimitedBan             string        `json:"is_limited_bn"`
	IsModerator              string        `json:"is_moderator"`
	IsNat                    string        `json:"is_nat"`
	IsRestricted             string        `json:"is_restricted"`
	IsSilenced               string        `json:"is_silenced"`
	AccountHistory           string        `json:"account_history"`
	ActiveTournamentBanner   string        `json:"active_tournament_banner"`
	Badges                   []Badge       `json:"badges"`
	CommentsCount            string        `json:"comments_count"`
	BeatmapPlaycountsCount   string        `json:"beatmap_playcounts_count"`
	FavoriteBeatmapsetCount  string        `json:"favorite_beatmapset_count"`
	FollowerCount            string        `json:"follower_count"`
	GraveyardBeatmapsetCount string        `json:"graveyard_beatmapset_count"`
	Groups                   string        `json:"groups"`
	GuestBeatmapsetCount     string        `json:"guest_beatmapset_count"`
	LovedBeatmapsetCount     string        `json:"loved_beatmapset_count"`
	MappingFollowerCount     string        `json:"mapping_follower_count"`
	MonthlyPlaycounts        []Count       `json:"monthly_playcounts"`
	PendingBeatmapsetCount   string        `json:"pending_beatmapset_count"`
	Names                    string        `json:"previous_usernames"`
	RankedBeatmapsetCount    string        `json:"ranked_beatmapset_count"`
	ReplaysWatchedCount      []Count       `json:"replays_watched_counts"`
	ScoresBestCount          string        `json:"scores_best_count"`
	ScoresFirstCount         string        `json:"scores_first_count"`
	ScoresPinnedCount        string        `json:"scores_pinned_count"`
	ScoresRecentCount        string        `json:"scores_recent_count"`
	Level                    string        `json:"level"`
	GlobalRank               string        `json:"global_rank"`
	PP                       string        `json:"pp"`
	RankedScore              string        `json:"ranked_score"`
	Accuracy                 string        `json:"accuracy"`
	PlayCount                string        `json:"play_count"`
	PlayTime                 string        `json:"play_time"`
	PlayTimeSeconds          string        `json:"play_time_seconds"`
	TotalScore               string        `json:"total_score"`
	TotalHits                string        `json:"total_hits"`
	MaximumCombo             string        `json:"maximum_combo"`
	Replays                  string        `json:"replays"`
	IsRanked                 string        `json:"is_ranked"`
	SS                       string        `json:"ss"`
	SSH                      string        `json:"ssh"`
	S                        string        `json:"s"`
	SH                       string        `json:"sh"`
	A                        string        `json:"a"`
	CountryRank              string        `json:"country_rank"`
	SupportLvl               string        `json:"support_level"`
	Achievements             []Achievement `json:"achievements"`
	RankHistory              History       `json:"rank_history"`
	RankedAndApprovedCount   string        `json:"ranked_and_approved_beatmapset_count"`
	UnrankedBeatmapsetCount  string        `json:"unranked_beatmapset_count"`
	ScoresBest               []Score       `json:"scores_best"`
	ScoresFirst              []Score       `json:"scores_first"`
	ScoresPinned             []Score       `json:"scores_pinned"`
}

// Ковёр пользователя
type Cover struct {
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

// Структура для подсчёта
type Count struct {
	StartDate string `json:"start_date"`
	Count     string `json:"count"`
}

// Достижение
type Achievement struct {
	AchievedAt    string `json:"achieved_at"`
	AchievementId string `json:"achievement_id"`
}

// Структура для истории рейтинга
type History struct {
	Mode string `json:"mode"`
	Data string `json:"data"`
}

// Рекорд
type Score struct {
	Accuracy              string     `json:"accuracy"`
	BeatMapId             string     `json:"beatmap_id"`
	BuildId               string     `json:"build_id"`
	EndedAt               string     `json:"ended_at"`
	MaximumCombo          string     `json:"maximum_combo"`
	Mods                  []string   `json:"mods"`
	Passed                string     `json:"passed"`
	Rank                  string     `json:"rank"`
	RulesetId             string     `json:"ruleset_id"`
	StartedAt             string     `json:"started_at"`
	Statistics            string     `json:"statistics"`
	TotalScore            string     `json:"total_score"`
	UserId                string     `json:"user_id"`
	BestId                string     `json:"best_id"`
	Id                    string     `json:"id"`
	LegacyPerfect         string     `json:"legacy_perfect"`
	PP                    string     `json:"pp"`
	Replay                string     `json:"replay"`
	Type                  string     `json:"type"`
	CurrentUserAttributes string     `json:"current_user_attributes"`
	BeatMap               BeatMap    `json:"beatmap"`
	BeatMapSet            BeatMapSet `json:"beatmapset"`
	Weight                Weight     `json:"weight"`
}

// Мапа
type BeatMap struct {
	BeatMapSetId     string `json:"beatmapset_id"`
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

// Мап сет
type BeatMapSet struct {
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

// Статистика
type Weight struct {
	Percentage string `json:"percentage"`
	PP         string `json:"pp"`
}

// Статуса пользователя
type OnlineInfo struct {
	Error  string `json:"error"`
	Status string `json:"is_online"`
}

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

	return "", 0
}

// Облегчённая функция поиска. Возвращает только искомое значение
func find(str, subStr, stopChar string) string {

	// Проверка на существование нужной строки
	if strings.Contains(str, subStr) {

		// Обрезка левой части
		str = str[strings.Index(str, subStr)+len(subStr):]

		// Обрезка правой части и вывод результата
		return str[:strings.Index(str, stopChar)]
	}

	return ""
}

// Функция для парсинга рекорда
func parseScore(pageStr string, left int, scoreType string) (Score, int) {

	var result Score

	result.Accuracy, left = findWithIndex(pageStr, "accuracy :", ",", left)
	result.BeatMapId, left = findWithIndex(pageStr, "beatmap_id :", ",", left)
	result.BuildId, left = findWithIndex(pageStr, "build_id :", ",", left)
	result.EndedAt, left = findWithIndex(pageStr, "ended_at : ", " ", left)
	result.MaximumCombo, left = findWithIndex(pageStr, "max_combo :", ",", left)

	// Цикл для обработки модов
	for c := 0; pageStr[c] != ']'; c++ {
		if pageStr[c:c+10] == "acronym : " {
			result.Mods = append(result.Mods, pageStr[c+10:c+12])
		}
	}

	result.Passed, left = findWithIndex(pageStr, "passed :", ",", left)
	result.Rank, left = findWithIndex(pageStr, "rank : ", " ", left)
	result.RulesetId, left = findWithIndex(pageStr, "ruleset_id :", ",", left)
	result.StartedAt, left = findWithIndex(pageStr, "started_at :", ",", left)
	result.Statistics, left = findWithIndex(pageStr, "statistics :{ ", "}", left)
	result.TotalScore, left = findWithIndex(pageStr, "total_score :", ",", left)
	result.UserId, left = findWithIndex(pageStr, " user_id :", ",", left)
	result.BestId, left = findWithIndex(pageStr, " best_id :", ",", left)
	result.Id, left = findWithIndex(pageStr, " id :", ",", left)
	result.LegacyPerfect, left = findWithIndex(pageStr, "legacy_perfect :", ",", left)
	result.PP, left = findWithIndex(pageStr, "pp :", ",", left)
	result.Replay, left = findWithIndex(pageStr, "replay :", ",", left)
	result.Type, left = findWithIndex(pageStr, "type : ", " ", left)
	result.CurrentUserAttributes, left = findWithIndex(pageStr, "current_user_attributes :{ ", "},", left)

	result.BeatMap.BeatMapSetId, left = findWithIndex(pageStr, "beatmapset_id :", ",", left)
	result.BeatMap.DifficultyRating, left = findWithIndex(pageStr, "difficulty_rating :", ",", left)
	result.BeatMap.Id, left = findWithIndex(pageStr, "id :", ",", left)
	result.BeatMap.Mode, left = findWithIndex(pageStr, "mode : ", " ", left)
	result.BeatMap.Status, left = findWithIndex(pageStr, "status : ", " ", left)
	result.BeatMap.TotalLength, left = findWithIndex(pageStr, "total_length :", ",", left)
	result.BeatMap.UserId, left = findWithIndex(pageStr, " user_id :", ",", left)
	result.BeatMap.Version, left = findWithIndex(pageStr, "version : ", " , accuracy", left)
	result.BeatMap.Accuracy, left = findWithIndex(pageStr, "accuracy :", ",", left)
	result.BeatMap.Ar, left = findWithIndex(pageStr, "ar :", ",", left)
	result.BeatMap.Bpm, left = findWithIndex(pageStr, "bpm :", ",", left)
	result.BeatMap.Convert, left = findWithIndex(pageStr, "convert :", ",", left)
	result.BeatMap.CountCircles, left = findWithIndex(pageStr, "count_circles :", ",", left)
	result.BeatMap.CountSliders, left = findWithIndex(pageStr, "count_sliders :", ",", left)
	result.BeatMap.CountSpinners, left = findWithIndex(pageStr, "count_spinners :", ",", left)
	result.BeatMap.Cs, left = findWithIndex(pageStr, " cs :", ",", left)
	result.BeatMap.DeletedAt, left = findWithIndex(pageStr, "deleted_at :", ",", left)
	result.BeatMap.Drain, left = findWithIndex(pageStr, "drain :", ",", left)
	result.BeatMap.HitLength, left = findWithIndex(pageStr, "hit_length :", ",", left)
	result.BeatMap.IsScoreable, left = findWithIndex(pageStr, "is_scoreable :", ",", left)
	result.BeatMap.LastUpdated, left = findWithIndex(pageStr, "last_updated : ", " ", left)
	result.BeatMap.ModeInt, left = findWithIndex(pageStr, "mode_int :", ",", left)
	result.BeatMap.PassCount, left = findWithIndex(pageStr, "passcount :", ",", left)
	result.BeatMap.PlayCount, left = findWithIndex(pageStr, "playcount :", ",", left)
	result.BeatMap.Ranked, left = findWithIndex(pageStr, "ranked :", ",", left)
	result.BeatMap.Url, left = findWithIndex(pageStr, "url : ", " ", left)
	result.BeatMap.Url = strings.ReplaceAll(result.BeatMap.Url, "\\", "")
	result.BeatMap.Checksum, left = findWithIndex(pageStr, "checksum : ", " ", left)

	result.BeatMapSet.Artist, left = findWithIndex(pageStr, "artist : ", " , artist_", left)
	result.BeatMapSet.ArtistUnicode, left = findWithIndex(pageStr, "artist_unicode : ", " ,", left)

	result.BeatMapSet.Covers.Cover, left = findWithIndex(pageStr, "cover : ", " , cover", left)
	result.BeatMapSet.Covers.Cover = strings.ReplaceAll(result.BeatMapSet.Covers.Cover, "\\", "")
	result.BeatMapSet.Covers.Cover2X, left = findWithIndex(pageStr, "cover@2x : ", " ,", left)
	result.BeatMapSet.Covers.Cover2X = strings.ReplaceAll(result.BeatMapSet.Covers.Cover2X, "\\", "")
	result.BeatMapSet.Covers.Card, left = findWithIndex(pageStr, "card : ", " , card@2x", left)
	result.BeatMapSet.Covers.Card = strings.ReplaceAll(result.BeatMapSet.Covers.Card, "\\", "")
	result.BeatMapSet.Covers.Card2X, left = findWithIndex(pageStr, "card@2x : ", " ,", left)
	result.BeatMapSet.Covers.Card2X = strings.ReplaceAll(result.BeatMapSet.Covers.Card2X, "\\", "")
	result.BeatMapSet.Covers.List, left = findWithIndex(pageStr, "list : ", " ,", left)
	result.BeatMapSet.Covers.List = strings.ReplaceAll(result.BeatMapSet.Covers.List, "\\", "")
	result.BeatMapSet.Covers.List2X, left = findWithIndex(pageStr, "list@2x : ", " ,", left)
	result.BeatMapSet.Covers.List2X = strings.ReplaceAll(result.BeatMapSet.Covers.List2X, "\\", "")
	result.BeatMapSet.Covers.SlimCover, left = findWithIndex(pageStr, "slimcover : ", " , slimcover", left)
	result.BeatMapSet.Covers.SlimCover = strings.ReplaceAll(result.BeatMapSet.Covers.SlimCover, "\\", "")
	result.BeatMapSet.Covers.SlimCover2X, left = findWithIndex(pageStr, "slimcover@2x : ", " }", left)
	result.BeatMapSet.Covers.SlimCover2X = strings.ReplaceAll(result.BeatMapSet.Covers.SlimCover2X, "\\", "")

	result.BeatMapSet.Creator, left = findWithIndex(pageStr, "creator : ", " ", left)
	result.BeatMapSet.FavoriteCount, left = findWithIndex(pageStr, "favourite_count :", ",", left)
	result.BeatMapSet.Hype, left = findWithIndex(pageStr, "hype :", ",", left)
	result.BeatMapSet.Id, left = findWithIndex(pageStr, "id :", ",", left)
	result.BeatMapSet.Nsfw, left = findWithIndex(pageStr, "nsfw :", ",", left)
	result.BeatMapSet.Offset, left = findWithIndex(pageStr, "offset :", ",", left)
	result.BeatMapSet.PlayCount, left = findWithIndex(pageStr, "play_count :", ",", left)
	result.BeatMapSet.PreviewUrl, left = findWithIndex(pageStr, "preview_url : \\/\\/", " , source", left)
	result.BeatMapSet.PreviewUrl = strings.ReplaceAll(result.BeatMapSet.PreviewUrl, "\\", "")
	result.BeatMapSet.Source, left = findWithIndex(pageStr, "source :", " ", left)
	result.BeatMapSet.Spotlight, left = findWithIndex(pageStr, "spotlight :", ",", left)
	result.BeatMapSet.Status, left = findWithIndex(pageStr, "status : ", " ,", left)
	result.BeatMapSet.Title, left = findWithIndex(pageStr, "title : ", " , title_unicode", left)
	result.BeatMapSet.TitleUnicode, left = findWithIndex(pageStr, "title_unicode : ", " ,", left)
	result.BeatMapSet.TrackId, left = findWithIndex(pageStr, "track_id :", ",", left)
	result.BeatMapSet.UserId, left = findWithIndex(pageStr, "user_id :", ",", left)
	result.BeatMapSet.Video, left = findWithIndex(pageStr, "video :", "}", left)

	if scoreType == "best" {
		result.Weight.Percentage, left = findWithIndex(pageStr, "percentage :", ",", left)
		result.Weight.PP, left = findWithIndex(pageStr, "pp :", "}", left)
		return result, left
	}

	// Постановка указателя на конец рекорда
	_, left = findWithIndex(pageStr, "}", "},", left)

	return result, left
}

// Функция получения информации о пользователе
func getUserInfo(id, mode string) UserInfo {

	// Если пользователь не ввёл id, по умолчанию ставит мой id
	if id == "" {
		id = "29829158"
	}

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id + "/" + mode)
	if err != nil {
		return UserInfo{
			Error: "http.Get error",
		}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// HTML полученной страницы в формате string
	pageStr := string(body)[90000:]

	// Проверка на страницу пользователя
	if !strings.Contains(pageStr, "js-react--profile") {
		return UserInfo{
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
	result := UserInfo{}

	// Крайняя левая граница поиска
	left := 0

	//---------------------------- Лучшие рекорды --------------------------------

	// Проверка на существование лучших рекордов
	if !strings.Contains(pageStr, "scoresBest :[]") {

		// Индекс конца лучших рекордов
		end := strings.Index(pageStr, "], scoresFirsts") - 40

		// Пока левая граница не пересечёт конец лучших рекордов
		for s := 0; left < end; s++ {

			// Структура для записи рекорда
			var score Score

			// Получение и запись рекорда
			score, left = parseScore(pageStr, left, "best")

			// Добавление рекорда к результату
			result.ScoresBest = append(result.ScoresBest, score)

		}

	}

	//---------------------------- Первые места --------------------------------

	// Проверка на существование Первых мест
	if !strings.Contains(pageStr, "scoresFirsts :[]") {

		// Индекс конца первых мест
		end := strings.Index(pageStr, "], scoresPinned") - 40

		// Пока левая граница не пересечёт конец первых мест
		for s := 0; left < end; s++ {

			// Структура для записи первого места
			var score Score

			// Получение и запись первого места
			score, left = parseScore(pageStr, left, "first")

			// Добавление первого места к результату
			result.ScoresFirst = append(result.ScoresFirst, score)

		}

	}

	//---------------------------- Закреплённые рекорды --------------------------------

	// Проверка на существование закреплённых рекордов
	if !strings.Contains(pageStr, "scoresPinned :[]") {

		// Индекс конца закреплённых рекордов
		end := strings.Index(pageStr, "scoresRecent :[") - 40

		// Пока левая граница не пересечёт конец закреплённых рекордов
		for s := 0; left < end; s++ {

			// Структура для записи рекорда
			var score Score

			// Получение и запись рекорда
			score, left = parseScore(pageStr, left, "pinned")

			// Добавление закреплённого рекорда к результату
			result.ScoresPinned = append(result.ScoresPinned, score)

		}

	}

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
	result.ProfileColor, left = findWithIndex(pageStr, "profile_colour :", ",", left)
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
	result.Playstyle, left = findWithIndex(pageStr, "playstyle :[ ", " ], ", left)
	result.PostCount, left = findWithIndex(pageStr, "post_count :", ",", left)
	result.ProfileOrder, left = findWithIndex(pageStr, "profile_order :[ ", " ],", left)
	result.Title, left = findWithIndex(pageStr, "title :", ",", left)
	result.TitleUrl, left = findWithIndex(pageStr, "title_url :", ",", left)
	result.Twitter, left = findWithIndex(pageStr, "twitter : ", " ,", left)
	result.Website, left = findWithIndex(pageStr, "website : ", " ,", left)
	result.Website = strings.ReplaceAll(result.Website, "\\", "")
	result.CountyName, left = findWithIndex(pageStr, "name : ", " }", left)

	result.UserCover.CustomUrl, left = findWithIndex(pageStr, "custom_url :", ",", left)
	result.UserCover.Url, left = findWithIndex(pageStr, "url : ", " ,", left)
	result.UserCover.Id, left = findWithIndex(pageStr, "id : ", "}", left)

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
	for c := strings.Index(pageStr, "badges :["); pageStr[c] != ']'; c++ {
		if pageStr[c:c+13] == "awarded_at : " {
			result.Badges = append(result.Badges, Badge{
				AwardedAt:   find(pageStr[c:], "awarded_at : ", " "),
				Description: find(pageStr[c:], "description : ", " ,"),
				ImageUrl:    strings.ReplaceAll(find(pageStr[c:], "image_url : ", " "), "\\", ""),
			})
		}
	}

	result.BeatmapPlaycountsCount, left = findWithIndex(pageStr, "beatmap_playcounts_count :", ",", left)
	result.CommentsCount, left = findWithIndex(pageStr, "comments_count :", ",", left)
	result.FavoriteBeatmapsetCount, left = findWithIndex(pageStr, "favourite_beatmapset_count :", ",", left)
	result.FollowerCount, left = findWithIndex(pageStr, "follower_count :", ",", left)
	result.GraveyardBeatmapsetCount, left = findWithIndex(pageStr, "graveyard_beatmapset_count :", ",", left)

	// Принадлежность к группам
	for c := strings.Index(pageStr, "groups :["); pageStr[c] != ']'; c++ {
		if pageStr[c] == '{' {
			result.Groups += find(pageStr[c:], "name : ", " ,") + ", "
		}
	}
	if result.Groups != "" {
		result.Groups = result.Groups[:len(result.Groups)-2]
	}

	result.GuestBeatmapsetCount, left = findWithIndex(pageStr, "guest_beatmapset_count :", ",", left)
	result.LovedBeatmapsetCount, left = findWithIndex(pageStr, "loved_beatmapset_count :", ",", left)
	result.MappingFollowerCount, left = findWithIndex(pageStr, "mapping_follower_count :", ",", left)

	// Проверка на наличие достижений
	if !strings.Contains(pageStr, "monthly_playcounts :[]") {

		// Конец части со статистикой
		end := strings.Index(pageStr, "pending_beatmapset_count") - 32

		// Цикл обработки статистики
		for left < end {

			// Инициализация структуры подсчета
			var count Count

			// Генерация подсчета
			count.StartDate, left = findWithIndex(pageStr, "start_date : ", " ", left)
			count.Count, left = findWithIndex(pageStr, "count :", "}", left)

			// Добавление статистики
			result.MonthlyPlaycounts = append(result.MonthlyPlaycounts, count)

		}
	}

	result.PendingBeatmapsetCount, left = findWithIndex(pageStr, "pending_beatmapset_count :", ",", left)
	result.Names, left = findWithIndex(pageStr, "previous_usernames :[ ", " ],", left)
	result.RankedBeatmapsetCount, left = findWithIndex(pageStr, "ranked_beatmapset_count :", ",", left)


	result.ScoresBestCount, left = findWithIndex(pageStr, "scores_best_count :", ",", left)
	result.ScoresFirstCount, left = findWithIndex(pageStr, "scores_first_count :", ",", left)
	result.ScoresPinnedCount, left = findWithIndex(pageStr, "scores_pinned_count :", ",", left)
	result.ScoresRecentCount, left = findWithIndex(pageStr, "scores_recent_count :", ",", left)
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
	result.IsRanked, left = findWithIndex(pageStr, "is_ranked:", ",", left)
	result.SS, left = findWithIndex(pageStr, "grade_counts :{ ss :", ",", left)
	result.SSH, left = findWithIndex(pageStr, "ssh :", ",", left)
	result.S, left = findWithIndex(pageStr, "s :", ",", left)
	result.SH, left = findWithIndex(pageStr, "sh :", ",", left)
	result.A, left = findWithIndex(pageStr, "a :", "}", left)
	result.CountryRank, left = findWithIndex(pageStr, "country_rank :", ",", left)
	result.SupportLvl, left = findWithIndex(pageStr, "support_level :", ",", left)

	// Проверка на наличие достижений
	if !strings.Contains(pageStr, "user_achievements :[]") {

		// Конец блока достижений
		end := strings.Index(pageStr, "rank_history :{") - 40

		// Цикл обработки достижений
		for left < end {

			// Инициализация достижения
			var achieve Achievement

			// Генерация достижения
			achieve.AchievedAt, left = findWithIndex(pageStr, "achieved_at : ", " ,", left)
			achieve.AchievementId, left = findWithIndex(pageStr, "achievement_id :", "}", left)

			// Добавление достижения
			result.Achievements = append(result.Achievements, achieve)

		}
	}

	result.RankHistory.Mode, left = findWithIndex(pageStr, "mode : ", " ,", left)
	result.RankHistory.Data, left = findWithIndex(pageStr, "data :[", "]", left)

	result.RankedAndApprovedCount, left = findWithIndex(pageStr, "ranked_and_approved_beatmapset_count :", ",", left)
	result.UnrankedBeatmapsetCount, _ = findWithIndex(pageStr, "unranked_beatmapset_count :", "}", left)

	return result
}

// Функция получения информации о пользователе
func getOnlineInfo(id string) OnlineInfo {

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

// Функция отправки информации о пользователе
func sendUserInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(getUserInfo(mux.Vars(request)["id"], mux.Vars(request)["mode"]))
}

// Функция отправки информации о статусе пользователя
func sendOnlineInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(getOnlineInfo(mux.Vars(request)["id"]))
}

func main() {

	// Вывод времени начала работы
	fmt.Println("API Start: " + string(time.Now().Format("2006-01-02 15:04:05")))
	fmt.Println("Port:\t" + os.Getenv("PORT"))

	/*	Сетап для тестов
		var sd int64
		for i := 0; i < 100; i++ {
			t := time.Now()
			getUserInfo("29829158", "")
			sd += time.Since(t).Milliseconds()
			fmt.Println("{", i, "}cur: \t", sd/(int64(i)+1))
		}
		println("fin:\t", sd/100)
	*/

	// Роутер
	router := mux.NewRouter()

	// Маршруты

	router.HandleFunc("/user", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/", sendUserInfo).Methods("GET")

	router.HandleFunc("/user/{id}", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/{mode}", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/{mode}/", sendUserInfo).Methods("GET")

	router.HandleFunc("/online", sendOnlineInfo).Methods("GET")
	router.HandleFunc("/online/", sendOnlineInfo).Methods("GET")

	router.HandleFunc("/online/{id}", sendOnlineInfo).Methods("GET")
	router.HandleFunc("/online/{id}/", sendOnlineInfo).Methods("GET")

	// Запуск API
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
