package book

import (
	"pgsql-sqlx/database"

	response "github.com/WahidinAji/web-response"
	"github.com/gofiber/fiber/v2"
)

func GetAll(ctx *fiber.Ctx) error {
	script := "SELECT id, author, title, description FROM books"
	rows, err := database.Db.QueryxContext(ctx.Context(), script)
	if err != nil {
		return err
	}
	result := []Book{}
	// var result []Book
	for rows.Next() {
		book := Book{}
		// if err := rows.StructScan(&book); err != nil {
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Description)
		if err != nil {
			return err
		}
		result = append(result, book)
	}
	ctx.Status(200)
	return ctx.JSON(
		response.WebResponse(fiber.StatusOK, "OK", result),
	)
}