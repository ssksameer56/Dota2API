package handlers

import (
	"context"

	odmodels "github.com/ssksameer56/Dota2API/models/opendota"
	"github.com/ssksameer56/Dota2API/opendota-client"
)

type Dota2Handler struct {
	dota2service *opendota.OpenDotaService
}

func (dh *Dota2Handler) GetAllHeroes() []odmodels.Hero {
	return dh.dota2service.GetAllHeroes(context.TODO())
}

func (dh *Dota2Handler) GetAllItems() []odmodels.Item {
	return dh.dota2service.GetAllItems(context.TODO())
}

func (dh *Dota2Handler) GetHero(id int) odmodels.Hero {
	return odmodels.Hero{}
}

func (dh *Dota2Handler) GetItem(id int) odmodels.Item {
	return odmodels.Item{}
}
