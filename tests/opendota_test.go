package tests

import (
	"context"
	"testing"

	"github.com/ssksameer56/Dota2API/opendota"
)

func TestGetAllHeroes(t *testing.T) {
	odClient := opendota.NewOpenDotaService(false)
	data := odClient.GetAllHeroes(context.Background())
	if data == nil {
		t.Errorf("Could not fetch the heroes")
		t.FailNow()
	}
	test_hero := (*data)[1]
	if test_hero.HeroName == "" {
		t.Errorf("Could not fetch the heroes")
		t.Fail()
	}
	if len(test_hero.Abilities) == 0 {
		t.Errorf("Could not fetch the abilities for hero")
		t.FailNow()
	}
	if test_hero.Abilities[0].Name == "" {
		t.Errorf("Could not fetch the abilities")
		t.Fail()
	}
	if len(test_hero.Roles) == 0 {
		t.Errorf("could not fetch herodata")
		t.FailNow()
	}
}
func TestGetLiveMatches(t *testing.T) {
	odClient := opendota.NewOpenDotaService(false)
	ids, err := odClient.GetLatestMatches(context.Background())
	if err != nil {
		t.FailNow()
	}
	if len(ids) < 1 {
		t.FailNow()
	}
}
func TestGetMatchDetails(t *testing.T) {
	odClient := opendota.NewOpenDotaService(false)
	matchID := 4080778303 //OG LGD TI8 Game 4
	data, err := odClient.GetMatchDetails(context.Background(), int(matchID))
	if err != nil {
		t.Errorf("Couldn't get match details:" + err.Error())
		t.FailNow()
	}
	if data.MatchID != int(matchID) {
		t.FailNow()
	}
	if len(data.PlayerDetails) < 10 {
		t.FailNow()
	}
}
