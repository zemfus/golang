package configurations

import (
	"context"
	"strconv"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type campusUpdateExec struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewCampusUpdateExec(opts *chainer.Opts) chainer.Chainer {
	return &campusUpdateExec{opts: opts}
}

func (r *campusUpdateExec) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r campusUpdateExec) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.CfgCampusUpdateExecStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	campusID, _ := strconv.Atoi(user.LastMsg)
	campus := &models.Campus{
		ID:   campusID,
		Name: r.opts.Update.Message.Text,
	}

	err := r.opts.RootRepo.UpdateCampus(ctx, campus)
	if err != nil {
		return nil, err
	}
	var msgReply tg.MessageConfig
	msgReply.Text = "Кампус успешно переименован."

	user.HandleStep = int(chainer.NonStep)
	user.LastMsg = ""
	err = r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
