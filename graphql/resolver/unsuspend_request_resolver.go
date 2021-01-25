package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/unsuspend_request"
)

func GetAllSuspendedRequests (p graphql.ResolveParams) (i interface{}, e error){
	requests := unsuspend_request.GetAll()
	return requests, nil
}

func CreateUnsuspendRequest(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	reason := p.Args["reason"].(string)
	req := unsuspend_request.Create(userId, reason)
	return req, nil
}

func DenyUnsuspendRequest(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	isSuccess := unsuspend_request.Delete(userId)
	return isSuccess, nil
}

