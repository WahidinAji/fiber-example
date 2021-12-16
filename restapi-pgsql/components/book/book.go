package book

import "github.com/jmoiron/sqlx"

type Book struct {
	ID          int    `json:"id"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Dependency struct {
	DB *sqlx.DB
}
