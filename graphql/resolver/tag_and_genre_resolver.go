package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/game_genre"
	"github.com/lionelritchie29/staem-backend/models/game_tag"
)

func GetTags(p graphql.ResolveParams) (i interface{}, e error){
	tags := game_tag.GetAll()
	return tags, nil
}

func GetGenre(p graphql.ResolveParams) (i interface{}, e error){
	genres := game_genre.GetAll()
	return genres, nil
}
