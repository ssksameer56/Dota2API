package graphapi

import (
	"fmt"
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

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	utils.LogInfo(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port), "Graph Server")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		utils.LogFatal(err.Error(), "GraphServer")
		wg.Done()
	}
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			utils.LogInfo(fmt.Sprintf("Request start: %s\n", r.URL.Path), "HTTP")
			t := time.Now()
			next.ServeHTTP(w, r)
			utils.LogInfo(fmt.Sprintf("Response Done in : %d\n", time.Since(t).Milliseconds()), "HTTP")
		},
	)
}
