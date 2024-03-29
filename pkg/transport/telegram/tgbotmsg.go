package telegram

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sabitvrustam/new/pkg/database/orders"
)

type Telegram struct {
	db    *sql.DB
	order *orders.Order
	log   *logrus.Logger
}

func NewTelegram(db *sql.DB, log *logrus.Logger) *Telegram {
	return &Telegram{db: db,
		order: orders.NewOrder(db, log),
		log:   log}
}

func (d *Telegram) Tgbot() {
	// подключаемся к боту с помощью токена
	key := os.Getenv("tbotapi")
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		d.log.Error(err)
	}
	bot.Debug = true
	d.log.Infof("Authorized on account %s", bot.Self.UserName)
	tgbotapi.SetLogger(d.log)
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
		msg.Text = ("нажмите /start")
		if update.Message == nil {
			continue
		}
		text := update.Message.Text
		if text == "/start" {
			hello := (`Здравствуйте я виртуальный помошник Рустам. чтобы узнать состояние заказа перейдите по ссылке /AKT, и введите номер квитанции`)
			msg.Text = hello
		}
		if text == "/AKT" {
			nomberAkt := (`Введите номер квитанции`)
			msg.Text = nomberAkt
		}
		if _, err := strconv.Atoi(text); err == nil {
			nom, _ := strconv.ParseInt(text, 10, 64)
			result, err := d.order.GetOrderByID(nom)
			if err != nil {
				d.log.Error(err)
			}
			if result != nil {
				msg.Text = fmt.Sprintf(" Акт № %d\n клиент - %s %s %s\n статус ремонта %s", result.IdOrder, result.User.FirstName,
					result.User.LastName, result.User.MidlName, result.StatusOrder)
			} else {
				msg.Text = ("такой квитанции нет")
			}

		}

		bot.Send(msg)
	}
}
