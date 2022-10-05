package configurations

import (
	"context"
	"fmt"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type campusEdit struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewCampusEdit(opts *chainer.Opts) chainer.Chainer {
	return &campusEdit{opts: opts}
}

func (r *campusEdit) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r campusEdit) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.CfgCampusEditStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	var text string
	campuses, err := r.opts.RootRepo.GetAllCampuses(ctx)
	if err != nil {
		return nil, err
	}
	if len(campuses) == 0 {
		text = "Список кампусов пуст"
	} else {
		text = "Редактирование Кампуса:\n ✏️ изменить имя\n❌ удалить кампус"
	}
	rows := make([][]tg.InlineKeyboardButton, 0, len(campuses)+1)
	for _, campus := range campuses {
		var row = tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(campus.Name, fmt.Sprint(chainer.NonStep)),
			tg.NewInlineKeyboardButtonData("✏️", fmt.Sprint(chainer.CfgCampusUpdateStep, "$", campus.ID)),
			tg.NewInlineKeyboardButtonData("❌", fmt.Sprint(chainer.CfgCampusDeleteStep, "$", campus.ID)),
		)
		rows = append(rows, row)
	}
	rows = append(rows, tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Назад", fmt.Sprint(chainer.CfgCampusStep))))

	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, text, tg.NewInlineKeyboardMarkup(rows...))

	return &msgReply, nil
}
