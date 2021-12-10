package book

import (
	"pgsql-sqlx/database"
	"pgsql-sqlx/helpers"

	"github.com/gofiber/fiber/v2"
)

func Finder(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	querySql := `SELECT id, author, title, description FROM books WHERE id = $1`
	// rows, err := database.Db.QueryxContext(ctx.Context(), querySql, id)
	rows, err := database.Db.QueryxContext(ctx.Context(), querySql, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	result := []Book{}
	if rows.Next() {
		book := Book{}
		// err := rows.Scan(&book.ID, &book.Author,&book.Title, &book.Description)
		err := rows.StructScan(&book)
		if err != nil {
			return err
		}
		result = append(result, book)
		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
			"code": ctx.Response().StatusCode(),
			"status": "OK",
			"data": result,
		})
	} else {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(helpers.RespApi(ctx.Response().StatusCode(),"Not Found", result))
	}
}