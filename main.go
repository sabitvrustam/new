package main

import (
	"fmt"
	"io"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "start":
			file, err := os.Open("start.txt")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()
			data := make([]byte, 256)
			for {
				n, err := file.Read(data)
				if err == io.EOF {
					break
				}
				msg.Text = string(data[:n])
			}
		case "help":
			msg.Text = "I understand / sayhi and / status."
		case "read":
			file, err := os.Open("Potter.txt")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()
			data := make([]byte, 128)
			for {
				n, err := file.Read(data)
				if err == io.EOF {
					break
				}

				msg.Text = string(data[:n])
			}
		case "write":
			msg.Text = "Введите текс для сохранения"
			for updates == nil {

				reply := "1"

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)
			}

		default:
			msg.Text = "я не знаю такой команды"
		}
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
