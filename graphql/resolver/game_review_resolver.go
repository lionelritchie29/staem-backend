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

func CreateGameReview(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	gameId := p.Args["gameId"].(int)
	content := p.Args["content"].(string)
	isRecommended := p.Args["isRecommended"].(bool)

	isSuccess := game_review.Create(userId, gameId, content, isRecommended)
	return isSuccess, nil
}

func IncreaseReviewUpvoteCount(p graphql.ResolveParams) (i interface{}, e error){
	reviewId := p.Args["reviewId"].(int)
	isSuccess := game_review.IncreaseUpvoteCount(reviewId)
	return isSuccess, nil
}

func IncreaseReviewDownvoteCount(p graphql.ResolveParams) (i interface{}, e error){
	reviewId := p.Args["reviewId"].(int)
	isSuccess := game_review.IncreaseDownvoteCount(reviewId)
	return isSuccess, nil
}

