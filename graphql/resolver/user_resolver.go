package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/user"
)

func GetUsers(p graphql.ResolveParams) (i interface{}, e error){
	users := user.GetAll()
	return users, nil
}