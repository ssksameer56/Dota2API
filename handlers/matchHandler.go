package handlers

import (
	"context"

	odmodels "github.com/ssksameer56/Dota2API/models/opendota"
	"github.com/ssksameer56/Dota2API/opendota"
)

type MatchDataHandler struct {
	dota2service opendota.OpenDotaService
}

func (mh *MatchDataHandler) GetLiveMatchIDs(pctx context.Context) []int {
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	matches := mh.dota2service.GetLatestMatches(ctx)
	return matches
}

func (mh *MatchDataHandler) GetMatchDetails(pctx context.Context, matchID int) odmodels.MatchDetails {
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	data := mh.dota2service.GetMatchDetails(ctx, matchID)
	return data
}
