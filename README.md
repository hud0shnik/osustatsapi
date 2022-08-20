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
	Names                    string  `json:"previous_usernames"`
	Badges                   []Badge `json:"badges"`
	AvatarUrl                string  `json:"avatar_url"`
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
	TotalScore               string  `json:"total_score"`
	TotalHits                string  `json:"total_hits"`
	MaximumCombo             string  `json:"maximum_combo"`
	Replays                  string  `json:"replays"`
	Level                    string  `json:"level"`
	SupportLvl               string  `json:"support_level"`
	FollowerCount            string  `json:"follower_count"`
	DefaultGroup             string  `json:"default_group"`
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
	BestBeatMap              beatMap `json:"best_beat_map"`
}
```


```Go
type beatMap struct {
	Title            string   `json:"title"`
	DifficultyRating string   `json:"difficulty_rating"`
	Id               string   `json:"id"`
	BuildId          string   `json:"build_id"`
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
	Checksum         string   `json:"checksum"`
	Creator          string   `json:"creator"`
	FavoriteCount    string   `json:"favorite_count"`
	Hype             string   `json:"hype"`
	Nsfw             string   `json:"nsfw"`
	Offset           string   `json:"offset"`
	Spotlight        string   `json:"spotlight"`
	RulesetId        string   `json:"ruleset_id"`
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
username | string |
previous_usernames | string |
badges | []Badge |
avatar_url | string |
id | string |
country_code | string | like "RU" or "JP"
global_rank | string |
country_rank | string |
pp | string | float value, 4 decimals
play_time | string | 
play_time_seconds | string |
ssh | string |
ss | string |
sh | string |
s | string |
a | string |
ranked_score | string |
accuracy | string | 
play_count | string |
total_score | string |
total_hits | string |
maximum_combo | string |
replays | string | Replays Watched by Others
level | string |
support_level | string |
default_group | string |
is_online | string |
is_active | string |
is_deleted | string |
is_nat | string |
is_moderator | string |
is_bot | string |
is_silenced | string |
is_restricted | string |
is_limited_bn | string |
is_supporter | string |
last_visit | string |
profile_color | string |
ranked_beatmapset_count | string |
pending_beatmapset_count | string |
pm_friends_only | string |
graveyard_beatmapset_count | string |
beatmap_playcounts_count | string |
comments_count | string |
favorite_beatmapset_count | string |
guest_beatmapset_count | string |
follower_count | string |
best_beat_map | beatMap |


<h4>Beat map</h4>

Field | Type | Description
------|------|------------
title | string | song name
difficulty_rating | string | 
id | string | 
build_id | string | 
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
ar | string | Approach Rate 
bpm | string | 
convert | string | 
count_circles | string | 
count_sliders | string | 
count_spinners | string | 
cs | string | Circle size value
deleted_at | string | 
drain | string | Health drain
hit_length | string | seconds from first note to last note not including breaks
is_scoreable | string | 
last_updated | string | 
mode_int | string | 
pass_count | string | Number of times the beatmap was passed, completed
play_count | string | Number of times the beatmap was played
ranked | string | 
url | string | 
checksum | string | 
creator | string | 
favorite_count | string | Number of times the beatmap was favourited
hype | string | 
nsfw | string | 
offset | string | 
spotlight | string | 
ruleset_id | string | 

<h4>Badge</h4>

Field | Type | Description
------|------|------------
awarded_at | string | 
description | string | Badge name
image_url | string | Badge image


[![License - BSD 3-Clause](https://img.shields.io/static/v1?label=License&message=BSD+3-Clause&color=%239a68af&style=for-the-badge)](/LICENSE)
