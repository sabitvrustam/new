package main

import (
	"github.com/sabitvrustam/new/pkg/database"
	"github.com/sabitvrustam/new/pkg/transport"
)

func main() {
	go transport.Tgbot()
	transport.Handler()
	go database.DbTest()

}
