package game

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
)

func GetAll() []models.Game{
	db := database.GetInstance()
	var games []models.Game

	db.Find(&games)

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