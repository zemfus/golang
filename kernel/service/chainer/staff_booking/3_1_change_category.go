package staffBooking

import (
	"context"
	"fmt"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type changeCategory struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewChangeCategory(opts *chainer.Opts) chainer.Chainer {
	return &changeCategory{opts: opts}
}

func (r *changeCategory) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r changeCategory) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.StaffChangeCategoryStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	bookType := r.opts.Update.CallbackQuery.Data

	categories, err := r.opts.RootRepo.GetAllCategory(ctx)
	if err != nil {
		return nil, err
	}
	if len(categories) == 0 {
		return tg.NewMessage(chatID, "Нет категорий"), nil
	}

	rows := make([][]tg.InlineKeyboardButton, 0, len(categories))
	for _, category := range categories {
		data := fmt.Sprintf("%s$%d", bookType, category.ID)
		row := tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(category.Name, data),
		)
		rows = append(rows, row)
	}
	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, "Выбери Категорию:", tg.NewInlineKeyboardMarkup(rows...))

	user.HandleStep = int(chainer.StaffChangeObjectStep)
	err = r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
