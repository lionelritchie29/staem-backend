package user

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/helpers"
	"github.com/lionelritchie29/staem-backend/input_models"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
	"time"
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

func GetByCode(code string) models.UserAccount {
	db := database.GetInstance()
	var user models.UserAccount
	db.Where("code = ?", code).Find(&user)

	return user
}

func GetByAccountName(accountName string, isLogin bool) models.UserAccount {
	db := database.GetInstance()
	var user models.UserAccount
	db.Where("account_name = ?", accountName).First(&user)

	if isLogin {
		user.Status = "online"
		db.Save(&user)
	}

	return user
}

func Logout(userId int) bool{
	db := database.GetInstance()
	var user models.UserAccount
	db.Find(&user, userId)
	user.Status = "offline"
	db.Save(&user)

	if db.Error != nil {
		return false
	}

	return true
}

func GetByCustomUrl(customUrl string) models.UserAccount {
	db := database.GetInstance()
	var profile models.UserProfile
	var user models.UserAccount

	db.Where("custom_url = ?", customUrl).First(&profile)
	db.First(&user, profile.ID)

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

func GetFriends(userId int) []models.UserAccount {
	db := database.GetInstance()
	var friends []models.UserAccount
	db.Raw("SELECT * FROM user_accounts WHERE id IN (" +
				"SELECT friend_id FROM friends WHERE user_id = " + strconv.FormatInt(int64(userId), 10) +
			") ").Scan(&friends)
	return friends
}

func GetGames(userId int) []models.Game {
	db := database.GetInstance()
	var games []models.Game

	db.Raw("SELECT * FROM games WHERE id IN ( " +
				"SELECT game_id FROM user_games WHERE user_id = " + strconv.FormatInt(int64(userId), 10) +
		")").Scan(&games)

	return games
}

func GetComments(userId int) []models.UserComment {
	db := database.GetInstance()
	var comments []models.UserComment
	db.Where("dest_user_id = ?", userId).Find(&comments)
	return comments
}

func Create(newUser input_models.NewUserAccount) bool {
	db := database.GetInstance()

	hashedPassword, _ := helpers.HashPassword(newUser.Password)

	userAccount := &models.UserAccount{
		AccountName:  newUser.AccountName,
		Email:        newUser.Email,
		Password:     hashedPassword,
		WalletAmount: 0,
		Role:         1, //user default
		SuspendedAt:  time.Time{},
		Status:       "offline",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	}

	resAccount := db.Create(userAccount)

	resProfile := db.Create(&models.UserProfile{
		ID:                       userAccount.ID,
		Level:                    1,
		Point:                    1,
		DisplayName:              userAccount.AccountName,
		CustomURL:                userAccount.AccountName,
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "default.jpg",
		ProfileBackgroundUrl:     "default.jpg",
		MiniProfileBackgroundUrl: "default.jpg",
		Theme:                    "default.jpg",
		FeaturedBadgeUrl:         "default.jpg",
		Country:                  newUser.Country,
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	if resAccount.Error != nil || resProfile.Error != nil  {
		return false
	} else {
		return true
	}
}

func Suspend(id int) bool {
	db := database.GetInstance()
	var user models.UserAccount
	db.Find(&user, id)
	user.SuspendedAt = time.Now()
	res := db.Save(&user)

	if res.Error != nil {
		return false
	}

	return true
}

func Unsuspend(userId int) bool {
	db := database.GetInstance()
	db.Where("user_id = ?", userId).Unscoped().Delete(&models.UnsuspendRequest{})
	db.Where("user_id = ?", userId).Unscoped().Delete(&models.UserReport{})

	var user models.UserAccount
	db.Find(&user, userId)
	user.SuspendedAt = time.Time{}
	db.Save(&user)

	return true
}