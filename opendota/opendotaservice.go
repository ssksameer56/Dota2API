package opendota

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ssksameer56/Dota2API/models/opendota"
	"github.com/ssksameer56/Dota2API/utils"
)

type OpenDotaService struct {
	isPremium bool
	client    utils.HttpClient
	Name      string
}

var baseURL = "https://api.opendota.com/api/"

func (od *OpenDotaService) GetAllHeroes(pctx context.Context) *map[int](opendota.Hero) {
	query := "constants/heroes"
	data, err := od.client.GetData(pctx, query)
	finalHeroes := map[int]opendota.Hero{}
	if err != nil {
		utils.LogError("Error when getting heroes: "+err.Error(), "GetAllHeroes")
		return &finalHeroes
	}
	heroes := opendota.Heroes{}
	err = json.Unmarshal(data, &heroes)
	if err != nil {
		utils.LogError("Error when parsing heroes: "+err.Error(), "GetAllHeroes")
		return &finalHeroes
	}

	for _, heroData := range heroes {
		finalHeroes[heroData.Id] = heroData
	}

	heroAbilities, err := od.GetHeroAbilities(pctx)

	if err != nil {
		utils.LogError("Error when getting abilities: "+err.Error(), "GetAllHeroes")
		return &finalHeroes
	}
	abilities, err := od.GetAbilities(pctx)

	if err != nil {
		utils.LogError("Error when getting abilities: "+err.Error(), "GetAllHeroes")
		return &finalHeroes
	}

	for _, abilites := range heroAbilities {
		for _, hero := range heroes {
			if hero.NPCName == abilites.NPCName {
				for ability := range abilities {
					data := abilities[ability]
					hero.Abilities = append(hero.Abilities, data)
				}
			}
		}
	}
	return &finalHeroes
}

func (od *OpenDotaService) GetAllItems(pctx context.Context) *map[int]opendota.Item {
	query := "constants/items"
	data, err := od.client.GetData(pctx, query)
	rawItems := opendota.Items{}
	items := map[int]opendota.Item{}
	if err != nil {
		utils.LogError("Error when getting items: "+err.Error(), "GetAllItems")
		return &items
	}

	err = json.Unmarshal(data, &rawItems)
	if err != nil {
		utils.LogError("Error when parsing heroes: "+err.Error(), "GetAllItems")
		return &items
	}
	for _, data := range rawItems {
		items[data.Id] = data
	}
	return &items
}

func (od *OpenDotaService) GetHeroAbilities(pctx context.Context) ([]opendota.HeroAbilityData, error) {
	query := "constants/hero_abilities"
	abilities := []opendota.HeroAbilityData{}
	data, err := od.client.GetData(pctx, query)
	if err != nil {
		utils.LogError("Error when getting heroes: "+err.Error(), "GetAllItems")
		return abilities, err
	}
	err = json.Unmarshal(data, &abilities)
	if err != nil {
		utils.LogError("Error when parsing heroes: "+err.Error(), "GetAllItems")
		return abilities, err
	}
	return abilities, nil
}

func (od *OpenDotaService) GetAbilities(pctx context.Context) (opendota.AbilityData, error) {
	query := "constants/abilites"
	abilities := opendota.AbilityData{}
	data, err := od.client.GetData(pctx, query)
	if err != nil {
		utils.LogError("Error when getting heroes: "+err.Error(), "GetAllItems")
		return abilities, err
	}
	err = json.Unmarshal(data, &abilities)
	if err != nil {
		utils.LogError("Error when parsing heroes: "+err.Error(), "GetAllItems")
		return abilities, err
	}
	return abilities, nil
}

func (od *OpenDotaService) GetLatestMatches(pctx context.Context) ([]int, error) {
	query := "live"
	data, err := od.client.GetData(pctx, query)
	matches := []opendota.Match{}
	matchIDs := []int{}
	if err != nil {
		utils.LogError("Error when getting matches: "+err.Error(), "GetLatestMatches")
		return matchIDs, err
	}
	err = json.Unmarshal(data, &matches)
	if err != nil {
		utils.LogError("Error when parsing match details: "+err.Error(), "GetLatestMatches")
		return matchIDs, err
	}

	for _, v := range matches {
		matchIDs = append(matchIDs, v.MatchID)
	}
	return matchIDs, nil
}

func (od *OpenDotaService) GetMatchDetails(pctx context.Context, matchID int) (opendota.MatchDetails, error) {
	query := "matches/%d"
	fmtQuery := fmt.Sprintf(query, matchID)
	data, err := od.client.GetData(pctx, fmtQuery)
	matchDetails := opendota.MatchDetails{}
	if err != nil {
		utils.LogError("Error when getting match details: "+err.Error(), "GetMatchDetails")
		return matchDetails, err
	}
	err = json.Unmarshal(data, &matchDetails)
	if err != nil {
		utils.LogError("Error when parsing match details: "+err.Error(), "GetMatchDetails")
		return matchDetails, err
	}
	return matchDetails, nil
}

func NewOpenDotaService(isPremium bool) *OpenDotaService {
	return &OpenDotaService{
		Name:      "OpenDotaService",
		isPremium: isPremium,
		client:    *utils.NewHttpClient(baseURL, 60),
	}
}
