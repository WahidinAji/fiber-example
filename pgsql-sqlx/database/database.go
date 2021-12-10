package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

// Connect function
func Connect() (err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	Db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}
	if err = Db.Ping(); err != nil {
		return err
	}
	return nil
}