package staffBooking

import (
	"context"
	"fmt"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type changeType struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewChangeType(opts *chainer.Opts) chainer.Chainer {
	return &changeType{opts: opts}
}

func (r *changeType) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r changeType) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.StaffChangeTypeStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	var keyboardBookingType = tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Помещение", fmt.Sprintf("%d$%s", chainer.StaffChangeCategoryStep, models.PlacesType)),
			tg.NewInlineKeyboardButtonData("Инвентарь", fmt.Sprintf("%d$%s", chainer.StaffChangeCategoryStep, models.InventoryType)),
		),
	)

	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, "Выбери тип:", keyboardBookingType)

	//user.HandleStep = int(chainer.StaffChangeCategoryStep)
	//err := r.opts.UserRepo.Update(ctx, user)
	//if err != nil {
	//	return nil, err
	//}

	return &msgReply, nil
}
