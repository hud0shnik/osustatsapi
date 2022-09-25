# 🎨 OsuStatsApi 🎶

Osu Stats Api provides fast and powerful
access to player statistics,
including real PP count and real Accuracy percentage

<h2>Structures</h2>

<h3>/user/</h3>
<h4>Request sample </h4>
  
   ``` Elixir
   GET https://osustatsapi.herokuapp.com/user/hud0shnik
   ```

<h4>UserInfo</h4>

Field                      |   Type   | Description
---------------------------|----------|------------
error                      |  string  | 
username                   |  string  |
groups                     |  string  | like "Developers"
active_tournament_banner   |  string  |
previous_usernames         |  string  |
badges                     | []Badge  |
avatar_url                 |  string  |
cover_url                  |  string  |
id                         |  string  |
playmode                   |  string  |
country_code               |  string  | like "RU" or "JP"
country_name               |  string  | like "Japan"
global_rank                |  string  |
country_rank               |  string  |
pp                         |  string  | float value, 4 decimals
play_time                  |  string  | like "202h24m22s"
play_time_seconds          |  string  |
ssh                        |  string  | silver ss
ss                         |  string  |
sh                         |  string  | silver s
s                          |  string  |
a                          |  string  |
ranked_score               |  string  |
accuracy                   |  string  | like "97.132"
play_count                 |  string  |
scores_best_count          |  string  |
scores_first_count         |  string  | first place ranks
scores_pinned_count        |  string  |
scores_recent_count        |  string  |
total_score                |  string  |
total_hits                 |  string  |
maximum_combo              |  string  |
replays                    |  string  | replays watched by others
level                      |  string  |
kudosu                     |  string  |
playstyle                  |  string  |
ocupation                  |  string  |
location                   |  string  |
post_count                 |  string  |
support_level              |  string  |
follower_count             |  string  |
default_group              |  string  |
discord                    |  string  |
interests                  |  string  |
has_supported              |  string  |
is_online                  |  string  |
is_active                  |  string  |
is_admin                   |  string  |
is_moderator               |  string  |
is_nat                     |  string  | Nomination Assessment Team
is_gmt                     |  string  | Global Moderation Team
is_bng                     |  string  | Beatmap Nominators Group
is_bot                     |  string  |
is_silenced                |  string  |
is_deleted                 |  string  |
is_restricted              |  string  | timeout from the community
is_limited_bn              |  string  |
is_full_bn                 |  string  |
is_supporter               |  string  |
last_visit                 |  string  |
profile_color              |  string  |
ranked_beatmapset_count    |  string  |
pending_beatmapset_count   |  string  |
pm_friends_only            |  string  | PM allowed only for friend
graveyard_beatmapset_count |  string  |
beatmap_playcounts_count   |  string  |
comments_count             |  string  |
favorite_beatmapset_count  |  string  |
guest_beatmapset_count     |  string  |
loved_beatmapset_count     |  string  |
mapping_follower_count     |  string  |
profile_order              |  string  |
join_date                  |  string  | like "2022-05-01T19:27:43+00:00"
website                    |  string  |
twitter                    |  string  |
max_friends                |  string  |
max_block                  |  string  |
title                      |  string  |
title_url                  |  string  |
scores_best                | []Score  |
scores_first               | []Score  |
scores_pinned              | []Score  |


<h4>Badge</h4>

Field       |    Type    | Description
------------|------------|------------
awarded_at  |   string   | 
description |   string   | 
image_url   |   string   |


<h4>Score</h4>

Field                   |    Type    | Description
------------------------|------------|------------
accuracy                |   string   |
beatmap_id              |   string   |
build_id                |   string   |
ended_at                |   string   |
maximum_combo           |   string   |
mods                    |  []string  |
passed                  |   string   |
rank                    |   string   |
ruleset_id              |   string   |
started_at              |   string   |
statistics              |   string   |
total_score             |   string   |
user_id                 |   string   |
best_id                 |   string   |
id                      |   string   |
legacy_perfect          |   string   |
pp                      |   string   |
replay                  |   string   |
type                    |   string   |
current_user_attributes |   string   |
beatmap                 |   BeatMap  |
beatmapset              | BeatMapSet |
weight                  |   Weight   | 


<h4>Beatmap</h4>

Field             |    Type    | Description
------------------|------------|------------
beatmapset_id     |   string   |
difficulty_rating |   string   |
id                |   string   |
mode              |   string   |
status            |   string   |
total_length      |   string   |
user_id           |   string   |
version           |   string   |
accuracy          |   string   |
ar                |   string   |
bpm               |   string   |
convert           |   string   |
count_circles	  |   string   |
count_sliders     |   string   |
count_spinners    |   string   |
cs                |   string   |
deleted_at        |   string   |
drain             |   string   |
hit_length        |   string   |
is_scoreable      |   string   |
last_updated      |   string   |
mode_int          |   string   |
pass_count        |   string   |
play_count        |   string   |
ranked            |   string   |
url               |   string   |
checksum          |   string   |


<h4>BeatmapSet</h4>

Field          |    Type    | Description
---------------|------------|------------
artist         |   string   |
artist_unicode |   string   |
covers         |   Covers   |
creator        |   string   |
favorite_count |   string   |
hype           |   string   |
id             |   string   |
nsfw           |   string   |
offset         |   string   |
play_count     |   string   |
preview_url    |   string   |
source         |   string   |
spotlight      |   string   |
status         |   string   |
title          |   string   |
title_unicode  |   string   |
track_id       |   string   |
userId         |   string   |
video          |   string   |


<h4>Covers</h4>

Field        |    Type    | Description
-------------|------------|------------
cover        |   string   |
cover@2x     |   string   |
card         |   string   |
card@2x      |   string   |
list         |   string   |
list@2x      |   string   |
slimcover    |   string   |
slimcover@2x |   string   |


<h4>Weight</h4>

Field      |    Type    | Description
-----------|------------|------------
percentage |   string   |
pp         |   string   |


<h3>/online/</h3>
<h4>Request sample </h4>
  
   ``` Elixir
   GET https://osustatsapi.herokuapp.com/online/hud0shnik
   ```
   
<h4>OnlineInfo</h4>

Field      |    Type    | Description
-----------|------------|------------
error      |   string   |
is_online  |   string   |


[![License - BSD 3-Clause](https://img.shields.io/static/v1?label=License&message=BSD+3-Clause&color=%239a68af&style=for-the-badge)](/LICENSE)
