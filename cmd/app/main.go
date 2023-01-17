package main

import (
	"github.com/sabitvrustam/new/pkg/database"
	transport "github.com/sabitvrustam/new/pkg/transport/http"
)

func main() {
	go transport.Tgbot()
	transport.Handler()
	go database.DbTest()

}
