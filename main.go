package main

import (
	"myGram/database"
	"myGram/routers"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")

	database.StartDB()

	server := router.StartApp()
	server.Run(PORT)
}
