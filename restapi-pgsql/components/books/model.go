package books

func ToBookResponse(book Book) BookResponse {
	return BookResponse{
		ID:          book.ID,
		Author:      book.Author,
		Title:       book.Title,
		Description: book.Description,
	}
}
func ToBookResponses(books []Book) []BookResponse {
	var bookResponses []BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}
	return bookResponses
}