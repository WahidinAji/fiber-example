package repository

import (
	"context"
	"database/sql"
	"restapi_pgsql/entity"
)

type BookRepository interface {
	GetAll(ctx context.Context) ([]entity.Book, error)
}

type bookRepositoryImpl struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository{
	return &bookRepositoryImpl{DB: db}
}

func (repository *bookRepositoryImpl) GetAll(ctx context.Context) ([]entity.Book, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx,script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []entity.Book
	for rows.Next(){
		book := entity.Book{}
		rows.Scan(&book.ID,&book.Author,&book.Title,&book.Author)
		books = append(books, book)
	}
	return books, nil
}