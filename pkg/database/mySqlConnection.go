package database

import (
	"database/sql"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func MySqlConnect() (db *sql.DB, err error) {

	var dbuser string = os.Getenv("bduser")
	var dbpass string = os.Getenv("bdpass")
	var pass string = fmt.Sprintf("%s:%s@tcp(127.0.0.1)/my_service", dbuser, dbpass)
	db, err = sql.Open("mysql", pass)
	if err != nil {
		log.Error("не удалось подключиться к базе данных", err)
	} else {
		log.Info("Подключение к базе данных")
	}
	return db, err

}
