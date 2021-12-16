package service

import (
	"database/sql"
	"restapi_pgsql/helpers"
	"restapi_pgsql/repository"
	"restapi_pgsql/web"

	"github.com/gofiber/fiber/v2"
)

type BookService interface {
	GetAll(ctx *fiber.Ctx) []web.BookResponse
}

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB *sql.DB
}

func NewBookServiceImpl(bookRepository repository.BookRepository, DB *sql.DB) BookService{
	return &BookServiceImpl{
		BookRepository: bookRepository,
		DB: DB,
	}
}

func (s *BookServiceImpl) GetAll(ctx *fiber.Ctx) []web.BookResponse{
	books, err:= s.BookRepository.GetAll(ctx.Context())
	helpers.PanicError(err)
	return helpers.ToBookResponses(books)
}