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

func GetByAccountName(accountName string) models.UserAccount {
	db := database.GetInstance()
	var user models.UserAccount
	db.Where("account_name = ?", accountName).First(&user)
	return user
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
		AvatarFrameUrl:           "default.jpg",
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
