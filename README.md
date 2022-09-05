# 🎨 OsuStatsApi 🎶

Osu Stats Api provides fast and powerful
access to player statistics,
including real PP count and real Accuracy percentage

<h3>Structures</h3>

<details>
<summary> 📘 Go structures </summary>
</br>
	
```Go
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
	CountryCode              string  `json:"country_code"`
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
	SupportLvl               string  `json:"support_level"`
	FollowerCount            string  `json:"follower_count"`
	DefaultGroup             string  `json:"default_group"`
	Discord                  string  `json:"discord"`
	Interests                string  `json:"interests"`
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
	JoinDate                 string  `json:"join_date"`
	BestBeatMap              beatMap `json:"best_beat_map"`
}
```


```Go
type beatMap struct {
	Title            string   `json:"title"`
	Card             string   `json:"card"`
	Version          string   `json:"version"`
	PreviewUrl       string   `json:"preview_url"`
	TrackId          string   `json:"track_id"`
	DifficultyRating string   `json:"difficulty_rating"`
	Id               string   `json:"id"`
	BuildId          string   `json:"build_id"`
	Cover            string   `json:"cover"`
	SlimCover        string   `json:"slimcover"`
	Statistics       string   `json:"statistics"`
	Rank             string   `json:"rank"`
	Mods             []string `json:"mods"`
	EndedAt          string   `json:"ended_at"`
	StartedAt        string   `json:"started_at"`
	Accuracy         string   `json:"accuracy"`
	MaximumCombo     string   `json:"maximum_combo"`
	PP               string   `json:"pp"`
	Passed           string   `json:"passed"`
	TotalScore       string   `json:"total_score"`
	LegacyPerfect    string   `json:"legacy_perfect"`
	Replay           string   `json:"replay"`
	Mode             string   `json:"mode"`
	Status           string   `json:"status"`
	TotalLength      string   `json:"total_length"`
	Ar               string   `json:"ar"`
	Bpm              string   `json:"bpm"`
	Convert          string   `json:"convert"`
	CountCircles     string   `json:"count_circles"`
	CountSliders     string   `json:"count_sliders"`
	CountSpinners    string   `json:"count_spinners"`
	Cs               string   `json:"cs"`
	DeletedAt        string   `json:"deleted_at"`
	Drain            string   `json:"drain"`
	HitLength        string   `json:"hit_length"`
	IsScoreable      string   `json:"is_scoreable"`
	LastUpdated      string   `json:"last_updated"`
	ModeInt          string   `json:"mode_int"`
	PassCount        string   `json:"pass_count"`
	PlayCount        string   `json:"play_count"`
	Ranked           string   `json:"ranked"`
	Url              string   `json:"url"`
	Artist           string   `json:"artist"`
	Checksum         string   `json:"checksum"`
	Creator          string   `json:"creator"`
	FavoriteCount    string   `json:"favorite_count"`
	Hype             string   `json:"hype"`
	Nsfw             string   `json:"nsfw"`
	Offset           string   `json:"offset"`
	Spotlight        string   `json:"spotlight"`
	RulesetId        string   `json:"ruleset_id"`
	BeatMapSetId     string   `json:"beatmapset_id"`
}
```

```Go
type Badge struct {
	AwardedAt   string `json:"awarded_at"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}
```


<h4>/online/</h4>

```Go
type OnlineInfo struct {
	Error  string `json:"error"`
	Status string `json:"is_online"`
}
```
</details>

<h4>User</h4>

Field | Type | Description
------|------|------------
error | string | 
username | string |
groups | string | like "Developers"
active_tournament_banner | string |
previous_usernames | string |
badges | []Badge |
avatar_url | string |
cover_url | string |
id | string |
country_code | string | like "RU" or "JP"
global_rank | string |
country_rank | string |
pp | string | float value, 4 decimals
play_time | string | like "202h24m22s"
play_time_seconds | string |
ssh | string | silver ss
ss | string |
sh | string | silver s
s | string |
a | string |
ranked_score | string |
accuracy | string | like "97.132"
play_count | string |
scores_best_count | string |
scores_first_count | string | first place ranks
scores_pinned_count | string |
scores_recent_count | string |
total_score | string |
total_hits | string |
maximum_combo | string |
replays | string | replays watched by others
level | string |
support_level | string |
follower_count | string |
default_group | string |
is_online | string |
is_active | string |
is_admin | string |
is_moderator | string |
is_nat | string | Nomination Assessment Team
is_gmt | string | Global Moderation Team
is_bng | string | Beatmap Nominators Group
is_bot | string |
is_silenced | string |
is_deleted | string |
is_restricted | string | timeout from the community
is_limited_bn | string |
is_full_bn | string |
is_supporter | string |
last_visit | string |
profile_color | string |
ranked_beatmapset_count | string |
pending_beatmapset_count | string |
pm_friends_only | string | PM allowed only for friend
graveyard_beatmapset_count | string |
beatmap_playcounts_count | string |
comments_count | string |
favorite_beatmapset_count | string |
guest_beatmapset_count | string |
join_date | string | like "2022-05-01T19:27:43+00:00"
best_beat_map | beatMap |


<h4>Beat map</h4>

Field | Type | Description
------|------|------------
title | string | song name
сard | string |
version | string |
preview_url | string |   
track_id | string | 
difficulty_rating | string | 
id | string | 
build_id | string |
cover | string |
slimcover | string | 
statistics | string | 
rank | string | 
mods | []string | 
ended_at | string | 
started_at | string | 
accuracy | string | 
maximum_combo | string | 
pp | string | 
passed | string | 
total_score | string | 
legacy_perfect | string | 
replay | string | 
mode | string | 
status | string | 
total_length | string | seconds from first note to last note including breaks 
ar | string | approach rate 
bpm | string | 
convert | string | 
count_circles | string | 
count_sliders | string | 
count_spinners | string | 
cs | string | circle size
deleted_at | string | 
drain | string | Health drain
hit_length | string | seconds from first note to last note not including breaks
is_scoreable | string | 
last_updated | string | 
mode_int | string | 
pass_count | string | number of times the beatmap was passed, completed
play_count | string | number of times the beatmap was played
ranked | string | 
url | string |  
artist | string | 
checksum | string | 
creator | string | 
favorite_count | string | Number of times the beatmap was favourited
hype | string | 
nsfw | string | 
offset | string | 
spotlight | string | 
ruleset_id | string |
beatmapset_id | string | 

<h4>Badge</h4>

Field | Type | Description
------|------|------------
awarded_at | string | 
description | string | badge name
image_url | string | badge image


[![License - BSD 3-Clause](https://img.shields.io/static/v1?label=License&message=BSD+3-Clause&color=%239a68af&style=for-the-badge)](/LICENSE)
