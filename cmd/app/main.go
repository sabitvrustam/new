package main

import (
	"github.com/sabitvrustam/new/pkg/database"
	"github.com/sabitvrustam/new/pkg/logger"
	"github.com/sabitvrustam/new/pkg/transport/http"
	//"github.com/sabitvrustam/new/pkg/transport/telegram"
)

func main() {
	log := logger.Init()
	db, err := database.MySqlConnect(log)
	if err != nil {
		panic(err)
	}

	//telegramm := telegram.NewTelegram(db, log)
	database.Migrate(db, log)
	//go telegramm.Tgbot()
	http.StartHandler(db, log)

}
