package opendota

import (
	"github.com/ssksameer56/Dota2API/models/opendota"
	"github.com/ssksameer56/Dota2API/utils"
)

type OpenDotaService struct {
	isPremium bool
	client    utils.HttpClient
	Name      string
}

func (od *OpenDotaService) GetAllHeroes() [](opendota.Hero) {
	return []opendota.Hero{}
}

func (od *OpenDotaService) GetAllItems() []opendota.Item {
	return []opendota.Item{}
}

func (od *OpenDotaService) GetHeroAbilities(heroName string) []opendota.Ability {
	return []opendota.Ability{}
}

func (od *OpenDotaService) GetHero(id int) opendota.Hero {
	return opendota.Hero{}
}
func (od *OpenDotaService) GetItem(id int) opendota.Item {
	return opendota.Item{}
}

func NewOpenDotaService(bURL string, isPremium bool) *OpenDotaService {
	return &OpenDotaService{
		Name:      "OpenDotaService",
		isPremium: isPremium,
		client:    *utils.NewHttpClient(bURL),
	}
}
