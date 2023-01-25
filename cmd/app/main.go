package main

import (
	"github.com/sabitvrustam/new/pkg/database"
	"github.com/sabitvrustam/new/pkg/transport/http"
	"github.com/sabitvrustam/new/pkg/transport/telegram"
)

func main() {

	db, err := database.MySqlConnect()
	if err != nil {
		panic(err)
	}
	database.Migrate(db)
	go telegram.Tgbot()
	http.StartHandler(db)

}
