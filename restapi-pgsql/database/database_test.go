package database

import (
	"fmt"
	"testing"
)

func TestGetConnection(t *testing.T) {
	db := GetConnection()
	fmt.Println(db)
	db.Close()
}