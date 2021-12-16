package helpers

import (
	"restapi_pgsql/entity"
	"restapi_pgsql/web"
)

func ToCategoryResponse(book entity.Book) web.BookResponse{
	return web.BookResponse{
		ID: book.ID,
		Author: book.Author,
		Title: book.Title,
		Description: book.Description,
	}
}

func ToBookResponses(book []entity.Book) []web.BookResponse{
	var bookResponses []web.BookResponse
	for _, book := range book {
		bookResponses = append(bookResponses, ToCategoryResponse(book))
	}
	return bookResponses
}