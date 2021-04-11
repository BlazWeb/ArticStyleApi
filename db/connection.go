package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user     = "root"
	db_password = ""
	db_addr     = "127.0.0.1"
	db_name     = "articstyle"
)

func GetDB() (*sql.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_addr, db_name)
	fmt.Println(s)
	db, err := sql.Open("mysql", s)

	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
