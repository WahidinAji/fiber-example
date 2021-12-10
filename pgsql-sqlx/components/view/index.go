package view

import "github.com/gofiber/fiber/v2"

func Index(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{
		"Title": "Restfull API | PgSql",
		"CardTitle": "Just Card Title, Nothing More!!",
		"Body": "Hi there, this is just simple create, read and delete.",
		"Message": "i didn't implement update bcz i'm too lazy now ^^",
	})
}