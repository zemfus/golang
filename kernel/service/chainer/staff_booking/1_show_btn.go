package staffBooking

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
	if int(chainer.StaffShowBtnBookingsStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	text := "Бронирование:"
	var staffKeyboard = tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Создать", fmt.Sprintf("%d$%d", chainer.StaffProxyCreateVSShow, chainer.StaffChangeTypeStep)),
			tg.NewInlineKeyboardButtonData("Просмотреть", fmt.Sprintf("%d$%d", chainer.StaffProxyCreateVSShow, chainer.StaffShowBookingsStep)),
		),
	)

	if r.opts.Update.Message == nil {
		chatID := r.opts.Update.CallbackQuery.From.ID
		msgID := r.opts.Update.CallbackQuery.Message.MessageID
		msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, text, staffKeyboard)
		return msgReply, nil
	}

	var msgReply tg.MessageConfig
	msgReply.Text = text
	msgReply.ReplyMarkup = staffKeyboard

	//user.HandleStep = int(chainer.StaffProxyCreateVSShow)
	//err := r.opts.UserRepo.Update(ctx, user)
	//if err != nil {
	//	return nil, err
	//}

	return &msgReply, nil
}
