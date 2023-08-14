package databases

import (
	"github.com/ihksanghazi/api-online-course/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := utils.GetToEnv("DB_HOST")
	user := utils.GetToEnv("DB_USER")
	pass := utils.GetToEnv("DB_PASSWORD")
	db_name := utils.GetToEnv("DB_NAME")
	port := utils.GetToEnv("DB_PORT")
	time_zone := utils.GetToEnv("DB_TIMEZONE")

	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + db_name + " port=" + port + " sslmode=disable TimeZone=" + time_zone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error : " + err.Error())
	}

	DB = db
}
