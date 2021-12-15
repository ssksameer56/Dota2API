package handlers

import (
	odmodels "github.com/ssksameer56/Dota2API/models/opendota"
	"github.com/ssksameer56/Dota2API/opendota-client"
)

type Dota2Handler struct {
	dota2service *opendota.OpenDotaService
}

func (dh *Dota2Handler) GetAllHeroes() []odmodels.Hero {
	return dh.dota2service.GetAllHeroes()
}

func (dh *Dota2Handler) GetAllItems() []odmodels.Item {
	return dh.dota2service.GetAllItems()
}

func (dh *Dota2Handler) GetHero(id int) odmodels.Hero {
	return dh.dota2service.GetHero(id)
}

func (dh *Dota2Handler) GetItem(id int) odmodels.Item {
	return dh.dota2service.GetItem(id)
}
