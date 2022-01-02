package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	server "github.com/ssksameer56/Dota2API/graph-api/api"
	model "github.com/ssksameer56/Dota2API/models/graph"
	"github.com/ssksameer56/Dota2API/utils"
)

func (r *mutationResolver) MarkHeroAsFavourite(ctx context.Context, heroID int, userID int) (bool, error) {
	_, err := r.FavouritesService.MarkFavouritesForAUser(ctx, userID, []int{heroID})
	if err != nil {
		utils.LogError("could not mark hero as favourite:"+err.Error(), "Graph Resolver")
		return false, errors.New("could not mark hero as favourite")
	}
	return true, err
}

func (r *mutationResolver) UnMarkHeroAsFavourite(ctx context.Context, heroID int, userID int) (bool, error) {
	existingIDs, err := r.FavouritesService.QueryFavouritesOfAUser(ctx, userID)
	if err != nil {
		utils.LogError("could not unmark hero as favourite"+err.Error(), "Graph Resolver")
		return false, errors.New("could not unmark hero as favourite")
	} else if len(existingIDs) == 0 {
		utils.LogError("no favourites exist for this user"+err.Error(), "Graph Resolver")
		return false, errors.New("no favourites exist for this user")
	}
	removedData := []int{}
	for _, v := range existingIDs {
		if v == heroID {
			continue
		}
		removedData = append(removedData, v)
	}
	_, err = r.FavouritesService.MarkFavouritesForAUser(ctx, userID, removedData)
	if err != nil {
		utils.LogError("couldn't update favourites for the user"+err.Error(), "Graph Resolver")
		return false, err
	}
	return true, err
}

func (r *queryResolver) GetAllHeroes(ctx context.Context) ([]*model.Hero, error) {
	allHeroes := r.ConstantDataService.GetAllHeroes()
	if len(allHeroes) == 0 {
		hero := model.Hero{}
		utils.LogError("no heroes present", "Graph Resolver")
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
	rawHero, err := r.ConstantDataService.GetHero(*name)
	if err != nil {
		utils.LogError(fmt.Sprintf("no hero present with name %s : %s", *name, err.Error()), "Graph Resolver")
		return &model.Hero{}, err
	}
	return TransformHero(rawHero), nil
}

func (r *queryResolver) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	rawItems := r.ConstantDataService.GetAllItems()
	items := []*model.Item{}
	if len(rawItems) == 0 {
		utils.LogError("no items present", "Graph Resolver")
		return []*model.Item{}, errors.New("no items present")
	}
	for _, rawItem := range rawItems {
		item := TransformItem(rawItem)
		items = append(items, item)
	}
	return items, nil
}

func (r *queryResolver) GetItem(ctx context.Context, name *string) (*model.Item, error) {
	rawItem, err := r.ConstantDataService.GetItem(*name)
	if err != nil {
		utils.LogError(fmt.Sprintf("no item present with name %s : %s", *name, err.Error()), "Graph Resolver")
		return &model.Item{}, err
	}
	return TransformItem(rawItem), nil
}

func (r *queryResolver) GetMatchDetails(ctx context.Context, ids []int) ([]*model.Match, error) {
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()
	matchesData := []*model.Match{}
	allHeroes, err := r.ConstantDataService.GetHeroDictionary()
	if err != nil {
		utils.LogError("couldnt fetch heroes"+err.Error(), "Graph Resolver")
		return matchesData, err
	}
	allItems, err := r.ConstantDataService.GetItemDictionary()
	if err != nil {
		utils.LogError("couldnt fetch items"+err.Error(), "Graph Resolver")
		return matchesData, err
	}
	for _, id := range ids {
		data, err := r.MatchDataService.GetMatchDetails(cctx, id)
		if err != nil {
			utils.LogError(fmt.Sprintf("couldnt fetch details for matchID %d : %s", id, err.Error()), "Graph Resolver")
			return []*model.Match{}, err
		}
		finalMatch := TransformMatch(&data, allItems, allHeroes)
		matchesData = append(matchesData, finalMatch)
	}
	return matchesData, nil
}

func (r *subscriptionResolver) GetLiveMatchIDs(ctx context.Context) (<-chan []int, error) {
	cctx, cancel := context.WithCancel(ctx)
	matchIDchan := make(chan []int)
	go func() {
		for {
			data, err := r.MatchDataService.GetLiveMatchIDs(cctx)
			if err != nil {
				utils.LogError("couldnt fetch live match IDs"+err.Error(), "Graph Resolver")
				cancel()
				return
			}
			matchIDs := []int{}
			for _, val := range data {
				matchIDs = append(matchIDs, int(val))
			}
			matchIDchan <- matchIDs
			time.Sleep(time.Second * 2)
		}
	}()
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
