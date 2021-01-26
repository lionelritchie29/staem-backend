package game

import (
	"fmt"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/helpers"
	"github.com/lionelritchie29/staem-backend/input_models"
	"github.com/lionelritchie29/staem-backend/models"
	"os"
	"strconv"
	"syreclabs.com/go/faker"
	"time"
)

func GetAll() []models.Game {
	db := database.GetInstance()
	var games []models.Game

	db.Find(&games)

	return games
}

func GetByTitle(query string) []models.Game {
	db := database.GetInstance()
	var games []models.Game
	db.Raw("SELECT * FROM games WHERE LOWER(title) LIKE LOWER('%" + query + "%')").Scan(&games)
	return games
}

func GetByTitleLimit5(query string) []models.Game {
	db := database.GetInstance()
	var games []models.Game
	db.Raw("SELECT * FROM games WHERE LOWER(title) LIKE LOWER('%" + query + "%') LIMIT 5").Scan(&games)
	return games
}

func Get(id int) models.Game {
	db := database.GetInstance()
	var game models.Game

	db.Find(&game, id)

	return game
}

func GetGenres(gameId int) []models.GameGenre {
	db := database.GetInstance()
	var genreNumbers []models.GameDetailGenre
	db.Find(&genreNumbers, "game_id = ?", gameId)

	var genres []models.GameGenre

	for _, genreNumber := range genreNumbers {
		var genre models.GameGenre
		db.Find(&genre, genreNumber.GameGenreID)
		genres = append(genres, genre)
	}

	return genres
}

func GetTags(gameId int) []models.GameTag {
	db := database.GetInstance()
	var gameTagDetails []models.GameDetailTag
	db.Find(&gameTagDetails, "game_id = ?", gameId)

	var tags []models.GameTag

	for _, gameTagDetail := range gameTagDetails {
		var tag models.GameTag
		db.Find(&tag, gameTagDetail.GameTagID)
		tags = append(tags, tag)
	}

	return tags
}

func GetCategories(gameId int) []models.GameCategory {
	db := database.GetInstance()
	var gameCategoryDetails []models.GameDetailCategory
	db.Find(&gameCategoryDetails, "game_id = ?", gameId)

	var categories []models.GameCategory

	for _, gameCategoryDetail := range gameCategoryDetails {
		var category models.GameCategory
		db.Find(&category, gameCategoryDetail.GameCategoryID)
		categories = append(categories, category)
	}

	return categories
}

func GetImages(gameId int) []models.GameImage {
	db := database.GetInstance()
	var gameImages []models.GameImage
	db.Find(&gameImages, "game_id = ?", gameId)
	return gameImages
}

func GetSystemRequirements(gameId int) []models.GameSystemRequirement {
	db := database.GetInstance()
	var req []models.GameSystemRequirement
	db.Raw("SELECT * FROM game_system_requirements WHERE game_id = " + strconv.FormatInt(int64(gameId), 10)).Scan(&req)
	return req
}

func GetFeaturedAndRecommendedGame() []models.Game {
	db := database.GetInstance()
	var featuredGames []models.Game

	db.Raw("SELECT * FROM games WHERE id IN (" +
		"SELECT game_id FROM game_playtimes " +
		"WHERE TO_DATE(date, 'YYYY-MM-DD') >= CURRENT_DATE - INTERVAL '7 day' " +
		"GROUP BY game_id LIMIT 5" +
		")").Scan(&featuredGames)

	return featuredGames
}

func GetCommunityRecommends() []models.Game {
	db := database.GetInstance()
	var communityRecommendedGames []models.Game
	db.Raw("SELECT * FROM games WHERE id IN (" +
		"SELECT game_id FROM game_reviews " +
		"GROUP BY game_id ORDER BY SUM(upvote_count) " +
		"DESC LIMIT 5" +
		")").Scan(&communityRecommendedGames)
	return communityRecommendedGames
}

func GetGameOnSale() []models.Game {
	db := database.GetInstance()
	var gamesOnSale []models.Game

	db.Model(&models.Game{}).Joins("LEFT JOIN game_sales ON game_sales.game_id = games.id").Where("game_sales.game_id = games.id").Scan(&gamesOnSale)

	return gamesOnSale
}

func GetTopSellers() []models.Game {
	db := database.GetInstance()
	var topSellerGames []models.Game
	db.Raw("SELECT * FROM games WHERE id IN (" +
		"SELECT game FROM game_transaction_details d " +
		"JOIN game_transaction_headers h ON d.game_transaction_id = h.id " +
		"WHERE h.transaction_date::date >= CURRENT_DATE - INTERVAL '7 day' " +
		"GROUP BY game ORDER BY SUM(quantity) DESC" +
		")").Scan(&topSellerGames)
	return topSellerGames
}

func GetRecentlyPublished() []models.Game {
	db := database.GetInstance()
	var games []models.Game
	//db.Find(&games).Where("TO_DATE(release_date, 'YYYY-MM-DD') >= CURRENT_DATE - INTERVAL '365 day'")
	db.Raw("SELECT * FROM games WHERE TO_DATE(release_date, 'YYYY-MM-DD') >= CURRENT_DATE - INTERVAL '365 day' ").
		Scan(&games)
	return games
}

