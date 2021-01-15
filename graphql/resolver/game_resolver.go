package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/game"
)

func GetGames(p graphql.ResolveParams) (i interface{}, e error) {
	games := game.GetAll()
	return games, nil
}

func GetGame(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	game := game.Get(id)
	return game, nil
}