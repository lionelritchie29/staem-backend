package user

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
)

func GetAll() []models.UserAccount{
	db := database.GetInstance()
	var users []models.UserAccount
	db.Find(&users)
	return users
}

func Get(userId int) models.UserAccount {
	db := database.GetInstance()
	var user models.UserAccount
	db.Find(&user, userId)
	return user
}

func GetRole(roleId int) models.UserRole {
	db := database.GetInstance()
	var role models.UserRole
	db.Find(&role, roleId)
	return role
}

func GetProfile(userId int) models.UserProfile {
	db := database.GetInstance()
	var profile models.UserProfile
	db.Find(&profile, userId)
	return profile
}
