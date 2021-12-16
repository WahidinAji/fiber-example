package book

import (
	"restapi_pgsql/helpers"

	"github.com/gofiber/fiber/v2"
	// "github.com/jmoiron/sqlx"
)

func Handler(db *bookRepoImpl) *fiber.App{
	app := fiber.New()
	app.Get("/books", func(c *fiber.Ctx) error {
		books, err := db.GetAll(c.Context())
		helpers.PanicError(err)
		c.Status(fiber.StatusOK)
		return c.JSON(books)
	})
	return app
}
