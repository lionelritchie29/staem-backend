package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/playtime"
)

func GetPlaytimes(p graphql.ResolveParams) (i interface{}, e error) {
	playtimes := playtime.GetAll()
	return playtimes, nil
}

func GetPlaytimesByGameId(p graphql.ResolveParams) (i interface{}, e error) {
	gameId := p.Args["id"].(int)
	playtimes := playtime.GetAllByGameId(gameId)
	return playtimes, nil
}