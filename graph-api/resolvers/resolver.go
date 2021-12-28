package resolvers

// This file will not be regenerated automatically.
//go:generate go run github.com/99designs/gqlgen
// It serves as dependency injection for your app, add any dependencies you require here.

import "github.com/ssksameer56/Dota2API/handlers"

type Resolver struct {
	favouritesService   handlers.FavouritesHandler
	constantDataService handlers.Dota2Handler
	matchDataService    handlers.MatchDataHandler
}
