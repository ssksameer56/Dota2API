package handlers

import (
	"context"
	"errors"

	"github.com/ssksameer56/Dota2API/models/common"
	odmodels "github.com/ssksameer56/Dota2API/models/opendota"
	"github.com/ssksameer56/Dota2API/opendota-client"
)

type Dota2Handler struct {
	dota2service *opendota.OpenDotaService
	GameData     *common.Dota2GameInfo
}

func (dh *Dota2Handler) GetAllHeroes() []odmodels.Hero {
	heroes := []odmodels.Hero{}
	for _, val := range *dh.GameData.Heroes {
		heroes = append(heroes, val)
	}
	return heroes
}

func (dh *Dota2Handler) GetAllItems() []odmodels.Item {
	items := []odmodels.Item{}
	for _, val := range *dh.GameData.Items {
		items = append(items, val)
	}
	return items
}

func (dh *Dota2Handler) GetHero(id int) (odmodels.Hero, error) {
	if len(*(dh.GameData.Heroes)) == 0 {
		return odmodels.Hero{}, errors.New("hero with that ID doesn't exist")
	}
	if val, ok := (*dh.GameData.Heroes)[id]; ok {
		return val, nil
	}
	return odmodels.Hero{}, errors.New("hero with that ID doesn't exist")
}

func (dh *Dota2Handler) GetItem(id int) (odmodels.Item, error) {
	if len(*(dh.GameData.Items)) == 0 {
		return odmodels.Item{}, errors.New("item with that ID doesn't exist")
	}
	if val, ok := (*dh.GameData.Items)[id]; ok {
		return val, nil
	}
	return odmodels.Item{}, errors.New("item with that ID doesn't exist")
}

func (dh *Dota2Handler) PopulateStaticData() error {
	dh.GameData.Heroes = dh.dota2service.GetAllHeroes(context.TODO())
	if len(*(dh.GameData.Heroes)) == 0 {
		return errors.New("got Empty Hero Data from Dota2 Service")
	}
	dh.GameData.Items = dh.dota2service.GetAllItems(context.TODO())
	if len(*(dh.GameData.Items)) == 0 {
		return errors.New("got Empty Item Data from Dota2 Service")
	}
	return nil
}
