package opendota

type MatchDetails struct {
	MatchID              int            `json:"match_id"`
	DireKills            int            `json:"dire_score"`
	RadiantKills         int            `json:"radiant_score"`
	Duration             int            `json:"duration"`
	GameMode             string         `json:"game_mode"`
	RadiantGoldAdvantage int            `json:"radiant_gold_adv"`
	RadiantXPAdvantage   int            `json:"radiant_xp_adv"`
	DireTowersKilled     int            `json:"tower_status_dire"`
	RadiantTowersKilled  int            `json:"tower_status_radiant"`
	DireBarracksStatus   int            `json:"barracks_status_dire"`
	RadiantBarrackStatus int            `json:"barracks_status_radiant"`
	PlayerDetails        []InGamePlayer `json:"players"`
}

type InGamePlayer struct {
	IsRadiant bool `json:"isRadiant"`
	Gold      int  `json:"gold"`
	Xp        int  `json:"xp_per_min"`
	Level     int  `json:"level"`
	Kills     int  `json:"kills"`
	Deaths    int  `json:"deaths"`
	Assists   int  `json:"assists"`
	Item1ID   int  `json:"item_0"`
	Item2ID   int  `json:"item_1"`
	Item3ID   int  `json:"item_2"`
	Item4ID   int  `json:"item_3"`
	Item5ID   int  `json:"item_4"`
	Item6ID   int  `json:"item_5"`
	Player
}

type Player struct {
	HeroID    int `json:"hero_id"`
	ExtraData interface{}
}
type Match struct {
	MatchID        int      `json:"match_id"`
	AccountDetails []Player `json:"players"`
	ExtraData      interface{}
}
