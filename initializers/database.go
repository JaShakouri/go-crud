package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {
	log.Println("Connecting to the database...")
	var err error
	DB, err = gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	log.Println("Database connected")
}
