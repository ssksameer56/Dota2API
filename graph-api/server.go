package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	server "github.com/ssksameer56/Dota2API/graph-api/api"
	"github.com/ssksameer56/Dota2API/graph-api/resolvers"
	"github.com/ssksameer56/Dota2API/models/common"
	"github.com/ssksameer56/Dota2API/utils"
)

const defaultPort = "8080"

var config common.ConfigData

func Init() {
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
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = config["graphAPIPort"]
	}
	srv := handler.NewDefaultServer(server.NewExecutableSchema(server.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	utils.LogInfo(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port), "Graph Server")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
