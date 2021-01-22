package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/friend_request"
)

func CreateFriendRequest(p graphql.ResolveParams) (i interface{}, e error) {
	fromId := p.Args["fromId"].(int)
	toId := p.Args["toId"].(int)

	return friend_request.Create(fromId, toId), nil
}

func AcceptFriendRequest(p graphql.ResolveParams) (i interface{}, e error) {
	fromId := p.Args["fromId"].(int)
	toId := p.Args["toId"].(int)

	return friend_request.Accept(fromId, toId), nil
}

func RejectFriendRequest(p graphql.ResolveParams) (i interface{}, e error) {
	fromId := p.Args["fromId"].(int)
	toId := p.Args["toId"].(int)

	return friend_request.Reject(fromId, toId), nil
}

func IgnoreFriendRequest(p graphql.ResolveParams) (i interface{}, e error) {
	fromId := p.Args["fromId"].(int)
	toId := p.Args["toId"].(int)

	return friend_request.Ignore(fromId, toId), nil
}
