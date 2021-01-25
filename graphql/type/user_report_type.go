package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/user"
)

var userReportType *graphql.Object

func GetUserReportType() *graphql.Object {
	if userReportType == nil {
		userReportType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "userReport",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.UserReport)
						reportedUser := user.Get(int(parent.UserID))
						return reportedUser, nil
					},
				},
				"reporter": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.UserReport)
						reporterUser := user.Get(int(parent.ReporterID))
						return reporterUser, nil
					},
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