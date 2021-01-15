package publisher

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
)

func GetAll() []models.Publisher{
	db := database.GetInstance()

	var publishers []models.Publisher
	db.Find(&publishers)

	return publishers
}

func Get(id int) models.Publisher {
	db := database.GetInstance()

	var pub models.Publisher
	db.Find(&pub, id)

	return pub
}