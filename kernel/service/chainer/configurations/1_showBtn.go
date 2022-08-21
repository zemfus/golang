package configurations

import (
	"context"
	"fmt"

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

func (r showBtn) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.CfgShowBtn) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	var msgReply tg.MessageConfig

	msgReply.Text = "Бронирование:"
	var staffKeyboard = tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Кампусы", fmt.Sprint(chainer.StaffChangeTypeStep)),
			tg.NewInlineKeyboardButtonData("Категории", fmt.Sprint(chainer.StaffShowBookingsStep)),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Помещения", fmt.Sprint(chainer.StaffChangeTypeStep)),
			tg.NewInlineKeyboardButtonData("Инвентарь", fmt.Sprint(chainer.StaffShowBookingsStep)),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Студенты", fmt.Sprint(chainer.StaffChangeTypeStep)),
		),
	)

	msgReply.ReplyMarkup = staffKeyboard

	user.HandleStep = int(chainer.CfgProxyItems)
	err := r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
