package configurations

import (
	"context"
	"fmt"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type campusProxy struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewCampusProxy(opts *chainer.Opts) chainer.Chainer {
	return &campusProxy{opts: opts}
}

func (r *campusProxy) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r campusProxy) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.CfgCampusStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	var campusProxyKeyboard = tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Создать ✏️",
				fmt.Sprint(chainer.CfgCampusCreateStep)),

			tg.NewInlineKeyboardButtonData("Редактировать ⚙️",
				fmt.Sprint(chainer.CfgCampusEditStep)),
		),

		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Назад", fmt.Sprint(chainer.CfgShowBtnStep)),
		),
	)

	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, "Конфигурация Кампуса:", campusProxyKeyboard)

	return &msgReply, nil
}
