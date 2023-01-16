package main

func main() {
	go Tgbot()
	handler()
	go DbTest()

}
