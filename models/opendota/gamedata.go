package opendota

type Ability = string

type Hero struct {
	HeroName  string
	Abilities Ability
	NPCName   string
	Talents   Ability
}

type Item struct {
}

type InGameHero struct {
}

//Intermediate struct to handle data returned by OpenDota service
type RawAbilityData struct {
	NPCName  string
	HeroData map[string]Ability
}
