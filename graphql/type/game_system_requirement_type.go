package _type

import "github.com/graphql-go/graphql"

var systemReq *graphql.Object

func GetGameSystemReqType() *graphql.Object {
	if systemReq == nil {
		systemReq = graphql.NewObject(graphql.ObjectConfig{
			Name: "system_requirement",
			Fields: graphql.Fields{
				"gameId": &graphql.Field{
					Type: graphql.Int,
				},
				"isRecommended": &graphql.Field{
					Type: graphql.Boolean,
				},
				"note": &graphql.Field{
					Type: graphql.String,
				},
				"operatingSystem": &graphql.Field{
					Type: graphql.String,
				},
				"processor": &graphql.Field{
					Type: graphql.String,
				},
				"memory": &graphql.Field{
					Type: graphql.String,
				},
				"graphics": &graphql.Field{
					Type: graphql.String,
				},
				"directX": &graphql.Field{
					Type: graphql.String,
				},
				"storage": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return systemReq
}
