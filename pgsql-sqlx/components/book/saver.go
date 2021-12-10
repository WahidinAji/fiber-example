package book

import (
	"database/sql"
	"fmt"
	"log"
	"pgsql-sqlx/database"

	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	newBook := new(Book)
	err := ctx.BodyParser(newBook)
	if err != nil {
		return ctx.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}
	querySql := "INSERT INTO books (author, title,description) VALUES ($1, $2, $3) RETURNING id"
	tx, err := database.Db.BeginTxx(ctx.Context(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	var id int
	row, execErr := tx.QueryxContext(ctx.Context(), querySql, &newBook.Author, &newBook.Title, &newBook.Description)
	if execErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return execErr
	}
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}
	newBook.ID = int(id)
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(id)
	return ctx.Status(fiber.StatusOK).JSON(
		Resp(fiber.StatusOK, "OK", newBook),
	)
}

//non transaction version
func Insert(c *fiber.Ctx) error {
	// New Employee struct
	u := new(Book)

	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var id int
	// Insert Employee into database
	res := database.Db.QueryRow("INSERT INTO books (author, title, description)VALUES ($1, $2, $3) RETURNING id", u.Author, u.Title, u.Description)
	if err := res.Scan(&id); err != nil {
		return err
	}

	// Print result
	log.Println(res, id)
	u.ID = id

	// Return Employee in JSON format
	return c.JSON(u)
}