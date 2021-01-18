package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/game"
)

func GetGames(p graphql.ResolveParams) (i interface{}, e error) {
	games := game.GetAll()
	return games, nil
}

func GetGamesByTitle(p graphql.ResolveParams) (i interface{}, e error) {
	searchQuery := p.Args["query"].(string)
	games := game.GetByTitle(searchQuery)
	return games, nil
}

func GetGame(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	game := game.Get(id)
	return game, nil
}

func GetSpecialOfferGames(p graphql.ResolveParams) (i interface{}, e error) {
	gamesOnSale := game.GetGameOnSale()
	return gamesOnSale, nil
}

func GetFeaturedRecommendedGames(p graphql.ResolveParams) (i interface{}, e error) {
	featuredGames := game.GetFeaturedAndRecommendedGame()
	return featuredGames, nil
}

func GetCommunityRecommendedGames(p graphql.ResolveParams) (i interface{}, e error) {
	commRecommendedGames := game.GetCommunityRecommends()
	return commRecommendedGames, nil
}

func GetTopSellerGames(p graphql.ResolveParams) (i interface{}, e error) {
	topSellerGames := game.GetTopSellers()
	return topSellerGames, nil
}

func GetRecentlyPublishedGames(p graphql.ResolveParams) (i interface{}, e error) {
	recentlyPublishedGames := game.GetRecentlyPublished()
	return recentlyPublishedGames, nil
}

func GetSpecialCategoryGames(p graphql.ResolveParams) (i interface{}, e error) {
	specialGames := game.GetSpecialCategory()
	return specialGames, nil
}