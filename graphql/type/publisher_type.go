package _type

import "github.com/graphql-go/graphql"

var publisherType *graphql.Object

func GetPublisherType() *graphql.Object{
	if(publisherType == nil) {
		publisherType = graphql.NewObject(graphql.ObjectConfig{
			Name: "publisher",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return publisherType
}
