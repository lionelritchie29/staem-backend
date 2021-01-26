package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/user_wishlist"
)

func CreateWishlist(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	gameId := p.Args["gameId"].(int)
	wishlist := user_wishlist.Create(userId, gameId)
	return wishlist, nil
}

func DeleteWishlist(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	gameId := p.Args["gameId"].(int)
	isSuccess := user_wishlist.Delete(userId, gameId)
	return isSuccess, nil
}

func GetWishlistByUserId(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	wishlist := user_wishlist.GetById(userId)
	return wishlist, nil
}

func GetWishlists(p graphql.ResolveParams) (i interface{}, e error){
	wishlist := user_wishlist.GetAll()
	return wishlist, nil
}
