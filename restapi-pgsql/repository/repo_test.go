package repository

import (
	"context"
	"fmt"
	"restapi_pgsql/database"
	"restapi_pgsql/helpers"
	"testing"
)

func TestGetAll(t *testing.T) {
	commentRepository := NewBookRepository(database.GetConnection())

	comments, err := commentRepository.GetAll(context.Background())
	helpers.PanicError(err)
	for _, comment := range comments {
		fmt.Println(comment)
	}
}