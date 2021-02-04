package main

import (
	"fmt"
	"github.com/functionalfoundry/graphqlws"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/lionelritchie29/staem-backend/api"
	"github.com/lionelritchie29/staem-backend/globals"
	"github.com/lionelritchie29/staem-backend/graphql/mutation"
	"github.com/lionelritchie29/staem-backend/graphql/query"
	"github.com/lionelritchie29/staem-backend/graphql/subscription"
	"github.com/lionelritchie29/staem-backend/middleware"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var errSchema error

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println(errEnv)
	}
	globals.Schema, errSchema = graphql.NewSchema(graphql.SchemaConfig{
		Query:        query.GetRoot(),
		Mutation:     mutation.GetRoot(),
		Subscription: subscription.GetRoot(),
	})

	if errSchema != nil {
		panic(errSchema)
	}

	/// WS ///
	globals.SubscriptionManager = graphqlws.NewSubscriptionManager(&globals.Schema)
	fmt.Println(globals.SubscriptionManager)
	websocketHandler := graphqlws.NewHandler(graphqlws.HandlerConfig{
		SubscriptionManager: globals.SubscriptionManager,
		Authenticate: func(token string) (interface{}, error) {
			return "Default user", nil
		},
	})
	///    ///


	h := handler.New(&handler.Config{
		Schema:           &globals.Schema,
		Pretty:           true,
		GraphiQL:         false,
		Playground:       true,
	})

	wrapped := middleware.CorsMiddleware(h)

	router := api.NewRouter()

	log.Println(os.Getenv("PORT"))

	// Choose the folder to serve
	staticDir := "images"
	path, _ := filepath.Abs(staticDir)
	log.Println(path)

	// Create the route
	router.
		PathPrefix("/" + staticDir + "/").
		Handler(http.StripPrefix("/" + staticDir + "/", http.FileServer(http.Dir(path))))

	router.Handle("/api", wrapped)
	router.Handle("/subscriptions", websocketHandler)

	log.Fatalln(http.ListenAndServe(os.Getenv("PORT"), router))
}
