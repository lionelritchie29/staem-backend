package game_item

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
)

func GetByUserId(userId int) ([]models.GameItem){
	db := database.GetInstance()
	var items []models.GameItem
	userIdStr := strconv.FormatInt(int64(userId), 10)
	db.Raw("SELECT * FROM game_items WHERE id IN( " +
				"SELECT game_item_id FROM user_game_items " +
				"WHERE user_id = " + userIdStr +
			")").Scan(&items)
	return items
}
