package main

import (
	"github.com/sabitvrustam/new/pkg/database"
	"github.com/sabitvrustam/new/pkg/transport/http"
	"github.com/sabitvrustam/new/pkg/transport/telegram"
)

func main() {
	go telegram.Tgbot()
	http.Handler()
	go database.DbTest()

}
