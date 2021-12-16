package book

import (
	"context"
	"database/sql"
	"restapi_pgsql/helpers"
)

type Repo interface {
	GetAll(ctx context.Context) ([]Book, error)
}

type bookRepoImpl struct {
	DB *sql.DB
}

func NewBookRepo(db *sql.DB) Repo  {
	return &bookRepoImpl{DB: db}
}

func (repo *bookRepoImpl) GetAll(ctx context.Context) ([]Book, error) {
	script := "SELECT id, author, title, description FROM books"
	rows, err := repo.DB.QueryContext(ctx, script)
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