package book

import (
	"database/sql"
	"log"
	"pgsql-sqlx/database"

	"github.com/gofiber/fiber/v2"
)


func Delete(ctx *fiber.Ctx) error{
	id, err := ctx.ParamsInt("id")
	if err != nil{
		return err
	}
	querySql := "DELETE FROM books WHERE id=$1"
	tx, err := database.Db.BeginTxx(ctx.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	res, err := tx.ExecContext(ctx.Context(), querySql, id)
	if err != nil {
		tx.Rollback()
		log.Panic(res)
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"code": ctx.Response().StatusCode(),
		"status": "OK",
	})
}