package configurations

import (
	"context"
	"fmt"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type campusGet struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewCampusGet(opts *chainer.Opts) chainer.Chainer {
	return &campusGet{opts: opts}
}

func (r *campusGet) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r campusGet) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.CfgGetCampusNameStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	var campusGetKeyboard = tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Назад", fmt.Sprint(chainer.CfgCampusStep)),
		),
	)

	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, "Введите название нового Кампуса:", campusGetKeyboard)

	user.HandleStep = int(chainer.CfgSetCampusNameStep)
	err := r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
