package common

import "github.com/ssksameer56/Dota2API/models/opendota"

type Dota2GameInfo struct {
	Heroes []*opendota.Hero
	Items  []*opendota.Item
}

type LatestMatches struct {
	Matches []*opendota.MatchDetails
}
