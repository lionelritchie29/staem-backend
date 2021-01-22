package resolver

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/helpers"
	"github.com/lionelritchie29/staem-backend/models/user"
)

func LoginResolver(p graphql.ResolveParams) (i interface{}, e error){
	accountName := p.Args["accountName"].(string)
	password := p.Args["password"].(string)

	findUser := user.GetByAccountName(accountName, true)

	fmt.Println(findUser)

	if helpers.CheckPassword(findUser.Password, password) {
		token, _ := helpers.GenerateToken(int(findUser.ID))
		return token, nil
	}

	return nil, nil
}

func LogoutResolver(p graphql.ResolveParams) (i interface{}, e error){
	id := p.Args["id"].(int)
	if user.Logout(id) {
		return true, nil
	}

	return false, nil
}