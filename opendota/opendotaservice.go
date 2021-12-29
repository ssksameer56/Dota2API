package opendota

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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
	utils.LogInfo(fmt.Sprintf("Fetched %d heroes", len(heroes)), "GetAllHeroes")
	for _, heroData := range heroes {
		finalHeroes[heroData.Id] = heroData
	}

	heroAbilities, err := od.getHeroAbilities(pctx)
	utils.LogInfo(fmt.Sprintf("Fetched %d abilities", len(heroAbilities)), "GetAllHeroes")

	if err != nil {
		utils.LogError("Error when getting abilities: "+err.Error(), "GetAllHeroes")
		return &finalHeroes
	}
	abilities, err := od.getAbilities(pctx)
	utils.LogInfo("Fetched all abilities data", "GetAllHeroes")

	if err != nil {
		utils.LogError("Error when getting abilities: "+err.Error(), "GetAllHeroes")
		return &finalHeroes
	}
	for _, abilites := range heroAbilities {
		for _, hero := range heroes {
			if hero.NPCName == abilites.NPCName {
				for _, abilityName := range abilites.AbilityNames {
					data := abilities[abilityName]
					hero.Abilities = append(hero.Abilities, *data)
				}
				finalHeroes[hero.Id] = hero
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
		utils.LogError("Error when parsing items: "+err.Error(), "GetAllItems")
		return &items
	}
	for _, data := range rawItems {
		items[data.Id] = data
	}
	utils.LogInfo(fmt.Sprintf("Got %d items", len(items)), "GetAllItems")
	return &items
}

func (od *OpenDotaService) getHeroAbilities(pctx context.Context) ([]opendota.HeroAbilityData, error) {
	query := "constants/hero_abilities"
	rawAbilityData := make(map[string]opendota.HeroAbilityData)
	abilities := []opendota.HeroAbilityData{}
	data, err := od.client.GetData(pctx, query)
	if err != nil {
		utils.LogError("Error when getting hero abilities: "+err.Error(), "GetHeroAbilities")
		return abilities, err
	}
	err = json.Unmarshal(data, &rawAbilityData)
	for npcName, rawData := range rawAbilityData {
		ability := opendota.HeroAbilityData{
			NPCName:      npcName,
			AbilityNames: rawData.AbilityNames,
		}
		abilities = append(abilities, ability)
	}
	if err != nil {
		utils.LogError("Error when parsing hero abilites: "+err.Error(), "GetHeroAbilities")
		return abilities, err
	}
	utils.LogInfo(fmt.Sprintf("Got %d abilities", len(abilities)), "GetHeroAbilities")
	return abilities, nil
}

func (od *OpenDotaService) getAbilities(pctx context.Context) (opendota.AbilityData, error) {
	query := "constants/abilities"
	abilities := opendota.AbilityData{}
	data, err := od.client.GetData(pctx, query)
	if err != nil {
		utils.LogError("Error when getting abilities: "+err.Error(), "GetAbilities")
		return abilities, err
	}
	err = json.Unmarshal(data, &abilities)
	if err != nil {
		utils.LogError("Error when parsing abilities: "+err.Error(), "GetAbilities")
		return abilities, err
	}
	for _, ability := range abilities {
		if len(ability.RawAbilityBehavior) > 0 {
			first := ability.RawAbilityBehavior[0]
			if first == '"' {
				var data string
				json.Unmarshal(ability.RawAbilityBehavior, &data)
				ability.AbilityBehavior = data
			} else if first == '[' {
				strArray := []string{}
				json.Unmarshal(ability.RawAbilityBehavior, &strArray)
				ability.AbilityBehavior = strings.Join(strArray, ",")
			}
		}
	}

	utils.LogInfo(fmt.Sprintf("Got %d abilities", len(abilities)), "GetAbilities")
	return abilities, nil
}

func (od *OpenDotaService) GetLatestMatches(pctx context.Context) ([]int64, error) {
	query := "live"
	data, err := od.client.GetData(pctx, query)
	matches := []opendota.Match{}
	matchIDs := []int64{}
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
		val, _ := strconv.Atoi(v.MatchID)
		matchIDs = append(matchIDs, int64(val))
	}
	utils.LogInfo(fmt.Sprintf("Got %d matches being played currently", len(matchIDs)), "GetLatestMatches")

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
