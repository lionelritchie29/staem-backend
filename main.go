package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/lionelritchie29/staem-backend/api"
	"github.com/lionelritchie29/staem-backend/graphql/mutation"
	"github.com/lionelritchie29/staem-backend/graphql/query"
	"github.com/lionelritchie29/staem-backend/helpers"
	"github.com/lionelritchie29/staem-backend/middleware"
	"log"
	"net/http"
	"os"
)

func main(){
	helpers.SetEnv()
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: query.GetRoot(),
		Mutation: mutation.GetRoot(),
	})

	if err != nil{
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
		Playground: true,
	})

	wrapped:= middleware.CorsMiddleware(h)

	router := api.NewRouter()

	// Choose the folder to serve
	staticDir := "/images/"
	path := os.Getenv("IMAGE_PATH")

	// Create the route
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir(path))))

	router.Handle("/api", wrapped)

	log.Fatalln(http.ListenAndServe(":2000", router))
}

