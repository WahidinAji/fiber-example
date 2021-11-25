package routes

import "github.com/gofiber/fiber/v2"

func book(app *fiber.App){
	api := app.Group("/api")
	api.Get("/books")
	api.Post("/books")
	api.Get("/books/:id")
	api.Put("/books/:id")
	api.Delete("/books/:id")
}