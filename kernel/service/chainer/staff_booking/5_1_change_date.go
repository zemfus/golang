package staffBooking

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type changeDate struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewChangeDate(opts *chainer.Opts) chainer.Chainer {
	return &changeDate{opts: opts}
}

func (r *changeDate) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

var patternDate = "2/1/2006"

func (r changeDate) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.StaffChangeDateStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	bookTypeAndCatAndObj := r.opts.Update.CallbackQuery.Data
	//bookTypeAndCatAndObjSL := strings.Split(r.opts.Update.CallbackQuery.Data, "$")
	//categoryID, _ := strconv.Atoi(bookTypeAndCatAndObjSL[1])
	//objID, _ := strconv.Atoi(bookTypeAndCatAndObjSL[2])

	now := time.Now()
	rows := make([][]tg.InlineKeyboardButton, 0, 5)

	for i := 0; i < 5; i++ {
		row := tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(strconv.Itoa(now.Day()), fmt.Sprintf("%s$%s",
				bookTypeAndCatAndObj, now.Format(patternDate))),
			tg.NewInlineKeyboardButtonData(strconv.Itoa(now.Add(time.Hour*24).Day()), fmt.Sprintf("%s$%s",
				bookTypeAndCatAndObj, now.Add(time.Hour*24).Format(patternDate))),
			tg.NewInlineKeyboardButtonData(strconv.Itoa(now.Add(time.Hour*48).Day()), fmt.Sprintf("%s$%s",
				bookTypeAndCatAndObj, now.Add(time.Hour*48).Format(patternDate))),
			tg.NewInlineKeyboardButtonData(strconv.Itoa(now.Add(time.Hour*72).Day()), fmt.Sprintf("%s$%s",
				bookTypeAndCatAndObj, now.Add(time.Hour*72).Format(patternDate))),
			tg.NewInlineKeyboardButtonData(strconv.Itoa(now.Add(time.Hour*96).Day()), fmt.Sprintf("%s$%s",
				bookTypeAndCatAndObj, now.Add(time.Hour*96).Format(patternDate))),
			tg.NewInlineKeyboardButtonData(strconv.Itoa(now.Add(time.Hour*120).Day()), fmt.Sprintf("%s$%s",
				bookTypeAndCatAndObj, now.Add(time.Hour*120).Format(patternDate))),
			tg.NewInlineKeyboardButtonData(strconv.Itoa(now.Add(time.Hour*144).Day()), fmt.Sprintf("%s$%s",
				bookTypeAndCatAndObj, now.Add(time.Hour*144).Format(patternDate))),
		)
		now = now.Add(time.Hour * 144)
		rows = append(rows, row)
	}
	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, "Выбери число:", tg.NewInlineKeyboardMarkup(rows...))

	user.HandleStep = int(chainer.StaffChangeTimeStep)
	err := r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
