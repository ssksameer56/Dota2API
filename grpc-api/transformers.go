package grpcapi

import (
	"github.com/ssksameer56/Dota2API/models/grpc"
	"github.com/ssksameer56/Dota2API/models/opendota"
)

func TransformHero(hero opendota.Hero) *grpc.Hero {
	roles := []grpc.Role{}
	for _, v := range hero.Roles {
		roles = append(roles, grpc.Role(grpc.Role_value[v]))
	}
	abilities := TransformAbilities(&hero)

	return &grpc.Hero{
		Name:          hero.NPCName,
		HeroName:      hero.HeroName,
		HeroAttribute: grpc.Attribute(grpc.Attribute_value[hero.PrimaryAttribute]),
		Role:          roles,
		Abilities:     abilities,
	}
}

func TransformItem(item opendota.Item) *grpc.Item {
	return &grpc.Item{
		Name: item.Name,
		Hint: item.Hint,
		Cost: int32(item.Cost),
		Id:   int32(item.Id),
	}
}

func TransformAbilities(hero *opendota.Hero) []*grpc.HeroAbility {
	rawAbilities := hero.Abilities
	abilities := []*grpc.HeroAbility{}
	for _, rawAbility := range rawAbilities {
		ability := grpc.HeroAbility{
			Name:            rawAbility.Name,
			Description:     rawAbility.Description,
			DamageType:      grpc.DamageType(grpc.DamageType_value[rawAbility.DamageType]),
			AbilityBehavior: grpc.AbilityBehavior(grpc.AbilityBehavior_value[rawAbility.AbilityBehavior]),
		}
		abilities = append(abilities, &ability)
	}
	return abilities
}

func TransformMatchDetails(match *opendota.MatchDetails, odItems map[int]opendota.Item, odHeroes map[int]opendota.Hero) *grpc.MatchDetailsResponse {
	finalMatch := grpc.MatchDetailsResponse{}
	radiantHeroes := []*grpc.InGameHero{}
	direHeroes := []*grpc.InGameHero{}
	for _, hero := range match.PlayerDetails {
		igHero := grpc.InGameHero{}
		items := []*grpc.Item{}
		items = append(items, TransformItem(odItems[hero.Item1ID]))
		items = append(items, TransformItem(odItems[hero.Item2ID]))
		items = append(items, TransformItem(odItems[hero.Item3ID]))
		items = append(items, TransformItem(odItems[hero.Item4ID]))
		items = append(items, TransformItem(odItems[hero.Item5ID]))
		items = append(items, TransformItem(odItems[hero.Item6ID]))
		igHero.Items = items
		igHero.Hero = TransformHero(odHeroes[hero.HeroID])
		igHero.Gold = int32(hero.Gold)
		igHero.XP = int32(hero.Xp)
		igHero.Kills = int32(hero.Kills)
		igHero.Deaths = int32(hero.Deaths)
		igHero.Assists = int32(hero.Assists)
		igHero.Level = int32(hero.Level)
		if hero.IsRadiant {
			radiantHeroes = append(radiantHeroes, &igHero)
		} else {
			direHeroes = append(direHeroes, &igHero)
		}
	}
	finalMatch.MatchID = int32(match.MatchID)
	finalMatch.DireKills = int32(match.DireKills)
	finalMatch.RadiantKills = int32(match.RadiantKills)
	finalMatch.Duration = float32(match.Duration)
	finalMatch.GameMode = int32(match.GameMode)
	finalMatch.GoldAdvantage = int32(match.RadiantCurrentGoldAdvantage)
	finalMatch.XPAdvantage = int32(match.RadiantCurrentXPAdvantage)
	//TODO : Bitmask tower kills
	finalMatch.RadiantTowersKilled = int32(match.RadiantTowersKilled)
	finalMatch.DireTowersKilled = int32(match.DireTowersKilled)
	finalMatch.RadiantBarracksKilled = int32(match.RadiantBarrackStatus)
	finalMatch.DireBaracksKilled = int32(match.DireBarracksStatus)
	finalMatch.RadiantHeroes = radiantHeroes
	finalMatch.DireHeros = direHeroes
	return &finalMatch
}
