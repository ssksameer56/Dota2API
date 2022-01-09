package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	grpcclient "github.com/ssksameer56/Dota2API/clients/grpc-client"
	"github.com/ssksameer56/Dota2API/models"
	"github.com/ssksameer56/Dota2API/utils"

	dota2grpc "github.com/ssksameer56/Dota2API/models/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var config models.Configuration
	loadConfig(&config)

	utils.LogFilePath = config.ClientLogFile
	utils.InitializeLogging()
	creds, err := credentials.NewClientTLSFromFile("../"+config.SSLCertificateLocation, "localhost")
	if err != nil {
		utils.LogFatal(fmt.Sprintf("failed to load credentials: %v", err), "GRPC Client")
		return
	}
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(UnaryLogger),
		grpc.WithStreamInterceptor(StreamLogger),
	}
	conn, err := grpc.Dial("localhost:"+config.GrpcAPIPort, options...)
	if err != nil {
		utils.LogInfo("Error when starting client", "GRPC Client")
	}
	rpcclient := dota2grpc.NewDota2ServiceClient(conn)

	dota2client := grpcclient.Dota2Client{
		RPCClient:  rpcclient,
		RPCOptions: []grpc.CallOption{},
	}
	//for loop for continous use
	opToDo, _ := strconv.Atoi(os.Args[1])
	switch opToDo {
	case 0:
		data, err := dota2client.GetAllHeroes()
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("%v", data)
		}
	case 1:
		heroName := os.Args[2]
		if heroName == "" {
			fmt.Println("Provide heroname as argument")
			return
		}
		data, err := dota2client.GetHero(heroName)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("%v", data)
		}
	case 2:
		data, err := dota2client.GetAllItems()
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("%v", data)
		}
	case 3:
		itemName := os.Args[2]
		if itemName == "" {
			fmt.Println("Provide itemName as argument")
			return
		}
		data, err := dota2client.GetItem(itemName)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("%v", data)
		}
	case 4:
		stream, err := dota2client.GetLiveMatches()
		if err != nil {
			fmt.Println(err)
			return
		} else {
			for id := range stream {
				fmt.Printf("%d, ", id)
			}
		}
	case 5:
		matchID, _ := strconv.Atoi(os.Args[2])
		if matchID == 0 {
			fmt.Println("Provide itemName as argument")
			return
		}
		stream, err := dota2client.StreamMatchDetails(int64(matchID))
		if err != nil {
			fmt.Println(err)
			return
		} else {
			for details := range stream {
				fmt.Printf("%d %d %d %f ", details.MatchID, details.DireKills, details.RadiantKills, details.Duration)
			}
		}
	}
}

func UnaryLogger(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	utils.LogInfo(fmt.Sprintf("Sending request to: %s - %v", method, req), "GRPC Client")
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func StreamLogger(ctx context.Context, desc *grpc.StreamDesc,
	cc *grpc.ClientConn, method string,
	streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	utils.LogInfo(fmt.Sprintf("Sending request to: %s - %v", method, desc), "GRPC Client")
	stream, err := streamer(ctx, desc, cc, method, opts...)
	return stream, err
}

func loadConfig(config *models.Configuration) {
	file, err := os.Open("../../config.json")
	if err != nil {
		fmt.Println("Cant open config" + err.Error())
		panic(1)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Cant read config" + err.Error())
		panic(1)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Cant parse config" + err.Error())
		panic(1)
	}
	utils.LogInfo("Initialised the Config", "loadConfig")
}
