# 🎨 Samples 🎶
<h2> /user/ </h2>
<h3>Request sample </h3>
  
   ``` Elixir
   GET https://osustatsapi.herokuapp.com/user/whitecat
   ```
   
 <h3>Response sample </h3>
  
   ``` Json
  {
   "error":"",
   "username":"WhiteCat",
   "previous_usernames":"eZHG",
   "badges":[
      {
         "awarded_at":"2022-08-02T18:15:26+00:00",
         "description":"The Perennial 2022 2nd Place",
         "image_url":"https://assets.ppy.sh/profile-badges/tp-2-2022.png"
      },
      {
         "awarded_at":"2021-11-29T16:11:23+00:00",
         "description":"osu! World Cup 2021 2nd Place (Germany)",
         "image_url":"https://assets.ppy.sh/profile-badges/badge_owc2021_2nd_place.png"
      },
      {
         "awarded_at":"2021-11-29T15:46:55+00:00",
         "description":"Corsace Open 2021 2nd Place",
         "image_url":"https://assets.ppy.sh/profile-badges/corsace-2nd-2021.png"
      },
      {
         "awarded_at":"2021-04-27T20:45:02+00:00",
         "description":"Bundesl\\u00e4nder Battle 2021 by Second Circle Winning Team",
         "image_url":"https://assets.ppy.sh/profile-badges/blb-2021.png"
      },
      {
         "awarded_at":"2020-12-06T19:38:14+00:00",
         "description":"osu! World Cup 2020 2nd Place (Germany)",
         "image_url":"https://assets.ppy.sh/profile-badges/badge_owc2020_2nd.png"
      }
   ],
   "avatar_url":"https://a.ppy.sh/4504101?1636063550.png",
   "id":"4504101",
   "country_code":"DE",
   "global_rank":"5",
   "country_rank":"1",
   "pp":"19018.6",
   "play_time":"649h46m3s",
   "play_time_seconds":"2339163",
   "ssh":"9",
   "ss":"38",
   "sh":"204",
   "s":"221",
   "a":"729",
   "ranked_score":"26280623252",
   "accuracy":"98.6339",
   "play_count":"37768",
   "total_score":"186732253136",
   "total_hits":"13183828",
   "maximum_combo":"7538",
   "replays":"3511463",
   "level":"101",
   "support_level":"3",
   "follower_count":"89569",
   "default_group":"default",
   "is_online":"false",
   "is_active":"true",
   "is_admin":"false",
   "is_moderator":"false",
   "is_nat":"false",
   "is_gmt":"false",
   "is_bng":"false",
   "is_bot":"false",
   "is_silenced":"false",
   "is_deleted":"false",
   "is_restricted":"false",
   "is_limited_bn":"false",
   "is_full_bn":"false",
   "is_supporter":"true",
   "last_visit":"2022-08-20T00:46:10+00:00",
   "profile_color":"null",
   "ranked_beatmapset_count":"0",
   "pending_beatmapset_count":"0",
   "pm_friends_only":"true",
   "graveyard_beatmapset_count":"5",
   "beatmap_playcounts_count":"2931",
   "comments_count":"0",
   "favorite_beatmapset_count":"1",
   "guest_beatmapset_count":"0",
   "best_beat_map":{
      "title":"Imagination (TV Size)",
      "difficulty_rating":"6.2",
      "id":"2097898",
      "build_id":"null",
      "statistics":"great :308, ok :8",
      "rank":"SH",
      "mods":[ "HD", "HR", "DT" ],
      "ended_at":"2019-10-10T19:30:20+00:00",
      "started_at":"null",
      "accuracy":"0.9831223628691983",
      "maximum_combo":"427",
      "pp":"1043.31",
      "passed":"true",
      "total_score":"6547213",
      "legacy_perfect":"true",
      "replay":"true",
      "mode":"osu",
      "status":"ranked",
      "total_length":"90",
      "ar":"9.4",
      "bpm":"204",
      "convert":"false",
      "count_circles":"188",
      "count_sliders":"144",
      "count_spinners":"1",
      "cs":"4",
      "deleted_at":"null",
      "drain":"6",
      "hit_length":"89",
      "is_scoreable":"true",
      "last_updated":"2020-05-18T23:36:30+00:00",
      "mode_int":"0",
      "pass_count":"578484",
      "play_count":"4951422",
      "ranked":"1",
      "url":"https://osu.ppy.sh/beatmaps/2444148",
      "checksum":"883ee55c1be0cfbea213b7dc0bba4caa",
      "creator":"browiec",
      "favorite_count":"4032",
      "hype":"null",
      "nsfw":"false",
      "offset":"0",
      "spotlight":"false",
      "ruleset_id":"0"
   }
}
   ```

<h2> /online/ </h2>
<h3>Request sample </h3>
  
   ``` Elixir
   GET https://osustatsapi.herokuapp.com/online/hud0shnik
   ```
   
 <h3>Response sample </h3>
  
   ``` Json
   {
  "error":     "",
  "is_online": "false"
   }
   ```
