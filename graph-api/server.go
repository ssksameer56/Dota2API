package graphapi

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	server "github.com/ssksameer56/Dota2API/graph-api/api"
	"github.com/ssksameer56/Dota2API/graph-api/resolvers"
	"github.com/ssksameer56/Dota2API/handlers"
	"github.com/ssksameer56/Dota2API/models"
	"github.com/ssksameer56/Dota2API/utils"
)

func StartGraphServer(config models.Configuration, dota2Handler *handlers.Dota2Handler, favHandler *handlers.FavouritesHandler,
	matchHandler *handlers.MatchDataHandler, wg sync.WaitGroup) {
	port := "8080"
	if config.GraphAPIPort != "" {
		port = config.GraphAPIPort
	}
	srv := handler.NewDefaultServer(server.NewExecutableSchema(server.Config{Resolvers: &resolvers.Resolver{
		FavouritesService:   favHandler,
		ConstantDataService: dota2Handler,
		MatchDataService:    matchHandler,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	utils.LogInfo(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port), "Graph Server")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
		wg.Done()
	}
}
