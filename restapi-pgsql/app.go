package main

import (
	"log"

	"github.com/WahidinAji/fiber-example/restapi-pgsql/components"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Initialize standard Go html template engine
    engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
        Views: engine,
    })
    app.Get("/", components.Index)

    log.Fatal(app.Listen(":3000"))
}