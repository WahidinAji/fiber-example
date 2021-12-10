package main

import (
	"log"
	"pgsql-sqlx/components/book"
	"pgsql-sqlx/components/view"
	"pgsql-sqlx/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {	
	// Connect with database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	// Initialize standard Go html template engine
    engine := html.New("./views", ".html")
	// Create a Fiber app
	app := fiber.New(fiber.Config{
        Views: engine,
    })
	app.Get("/", view.Index)
	book.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

