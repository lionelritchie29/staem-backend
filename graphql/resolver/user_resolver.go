package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/input_models"
	"github.com/lionelritchie29/staem-backend/models/user"
	"github.com/lionelritchie29/staem-backend/models/wallet"
	"github.com/mitchellh/mapstructure"
)

func GetUsers(p graphql.ResolveParams) (i interface{}, e error){
	users := user.GetAll()
	return users, nil
}

func GetUser(p graphql.ResolveParams) (i interface{}, e error){
	id := p.Args["id"].(int)
	user := user.Get(id)
	return user, nil
}

func GetFriends(p graphql.ResolveParams) (i interface{}, e error){
	id := p.Args["id"].(int)
	friends := user.GetFriends(id)
	return friends, nil
}

func GetUserByAccountName(p graphql.ResolveParams) (i interface{}, e error){
	customUrl := p.Args["url"].(string)
	user := user.GetByCustomUrl(customUrl)
	return user, nil
}

func CreateUser(p graphql.ResolveParams) (i interface{}, e error){
	newUserRaw := p.Args["newUser"]
	var newUser input_models.NewUserAccount
	mapstructure.Decode(newUserRaw, &newUser)
	success := user.Create(newUser)
	return success, nil
}

func RedeemWallet(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	code := p.Args["code"].(string)

	walletId := wallet.Redeem(userId, code)

	if walletId == -1{
		return nil, nil
	}

	wallet := wallet.Get(walletId)

	return wallet, nil
}