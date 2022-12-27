package main

func main() {
	go tgbot()
	handler()
	go dbTest()

}
