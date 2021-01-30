package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/game_item"
)

func GetGameItemById(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	gameItem := game_item.Get(id)
	return gameItem, nil
}
