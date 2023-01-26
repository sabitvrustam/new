package main

import (
	"os"

	"github.com/sabitvrustam/new/pkg/database"
	"github.com/sabitvrustam/new/pkg/transport/http"
	"github.com/sabitvrustam/new/pkg/transport/telegram"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.TraceLevel)
	log.SetOutput(os.Stdout)
}

func main() {
	db, err := database.MySqlConnect()
	if err != nil {
		panic(err)
	}

	telegramm := telegram.NewTelegram(db)
	database.Migrate(db)
	go telegramm.Tgbot()
	http.StartHandler(db)

}
