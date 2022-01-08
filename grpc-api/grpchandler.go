package grpcapi

import (
	"context"

	"github.com/ssksameer56/Dota2API/handlers"
	"github.com/ssksameer56/Dota2API/models/grpc"
)

type GrpcServer struct {
	grpc.UnimplementedDota2ServiceServer
	Dota2Handler      *handlers.Dota2Handler
	MatchDataHandler  *handlers.MatchDataHandler
	FavouritesHandler *handlers.FavouritesHandler
}

func (s *GrpcServer) GetAllHeroes(pctx context.Context, req *grpc.GetAllHeroesRequest) (*grpc.GetAllHeroesResponse, error) {
	rawData := s.Dota2Handler.GetAllHeroes()
	data := []*grpc.Hero{}
	for _, hero := range rawData {
		mHero := TransformHero(hero)
		data = append(data, mHero)
	}
	response := grpc.GetAllHeroesResponse{
		Hero: data,
	}
	return &response, nil
}
func (s *GrpcServer) GetHero(pctx context.Context, req *grpc.GetHeroRequest) (*grpc.Hero, error) {
	rawHero, err := s.Dota2Handler.GetHero(req.HeroName)
	if err != nil {
		return nil, err
	}
	return TransformHero(rawHero), nil
}
func (s *GrpcServer) GetAllItems(context.Context, *grpc.GetAllItemsRequest) (*grpc.GetAllItemsResponse, error) {
	rawData := s.Dota2Handler.GetAllItems()
	data := []*grpc.Item{}
	for _, item := range rawData {
		mItem := TransformItem(item)
		data = append(data, mItem)
	}
	response := grpc.GetAllItemsResponse{
		Items: data,
	}
	return &response, nil
}
func (s *GrpcServer) GetItem(pctx context.Context, req *grpc.GetItemRequest) (*grpc.Item, error) {
	rawItem, err := s.Dota2Handler.GetItem(req.Name)
	if err != nil {
		return nil, err
	}
	return TransformItem(rawItem), nil
}
func (s *GrpcServer) GetLiveMatches(req *grpc.MatchIDsRequest, stream grpc.Dota2Service_GetLiveMatchesServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	data, err := s.MatchDataHandler.GetLiveMatchIDs(ctx)
	if err != nil {
		stream.Send(&grpc.MatchIDsResponse{
			MatchID: 0,
		})
		return err
	}
	if stream.Context().Err() == context.Canceled {
		return nil
	}
	for _, val := range data {
		response := grpc.MatchIDsResponse{
			MatchID: int32(val),
		}
		stream.Send(&response)
	}
	return nil
}
func (s *GrpcServer) GetMatchDetails(req *grpc.MatchDetailsRequest, stream grpc.Dota2Service_GetMatchDetailsServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	itemData, err := s.Dota2Handler.GetItemDictionary()
	if err != nil {
		stream.Send(&grpc.MatchDetailsResponse{})
		return err
	}
	heroData, err := s.Dota2Handler.GetHeroDictionary()
	if err != nil {
		stream.Send(&grpc.MatchDetailsResponse{})
		return err
	}
	for {
		data, err := s.MatchDataHandler.GetMatchDetails(ctx, int(req.MatchID))
		if stream.Context().Err() == context.Canceled {
			return nil
		}
		if err != nil {
			stream.Send(&grpc.MatchDetailsResponse{})
			return err
		}
		stream.Send(TransformMatchDetails(&data, itemData, heroData))
	}
}
