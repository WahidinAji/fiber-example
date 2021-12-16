package books

import (
	"database/sql"
	"restapi_pgsql/helpers"

	"github.com/gofiber/fiber/v2"
)

type bookRepositoryImpl struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository  {
	// kenapa return pakai &, karena setiap method dibawahnya di implemenetasikan menggunakan pointer.
	return &bookRepositoryImpl{DB: db}
}

func (repository *bookRepositoryImpl) GetAll(ctx *fiber.Ctx) ([]Book, error){
	script := "SELECT id, author, title, description FROM books"
	rows, err := repository.DB.QueryContext(ctx.Context(), script)
	helpers.PanicError(err)
	defer rows.Close()
	var books []Book
	for rows.Next() {
		book := Book{}
		rows.Scan(&book.ID, &book.Author, &book.Title, &book.Description)
		books = append(books, book)
	}
	return books, nil
}