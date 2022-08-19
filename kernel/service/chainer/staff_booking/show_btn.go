package staffBooking

import (
	"context"
	"fmt"

	"boobot/kernel/domain/btn"
	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type showBtn struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewShowBtn(opts *chainer.Opts) chainer.Chainer {
	return &showBtn{opts: opts}
}

func (r *showBtn) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r showBtn) Handle(ctx context.Context, user *models.User) (*tg.MessageConfig, error) {
	if int(chainer.StaffShowBtnBookingsStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	var msgReply tg.MessageConfig
	msgReply.Text = btn.Booking

	var staffKeyboard = tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Создать", fmt.Sprint(chainer.StaffCreateBookingsStep)),
			tg.NewInlineKeyboardButtonData("Просмотреть", fmt.Sprint(chainer.StaffShowBookingsStep)),
		),
	)

	msgReply.ReplyMarkup = staffKeyboard

	//todo push steps
	//user.HandleStep = int(chainer.NonStep)

	return &msgReply, nil
}
