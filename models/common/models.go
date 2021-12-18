package common

import "github.com/ssksameer56/Dota2API/models/opendota"

type Dota2GameInfo struct {
	Heroes *map[int]opendota.Hero
	Items  *map[int]opendota.Item
}

type LatestMatches struct {
	Matches []*opendota.MatchDetails
}
