# 🎨 OsuStatsApi 🎶

Osu Stats Api provides fast and powerful
access to player statistics,
including real PP count and real Accuracy percentage

<h2>Structures</h2>

<h3>/user/</h3>
<h4>Request sample </h4>
  
   ``` Elixir
   https://osustatsapi.vercel.app/api/user?id=hud0shnik
   ```
  
<h4>UserInfo</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
success                     |        bool        | "true" or "false" 
error                       |       string       | api error response (default value= "")
avatar_url                  |       string       |
country_code                |       string       | like "RU" or "JP"
default_group               |       string       |
id                          |        int         |
is_active                   |        bool        |
is_bot                      |        bool        |
is_deleted                  |        bool        |
is_online                   |        bool        |
is_supporter                |        bool        |
last_visit                  |       string       | 
pm_friends_only             |        bool        | PM allowed only for
profile_color               |       string       |
username                    |       string       |
cover_url                   |       string       |
discord                     |       string       |
has_supported               |        bool        |
interests                   |       string       |
join_date                   |       string       | like "2022-05-01T19:27:43+00:00"
kudosu                      |        int         |
location                    |       string       |
max_friends                 |        int         |
max_block                   |        int         |
occupation                  |       string       |
playmode                    |       string       |
playstyle                   |      []string      |
post_count                  |        int         |
profile_order               |      []string      |
title                       |       string       |
title_url                   |       string       |
twitter                     |       string       |
website                     |       string       |
country_name                |       string       | like "Japan"
cover                       |       Cover        |
is_admin                    |        bool        |    
is_bng                      |        bool        | Beatmap Nominators Group
is_full_bn                  |        bool        | full ban
is_gmt                      |        bool        | Global Moderation Team
is_limited_bn               |        bool        | limited ban
is_moderator                |        bool        |
is_nat                      |        bool        | Nomination Assessment Team
is_restricted               |        bool        | timeout from the community
is_silenced                 |        bool        |
account_history             |       string       |
active_tournament_banner    |       string       |
badges                      |      []Badge       |
comments_count              |        int         |
follower_count              |        int         |
groups                      |       string       | like "Developers"
mapping_follower_count      |        int         |
pending_beatmapset_count    |        int         |
previous_usernames          |      []string      |
level                       |        int         |
global_rank                 |        int         |
pp                          |       float        | float value, 4 decimals
ranked_score                |        int         |
accuracy                    |       float        | like "97.132"
play_count                  |        int         |
play_time                   |       string       | like "202h24m22s"
play_time_seconds           |        int         |
total_score                 |        int         |
total_hits                  |        int         |
maximum_combo               |        int         |
replays                     |        int         | replays watched by others   
is_ranked                   |        bool        |
ss                          |        int         | 
ssh                         |        int         | silver ss 
s                           |        int         |
sh                          |        int         | silver s
a                           |        int         |
country_rank                |        int         |
support_level               |        int         |
achievements                |    []Achievement   |
medals                      |        int         |
rank_history                |       History      |
unranked_beatmapset_count   |        int         |
favorite_beatmaps           |     []Beatmap      |
graveyard_beatmaps          |     []Beatmap      |
guest_beatmaps              |     []Beatmap      |
loved_beatmaps              |     []Beatmap      |
ranked_beatmaps             |     []Beatmap      |
pending_beatmaps            |     []Beatmap      |
kudosu_items                |     []Kudosu       |
recent_activity             |    []Activity      |
best                        |      []Score       |
firsts                      |      []Score       | first places
pinned                      |      []Score       | pinned beatmaps
beatmap_playcounts          |    []PlayCount     |
monthly_playcounts          |      []Count       |
replays_watched_counts      |      []Count       |


<h4>Cover</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
custom_url                  |       string       |
url                         |       string       |
id                          |        int         |


<h4>Badge</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
awarded_at                  |       string       | like "2022-10-08T03:47:35+00:00"
description                 |       string       |
image_url                   |       string       |


<h4>Achievement</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
achieved_at                 |       string       | UTC format date (yyyy-mm-ddThh:ss:ssZ)
achievement_id              |       string       |


<h4>History</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
mode                        |       string       | like "osu"
data                        |       []int        |


<h4>Beatmap</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
artist                      |       string       |
artist_unicode              |       string       |
covers                      |       Cover        |
creator                     |       string       |
favorite_count              |        int         |
hype                        |       string       |
id                          |        int         |
nsfw                        |        bool        | Not Safe For Work
offset                      |        int         |
play_count                  |        int         |
preview_url                 |       string       | like "b.ppy.sh/preview/1730467.mp3"
source                      |       string       |
spotlight                   |        bool        |
status                      |       string       |
title                       |       string       |
title_unicode               |       string       |
track_id                    |       string       |
userId                      |        int         |
video                       |        bool        |
download_disabled           |        bool        |
bpm                         |       float        |
can_be_hyped                |        bool        |
discussion_enabled          |        bool        |
discussion_locked           |        bool        |
is_scoreable                |        bool        |
last_updated                |       string       | UTC date
legacy_thread_url           |       string       |
nominations_summary         | NominationsSummary |
ranked                      |        int         |
ranked_date                 |       string       |
storyboard                  |        bool        |
submitted_date              |       string       | UTC date
tags                        |      []string      |
beatmap                     |       Beatmaps     |


