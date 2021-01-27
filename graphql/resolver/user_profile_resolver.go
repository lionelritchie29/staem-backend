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

func UpdateProfileAvatar(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	imageString := p.Args["imageString"].(string)
	isSuccess := user_profile.UpdateAvatar(userId, imageString)
	return isSuccess, nil
}

func UpdateProfilAvatarFrame(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	avatarFrameUrl := p.Args["avatarFrameUrl"].(string)
	isSuccess := user_profile.UpdateAvatarFrame(userId, avatarFrameUrl)
	return isSuccess, nil
}

func UpdateProfileBackground(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	profileBackgroundUrl := p.Args["profileBackgroundUrl"].(string)
	isSuccess := user_profile.UpdateBackground(userId, profileBackgroundUrl)
	return isSuccess, nil
}

func UpdateMiniProfileBackground(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	miniProfileBackgroundUrl := p.Args["miniProfileBackgroundUrl"].(string)
	isSuccess := user_profile.UpdateMiniBackground(userId, miniProfileBackgroundUrl)
	return isSuccess, nil
}