package main

import (
	"fmt"
	"github.com/functionalfoundry/graphqlws"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/lionelritchie29/staem-backend/api"
	"github.com/lionelritchie29/staem-backend/globals"
	"github.com/lionelritchie29/staem-backend/graphql/mutation"
	"github.com/lionelritchie29/staem-backend/graphql/query"
	"github.com/lionelritchie29/staem-backend/graphql/subscription"
	"github.com/lionelritchie29/staem-backend/helpers"
	"github.com/lionelritchie29/staem-backend/middleware"
	"log"
	"net/http"
	"os"
)

var errSchema error

func main() {
	helpers.SetEnv()
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

	// Choose the folder to serve
	staticDir := "/images/"
	path := os.Getenv("IMAGE_PATH")

	// Create the route
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir(path))))

	router.Handle("/api", wrapped)
	router.Handle("/subscriptions", websocketHandler)

	log.Fatalln(http.ListenAndServe(":8080", router))
}
