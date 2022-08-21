package staffBooking

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type changeObj struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewChangeObj(opts *chainer.Opts) chainer.Chainer {
	return &changeObj{opts: opts}
}

func (r *changeObj) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r changeObj) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.StaffChangeObjectStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	CmdAndBookTypeAndCat := r.opts.Update.CallbackQuery.Data
	_, bookTypeAndCat, _ := strings.Cut(CmdAndBookTypeAndCat, "$")
	bookTypeAndCatSL := strings.Split(r.opts.Update.CallbackQuery.Data, "$")
	categoryID, _ := strconv.Atoi(bookTypeAndCatSL[2])

	var msgReply tg.EditMessageTextConfig
	if bookTypeAndCatSL[1] == string(models.PlacesType) {
		places, err := r.opts.RootRepo.GetAllPlacesByCampusIDAndCategoryID(ctx, *user.CampusID, categoryID)
		if err != nil {
			return nil, err
		}
		if len(places) == 0 {
			return tg.NewMessage(chatID, "Нет категорий"), nil
		}

		rows := make([][]tg.InlineKeyboardButton, 0, len(places))
		for _, place := range places {
			data := fmt.Sprintf("%d$%s$%d", chainer.StaffChangeDateStep, bookTypeAndCat, place.ID)
			row := tg.NewInlineKeyboardRow(
				tg.NewInlineKeyboardButtonData(place.Name, data),
			)
			rows = append(rows, row)
		}
		msgReply = tg.NewEditMessageTextAndMarkup(chatID, msgID, "Выбери помещение:", tg.NewInlineKeyboardMarkup(rows...))
	} else {
		inventories, err := r.opts.RootRepo.GetAllInventoryByCampusIDAndCategoryID(ctx, *user.CampusID, categoryID)
		if err != nil {
			return nil, err
		}
		if len(inventories) == 0 {
			return tg.NewMessage(chatID, "Нет категорий"), nil
		}

		rows := make([][]tg.InlineKeyboardButton, 0, len(inventories))
		for _, inventory := range inventories {
			data := fmt.Sprintf("%s$%d", bookTypeAndCat, inventory.ID)
			row := tg.NewInlineKeyboardRow(
				tg.NewInlineKeyboardButtonData(inventory.Name, data),
			)
			rows = append(rows, row)
		}
		msgReply = tg.NewEditMessageTextAndMarkup(chatID, msgID, "Выбери инвентарь:", tg.NewInlineKeyboardMarkup(rows...))
	}

	//user.HandleStep = int(chainer.StaffChangeDateStep)
	//err := r.opts.UserRepo.Update(ctx, user)
	//if err != nil {
	//	return nil, err
	//}

	return &msgReply, nil
}
