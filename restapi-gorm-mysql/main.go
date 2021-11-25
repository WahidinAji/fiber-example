package main

import (
	"github.com/WahidinAji/fiber-example/restapi-gorm-mysql/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
}