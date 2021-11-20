package main

import (
	"log"

	"github.com/WahidinAji/fiber-example/restapi-mysql/database"
	"github.com/WahidinAji/fiber-example/restapi-mysql/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)


func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}