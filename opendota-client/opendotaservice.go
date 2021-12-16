package opendota

import (
	"context"
	"encoding/json"

	"github.com/ssksameer56/Dota2API/models/opendota"
	"github.com/ssksameer56/Dota2API/utils"
)

type OpenDotaService struct {
	isPremium bool
	client    utils.HttpClient
	Name      string
}

var baseURL = "https://api.opendota.com/api/"

func (od *OpenDotaService) GetAllHeroes(pctx context.Context) [](opendota.Hero) {
	query := "heroes"
	data, err := od.client.GetData(pctx, query)
	if err != nil {
		utils.LogError("Error when getting heroes: "+err.Error(), "GetAllHeroes")
		return []opendota.Hero{}
	}
	heroes := []opendota.Hero{}
	err = json.Unmarshal(data, &heroes)
	if err != nil {
		utils.LogError("Error when parsing heroes: "+err.Error(), "GetAllHeroes")
		return []opendota.Hero{}
	}

	//Map all the abilities to hero
	abilities, err := od.GetHeroAbilities(pctx)
	if err != nil {
		utils.LogError("Error when getting abilities: "+err.Error(), "GetAllHeroes")
		return []opendota.Hero{}
	}
	for _, ability := range abilities {
		for _, hero := range heroes {
			if hero.NPCName == ability.NPCName {
				hero.Abilities = ability.HeroData["abilities"]
				hero.Talents = ability.HeroData["talents"]
			}
		}
	}
	return heroes
}

func (od *OpenDotaService) GetAllItems(pctx context.Context) []opendota.Item {
	query := "items"
	data, err := od.client.GetData(pctx, query)
	items := []opendota.Item{}
	if err != nil {
		utils.LogError("Error when getting heroes: "+err.Error(), "GetAllItems")
		return items
	}

	err = json.Unmarshal(data, &items)
	if err != nil {
		utils.LogError("Error when parsing heroes: "+err.Error(), "GetAllItems")
		return items
	}
	return items
}

func (od *OpenDotaService) GetHeroAbilities(pctx context.Context) ([]opendota.RawAbilityData, error) {
	query := "constants/hero_abilities"
	abilities := []opendota.RawAbilityData{}
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

func (od *OpenDotaService) GetLatestMatches(pctx context.Context) []int {
	query := "live_matches"
	data, err := od.client.GetData(pctx, query)
	matches := []int{}
	if err != nil {
		utils.LogError("Error when getting heroes: "+err.Error(), "GetAllHeroes")
		return matches
	}
	err = json.Unmarshal(data, &matches)
	if err != nil {
		utils.LogError("Error when parsing heroes: "+err.Error(), "GetAllHeroes")
		return matches
	}
	return matches
}

func NewOpenDotaService(isPremium bool) *OpenDotaService {
	return &OpenDotaService{
		Name:      "OpenDotaService",
		isPremium: isPremium,
		client:    *utils.NewHttpClient(baseURL, 60),
	}
}
