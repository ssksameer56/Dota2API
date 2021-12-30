package database

import "context"

type DBModelFavourites struct {
	UserID  int
	HeroIDs string
}

type Favourites interface {
	QueryFavouritesTable(context.Context, string) (*[]DBModelFavourites, error)
	ModifyFavouritesTable(context.Context, string) (int64, error)
}
