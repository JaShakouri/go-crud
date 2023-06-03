package main

import (
	"crm-api/initializers"
	"crm-api/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func migrationPost() {
	log.Println("Migrating post model...")
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Migration post model was failure => " + err.Error())
	}
	log.Println("Migrating post model Done")
}

func main() {

	log.Println("Migrating start...")
	log.Println("........................")
	migrationPost()

}
