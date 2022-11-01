package main

import "os"

func main() {
	os.Getenv("foo")
	go tgbot()
	handleFunc()

}
