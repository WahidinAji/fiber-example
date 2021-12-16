package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"restapi_pgsql/helpers"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/joho/godotenv/autoload" //this package only for database_test_connection
)

func GetConnection() *sql.DB {
	
	if confAppName := os.Getenv("APP_NAME"); confAppName == "" {
		log.Fatal("APP_NAME environment variable")
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("DB_USER environment variable")
	}
	dbPort := os.Getenv("SERVER_PORT")
	if dbPort == "" {
		log.Fatal("SERVER_PORT environment variable")
	}
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		log.Fatal("DB_PASS environment variable")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME environment variable")
	}
	dsn := fmt.Sprintf("postgresql://%s:%s@127.0.0.1:%s/%s",dbUser,dbPass,dbPort,dbName)
	
	db, err := sql.Open("pgx", dsn)
	helpers.PanicError(err)
	db.SetMaxIdleConns(10) //min 10 connection
	db.SetMaxOpenConns(100) //max 100 connection
	db.SetConnMaxIdleTime(5 * time.Minute) //if iin 5 minutes nothing happen? db will close the connection
	db.SetConnMaxLifetime(60 * time.Minute) //after 60 minutes, the connection will create new connection
	return db
}