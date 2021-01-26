package _type

import "github.com/graphql-go/graphql"

var loginReturnType *graphql.Object

func GetLoginReturnType() *graphql.Object {
	if loginReturnType == nil {
		loginReturnType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "loginReturn",
			Fields: graphql.Fields{
				"token": &graphql.Field{
					Type: graphql.String,
				},
				"role": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return loginReturnType
}