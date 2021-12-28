package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	server "github.com/ssksameer56/Dota2API/graph-api/api"
	model "github.com/ssksameer56/Dota2API/models/graph"
	"github.com/ssksameer56/Dota2API/utils"
)

func (r *mutationResolver) MarkHeroAsFavourite(ctx context.Context, heroID int, userID int) (bool, error) {
	_, err := r.favouritesService.MarkFavouritesForAUser(ctx, userID, []int{heroID})
	if err != nil {
		return false, errors.New("could not mark hero as favourite")
	}
	return true, err
}

func (r *mutationResolver) UnMarkHeroAsFavourite(ctx context.Context, heroID int, userID int) (bool, error) {
	existingIDs, err := r.favouritesService.QueryFavouritesOfAUser(ctx, userID)
	if err != nil {
		return false, errors.New("could not mark hero as favourite")
	} else if len(existingIDs) == 0 {
		return false, errors.New("no favourites exist for this user")
	}
	removedData := []int{}
	for _, v := range existingIDs {
		if v == heroID {
			continue
		}
		removedData = append(removedData, v)
	}
	_, err = r.favouritesService.MarkFavouritesForAUser(ctx, userID, removedData)
	if err != nil {
		return false, err
	}
	return true, err
}

func (r *queryResolver) GetAllHeroes(ctx context.Context) ([]*model.Hero, error) {
	allHeroes := r.constantDataService.GetAllHeroes()
	if len(allHeroes) == 0 {
		hero := model.Hero{}
		return []*model.Hero{&hero}, errors.New("no heroes present")
	}
	heroData := []*model.Hero{}
	for _, rawHero := range allHeroes {
		hero := TransformHero(rawHero)
		heroData = append(heroData, hero)
	}
	return heroData, nil
}

func (r *queryResolver) GetHero(ctx context.Context, name *string) (*model.Hero, error) {
	rawHero, err := r.constantDataService.GetHero(*name)
	if err != nil {
		return &model.Hero{}, err
	}
	return TransformHero(rawHero), nil
}

func (r *queryResolver) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	rawItems := r.constantDataService.GetAllItems()
	items := []*model.Item{}
	for _, rawItem := range rawItems {
		item := TransformItem(rawItem)
		items = append(items, item)
	}
	return items, nil
}

func (r *queryResolver) GetItem(ctx context.Context, name *string) (*model.Item, error) {
	rawItem, err := r.constantDataService.GetItem(*name)
	if err != nil {
		return &model.Item{}, err
	}
	return TransformItem(rawItem), nil
}

func (r *queryResolver) GetMatchDetails(ctx context.Context, ids []int) ([]*model.Match, error) {
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()
	matchesData := []*model.Match{}
	allHeroes, err := r.constantDataService.GetHeroDictionary()
	if err != nil {
		utils.LogError(err.Error(), "GraphAPI")
		return matchesData, err
	}
	allItems, err := r.constantDataService.GetItemDictionary()
	if err != nil {
		utils.LogError(err.Error(), "GraphAPI")
		return matchesData, err
	}
	for _, id := range ids {
		data := r.matchDataService.GetMatchDetails(cctx, id)
		finalMatch := TransformMatch(&data, allItems, allHeroes)
		matchesData = append(matchesData, finalMatch)
	}
	return matchesData, nil
}

func (r *subscriptionResolver) GetLiveMatchIDs(ctx context.Context) (<-chan []int, error) {
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()
	data := r.matchDataService.GetLiveMatchIDs(cctx)
	matchIDchan := make(chan []int)
	matchIDchan <- data
	return matchIDchan, nil
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

// Subscription returns server.SubscriptionResolver implementation.
func (r *Resolver) Subscription() server.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
