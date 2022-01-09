package grpcclient

import (
	"context"
	"errors"
	"io"
	"time"

	dota2grpc "github.com/ssksameer56/Dota2API/models/grpc"
	"github.com/ssksameer56/Dota2API/utils"
	"google.golang.org/grpc"
)

type Dota2Client struct {
	RPCClient  dota2grpc.Dota2ServiceClient
	RPCOptions []grpc.CallOption
}

func (dc *Dota2Client) GetAllHeroes() (*dota2grpc.GetAllHeroesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	data, err := dc.RPCClient.GetAllHeroes(ctx, &dota2grpc.GetAllHeroesRequest{}, dc.RPCOptions...)
	if err != nil {
		utils.LogError("Error from RPC Server:"+err.Error(), "GRPC Client")
		return nil, errors.New("error when getting heroes")
	}
	return data, nil
}

func (dc *Dota2Client) GetHero(heroName string) (*dota2grpc.Hero, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := dota2grpc.GetHeroRequest{HeroName: heroName}
	data, err := dc.RPCClient.GetHero(ctx, &req, dc.RPCOptions...)
	if err != nil {
		utils.LogError("Error from RPC Server:"+err.Error(), "GRPC Client")
		return nil, errors.New("error when getting heroes")
	}
	return data, nil
}

func (dc *Dota2Client) GetAllItems() (*dota2grpc.GetAllItemsResponse, error) {
	panic(1)
}

func (dc *Dota2Client) GetItem(itemName string) (*dota2grpc.Item, error) {
	panic(1)
}

func (dc *Dota2Client) GetLiveMatches() (<-chan int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := dota2grpc.MatchIDsRequest{}
	data, err := dc.RPCClient.GetLiveMatches(ctx, &req, dc.RPCOptions...)
	if err != nil {
		utils.LogError("Error from RPC Server:"+err.Error(), "GRPC Client")
		return nil, errors.New("error when getting heroes")
	}
	dataChan := make(chan int, 200)
	go func() {
		for {
			data, err := data.Recv()
			if err == io.EOF {
				return
			}
			dataChan <- int(data.MatchID)
		}
	}()
	return dataChan, nil
}

func (dc *Dota2Client) StreamMatchDetails(matchID int64) (<-chan *dota2grpc.MatchDetailsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := dota2grpc.MatchDetailsRequest{MatchID: matchID}
	data, err := dc.RPCClient.GetMatchDetails(ctx, &req, dc.RPCOptions...)
	if err != nil {
		utils.LogError("Error from RPC Server:"+err.Error(), "GRPC Client")
		return nil, errors.New("error when getting heroes")
	}
	dataChan := make(chan *dota2grpc.MatchDetailsResponse, 200)
	go func() {
		for {
			data, err := data.Recv()
			if err == io.EOF {
				return
			}
			dataChan <- data
		}
	}()
	return dataChan, nil
}
