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

<h4>MapResponse</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
success                     |        bool        | "true" or "false" 
error                       |       string       | api error response 
artist                      |       string       |
artist_string               |       string       |
covers                      |       Covers       |
creator                     |       string       |
favorite_count              |        int         |
hype_current                |        int         |
hype_required               |        int         |
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
track_id                    |        int         |
user_id                     |        int         |
video                       |        bool        |
download_disabled           |        bool        |
bpm                         |        float       |
can_be_hyped                |        bool        |
discussion_enabled          |        bool        |
discussion_locked           |        bool        |
is_scoreable                |        bool        |
last_updated                |       string       |
legacy_thread_url           |       string       |
nominations_summary         | NominationsSummary |
ranked                      |        int         |
ranked_date                 |       string       |
storyboard                  |        bool        |
submitted_date              |       string       |
tags                        |      []string      |
beatmaps                    |       []Maps       |
converts                    |       []Maps       |
current_nominations         |[]CurrentNomination |
description                 |       string       |
genre_id                    |        int         |
genre_name                  |       string       |
language_id                 |        int         |
language_name               |       string       |
ratings                     |       []int        |
recent_favourites           |      []BmUser      |
related_users               |      []BmUser      |
user                        |       BmUser       |
comments                    |      []Comment     |
pinned_comments             |      []Comment     |
user_follow                 |        bool        |

  
<h4>Covers</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
cover                       |       string       |   
cover@2x                    |       string       |   
card                        |       string       |
card@2x                     |       string       |
list                        |       string       |
list@2x                     |       string       |
slimcover                   |       string       |
slimcover@2x                |       string       |

  
<h4>NominationsSummary</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
current                     |        int         |
required                    |        int         |


<h4>Maps</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
beatmapset_id               |        int         |
difficulty_rating           |       float        |
id                          |        int         |
mode                        |       string       |
status                      |       string       |
total_length                |        int         |
user_id                     |        int         |
version                     |       string       |
accuracy                    |       float        |
ar                          |       float        |
bpm                         |       float        |
convert                     |        bool        |
count_circles               |        int         |
count_sliders               |        int         |
count_spinners              |        int         |
cs                          |       float        |
deleted_at                  |       string       |
drain                       |       float        |
hit_length                  |        int         |
is_scoreable                |        bool        |
last_updated                |       string       |
mode_int                    |        int         |
pass_count                  |        int         |
play_count                  |        int         |
ranked                      |        int         |
url                         |       string       |
checksum                    |       string       |
failtimes                   |     Failtimes      |
max_combo                   |        int         |


<h4>Failtimes</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
Fail                        |       []int        |
Exit                        |       []int        |
  

<h4>CurrentNomination</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
beatmapset_id               |        int         |
rulesets                    |       string       |
reset                       |        bool        |
user_id                     |        int         |



<img src="https://wakatime.com/badge/user/ee2709af-fc5f-498b-aaa1-3ea47bf12a00/project/eeb27ba3-3b0a-487f-9650-64aefd7a8458.svg?style=for-the-badge">

[![License - BSD 3-Clause](https://img.shields.io/static/v1?label=License&message=BSD+3-Clause&color=%239a68af&style=for-the-badge)](/LICENSE)
