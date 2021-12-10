package main

import (
	"log"
	"pgsql-sqlx/components/book"
	"pgsql-sqlx/database"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {	
	// Connect with database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	// Create a Fiber app
	app := fiber.New()
	book.SetupRoutes(app)
	app.Listen(":3000")
}

