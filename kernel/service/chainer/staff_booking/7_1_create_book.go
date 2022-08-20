package staffBooking

import (
	"context"
	"fmt"
	"time"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type createBook struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewCreateBook(opts *chainer.Opts) chainer.Chainer {
	return &createBook{opts: opts}
}

func (r *createBook) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r createBook) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.StaffChangeTimeStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	bookTypeAndCatAndObjAndDay := r.opts.Update.CallbackQuery.Data
	//bookTypeAndCatAndObjAndDaySL := strings.Split(r.opts.Update.CallbackQuery.Data, "$")
	//categoryID, _ := strconv.Atoi(bookTypeAndCatAndObjAndDaySL[1])
	//objID, _ := strconv.Atoi(bookTypeAndCatAndObjAndDaySL[2])

	now := time.Now()
	rows := make([][]tg.InlineKeyboardButton, 0, 5)

	for i := 0; i < 5; i++ {
		row := tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(fmt.Sprintf("%s-%s", now.Format(timePattern), now.Add(time.Hour).Format(timePattern)),
				fmt.Sprintf("%s$%s", bookTypeAndCatAndObjAndDay,
					fmt.Sprintf("$%s$%s", now.Format(timePattern), now.Add(time.Hour).Format(timePattern)))),

			tg.NewInlineKeyboardButtonData(fmt.Sprintf("%s-%s", now.Add(time.Hour).Format(timePattern), now.Add(time.Hour*2).Format(timePattern)),
				fmt.Sprintf("%s$%s", bookTypeAndCatAndObjAndDay,
					fmt.Sprintf("$%s$%s", now.Add(time.Hour).Format(timePattern), now.Add(time.Hour*2).Format(timePattern)))),

			tg.NewInlineKeyboardButtonData(fmt.Sprintf("%s-%s", now.Add(time.Hour*2).Format(timePattern), now.Add(time.Hour*3).Format(timePattern)),
				fmt.Sprintf("%s$%s", bookTypeAndCatAndObjAndDay,
					fmt.Sprintf("$%s$%s", now.Add(time.Hour*2).Format(timePattern), now.Add(time.Hour*3).Format(timePattern)))),
		)
		now = now.Add(time.Hour * 3)
		rows = append(rows, row)
	}
	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, "Выбери время:", tg.NewInlineKeyboardMarkup(rows...))

	user.HandleStep = int(chainer.StaffChangeTimeStep)
	err := r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
