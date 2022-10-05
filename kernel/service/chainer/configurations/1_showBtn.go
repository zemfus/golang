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
	if int(chainer.CfgShowBtnStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	text := "Конфигурация:"
	var cfgItemsKeyboard = tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Кампусы 🏢",
				fmt.Sprint(chainer.CfgProxyItemsStep, "$", chainer.CfgCampusStep)),

			tg.NewInlineKeyboardButtonData("Категории 🗃",
				fmt.Sprint(chainer.CfgProxyItemsStep, "$", chainer.CfgCategoryStep)),
		),

		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Помещения 🚪",
				fmt.Sprint(chainer.CfgProxyItemsStep, "$", chainer.CfgPlaceStep)),

			tg.NewInlineKeyboardButtonData("Инвентарь 🛒",
				fmt.Sprint(chainer.CfgProxyItemsStep, "$", chainer.CfgInventoryStep)),
		),

		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Студенты 👩‍🎓", fmt.Sprint(chainer.CfgProxyItemsStep, "$", chainer.CfgStudentsStep)),
		),
	)

	if r.opts.Update.Message == nil {
		chatID := r.opts.Update.CallbackQuery.From.ID
		msgID := r.opts.Update.CallbackQuery.Message.MessageID
		msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, text, cfgItemsKeyboard)
		return msgReply, nil
	}

	var msgReply tg.MessageConfig
	msgReply.Text = text
	msgReply.ReplyMarkup = cfgItemsKeyboard

	return &msgReply, nil
}
