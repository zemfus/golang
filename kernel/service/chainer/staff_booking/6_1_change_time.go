package staffBooking

import (
	"context"
	"fmt"
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
	cmdBookTypeAndCatAndObjAndDaySL := strings.Split(r.opts.Update.CallbackQuery.Data, "$")
	//categoryID, _ := strconv.Atoi(cmdBookTypeAndCatAndObjAndDaySL[1])
	//objID, _ := strconv.Atoi(cmdBookTypeAndCatAndObjAndDaySL[2])

	now := time.Now()
	rows := make([][]tg.InlineKeyboardButton, 0, 5)

	for i := 0; i < 5; i++ {
		row := tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(fmt.Sprintf("%s-%s", now.Format(timePattern), now.Add(time.Hour).Format(timePattern)),
				fmt.Sprintf("%s$%s", cmdBookTypeAndCatAndObjAndDay,
					fmt.Sprintf("$%s$%s", now.Format(timePattern), now.Add(time.Hour).Format(timePattern)))),

			tg.NewInlineKeyboardButtonData(fmt.Sprintf("%s-%s", now.Add(time.Hour).Format(timePattern), now.Add(time.Hour*2).Format(timePattern)),
				fmt.Sprintf("%s$%s", cmdBookTypeAndCatAndObjAndDay,
					fmt.Sprintf("$%s$%s", now.Add(time.Hour).Format(timePattern), now.Add(time.Hour*2).Format(timePattern)))),

			tg.NewInlineKeyboardButtonData(fmt.Sprintf("%s-%s", now.Add(time.Hour*2).Format(timePattern), now.Add(time.Hour*3).Format(timePattern)),
				fmt.Sprintf("%s$%s", cmdBookTypeAndCatAndObjAndDay,
					fmt.Sprintf("$%s$%s", now.Add(time.Hour*2).Format(timePattern), now.Add(time.Hour*3).Format(timePattern)))),
		)
		now = now.Add(time.Hour * 3)
		rows = append(rows, row)
	}

	rows = append(rows, tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("Назад", fmt.Sprintf("%d$%s", chainer.StaffChangeDateStep,
			strings.Join(cmdBookTypeAndCatAndObjAndDaySL[1:len(cmdBookTypeAndCatAndObjAndDaySL)-1], "$"))),
	))

	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, "Выбери время:", tg.NewInlineKeyboardMarkup(rows...))

	//user.HandleStep = int(chainer.StaffChangeTimeStep)
	//err := r.opts.UserRepo.Update(ctx, user)
	//if err != nil {
	//	return nil, err
	//}

	return &msgReply, nil
}
