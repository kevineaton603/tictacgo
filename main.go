package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kevineaton603/tictacgo/handlers"
	"github.com/kevineaton603/tictacgo/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	connectionURL := os.Getenv("CONNECTION_URL")
	db, err := gorm.Open(postgres.Open(connectionURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Game{})

	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.Redirect(302, "game/")
	})

	app.GET("/game", func(c echo.Context) error {
		return c.Redirect(302, "game/")
	})

	gameHandler := handlers.NewGameHandler(app, db)

	gameHandler.Mount()

	app.Static("/assets", "assets")

	app.Logger.Fatal(app.Start(":1323"))
}
