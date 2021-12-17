package opendota

type MatchDetails struct {
	MatchID              int    `json:"match_id"`
	DireKills            int    `json:"dire_score"`
	RadiantKills         int    `json:"radiant_score"`
	Duration             int    `json:"duration"`
	GameMode             string `json:"game_mode"`
	RadiantGoldAdvantage int    `json:"radiant_gold_adv"`
	RadiantXPAdvantage   int    `json:"radiant_xp_adv"`
	DireTowersKilled     int    `json:"tower_status_dire"`
	RadiantTowersKilled  int    `json:"tower_status_radiant"`
	DireBarracksStatus   int    `json:"barracks_status_dire"`
	RadiantBarrackStatus int    `json:"barracks_status_radiant"`
}

type Account struct {
	Details map[string]string
}
type Match struct {
	MatchID        int                  `json:"match_id"`
	AccountDetails map[string][]Account `json:"players"`
	ExtraData      interface{}
}
