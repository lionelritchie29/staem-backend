package main

import (
	"github.com/lionelritchie29/staem-backend/api"
	"github.com/lionelritchie29/staem-backend/graphql/query"
	"github.com/lionelritchie29/staem-backend/middleware"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main(){
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: query.GetRoot(),

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

	router:= api.NewRouter()
	router.Handle("/api", wrapped)
	log.Fatalln(http.ListenAndServe(":2000", router))
}
