package register

import (
	"context"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type reqEmail struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewReqEmail(opts *chainer.Opts) chainer.Chainer {
	return &reqEmail{opts: opts}
}

func (r *reqEmail) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r reqEmail) Handle(ctx context.Context, user *models.User) (*tg.MessageConfig, error) {
	if int(chainer.StartRequestEmailStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	var msgReply tg.MessageConfig
	msgReply.Text = "Введи свою школьную почту, а затем пройди по ссылке в письме для авторизации."

	user.HandleStep = int(chainer.StartSendConfirmCodeStep)
	user.ID = int(r.opts.Update.Message.From.ID)
	user.Role = models.Unknown

	err := r.opts.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
