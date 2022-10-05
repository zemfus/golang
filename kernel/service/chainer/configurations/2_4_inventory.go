package configurations

import (
	"context"
	"fmt"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type inventory struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewInventory(opts *chainer.Opts) chainer.Chainer {
	return &inventory{opts: opts}
}

func (r *inventory) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r inventory) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.CfgInventoryStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	var categoryKeyboard = tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Создать ✏️",
				fmt.Sprint(chainer.CfgCampusStep, "$", chainer.CfgProxyItemsStep)),

			tg.NewInlineKeyboardButtonData("Редактировать ⚙️",
				fmt.Sprint(chainer.CfgCategoryStep, "$", chainer.CfgProxyItemsStep)),
		),

		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Назад", fmt.Sprint(chainer.CfgShowBtnStep)),
		),
	)
	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, "Конфигурация Инвентарь:", categoryKeyboard)

	return &msgReply, nil
}
