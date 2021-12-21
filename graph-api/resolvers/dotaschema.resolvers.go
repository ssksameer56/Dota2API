package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	server "github.com/ssksameer56/Dota2API/graph/api"
	model "github.com/ssksameer56/Dota2API/models/graph"
)

func (r *mutationResolver) MarkHeroAsFavourite(ctx context.Context, id *int) ([]*model.Hero, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UnMarkHeroAsFavourite(ctx context.Context, id *int) ([]*model.Hero, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllHeroes(ctx context.Context) ([]*model.Hero, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetHero(ctx context.Context, name *string) (*model.Hero, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetItem(ctx context.Context, name *string) (*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMatchDetails(ctx context.Context, ids []int) ([]*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) GetLiveMatchIDs(ctx context.Context) (<-chan []int, error) {
	panic(fmt.Errorf("not implemented"))
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