<h4>NominationsSummary</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
current                     |        int         |
required                    |        int         |


<h4>Beatmaps</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
beatmapset_id               |        int         |
difficulty_rating           |       float        |
id                          |        int         |
mode                        |       string       | like "osu"
status                      |       string       | like "ranked"
total_length                |        int         |
user_id                     |        int         |
version                     |       string       |
accuracy                    |       float        |
ar                          |       float        | Approach Rate
bpm                         |       float        |
convert                     |        bool        |
count_circles               |        int         |
count_sliders               |        int         |
count_spinners              |        int         |
cs                          |       float        | Circle Size
deleted_at                  |       string       | "null" or UTC date
drain                       |       float        |
hit_length                  |        int         |
is_scoreable                |        bool        |
last_updated                |       string       | UTC date
mode_int                    |        int         |
pass_count                  |        int         |
play_count                  |        int         |
ranked                      |        int         |
url                         |       string       |
checksum                    |       string       |


<h4>Kudosu</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
id                          |        int         |
action                      |       string       |
amount                      |        int         |
model                       |       string       |
created_at                  |       string       |
giver                       |     KudosuGiver    |
post                        |     KudosuPost     |
details                     |       string       |


<h4>KudosuGiver</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
url                         |       string       |
username                    |       string       |


<h4>KudosuPost</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
url                         |       string       |
title                       |       string       |


<h4>Score</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
accuracy                    |       float        |
beatmap_id                  |        int         |
build_id                    |       string       |  
ended_at                    |       string       |
legacy_score_id             |       string       |
legacy_total_score          |       string       |
max_combo                   |        int         |
maximum_statistics          |     Statistics     |
mods                        |      []string      | 
passed                      |        bool        |
rank                        |       string       |
ruleset_id                  |        int         |
started_at                  |       string       |
statistics                  |     Statistics     |
total_score                 |        int         |
user_id                     |        int         |
best_id                     |        int         |
id                          |        int         |
legacy_perfect              |        bool        |
pp                          |       float        |
replay                      |        bool        |
type                        |       string       |
current_user_attributes     |       string       |
beatmap                     |      Beatmaps      |
beatmapset                  |     Beatmapset     |
weight                      |       Weight       |


<h4>Statistics</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
good                        |        int         |
great                       |        int         |
meh                         |        int         |
miss                        |        int         |
ok                          |        int         |
perfect                     |        int         |


<h4>Beatmapset</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
artist                      |       string       |
artist_unicode              |       string       |
covers                      |       Covers       |
creator                     |       string       |
favorite_count              |        int         |
hype                        |       string       |
id                          |        int         |
nsfw                        |        bool        |
offset                      |        int         |
play_count                  |        int         |
preview_url                 |       string       |
source                      |       string       |
spotlight                   |        bool        |
status                      |       string       |
title                       |       string       |
title_unicode               |       string       |
track_id                    |       string       |
userId                      |        int         |
video                       |        bool        |

<h4>Weight</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
percentage                  |       float        |
pp                          |       float        |


<h4>PlayCount</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
beatmapset_id               |        int         |
difficulty_rating           |       float        |
id                          |        int         |
status                      |       string       |
total_length                |        int         |
user_id                     |        int         |
version                     |       string       |


<h4>PlayCountBeatmap</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
beatmapset_id               |        int         |
difficulty_rating           |       float        |
id                          |        int         |
status                      |       string       |
total_length                |        int         |
user_id                     |        int         |
version                     |       string       |
 

<h3>/online/</h3>
<h4>Request sample </h4>
  
   ``` Elixir
   https://osustatsapi.vercel.app/api/online?id=hud0shnik
   ```
   
<h4>OnlineResponse</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
success                     |        bool        | 
error                       |       string       | api error response (default value= "")
status                      |       string       |

<h3>/map/</h3>
<h4>Request sample </h4>
  
   ``` Elixir
   https://osustatsapi.vercel.app/api/map?beatmapset=1607429&id=3477840
   ```




<img src="https://wakatime.com/badge/user/ee2709af-fc5f-498b-aaa1-3ea47bf12a00/project/eeb27ba3-3b0a-487f-9650-64aefd7a8458.svg?style=for-the-badge">

[![License - BSD 3-Clause](https://img.shields.io/static/v1?label=License&message=BSD+3-Clause&color=%239a68af&style=for-the-badge)](/LICENSE)
