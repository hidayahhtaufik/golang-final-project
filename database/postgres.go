package database

import (
	"fmt"
	"log"
	"myGram/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	_       = godotenv.Load()
	db      *gorm.DB
	err     error
	host    = os.Getenv("DB_HOST")
	port    = os.Getenv("DB_PORT")
	user    = os.Getenv("DB_USER")
	pass    = os.Getenv("DB_PASSWORD")
	dbname  = os.Getenv("DB_NAME")
	sslmode = os.Getenv("DB_SSL_MODE")
)

func StartDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, dbname, sslmode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	fmt.Println("Connected to database success")
}

func GetDB() *gorm.DB {
	return db
}
