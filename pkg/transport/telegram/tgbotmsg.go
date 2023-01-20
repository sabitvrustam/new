package telegram

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sabitvrustam/new/pkg/database"
)

func Tgbot() {
	// подключаемся к боту с помощью токена
	key := os.Getenv("tbotapi")
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		fmt.Println(err)
	}
	// читаем обновления из канала
	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message == nil {
			continue
		}
		text := update.Message.Text
		if text == "/start" {
			hello := (`Здравствуйте я виртуальный помошник Рустам. чтобы узнать состояние заказа перейдите по ссылке /AKT, и введите номер квитанции`)
			msg.Text = hello
			bot.Send(msg)
		}
		if text == "/AKT" {
			nomberAkt := (`Введите номер квитанции`)
			msg.Text = nomberAkt
			bot.Send(msg)
		}
		nom, _ := strconv.ParseInt(text, 10, 64)
		result, _ := database.ReadOrder(nom)
		msg.Text = fmt.Sprintf(" Акт № %d\n клиент - %s %s %s\n статус ремонта %s", result.IdOrder, result.User.FirstName, result.User.LastName, result.User.MidlName, result.StatusOrder)
		bot.Send(msg)
	}
}
