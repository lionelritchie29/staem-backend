package database


import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

const DB_NAME=""
const DB_HOST="127.0.0.1"
const DB_PORT="5432"				//port on installation
const DB_USER="postgres"			//default is postgres
const DB_PASS=""	//password on installation

func Connect()(*gorm.DB, error){

	return gorm.Open("postgres","host="+DB_HOST+" port="+DB_PORT+" user="+DB_USER+" dbname="+DB_NAME+" password="+DB_PASS+" sslmode=disable")
}