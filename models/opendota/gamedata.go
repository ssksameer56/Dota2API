package opendota

type ability struct {
	Name            string `json:"dname"`
	DamageType      string `json:"dmg_type"`
	Description     string `json:"desc"`
	AbilityBehavior string `json:"behavior"`
	ExtraData       interface{}
}

type Hero struct {
	Id        int    `json:"id"`
	HeroName  string `json:"localized_name"`
	Abilities []ability
	NPCName   string `json:"name"`
}
type Item struct {
	Name string   `json:"dname"`
	Id   int      `json:"id"`
	Cost int      `json:"cost"`
	Hint []string `json:"hint"`
}

type InGameHero struct {
}

//Intermediate struct to handle data returned by OpenDota service
type HeroAbilityData struct {
	NPCName      string
	AbilityNames []string `json:"abilities"`
	ExtraData    interface{}
}

//Abilities as returned by opendota in raw format
type AbilityData map[string]ability

//Items as returned by OpenDota in raw format
type Items map[string]Item

type Heroes map[string]Hero
