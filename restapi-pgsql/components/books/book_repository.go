package books

import "github.com/gofiber/fiber/v2"

type BookRepository interface {
	GetAll(ctx *fiber.Ctx) ([]Book, error)
}