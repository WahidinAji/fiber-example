package main

import (
	"fmt"
	"log"

	"restapi_pgsql/components"
	"restapi_pgsql/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	db := database.GetConnection()
	fmt.Println(db)
	defer db.Close()
	// Initialize standard Go html template engine
    engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
        Views: engine,
    })
    app.Get("/", components.Index)

    log.Fatal(app.Listen(":3000"))
}