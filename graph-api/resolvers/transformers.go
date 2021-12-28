package resolvers

import (
	model "github.com/ssksameer56/Dota2API/models/graph"
	"github.com/ssksameer56/Dota2API/models/opendota"
)

func TransformHero(hero opendota.Hero) *model.Hero {
	return &model.Hero{
		Name:             hero.HeroName,
		PrimaryAttribute: model.Attribute(hero.PrimaryAttribute),
		Abilities:        TransformAbilities(hero),
	}
}

func TransformAbilities(hero opendota.Hero) []*model.HeroAbility {
	rawAbilities := hero.Abilities
	abilities := []*model.HeroAbility{}
	for _, rawAbility := range rawAbilities {
		ability := model.HeroAbility{
			Name:        rawAbility.Name,
			Description: rawAbility.Description,
			DamageType:  model.DamageType(rawAbility.DamageType),
			Behaviour:   model.AbilityBehavior(rawAbility.AbilityBehavior),
		}
		abilities = append(abilities, &ability)
	}
	return abilities
}

func TransformItem(item opendota.Item) *model.Item {
	return &model.Item{
		Name: item.Name,
		Hint: item.Hint,
		Cost: item.Cost,
		ID:   item.Id,
	}
}

func TransformMatch(match *opendota.MatchDetails, odItems map[int]opendota.Item, odHeroes map[int]opendota.Hero) *model.Match {
	finalMatch := model.Match{}
	radiantHeroes := []*model.InGameHero{}
	direHeroes := []*model.InGameHero{}
	for _, hero := range match.PlayerDetails {
		igHero := model.InGameHero{}
		items := []*model.Item{}
		items = append(items, TransformItem(odItems[hero.Item1ID]))
		items = append(items, TransformItem(odItems[hero.Item2ID]))
		items = append(items, TransformItem(odItems[hero.Item3ID]))
		items = append(items, TransformItem(odItems[hero.Item4ID]))
		items = append(items, TransformItem(odItems[hero.Item5ID]))
		items = append(items, TransformItem(odItems[hero.Item6ID]))
		igHero.Items = items
		igHero.Hero = TransformHero(odHeroes[hero.HeroID])
		igHero.Gold = hero.Gold
		igHero.Xp = hero.Xp
		igHero.Kills = hero.Kills
		igHero.Deaths = hero.Deaths
		igHero.Assists = hero.Assists
		igHero.Level = hero.Level
		if hero.IsRadiant {
			radiantHeroes = append(radiantHeroes, &igHero)
		} else {
			direHeroes = append(direHeroes, &igHero)
		}
	}
	finalMatch.MatchID = match.MatchID
	finalMatch.DireKills = match.DireKills
	finalMatch.RadiantKills = match.RadiantKills
	finalMatch.Duration = match.Duration
	finalMatch.GameMode = match.GameMode
	finalMatch.GoldAdvantage = match.RadiantGoldAdvantage
	finalMatch.XpAdvantage = match.RadiantXPAdvantage
	//TODO : Bitmask tower kills
	finalMatch.RadiantTowersKilled = match.RadiantTowersKilled
	finalMatch.DireTowersKilled = match.DireTowersKilled
	finalMatch.RadiantBarracksKilled = match.RadiantBarrackStatus
	finalMatch.DireBarracksKilled = match.DireBarracksStatus
	finalMatch.RadiantHeroes = radiantHeroes
	finalMatch.DireHeroes = direHeroes
	return &finalMatch
}