func GetSpecialCategory() []models.Game {
	db := database.GetInstance()
	var gamesOnSale []models.Game

	db.Model(&models.Game{}).Joins("LEFT JOIN game_sales ON game_sales.game_id = games.id").Where("game_sales.game_id = games.id AND game_sales.discount >= 50").Scan(&gamesOnSale)

	return gamesOnSale
}

func Create(newGame input_models.NewGame) models.Game {
	db := database.GetInstance()

	game := models.Game{
		Publisher:   newGame.PublisherId,
		Developer:   newGame.DeveloperId,
		Title:       newGame.Title,
		Description: newGame.Description,
		Price:       newGame.Price,
		ReleaseDate: newGame.ReleaseDate,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	}

	db.Create(&game)

	for i := 0; i < len(newGame.TagIds); i++ {
		db.Create(&models.GameDetailTag{
			GameID:    game.ID,
			GameTagID: uint(newGame.TagIds[i]),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		})
	}

	for i := 0; i < len(newGame.GenreIds); i++ {
		db.Create(&models.GameDetailGenre{
			GameID:      game.ID,
			GameGenreID: uint(newGame.GenreIds[i]),
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   nil,
		})
	}

	db.Create(&models.GameImage{
		GameID:    game.ID,
		Url:       "header.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	helpers.MakeDirectoryIfNotExists(os.Getenv("IMAGE_PATH") + "/games/" + strconv.FormatInt(int64(game.ID), 10))
	helpers.SaveImage(newGame.GameHeaderImage, "games/"+strconv.FormatInt(int64(game.ID), 10)+"/header.jpg")

	for i := 0; i < len(newGame.GameImages); i++ {
		db.Create(&models.GameImage{
			GameID:    game.ID,
			Url:       strconv.FormatInt(int64(i+1), 10) + ".jpg",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		})

		filename := "games/" + strconv.FormatInt(int64(game.ID), 10) + "/" + strconv.FormatInt(int64(i+1), 10) + ".jpg"
		helpers.SaveImage(newGame.GameImages[i], filename)
	}

	db.Create(&models.GameSystemRequirement{
		GameID:          game.ID,
		IsRecommended:   true,
		Note:            faker.Lorem().Sentence(4),
		OperatingSystem: "Windows 10",
		Processor:       "FX-6350 or Equivalent; Core i5-3570 or Equivalent",
		Memory:          "8 GB RAM",
		Graphics:        "AMD: Radeon 7970/Radeon R9 280x or Equivalent; NVIDIA: GeForce GTX 760 or Equivalent",
		DirectX:         "Version 11",
		Storage:         "50 GB available space",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&models.GameSystemRequirement{
		GameID:          game.ID,
		IsRecommended:   false,
		Note:            faker.Lorem().Sentence(4),
		OperatingSystem: "Windows 10",
		Processor:       "Ryzen 3 1300X or Equivalent; Core i7-4790 or Equivalent",
		Memory:          "16 GB RAM",
		Graphics:        "AMD: Radeon RX 480 or Equivalent; NVIDIA: GeForce GTX 1060 or Equivalent",
		DirectX:         "Version 11",
		Storage:         "50 GB available space",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	return game
}

func Delete(id int) bool {
	db := database.GetInstance()
	var game models.Game
	res := db.Delete(game, id)

	if res.Error != nil {
		fmt.Println(res.Error)
		return false
	}

	return true
}

//func Update(newGame input_models.NewGame, id int) {
//	db := database.GetInstance()
//
//	db.Model(&models.Game{}).Where("id = ?", id).Update(models.Game{
//		Publisher:   newGame.PublisherId,
//		Developer:   newGame.DeveloperId,
//		Title:       newGame.Title,
//		Description: newGame.Description,
//		Price:       newGame.Price,
//		ReleaseDate: newGame.ReleaseDate,
//		CreatedAt:   time.Time{},
//		UpdatedAt:   time.Time{},
//		DeletedAt:   nil,
//	})
//
//	db.Create(&models.GameImage{
//		GameID:    game.ID,
//		Url:       "header.jpg",
//		CreatedAt: time.Time{},
//		UpdatedAt: time.Time{},
//		DeletedAt: nil,
//	})
//
//	helpers.SaveImage(newGame.GameHeaderImage, "games/" + strconv.FormatInt(int64(game.ID), 10) + "/header.jpg")
//
//	for i:=0; i<len(newGame.GameImages); i++ {
//		db.Create(&models.GameImage{
//			GameID:    game.ID,
//			Url:       strconv.FormatInt(int64(i+1), 10) + ".jpg",
//			CreatedAt: time.Time{},
//			UpdatedAt: time.Time{},
//			DeletedAt: nil,
//		})
//
//		filename := "games/" + strconv.FormatInt(int64(game.ID), 10) + "/" + strconv.FormatInt(int64(i+1), 10) + ".jpg"
//		helpers.SaveImage(newGame.GameImages[i], filename)
//	}
//
//	return game
//}
