package developer

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
)

func GetAll() []models.Developer {
	db := database.GetInstance()

	var devs []models.Developer
	db.Find(&devs)

	return devs
}

func Get(id int) models.Developer {
	db := database.GetInstance()
	var dev models.Developer

	db.Find(&dev, id)

	return dev
}
