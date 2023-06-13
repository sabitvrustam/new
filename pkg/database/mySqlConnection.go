package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func MySqlConnect(log *logrus.Logger) (db *sql.DB, err error) {

	var dbuser string = os.Getenv("bduser")
	var dbpass string = os.Getenv("bdpass")
	var pass string = fmt.Sprintf("%s:%s@tcp(localhost:3306)/my_service", dbuser, dbpass)
	db, err = sql.Open("mysql", pass)
	if err != nil {
		log.Error("не удалось подключиться к базе данных", err)
	} else {
		log.Info("Подключение к базе данных")
	}
	return db, err

}
