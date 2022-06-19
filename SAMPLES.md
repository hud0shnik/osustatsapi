# 🎨 Samples 🎶
<h2> /user/ </h2>
<h3>Request sample </h3>
  
   ``` Elixir
   GET https://osustatsapi.herokuapp.com/user/4504101
   ```
   
 <h3>Response sample </h3>
  
   ``` Json
  "username": "WhiteCat",
  "previous_usernames": "eZHG",
  "badges": [
    {
      "awarded_at": "2021-11-29T16:11:23+00:00",
      "description": "osu! World Cup 2021 2nd Place (Germany) ",
      "image_url": "https://assets.ppy.sh/profile-badges/badge_owc2021_2nd_place.png"
    },
    {
      "awarded_at": "2021-11-29T15:46:55+00:00",
      "description": "Corsace Open 2021 2nd Place ",
      "image_url": "https://assets.ppy.sh/profile-badges/corsace-2nd-2021.png"
    },
    {
      "awarded_at": "2021-04-27T20:45:02+00:00",
      "description": "Bundesl\\u00e4nder Battle 2021 by Second Circle Winning Team ",
      "image_url": "https://assets.ppy.sh/profile-badges/blb-2021.png"
    },
    {
      "awarded_at": "2020-12-06T19:38:14+00:00",
      "description": "osu! World Cup 2020 2nd Place (Germany) ",
      "image_url": "https://assets.ppy.sh/profile-badges/badge_owc2020_2nd.png"
    }
  ],
  "avatar_url": "https://a.ppy.sh/4504101?1636063550.png",
  "id": "4504101",
  "country_code": "DE",
  "global_rank": "4",
  "country_rank": "1",
  "pp": "19006.5",
  "play_time": "622h47m14s",
  "play_time_seconds": "2242034",
  "ssh": "9",
  "ss": "36",
  "sh": "203",
  "s": "203",
  "a": "718",
  "ranked_score": "25690983242",
  "accuracy": "98.6239",
  "play_count": "36448",
  "total_score": "179182430408",
  "total_hits": "12613574",
  "maximum_combo": "7538",
  "replays": "3401513",
  "level": "101",
  "support_level": "3",
  "default_group": "default",
  "is_online": "false",
  "is_active": "true",
  "is_deleted": "false",
  "is_nat": "false",
  "is_moderator": "false",
  "is_bot": "false",
  "is_silenced": "false",
  "is_restricted": "false",
  "is_limited_bn": "false",
  "is_supporter": "true",
  "last_visit": "2022-06-14T21:38:33+00:00",
  "profile_color": "null",
  "ranked_beatmapset_count": "0",
  "pending_beatmapset_count": "0",
  "pm_friends_only": "true",
  "graveyard_beatmapset_count": "5",
  "beatmap_playcounts_count": "2812",
  "comments_count": "0",
  "favorite_beatmapset_count": "1",
  "guest_beatmapset_count": "0",
  "follower_count": "88039",
  "best_beat_map": {
    "title": "Team Magma &amp; Aqua Leader Battle Theme (Unofficial) ",
    "difficulty_rating": "6.14",
    "id": "2097898",
    "build_id": "null",
    "statistics": "great :308, ok :8",
    "rank": "SH",
    "mods": [
      "HD",
      "HR",
      "DT"
    ],
    "ended_at": "2019-10-10T19:30:20+00:00",
    "started_at": "null",
    "accuracy": "0.9831223628691983",
    "maximum_combo": "427",
    "pp": "1076.33",
    "passed": "true",
    "total_score": "4870654",
    "legacy_perfect": "false",
    "replay": "true",
    "mode": "osu",
    "status": "ranked",
    "total_length": "74",
    "ar": "9.2",
    "bpm": "200.6",
    "convert": "false",
    "count_circles": "210",
    "count_sliders": "105",
    "count_spinners": "1",
    "cs": "4",
    "deleted_at": "null",
    "drain": "6.2",
    "hit_length": "74",
    "is_scoreable": "true",
    "last_updated": "2019-07-25T08:19:53+00:00",
    "mode_int": "0",
    "pass_count": "308502",
    "play_count": "4812291",
    "ranked": "1",
    "url": "https://osu.ppy.sh/beatmaps/2097898",
    "checksum": "d0158fbafb1507f5eaaf55b2af33eb92",
    "creator": "Sotarks",
    "favorite_count": "1185",
    "hype": "null",
    "nsfw": "false",
    "offset": "0",
    "spotlight": "false",
    "ruleset_id": "0"
  }
   ```

<h2> /online/ </h2>
<h3>Request sample </h3>
  
   ``` Elixir
   GET https://osustatsapi.herokuapp.com/online/29829158
   ```
   
 <h3>Response sample </h3>
  
   ``` Json
   "is_online": "true"
   ```
