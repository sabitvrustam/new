package database

import (
	"database/sql"
	"fmt"
	"os"
)

func MySqlConnect() (db *sql.DB, err error) {

	var dbuser string = os.Getenv("bduser")
	var dbpass string = os.Getenv("bdpass")
	var pass string = fmt.Sprintf("%s:%s@tcp(127.0.0.1)/my_service", dbuser, dbpass)
	db, err = sql.Open("mysql", pass)
	return db, err

}
