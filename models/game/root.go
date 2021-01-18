package game

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
)

func GetAll() []models.Game{
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

	db.Raw("SELECT " +
		"ga.id, ga.publisher, ga.developer, ga.title," +
		"ga.description, ga.price, ga.release_date," +
		"ga.created_at, ga.updated_at, ga.deleted_at " +
		"FROM game_playtimes gp JOIN games ga ON gp.game_id = ga.id " +
		"WHERE TO_DATE(gp.date, 'YYYY-MM-DD') >= CURRENT_DATE - INTERVAL '14 day' " +
		"GROUP BY " +
		"ga.id, ga.publisher, ga.developer, ga.title," +
		"ga.description, ga.price, ga.release_date," +
		"ga.created_at, ga.updated_at, ga.deleted_at " +
		"ORDER BY SUM(gp.play_hour) DESC " +
		"LIMIT 3").
		Scan(&featuredGames)

	return featuredGames
}

func GetCommunityRecommends() []models.Game {
	db := database.GetInstance()
	var communityRecommendedGames []models.Game
	db.Raw("SELECT * FROM games WHERE id IN (" +
		   		"SELECT game_id FROM game_reviews " +
				"GROUP BY game_id ORDER BY SUM(upvote_count) " +
				"DESC LIMIT 4" +
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