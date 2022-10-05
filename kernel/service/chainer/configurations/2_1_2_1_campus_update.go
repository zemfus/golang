package configurations

import (
	"context"
	"fmt"
	"strings"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type campusUpdate struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewCampusUpdate(opts *chainer.Opts) chainer.Chainer {
	return &campusUpdate{opts: opts}
}

func (r *campusUpdate) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r campusUpdate) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.CfgCampusUpdateStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	campusIDStr := strings.Split(r.opts.Update.CallbackQuery.Data, "$")[1]

	var campusUpdateKeyboard = tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Назад", fmt.Sprint(chainer.CfgCampusEditStep)),
		),
	)

	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, "Введите новое имя:", campusUpdateKeyboard)

	user.HandleStep = int(chainer.CfgCampusUpdateExecStep)
	user.LastMsg = campusIDStr
	err := r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return &msgReply, nil
}
