package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/game_review"
)

func GetGameReviewByGameId(p graphql.ResolveParams) (i interface{}, e error){
	id := p.Args["gameId"].(int)
	reviews := game_review.GetByGameId(id)
	return reviews, nil
}

func GetPastMonthGameReviewById(p graphql.ResolveParams) (i interface{}, e error){
	id := p.Args["gameId"].(int)
	reviews := game_review.GetHelpfulReviewByGameId(id)
	return reviews, nil
}

func GetRecentlyPostedReviewById(p graphql.ResolveParams) (i interface{}, e error){
	id := p.Args["gameId"].(int)
	reviews := game_review.GetRecentlyPostedbyGameId(id)
	return reviews, nil
}

