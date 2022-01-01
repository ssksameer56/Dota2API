package handlers

import (
	"context"

	odmodels "github.com/ssksameer56/Dota2API/models/opendota"
	"github.com/ssksameer56/Dota2API/opendota"
	"github.com/ssksameer56/Dota2API/utils"
)

type MatchDataHandler struct {
	Dota2service *opendota.OpenDotaService
}

//Get IDs of matches being currently played
func (mh *MatchDataHandler) GetLiveMatchIDs(pctx context.Context) ([]int64, error) {
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	matches, err := mh.Dota2service.GetLatestMatches(ctx)
	if err != nil {
		utils.LogError("Cant fetch live match IDs: "+err.Error(), "GetLiveMatchIDs")
		return []int64{}, err
	}
	return matches, nil
}

//Get details about a particular matchID
func (mh *MatchDataHandler) GetMatchDetails(pctx context.Context, matchID int) (odmodels.MatchDetails, error) {
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	data, err := mh.Dota2service.GetMatchDetails(ctx, matchID)
	data.RadiantCurrentGoldAdvantage = data.RadiantGoldAdvantage[len(data.RadiantGoldAdvantage)-1]
	data.RadiantCurrentXPAdvantage = data.RadiantXPAdvantage[len(data.RadiantXPAdvantage)-1]

	if err != nil {
		utils.LogError("Cant fetch match details: "+err.Error(), "GetMatchDetails")
		return odmodels.MatchDetails{}, err
	}
	return data, nil
}
