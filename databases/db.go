package databases

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	time_zone := os.Getenv("DB_TIMEZONE")

	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + db_name + " port=" + port + " sslmode=disable TimeZone=" + time_zone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error : " + err.Error())
	}

	DB = db
}
