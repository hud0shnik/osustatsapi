# 🎨 Samples 🎶
<h2> /user/ </h2>
<h3>Request sample </h3>
  
   ``` Elixir
   GET https://osustatsapi.herokuapp.com/user/4504101
   ```
   
 <h3>Response sample </h3>
  
   ``` Json
  "username": "filsdelama",
  "previous_usernames": "Igorr , Sub2PewDiePie",
  "badges": [
    {
      "awarded_at": "2021-01-01T10:28:56+00:00",
      "description": "osu! Collegiate League Fall 2020 Winning Team ",
      "image_url": "https://assets.ppy.sh/profile-badges/ocl-2020.png"
    },
    {
      "awarded_at": "2019-11-29T15:38:34+00:00",
      "description": "osu! French Sightread Contest Winning Team ",
      "image_url": "https://assets.ppy.sh/profile-badges/ofsc-2019.png"
    },
    {
      "awarded_at": "2018-07-18T15:30:35+00:00",
      "description": "Rising Nations 2018 Winning Team ",
      "image_url": "https://assets.ppy.sh/profile-badges/rsn-2018.png"
    }
  ],
  "avatar_url": "https://a.ppy.sh/2831793?1627054739.jpeg",
  "id": "2831793",
  "country_code": "FR",
  "global_rank": "136",
  "country_rank": "4",
  "pp": "14190.4",
  "play_time": "2412h32m13s",
  "play_time_seconds": "8685133",
  "ssh": "226",
  "ss": "90",
  "sh": "668",
  "s": "1363",
  "a": "1011",
  "ranked_score": "63246553624",
  "accuracy": "99.2593",
  "play_count": "164299",
  "total_score": "533363609548",
  "total_hits": "41037527",
  "maximum_combo": "8253",
  "replays": "788450",
  "level": "105",
  "support_level": "0",
  "default_group": "default",
  "is_online": "false",
  "is_active": "true",
  "is_deleted": "false",
  "is_nat": "false",
  "is_moderator": "false",
  "is_admin": "false",
  "is_bot": "false",
  "is_gmt": "false",
  "is_bng": "false",
  "is_full_bn": "false",
  "is_silenced": "false",
  "is_restricted": "false",
  "is_limited_bn": "false",
  "is_supporter": "false",
  "last_visit": "2022-06-19T15:56:20+00:00",
  "profile_color": "null",
  "ranked_beatmapset_count": "0",
  "pending_beatmapset_count": "0",
  "pm_friends_only": "false",
  "graveyard_beatmapset_count": "36",
  "beatmap_playcounts_count": "7420",
  "comments_count": "152",
  "favorite_beatmapset_count": "388",
  "guest_beatmapset_count": "0",
  "follower_count": "15943",
  "best_beat_map": {
    "title": "Yubi Bouenkyou (TV Size) ",
    "difficulty_rating": "6.11",
    "id": "2469345",
    "build_id": "null",
    "statistics": "great :334, ok :3",
    "rank": "S",
    "mods": [
      "DT"
    ],
    "ended_at": "2022-05-06T15:34:23+00:00",
    "started_at": "null",
    "accuracy": "0.9940652818991098",
    "maximum_combo": "464",
    "pp": "765.479",
    "passed": "true",
    "total_score": "4433799",
    "legacy_perfect": "true",
    "replay": "true",
    "mode": "osu",
    "status": "ranked",
    "total_length": "86",
    "ar": "9.4",
    "bpm": "180",
    "convert": "false",
    "count_circles": "209",
    "count_sliders": "126",
    "count_spinners": "2",
    "cs": "4",
    "deleted_at": "null",
    "drain": "5",
    "hit_length": "85",
    "is_scoreable": "true",
    "last_updated": "2020-06-21T04:02:53+00:00",
    "mode_int": "0",
    "pass_count": "504756",
    "play_count": "3729978",
    "ranked": "1",
    "url": "https://osu.ppy.sh/beatmaps/2469345",
    "checksum": "0ea7dca4e125121cb8b9913fdc443555",
    "creator": "fieryrage",
    "favorite_count": "1487",
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
