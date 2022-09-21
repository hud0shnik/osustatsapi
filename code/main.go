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

// Структура значка профиля
type Badge struct {
	AwardedAt   string `json:"awarded_at"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

// Структура для хранения полной информации о пользователе
type UserInfo struct {
	Error                    string  `json:"error"`
	Username                 string  `json:"username"`
	Groups                   string  `json:"groups"`
	ActiveTournamentBanner   string  `json:"active_tournament_banner"`
	Names                    string  `json:"previous_usernames"`
	Badges                   []Badge `json:"badges"`
	AvatarUrl                string  `json:"avatar_url"`
	CoverUrl                 string  `json:"cover_url"`
	UserID                   string  `json:"id"`
	Playmode                 string  `json:"playmode"`
	CountryCode              string  `json:"country_code"`
	CountyName               string  `json:"country_name"`
	GlobalRank               string  `json:"global_rank"`
	CountryRank              string  `json:"country_rank"`
	PP                       string  `json:"pp"`
	PlayTime                 string  `json:"play_time"`
	PlayTimeSeconds          string  `json:"play_time_seconds"`
	SSH                      string  `json:"ssh"`
	SS                       string  `json:"ss"`
	SH                       string  `json:"sh"`
	S                        string  `json:"s"`
	A                        string  `json:"a"`
	RankedScore              string  `json:"ranked_score"`
	Accuracy                 string  `json:"accuracy"`
	PlayCount                string  `json:"play_count"`
	ScoresBestCount          string  `json:"scores_best_count"`
	ScoresFirstCount         string  `json:"scores_first_count"`
	ScoresPinnedCount        string  `json:"scores_pinned_count"`
	ScoresRecentCount        string  `json:"scores_recent_count"`
	TotalScore               string  `json:"total_score"`
	TotalHits                string  `json:"total_hits"`
	MaximumCombo             string  `json:"maximum_combo"`
	Replays                  string  `json:"replays"`
	Level                    string  `json:"level"`
	Kudosu                   string  `json:"kudosu"`
	Playstyle                string  `json:"playstyle"`
	Occupation               string  `json:"occupation"`
	Location                 string  `json:"location"`
	PostCount                string  `json:"post_count"`
	SupportLvl               string  `json:"support_level"`
	FollowerCount            string  `json:"follower_count"`
	DefaultGroup             string  `json:"default_group"`
	Discord                  string  `json:"discord"`
	Interests                string  `json:"interests"`
	HasSupported             string  `json:"has_supported"`
	IsOnline                 string  `json:"is_online"`
	IsActive                 string  `json:"is_active"`
	IsAdmin                  string  `json:"is_admin"`
	IsModerator              string  `json:"is_moderator"`
	IsNat                    string  `json:"is_nat"`
	IsGmt                    string  `json:"is_gmt"`
	IsBng                    string  `json:"is_bng"`
	IsBot                    string  `json:"is_bot"`
	IsSilenced               string  `json:"is_silenced"`
	IsDeleted                string  `json:"is_deleted"`
	IsRestricted             string  `json:"is_restricted"`
	IsLimitedBan             string  `json:"is_limited_bn"`
	IsFullBan                string  `json:"is_full_bn"`
	IsSupporter              string  `json:"is_supporter"`
	LastVisit                string  `json:"last_visit"`
	ProfileColor             string  `json:"profile_color"`
	RankedBeatmapsetCount    string  `json:"ranked_beatmapset_count"`
	PendingBeatmapsetCount   string  `json:"pending_beatmapset_count"`
	PmFriendsOnly            string  `json:"pm_friends_only"`
	GraveyardBeatmapsetCount string  `json:"graveyard_beatmapset_count"`
	BeatmapPlaycountsCount   string  `json:"beatmap_playcounts_count"`
	CommentsCount            string  `json:"comments_count"`
	FavoriteBeatmapsetCount  string  `json:"favorite_beatmapset_count"`
	GuestBeatmapsetCount     string  `json:"guest_beatmapset_count"`
	ProfileOrder             string  `json:"profile_order"`
	JoinDate                 string  `json:"join_date"`
	Website                  string  `json:"website"`
	Twitter                  string  `json:"twitter"`
	MaxFriends               string  `json:"max_friends"`
	MaxBLock                 string  `json:"max_block"`
	Title                    string  `json:"title"`
	TitleUrl                 string  `json:"title_url"`
	ScoresBest               []Score `json:"scores_best"`
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

	// Сохранение html"ки в файл sample.html (для тестов)
	/*
		if err := os.WriteFile("sample.html", []byte(pageStr), 0666); err != nil {
			log.Fatal(err)
		}
	*/

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

	left := 0

	//--------------------------- Лучшая мапа ------------------------------

	if !strings.Contains(pageStr, "scoresBest :[]") {

		end := strings.Index(pageStr, "], scoresFirsts") - 10

		for s := 0; left < end; s++ {

			result.ScoresBest = append(result.ScoresBest, Score{})

			result.ScoresBest[s].Accuracy, left = findWithIndex(pageStr, "accuracy :", ",", left)
			result.ScoresBest[s].BeatMapId, left = findWithIndex(pageStr, "beatmap_id :", ",", left)
			result.ScoresBest[s].BuildId, left = findWithIndex(pageStr, "build_id :", ",", left)
			result.ScoresBest[s].EndedAt, left = findWithIndex(pageStr, "ended_at : ", " ", left)
			result.ScoresBest[s].MaximumCombo, left = findWithIndex(pageStr, "max_combo :", ",", left)

			// Цикл для обработки модов
			for c := 0; pageStr[c] != ']'; c++ {
				if pageStr[c:c+10] == "acronym : " {
					result.ScoresBest[s].Mods = append(result.ScoresBest[s].Mods, pageStr[c+10:c+12])
					// settings :{}
				}
			}

			result.ScoresBest[s].Passed, left = findWithIndex(pageStr, "passed :", ",", left)
			result.ScoresBest[s].Rank, left = findWithIndex(pageStr, "rank : ", " ", left)
			result.ScoresBest[s].RulesetId, left = findWithIndex(pageStr, "ruleset_id :", ",", left)
			result.ScoresBest[s].StartedAt, left = findWithIndex(pageStr, "started_at :", ",", left)
			result.ScoresBest[s].Statistics, left = findWithIndex(pageStr, "statistics :{ ", "}", left)
			result.ScoresBest[s].TotalScore, left = findWithIndex(pageStr, "total_score :", ",", left)
			result.ScoresBest[s].UserId, left = findWithIndex(pageStr, " user_id :", ",", left)
			result.ScoresBest[s].BestId, left = findWithIndex(pageStr, " best_id :", ",", left)
			result.ScoresBest[s].Id, left = findWithIndex(pageStr, " id :", ",", left)
			result.ScoresBest[s].LegacyPerfect, left = findWithIndex(pageStr, "legacy_perfect :", ",", left)
			result.ScoresBest[s].PP, left = findWithIndex(pageStr, "pp :", ",", left)
			result.ScoresBest[s].Replay, left = findWithIndex(pageStr, "replay :", ",", left)
			result.ScoresBest[s].Type, left = findWithIndex(pageStr, "type : ", " ", left)
			result.ScoresBest[s].CurrentUserAttributes, left = findWithIndex(pageStr, "current_user_attributes :{ ", "},", left)

			result.ScoresBest[s].BeatMap.BeatMapSetId, left = findWithIndex(pageStr, "beatmapset_id :", ",", left)
			result.ScoresBest[s].BeatMap.DifficultyRating, left = findWithIndex(pageStr, "difficulty_rating :", ",", left)
			result.ScoresBest[s].BeatMap.Id, left = findWithIndex(pageStr, "id :", ",", left)
			result.ScoresBest[s].BeatMap.Mode, left = findWithIndex(pageStr, "mode : ", " ", left)
			result.ScoresBest[s].BeatMap.Status, left = findWithIndex(pageStr, "status : ", " ", left)
			result.ScoresBest[s].BeatMap.TotalLength, left = findWithIndex(pageStr, "total_length :", ",", left)
			result.ScoresBest[s].BeatMap.UserId, left = findWithIndex(pageStr, " user_id :", ",", left)
			result.ScoresBest[s].BeatMap.Version, left = findWithIndex(pageStr, "version : ", " , accuracy", left)
			result.ScoresBest[s].BeatMap.Accuracy, left = findWithIndex(pageStr, "accuracy :", ",", left)
			result.ScoresBest[s].BeatMap.Ar, left = findWithIndex(pageStr, "ar :", ",", left)
			result.ScoresBest[s].BeatMap.Bpm, left = findWithIndex(pageStr, "bpm :", ",", left)
			result.ScoresBest[s].BeatMap.Convert, left = findWithIndex(pageStr, "convert :", ",", left)
			result.ScoresBest[s].BeatMap.CountCircles, left = findWithIndex(pageStr, "count_circles :", ",", left)
			result.ScoresBest[s].BeatMap.CountSliders, left = findWithIndex(pageStr, "count_sliders :", ",", left)
			result.ScoresBest[s].BeatMap.CountSpinners, left = findWithIndex(pageStr, "count_spinners :", ",", left)
			result.ScoresBest[s].BeatMap.Cs, left = findWithIndex(pageStr, " cs :", ",", left)
			result.ScoresBest[s].BeatMap.DeletedAt, left = findWithIndex(pageStr, "deleted_at :", ",", left)
			result.ScoresBest[s].BeatMap.Drain, left = findWithIndex(pageStr, "drain :", ",", left)
			result.ScoresBest[s].BeatMap.HitLength, left = findWithIndex(pageStr, "hit_length :", ",", left)
			result.ScoresBest[s].BeatMap.IsScoreable, left = findWithIndex(pageStr, "is_scoreable :", ",", left)
			result.ScoresBest[s].BeatMap.LastUpdated, left = findWithIndex(pageStr, "last_updated : ", " ", left)
			result.ScoresBest[s].BeatMap.ModeInt, left = findWithIndex(pageStr, "mode_int :", ",", left)
			result.ScoresBest[s].BeatMap.PassCount, left = findWithIndex(pageStr, "passcount :", ",", left)
			result.ScoresBest[s].BeatMap.PlayCount, left = findWithIndex(pageStr, "playcount :", ",", left)
			result.ScoresBest[s].BeatMap.Ranked, left = findWithIndex(pageStr, "ranked :", ",", left)
			result.ScoresBest[s].BeatMap.Url, left = findWithIndex(pageStr, "url : ", " ", left)
			result.ScoresBest[s].BeatMap.Url = strings.ReplaceAll(result.ScoresBest[s].BeatMap.Url, "\\", "")
			result.ScoresBest[s].BeatMap.Checksum, left = findWithIndex(pageStr, "checksum : ", " ", left)

			result.ScoresBest[s].BeatMapSet.Artist, left = findWithIndex(pageStr, "artist :", ", artist_", left)
			result.ScoresBest[s].BeatMapSet.ArtistUnicode, left = findWithIndex(pageStr, "artist_unicode :", ",", left)
			result.ScoresBest[s].BeatMapSet.Covers.Cover, left = findWithIndex(pageStr, "cover : ", " , cover", left)
			result.ScoresBest[s].BeatMapSet.Covers.Cover = strings.ReplaceAll(result.ScoresBest[s].BeatMapSet.Covers.Cover, "\\", "")
			result.ScoresBest[s].BeatMapSet.Covers.Cover2X, left = findWithIndex(pageStr, "cover@2x : ", " ,", left)
			result.ScoresBest[s].BeatMapSet.Covers.Cover2X = strings.ReplaceAll(result.ScoresBest[s].BeatMapSet.Covers.Cover2X, "\\", "")
			result.ScoresBest[s].BeatMapSet.Covers.Card, left = findWithIndex(pageStr, "card : ", " , card@2x", left)
			result.ScoresBest[s].BeatMapSet.Covers.Card = strings.ReplaceAll(result.ScoresBest[s].BeatMapSet.Covers.Card, "\\", "")
			result.ScoresBest[s].BeatMapSet.Covers.Card2X, left = findWithIndex(pageStr, "card@2x : ", " ,", left)
			result.ScoresBest[s].BeatMapSet.Covers.Card2X = strings.ReplaceAll(result.ScoresBest[s].BeatMapSet.Covers.Card2X, "\\", "")
			result.ScoresBest[s].BeatMapSet.Covers.List2X, left = findWithIndex(pageStr, "list@2x : ", " ,", left)
			result.ScoresBest[s].BeatMapSet.Covers.List2X = strings.ReplaceAll(result.ScoresBest[s].BeatMapSet.Covers.List2X, "\\", "")
			result.ScoresBest[s].BeatMapSet.Covers.SlimCover, left = findWithIndex(pageStr, "slimcover : ", " , slimcover", left)
			result.ScoresBest[s].BeatMapSet.Covers.SlimCover = strings.ReplaceAll(result.ScoresBest[s].BeatMapSet.Covers.SlimCover, "\\", "")
			result.ScoresBest[s].BeatMapSet.Covers.SlimCover2X, left = findWithIndex(pageStr, "slimcover@2x : ", " }", left)
			result.ScoresBest[s].BeatMapSet.Covers.SlimCover2X = strings.ReplaceAll(result.ScoresBest[s].BeatMapSet.Covers.SlimCover2X, "\\", "")

			result.ScoresBest[s].BeatMapSet.Creator, left = findWithIndex(pageStr, "creator : ", " ", left)
			result.ScoresBest[s].BeatMapSet.FavoriteCount, left = findWithIndex(pageStr, "favourite_count :", ",", left)
			result.ScoresBest[s].BeatMapSet.Hype, left = findWithIndex(pageStr, "hype :", ",", left)
			result.ScoresBest[s].BeatMapSet.Id, left = findWithIndex(pageStr, "id :", ",", left)
			result.ScoresBest[s].BeatMapSet.Nsfw, left = findWithIndex(pageStr, "nsfw :", ",", left)
			result.ScoresBest[s].BeatMapSet.Offset, left = findWithIndex(pageStr, "offset :", ",", left)
			result.ScoresBest[s].BeatMapSet.PlayCount, left = findWithIndex(pageStr, "play_count :", ",", left)
			result.ScoresBest[s].BeatMapSet.PreviewUrl, left = findWithIndex(pageStr, "preview_url : \\/\\/", " , source", left)
			result.ScoresBest[s].BeatMapSet.PreviewUrl = strings.ReplaceAll(result.ScoresBest[s].BeatMapSet.PreviewUrl, "\\", "")
			result.ScoresBest[s].BeatMapSet.Source, left = findWithIndex(pageStr, "source :", " ", left)
			result.ScoresBest[s].BeatMapSet.Spotlight, left = findWithIndex(pageStr, "spotlight :", ",", left)
			result.ScoresBest[s].BeatMapSet.Status, left = findWithIndex(pageStr, "status :", ",", left)
			result.ScoresBest[s].BeatMapSet.Title, left = findWithIndex(pageStr, "title : ", " , title_unicode", left)
			result.ScoresBest[s].BeatMapSet.TitleUnicode, left = findWithIndex(pageStr, "title_unicode :", ",", left)
			result.ScoresBest[s].BeatMapSet.TrackId, left = findWithIndex(pageStr, "track_id :", ",", left)
			result.ScoresBest[s].BeatMapSet.UserId, left = findWithIndex(pageStr, "user_id :", ",", left)
			result.ScoresBest[s].BeatMapSet.Video, left = findWithIndex(pageStr, "video :", "}", left)
			result.ScoresBest[s].Weight.Percentage, left = findWithIndex(pageStr, "percentage :", ",", left)
			result.ScoresBest[s].Weight.PP, left = findWithIndex(pageStr, "pp :", "}", left)

		}

	}

	//--------------------------- Статистика игрока ------------------------------

	// Ссылка на аватар
	result.AvatarUrl, left = findWithIndex(pageStr, "avatar_url : ", " ", left)
	result.AvatarUrl = strings.ReplaceAll(result.AvatarUrl, "\\", "")

	// Код страны
	result.CountryCode, left = findWithIndex(pageStr, "country_code : ", " ", left)

	// Группа
	result.DefaultGroup, left = findWithIndex(pageStr, "default_group : ", " ", left)

	// Айди
	result.UserID, left = findWithIndex(pageStr, " id :", ",", left)

	// Активность
	result.IsActive, left = findWithIndex(pageStr, "is_active :", ",", left)

	// Бот
	result.IsBot, left = findWithIndex(pageStr, "is_bot :", ",", left)

	// Удалённый профиль
	result.IsDeleted, left = findWithIndex(pageStr, "is_deleted :", ",", left)

	// Статус в сети
	result.IsOnline, left = findWithIndex(pageStr, "is_online :", ",", left)

	// Подписка
	result.IsSupporter, left = findWithIndex(pageStr, "is_supporter :", ",", left)

	// В последний раз был в сети
	result.LastVisit, left = findWithIndex(pageStr, "last_visit : ", " ", left)

	// Сообщения только от друзей
	result.PmFriendsOnly, left = findWithIndex(pageStr, "pm_friends_only :", ",", left)

	// Цвет профиля
	result.ProfileColor, left = findWithIndex(pageStr, "profile_colour :", ",", left)

	// Юзернейм
	result.Username, left = findWithIndex(pageStr, "username : ", " ", left)

	// Шапка профиля
	result.CoverUrl, left = findWithIndex(pageStr, "cover_url : ", " ", left)
	result.CoverUrl = strings.ReplaceAll(result.CoverUrl, "\\", "")

	// Дискорд
	result.Discord, left = findWithIndex(pageStr, "discord :", " ", left)

	// Спонсорка
	result.HasSupported, left = findWithIndex(pageStr, "has_supported :", ",", left)

	// Интересы
	result.Interests, left = findWithIndex(pageStr, "interests :", ", join_date", left)

	// Дата регистрации
	result.JoinDate, left = findWithIndex(pageStr, "join_date : ", " ,", left)

	// Кудосу
	result.Kudosu, left = findWithIndex(pageStr, "kudosu :{ total :", ",", left)

	// Локация
	result.Location, left = findWithIndex(pageStr, "location :", ",", left)

	// Размер черного списка
	result.MaxBLock, left = findWithIndex(pageStr, "max_blocks :", ",", left)

	// Максимальное количество друзей
	result.MaxFriends, left = findWithIndex(pageStr, "max_friends :", ",", left)

	// Род деятельности
	result.Occupation, left = findWithIndex(pageStr, "occupation :", ",", left)

	// Режим игры
	result.Playmode, left = findWithIndex(pageStr, "playmode : ", " ,", left)

	// Стиль игры
	result.Playstyle, left = findWithIndex(pageStr, "playstyle :[ ", " ], ", left)

	// Количество постов
	result.PostCount, left = findWithIndex(pageStr, "post_count :", ",", left)

	// Порядок карточек в профиле
	result.ProfileOrder, left = findWithIndex(pageStr, "profile_order :[ ", " ],", left)

	// Тайтл
	result.Title, left = findWithIndex(pageStr, "title :", ",", left)

	// Адрес тайтла
	result.TitleUrl, left = findWithIndex(pageStr, "title_url :", ",", left)

	// Твиттер
	result.Twitter, left = findWithIndex(pageStr, "twitter :", ",", left)

	// Ссылка на сайт
	result.Website, left = findWithIndex(pageStr, "website :", ",", left)
	result.Website = strings.ReplaceAll(result.Website, "\\", "")

	// Название страны
	result.CountyName, left = findWithIndex(pageStr, "name : ", " }", left)

	// Администрация
	result.IsAdmin, left = findWithIndex(pageStr, "is_admin :", ",", left)

	// Команда номинации
	result.IsBng, left = findWithIndex(pageStr, "is_bng :", ",", left)

	// Вечный бан
	result.IsFullBan, left = findWithIndex(pageStr, "is_full_bn :", ",", left)

	// Команда глобальной модерации
	result.IsGmt, left = findWithIndex(pageStr, "is_gmt :", ",", left)

	// Временный бан
	result.IsLimitedBan, left = findWithIndex(pageStr, "is_limited_bn :", ",", left)

	// Модератор
	result.IsModerator, left = findWithIndex(pageStr, "is_moderator :", ",", left)

	// Команда оценки номинаций
	result.IsNat, left = findWithIndex(pageStr, "is_nat :", ",", left)

	// Ограничение
	result.IsRestricted, left = findWithIndex(pageStr, "is_restricted :", ",", left)

	// Немота
	result.IsSilenced, left = findWithIndex(pageStr, "is_silenced :", ",", left)

	// Баннер текущего турнира
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

	// Количество сыгранных карт
	result.BeatmapPlaycountsCount, left = findWithIndex(pageStr, "beatmap_playcounts_count :", ",", left)

	// Количество комментариев
	result.CommentsCount, left = findWithIndex(pageStr, "comments_count :", ",", left)

	// Количество любимых карт
	result.FavoriteBeatmapsetCount, left = findWithIndex(pageStr, "favourite_beatmapset_count :", ",", left)

	// Подписчики
	result.FollowerCount, left = findWithIndex(pageStr, "follower_count :", ",", left)

	// Заброшенные карты
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

	// Карты с гостевым участием
	result.GuestBeatmapsetCount, left = findWithIndex(pageStr, "guest_beatmapset_count :", ",", left)

	// Карты на рассмотрении
	result.PendingBeatmapsetCount, left = findWithIndex(pageStr, "pending_beatmapset_count :", ",", left)

	// Юзернеймы
	result.Names, left = findWithIndex(pageStr, "previous_usernames :[ ", " ],", left)

	// Рейтинговые и одобренные карты
	result.RankedBeatmapsetCount, left = findWithIndex(pageStr, "ranked_beatmapset_count :", ",", left)

	// Лучшие рекорды
	result.ScoresBestCount, left = findWithIndex(pageStr, "scores_best_count :", ",", left)

	// Первые места
	result.ScoresFirstCount, left = findWithIndex(pageStr, "scores_first_count :", ",", left)

	// Закреплённые рекорды
	result.ScoresPinnedCount, left = findWithIndex(pageStr, "scores_pinned_count :", ",", left)

	// Недавние рекорды
	result.ScoresRecentCount, left = findWithIndex(pageStr, "scores_recent_count :", ",", left)

	// Уровень
	result.Level, left = findWithIndex(pageStr, "level :{ current :", ",", left)

	// Глобальный рейтинг
	result.GlobalRank, left = findWithIndex(pageStr, "global_rank :", ",", left)

	// PP-хи
	result.PP, left = findWithIndex(pageStr, "pp :", ",", left)

	// Всего очков
	result.RankedScore, left = findWithIndex(pageStr, "ranked_score :", ",", left)

	// Точность попаданий
	result.Accuracy, left = findWithIndex(pageStr, "hit_accuracy :", ",", left)

	// Количество игр
	result.PlayCount, left = findWithIndex(pageStr, "play_count :", ",", left)

	// Время в игре в секундах
	result.PlayTimeSeconds, left = findWithIndex(pageStr, "play_time :", ",", left)

	// Время в игре в часах
	duration, _ := time.ParseDuration(result.PlayTimeSeconds + "s")
	result.PlayTime = duration.String()

	// Рейтинговые очки
	result.TotalScore, left = findWithIndex(pageStr, "total_score :", ",", left)

	// Всего попаданий
	result.TotalHits, left = findWithIndex(pageStr, "total_hits :", ",", left)

	// Максимальное комбо
	result.MaximumCombo, left = findWithIndex(pageStr, "maximum_combo :", ",", left)

	// Реплеев просмотрено другими
	result.Replays, left = findWithIndex(pageStr, "replays_watched_by_others :", ",", left)

	// SS-ки
	result.SS, left = findWithIndex(pageStr, "grade_counts :{ ss :", ",", left)

	// SSH-ки
	result.SSH, left = findWithIndex(pageStr, "ssh :", ",", left)

	// S-ки
	result.S, left = findWithIndex(pageStr, "s :", ",", left)

	// SH-ки
	result.SH, left = findWithIndex(pageStr, "sh :", ",", left)

	// A-хи
	result.A, left = findWithIndex(pageStr, "a :", "}", left)

	// Рейтинг в стране
	result.CountryRank, left = findWithIndex(pageStr, "country_rank :", ",", left)

	// Уровень подписки
	result.SupportLvl, _ = findWithIndex(pageStr, "support_level :", ",", left)

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
