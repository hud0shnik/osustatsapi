# 🎨 OsuStatsApi 🎶

Osu Stats Api provides fast and powerful
access to player statistics,
including real PP count and real Accuracy percentage

<h3>Structures</h3>
<h4>/user/</h4>

```Go
type UserInfo struct {
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
	DefaultGroup             string  `json:"default_group"`
	IsOnline                 string  `json:"is_online"`
	IsActive                 string  `json:"is_active"`
	IsDeleted                string  `json:"is_deleted"`
	IsNat                    string  `json:"is_nat"`
	IsModerator              string  `json:"is_moderator"`
	IsBot                    string  `json:"is_bot"`
	IsSilenced               string  `json:"is_silenced"`
	IsRestricted             string  `json:"is_restricted"`
	IsLimitedBn              string  `json:"is_limited_bn"`
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
	FollowerCount            string  `json:"follower_count"`
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
	Status string `json:"is_online"`
}
```
