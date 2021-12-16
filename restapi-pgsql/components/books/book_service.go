package books

// import (
// 	"database/sql"
// 	"restapi_pgsql/helpers"

// 	"github.com/gofiber/fiber/v2"
// )

// type BookService interface {
// 	GetAll(ctx *fiber.Ctx) []BookResponse
// }

// type BookServiceImpl struct {
// 	BookRepository BookRepository
// 	DB *sql.DB
// }

// func NewBookServiceImpl(bookRepository BookRepository, DB *sql.DB) BookService{
// 	return &BookServiceImpl{
// 		BookRepository: bookRepository,
// 		DB: DB,
// 	}
// }

// func (s *BookServiceImpl) GetAll(ctx *fiber.Ctx) []BookResponse{
// 	books, err:= s.BookRepository.GetAll(ctx.Context())
// 	helpers.PanicError(err)
// 	return ToBookResponses(books)
// }