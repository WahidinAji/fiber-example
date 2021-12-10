package book

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	route := api.Group("/books")
	route.Get("/", GetAll)
	route.Post("/", Create)
	route.Get("/:id", Finder)
	route.Delete("/:id", Delete)
}