package staffBooking

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type changeTime struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewChangeTime(opts *chainer.Opts) chainer.Chainer {
	return &changeTime{opts: opts}
}

func (r *changeTime) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

var timePattern = "15:04"

func (r changeTime) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.StaffChangeTimeStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	cmdBookTypeAndCatAndObjAndDay := r.opts.Update.CallbackQuery.Data
	//_, bookTypeAndCatAndObjAndDay, _ := strings.Cut(r.opts.Update.CallbackQuery.Data, "$")
	cmdBookTypeAndCatAndObjAndDaySL := strings.Split(r.opts.Update.CallbackQuery.Data, "$")
	bookType := cmdBookTypeAndCatAndObjAndDaySL[1]
	categoryID, _ := strconv.Atoi(cmdBookTypeAndCatAndObjAndDaySL[2])
	objID, _ := strconv.Atoi(cmdBookTypeAndCatAndObjAndDaySL[3])
	ti, err := time.Parse("2/1/2006", cmdBookTypeAndCatAndObjAndDaySL[4])
	if err != nil {
		return nil, err
	}
	//err = r.opts.BookRepo.Create(ctx, &models.Booking{
	//	BookType:    "inventory",
	//	UserID:      234899515,
	//	InventoryID: &objID,
	//	PlacesID:    nil,
	//	Confirm:     false,
	//	StartAt:     time.Now().Add(time.Hour * 3),
	//	EndAt:       time.Now().Add(time.Hour * 5),
	//	Status:      false,
	//})
	//if err != nil {
	//	return nil, err
	//}
	var text string
	bookings, err := r.opts.BookRepo.GetActiveBookings(ctx, &models.Booking{
		ID:          0,
		BookType:    "",
		UserID:      user.ID,
		InventoryID: &objID,
		PlacesID:    nil,
		Confirm:     false,
		StartAt:     ti,
		EndAt:       ti,
		Status:      false,
	})
	if err != nil {
		return nil, err
	}

	for i, booking := range bookings {
		usr, err := r.opts.UserRepo.GetByID(ctx, booking.UserID)
		if err != nil {
			return nil, err
		}
		text += fmt.Sprintf("%d) %s, %s-%s\n\n", i+1, usr.Nickname,
			booking.StartAt.Format("15:04"),
			booking.EndAt.Format("15:04"),
		)
	}

	if len(text) != 0 {
		text = fmt.Sprintf("На текущую дату имеюся действуюшие бронирование:\n\n%s\nвыберите время исходя из времени забронированных слотов.", text)
	}

	println(cmdBookTypeAndCatAndObjAndDay, bookType, categoryID, objID)
	//now := time.Now()
	keyboard := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Начало", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData("Время брони", fmt.Sprint(chainer.NonStep)),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("+", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData(" ", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData("+", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData(" ", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData(" ", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData("+", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData(" ", fmt.Sprint(chainer.NonStep)),
		), tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("00", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData("ч", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData("00", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData("м", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData(" ", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData("60", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData("м", fmt.Sprint(chainer.NonStep)),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("-", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData(" ", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData("-", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData(" ", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData(" ", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData("-", fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData(" ", fmt.Sprint(chainer.NonStep)),
		),

		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Готово", fmt.Sprint(chainer.NonStep)),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Назад", fmt.Sprintf("%d$%s", chainer.StaffChangeDateStep,
				strings.Join(cmdBookTypeAndCatAndObjAndDaySL[1:len(cmdBookTypeAndCatAndObjAndDaySL)-1], "$"))),
		),
	)

	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, text, keyboard)

	//user.HandleStep = int(chainer.StaffChangeTimeStep)
	//err := r.opts.UserRepo.Update(ctx, user)
	//if err != nil {
	//	return nil, err
	//}

	return &msgReply, nil
}
