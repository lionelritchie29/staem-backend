package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/user_report"
)

func GetReportsByUserId (p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	reports := user_report.GetByUserId(userId)
	return reports, nil
}

func CreateReport(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	reporterId := p.Args["reporterId"].(int)
	reason := p.Args["reason"].(string)
	status := user_report.Create(userId, reporterId, reason)
	return status, nil
}
