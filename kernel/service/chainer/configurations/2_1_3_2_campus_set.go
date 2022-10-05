package configurations

import (
	"context"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type campusSet struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewCampusSet(opts *chainer.Opts) chainer.Chainer {
	return &campusSet{opts: opts}
}

func (r *campusSet) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r campusSet) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.CfgSetCampusNameStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	campus := &models.Campus{
		Name: r.opts.Update.Message.Text,
	}

	err := r.opts.RootRepo.CreateCampus(ctx, campus)
	if err != nil {
		return nil, err
	}
	var msgReply tg.MessageConfig
	msgReply.Text = "Кампус успешно создан."

	user.HandleStep = int(chainer.NonStep)
	err = r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
