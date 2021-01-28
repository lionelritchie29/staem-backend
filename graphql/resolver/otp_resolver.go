package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
)

func CreateOTP(p graphql.ResolveParams) (i interface{}, e error) {
	email := p.Args["email"].(string)
	accountName := p.Args["accountName"].(string)
	isSuccess := models.CreateOTP(email, accountName)
	return isSuccess, nil
}

func VerifyOTP(p graphql.ResolveParams) (i interface{}, e error) {
	code := p.Args["code"].(string)
	isSuccess := models.VerifiyOTP(code)
	return isSuccess, nil
}