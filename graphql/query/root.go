package query


import (
	"github.com/graphql-go/graphql"
	//res "github.com/lionelritchie29/staem-backend/graphql/query/resolver"
	//typ "github.com/lionelritchie29/staem-backend/graphql/type"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			// To Be Filled
		},
	})
}