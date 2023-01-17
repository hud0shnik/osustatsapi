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
