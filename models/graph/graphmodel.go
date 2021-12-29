// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// Hero defined as per Dota 2
type Hero struct {
	Name             string         `json:"name"`
	HeroName         string         `json:"heroName"`
	PrimaryAttribute Attribute      `json:"primaryAttribute"`
	Roles            []Role         `json:"roles"`
	Abilities        []*HeroAbility `json:"abilities"`
}

type HeroAbility struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	DamageType  DamageType      `json:"damageType"`
	Behaviour   AbilityBehavior `json:"behaviour"`
}

// Type to Handle stats between two heroes
type HeroMatchup struct {
	Hero1      *Hero `json:"hero1"`
	Hero2      *Hero `json:"hero2"`
	WinPercent int   `json:"winPercent"`
}

// Model to map in game hero
type InGameHero struct {
	Hero    *Hero   `json:"hero"`
	Items   []*Item `json:"items"`
	Gold    int     `json:"gold"`
	Xp      int     `json:"xp"`
	Level   int     `json:"level"`
	Kills   int     `json:"kills"`
	Deaths  int     `json:"deaths"`
	Assists int     `json:"assists"`
}

// Items available in Dota 2
type Item struct {
	Name string   `json:"name"`
	ID   int      `json:"id"`
	Cost int      `json:"cost"`
	Hint []string `json:"hint"`
}

// Basic attributes of a Dota 2 Match
type Match struct {
	MatchID               int           `json:"matchID"`
	DireKills             int           `json:"direKills"`
	RadiantKills          int           `json:"radiantKills"`
	Duration              int           `json:"duration"`
	GameMode              int           `json:"gameMode"`
	GoldAdvantage         []int         `json:"goldAdvantage"`
	XpAdvantage           []int         `json:"xpAdvantage"`
	RadiantTowersKilled   int           `json:"radiantTowersKilled"`
	DireTowersKilled      int           `json:"direTowersKilled"`
	RadiantBarracksKilled int           `json:"radiantBarracksKilled"`
	DireBarracksKilled    int           `json:"direBarracksKilled"`
	RadiantHeroes         []*InGameHero `json:"radiantHeroes"`
	DireHeroes            []*InGameHero `json:"direHeroes"`
}

type AbilityBehavior string

const (
	AbilityBehaviorPassive AbilityBehavior = "PASSIVE"
	AbilityBehaviorActive  AbilityBehavior = "ACTIVE"
	AbilityBehaviorChannel AbilityBehavior = "CHANNEL"
	AbilityBehaviorToggled AbilityBehavior = "TOGGLED"
)

var AllAbilityBehavior = []AbilityBehavior{
	AbilityBehaviorPassive,
	AbilityBehaviorActive,
	AbilityBehaviorChannel,
	AbilityBehaviorToggled,
}

func (e AbilityBehavior) IsValid() bool {
	switch e {
	case AbilityBehaviorPassive, AbilityBehaviorActive, AbilityBehaviorChannel, AbilityBehaviorToggled:
		return true
	}
	return false
}

func (e AbilityBehavior) String() string {
	return string(e)
}

func (e *AbilityBehavior) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AbilityBehavior(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AbilityBehavior", str)
	}
	return nil
}

func (e AbilityBehavior) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Attack Type for a hero
type AttackType string

const (
	AttackTypeMelee  AttackType = "MELEE"
	AttackTypeRanged AttackType = "RANGED"
)

var AllAttackType = []AttackType{
	AttackTypeMelee,
	AttackTypeRanged,
}

func (e AttackType) IsValid() bool {
	switch e {
	case AttackTypeMelee, AttackTypeRanged:
		return true
	}
	return false
}

func (e AttackType) String() string {
	return string(e)
}

func (e *AttackType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AttackType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AttackType", str)
	}
	return nil
}

func (e AttackType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Attribute type for a hero
type Attribute string

const (
	AttributeStrength     Attribute = "STRENGTH"
	AttributeAgility      Attribute = "AGILITY"
	AttributeIntelligence Attribute = "INTELLIGENCE"
)

var AllAttribute = []Attribute{
	AttributeStrength,
	AttributeAgility,
	AttributeIntelligence,
}

func (e Attribute) IsValid() bool {
	switch e {
	case AttributeStrength, AttributeAgility, AttributeIntelligence:
		return true
	}
	return false
}

func (e Attribute) String() string {
	return string(e)
}

func (e *Attribute) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Attribute(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Attribute", str)
	}
	return nil
}

func (e Attribute) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DamageType string

const (
	DamageTypePhysical DamageType = "PHYSICAL"
	DamageTypeMagical  DamageType = "MAGICAL"
	DamageTypePure     DamageType = "PURE"
)

var AllDamageType = []DamageType{
	DamageTypePhysical,
	DamageTypeMagical,
	DamageTypePure,
}

func (e DamageType) IsValid() bool {
	switch e {
	case DamageTypePhysical, DamageTypeMagical, DamageTypePure:
		return true
	}
	return false
}

func (e DamageType) String() string {
	return string(e)
}

func (e *DamageType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DamageType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DamageType", str)
	}
	return nil
}

func (e DamageType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The roles available for various heroes
type Role string

const (
	RoleCarry     Role = "CARRY"
	RoleNuker     Role = "NUKER"
	RoleInitiator Role = "INITIATOR"
	RoleDisabler  Role = "DISABLER"
	RoleDurable   Role = "DURABLE"
	RoleEscape    Role = "ESCAPE"
	RoleSupport   Role = "SUPPORT"
	RolePusher    Role = "PUSHER"
	RoleJungler   Role = "JUNGLER"
)

var AllRole = []Role{
	RoleCarry,
	RoleNuker,
	RoleInitiator,
	RoleDisabler,
	RoleDurable,
	RoleEscape,
	RoleSupport,
	RolePusher,
	RoleJungler,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleCarry, RoleNuker, RoleInitiator, RoleDisabler, RoleDurable, RoleEscape, RoleSupport, RolePusher, RoleJungler:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Type of shops where items are available
type Shoptype string

const (
	ShoptypeAncient    Shoptype = "ANCIENT"
	ShoptypeSecretshop Shoptype = "SECRETSHOP"
	ShoptypeSideshop   Shoptype = "SIDESHOP"
)

var AllShoptype = []Shoptype{
	ShoptypeAncient,
	ShoptypeSecretshop,
	ShoptypeSideshop,
}

func (e Shoptype) IsValid() bool {
	switch e {
	case ShoptypeAncient, ShoptypeSecretshop, ShoptypeSideshop:
		return true
	}
	return false
}

func (e Shoptype) String() string {
	return string(e)
}

func (e *Shoptype) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Shoptype(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SHOPTYPE", str)
	}
	return nil
}

func (e Shoptype) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
