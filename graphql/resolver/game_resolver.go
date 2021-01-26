package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/input_models"
	"github.com/lionelritchie29/staem-backend/models/game"
	"github.com/mitchellh/mapstructure"
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

func GetGamesByTitleLimit5(p graphql.ResolveParams) (i interface{}, e error) {
	searchQuery := p.Args["query"].(string)
	games := game.GetByTitleLimit5(searchQuery)
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

func CreateGame(p graphql.ResolveParams) (i interface{}, e error) {
	newGameRaw := p.Args["newGame"]
	var newGame input_models.NewGame
	mapstructure.Decode(newGameRaw, &newGame)
	gameCreated := game.Create(newGame)
	return gameCreated, nil
}

func DeleteGame(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	if game.Delete(id) {
		return true, nil
	}

	return false, nil
}