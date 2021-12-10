package view

import (
	"github.com/gofiber/fiber/v2"
)

func ViewRoute(app *fiber.App) {	
	app.Get("/", Index)
}