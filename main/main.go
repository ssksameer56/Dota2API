package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ssksameer56/Dota2API/database"
	graphapi "github.com/ssksameer56/Dota2API/graph-api"
	grpcapi "github.com/ssksameer56/Dota2API/grpc-api"
	"github.com/ssksameer56/Dota2API/handlers"
	"github.com/ssksameer56/Dota2API/models"
	"github.com/ssksameer56/Dota2API/models/common"
	"github.com/ssksameer56/Dota2API/opendota"
	"github.com/ssksameer56/Dota2API/utils"
)

func main() {
	var config models.Configuration
	var openDotaService *opendota.OpenDotaService
	loadConfig(&config)
	loadServices(openDotaService, config)
	dotaHandler, matchDataHandler, favouritesHandler, err := loadHandlers(openDotaService, config)
	if err != nil {
		utils.LogInfo("Cant initialize handlers", "MAIN")
		panic(1)
	}
	go graphapi.StartGraphServer(config, dotaHandler, favouritesHandler, matchDataHandler)
	go grpcapi.StartGrpcServer(config, dotaHandler, matchDataHandler)
}

//Load Config Data for the APIs
func loadConfig(config *models.Configuration) {
	file, err := os.Open("../config.json")
	if err != nil {
		fmt.Println("Cant open config")
		panic(1)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Cant read config")
		panic(1)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Cant parse config")
		panic(1)
	}
	utils.LogInfo("Initialised the Config", "loadConfig")
}

//Load Underlying Services such as OpenDota
func loadServices(odService *opendota.OpenDotaService, config models.Configuration) {
	odService = opendota.NewOpenDotaService(config.IsProduction)
	utils.LogInfo("Initialised the OpenDotaService", "loadServices")
}

//Load Handlers Logic that handle incoming requests
func loadHandlers(odService *opendota.OpenDotaService, config models.Configuration) (*handlers.Dota2Handler, *handlers.MatchDataHandler,
	*handlers.FavouritesHandler, error) {

	dataHandler := &handlers.Dota2Handler{
		Dota2service: odService,
		GameData:     &common.Dota2GameInfo{},
	}

	err := dataHandler.PopulateStaticData()
	if err != nil {
		utils.LogError("Couldn't initialize game data from service", "loadHandlers")
		panic(1)
	}
	matchHandler := &handlers.MatchDataHandler{
		Dota2service: odService,
	}

	sqlConnection := &database.SqlConnection{
		ConnectionString: config.ConnectionString,
		DatabaseName:     config.DatabaseName,
		DriverName:       "mysql",
	}
	favHandler := &handlers.FavouritesHandler{
		MysqlConn:       sqlConnection,
		FavouritesTable: "Favourites",
	}
	return dataHandler, matchHandler, favHandler, nil

}
