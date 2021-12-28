package handlers

import (
	"context"

	odmodels "github.com/ssksameer56/Dota2API/models/opendota"
	"github.com/ssksameer56/Dota2API/opendota"
	"github.com/ssksameer56/Dota2API/utils"
)

type MatchDataHandler struct {
	dota2service opendota.OpenDotaService
}

//Get IDs of matches being currently played
func (mh *MatchDataHandler) GetLiveMatchIDs(pctx context.Context) ([]int, error) {
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	matches, err := mh.dota2service.GetLatestMatches(ctx)
	if err != nil {
		utils.LogError("Cant fetch live match IDs: "+err.Error(), "GetLiveMatchIDs")
		return []int{}, err
	}
	return matches, nil
}

//Get details about a particular matchID
func (mh *MatchDataHandler) GetMatchDetails(pctx context.Context, matchID int) (odmodels.MatchDetails, error) {
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	data, err := mh.dota2service.GetMatchDetails(ctx, matchID)
	if err != nil {
		utils.LogError("Cant fetch match details: "+err.Error(), "GetMatchDetails")
		return odmodels.MatchDetails{}, err
	}
	return data, nil
}
