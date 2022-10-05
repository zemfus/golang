package configurations

import (
	"context"
	"strconv"
	"strings"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type campusDelete struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewCampusDelete(opts *chainer.Opts) chainer.Chainer {
	return &campusDelete{opts: opts}
}

func (r *campusDelete) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r campusDelete) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.CfgCampusDeleteStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	text := strings.Split(r.opts.Update.CallbackQuery.Data, "$")
	campusID, err := strconv.Atoi(text[1])
	if err != nil {
		return nil, err
	}

	ok, err := r.opts.UserRepo.ExistsUsersInCampusByID(ctx, campusID)
	if err != nil {
		return nil, err
	}

	var msgText string
	if ok {
		msgText = "В данном кампусе есть активные студенты, пожалуйста, переведите сначала студентов в другой кампус или удалите, и повторите заново."
	} else {
		err = r.opts.RootRepo.DeleteCampus(ctx, campusID)
		if err != nil {
			return nil, err
		}
		msgText = "Кампус успешно удален!"
	}

	callback := tg.NewCallbackWithAlert(r.opts.Update.CallbackQuery.ID, msgText)
	_, _ = r.opts.Bot.Request(callback)

	user.HandleStep = int(chainer.CfgCampusEditStep)
	return NewCampusEdit(r.opts).Handle(ctx, user)

}
