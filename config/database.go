package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error in the .env: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	var dsn string
	var dialector gorm.Dialector

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	dialector = postgres.Open(dsn)

	DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect DB: %v", err)
	}

	log.Println("Conected")

	// autoMigrate()

}

// func autoMigrate(){
// 	err := DB.AutoMigrate(
// 		&models.User{},
// 		&models.Favorite{},
// 		&models.Movie{},
// 		&models.Admin{},
// 		&models.Serie{},
// 		&models.Actor{},
// 	)
// 	if err != nil {
// 		log.Fatalf("Error en la migración de la base de datos: %v", err)
// 	}

// 	log.Println("Migración completada")
// }
