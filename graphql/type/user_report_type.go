package _type

import "github.com/graphql-go/graphql"

var userReportType *graphql.Object

func GetUserReportType() *graphql.Object {
	if userReportType == nil {
		userReportType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "userReport",
			Fields: graphql.Fields{
				"userId": &graphql.Field{
					Type: graphql.Int,
				},
				"reporterId": &graphql.Field{
					Type: graphql.Int,
				},
				"reason": &graphql.Field{
					Type: graphql.String,
				},
				"createdAt": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		})
	}

	return userReportType
}