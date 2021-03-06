package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

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

	runGraph := flag.Bool("graph", false, "Run GraphQL Server. False by Default")
	runGrpc := flag.Bool("grpc", false, "Run gRPC Server. False by default")
	flag.Parse()

	var config models.Configuration
	var openDotaService *opendota.OpenDotaService
	loadConfig(&config)

	utils.LogFilePath = config.LogFile
	utils.InitializeLogging()

	openDotaService = loadServices(config)
	dotaHandler, matchDataHandler, favouritesHandler, err := loadHandlers(openDotaService, config)
	if err != nil {
		utils.LogInfo("Cant initialize handlers", "MAIN")
		panic(1)
	}
	wg := sync.WaitGroup{}

	if *runGraph {
		wg.Add(1)
		utils.LogInfo("Graph Server to be run", "MAIN")
		go graphapi.StartGraphServer(config, dotaHandler, favouritesHandler, matchDataHandler, wg)
	}
	if *runGrpc {
		wg.Add(1)
		utils.LogInfo("GRPC Server to be run", "MAIN")
		go grpcapi.StartGrpcServer(config, dotaHandler, matchDataHandler, wg)
	}
	wg.Wait()
}

//Load Config Data for the APIs
func loadConfig(config *models.Configuration) {
	file, err := os.Open("../config.json")
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

//Load Underlying Services such as OpenDota
func loadServices(config models.Configuration) *opendota.OpenDotaService {
	isProduction, _ := strconv.ParseBool(config.IsProduction)
	odService := opendota.NewOpenDotaService(isProduction)
	utils.LogInfo("Initialised the OpenDotaService", "loadServices")
	return odService
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

	hostName := os.Getenv("DATABASE_HOST")
	abc := os.Environ()
	fmt.Println(abc)
	if hostName != "" {
		config.ConnectionString = strings.ReplaceAll(config.ConnectionString, "localhost", hostName)
		utils.LogInfo(fmt.Sprintf("Changed Connection string to %s", config.ConnectionString), "loadHandlers")
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
