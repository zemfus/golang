package service

import (
	"context"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	"boobot/kernel/service/chainer/register"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type start struct {
	opts  *Opts
	chain chainer.Chainer
}

func NewStart(opts *Opts) (Service, error) {
	//todo: check opts
	return &start{
		opts: opts,
	}, nil
}

func (s start) Execute(ctx context.Context, user *models.User) (*tg.MessageConfig, error) {
	user.HandleStep = chainer.CheckStepHandle(user.HandleStep, chainer.StartRequestEmailStep, // todo change StartRequestEmailStep
		chainer.StartSteps...)

	if s.opts.Update.Message != nil && s.opts.Update.Message.Text == string(models.Start) {
		user.HandleStep = int(chainer.StartRequestEmailStep)
	}

	opts := &chainer.Opts{
		UserRepo:    s.opts.UserRepo,
		Update:      s.opts.Update,
		SessionRepo: s.opts.SessionRepo,
		RootRepo:    s.opts.RootRepo,
	}

	chain := register.NewReqEmail(opts)
	chain.SetNext(register.NewSendConfirmURL(opts)).
		SetNext(register.NewCheckCode(opts)).
		SetNext(register.NewSetCampus(opts))

	msgReply, err := chain.Handle(ctx, user)
	if err != nil {
		return nil, err
	}
	//todo if callback

	if s.opts.Update.Message != nil {
		msgReply.ChatID = s.opts.Update.Message.From.ID
	} else {
		msgReply.ChatID = s.opts.Update.CallbackQuery.From.ID
	}

	return msgReply, nil
}
