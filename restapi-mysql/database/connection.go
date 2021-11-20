package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

const (
	host     = "localhost"
	port		= 5432
	user		= "root"
	password	= ""
	dbname		= "fiber_mysql"
)

func Connect() error {
	var err error
	//use dsn string to open
	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",user,password,dbname))
	if err != nil {
		return err
	}
	if err = Db.Ping(); err != nil {
		return err
	}
	return nil
}