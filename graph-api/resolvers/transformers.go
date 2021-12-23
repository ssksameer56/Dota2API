package resolvers

import (
	model "github.com/ssksameer56/Dota2API/models/graph"
	"github.com/ssksameer56/Dota2API/models/opendota"
)

func TransformHero(hero *opendota.Hero) *model.Hero {
	return &model.Hero{
		Name:             hero.HeroName,
		PrimaryAttribute: model.Attribute(hero.PrimaryAttribute),
		Abilities:        TransformAbilities(hero),
	}
}

func TransformAbilities(hero *opendota.Hero) []*model.HeroAbility {
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

func TransformItem(item *opendota.Item) *model.Item {
	return &model.Item{
		Name: item.Name,
		Hint: item.Hint,
		Cost: item.Cost,
		ID:   item.Id,
	}
}
