package graphapi

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

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

	http.Handle("/", LoggingMiddleWare(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", srv)

	utils.LogInfo(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port), "Graph Server")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
		wg.Done()
	}
}

func LoggingMiddleWare(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		utils.LogInfo(fmt.Sprintf("{r.Method: %s, r.URL.Path: %s", r.Method, r.URL.Path), "HTTP")
		handler(w, r)
		utils.LogInfo(fmt.Sprintf("Time Taken: %d", time.Since(t).Milliseconds()), "HTTP")
	}
}
