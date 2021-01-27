package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/input_models"
	"github.com/lionelritchie29/staem-backend/models/user_profile"
	"github.com/mitchellh/mapstructure"
)

func UpdateProfileInfo(p graphql.ResolveParams) (i interface{}, e error){
	newProfileRaw := p.Args["newProfile"]
	var newProfile input_models.NewProfileInfo
	mapstructure.Decode(newProfileRaw, &newProfile)
	isSuccess := user_profile.UpdateInfo(newProfile)
	return isSuccess, nil
}
