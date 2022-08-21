package btn

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	Start = "/start" // start

	Booking       = "Бронирование"
	MyBooking     = "Мои бронирование"
	Requests      = "Заявки"
	Configuration = "Конфигурация"
	Events        = "Мероприятия"
	Subscriptions = "Подписки"
	Status        = "Статус"
)

var AllCmds = map[string]bool{
	Start:         true,
	Booking:       true,
	MyBooking:     true,
	Requests:      true,
	Configuration: true,
	Events:        true,
	Subscriptions: true,
	Status:        true,
}

var Staff = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(Booking),
		tg.NewKeyboardButton(Requests),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(Configuration),
		tg.NewKeyboardButton(Events),
	),
)

var Student = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(Booking),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(Subscriptions),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(Status),
	),
)
