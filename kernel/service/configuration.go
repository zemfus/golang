package service

import (
	"context"

	"boobot/kernel/domain/btn"
	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	"boobot/kernel/service/chainer/configurations"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type configuration struct {
	opts  *Opts
	chain chainer.Chainer
}

func NewConfiguration(opts *Opts) (Service, error) {
	//todo: check opts
	return &configuration{
		opts: opts,
	}, nil
}

func (s configuration) Execute(ctx context.Context, user *models.User) (tg.Chattable, error) {
	user.HandleStep = chainer.CheckStepHandle(user.HandleStep, chainer.CfgShowBtnStep,
		chainer.CfgSteps...)

	if s.opts.Update.Message != nil && s.opts.Update.Message.Text == btn.Configuration {
		user.HandleStep = int(chainer.CfgShowBtnStep)
	}

	opts := &chainer.Opts{
		UserRepo:    s.opts.UserRepo,
		Update:      s.opts.Update,
		SessionRepo: s.opts.SessionRepo,
		RootRepo:    s.opts.RootRepo,
		Bot:         s.opts.Bot,
	}

	chain := configurations.NewShowBtn(opts)
	chain.SetNext(configurations.NewCfgProxy(opts)).
		SetNext(configurations.NewCampus(opts)).
		SetNext(configurations.NewInventory(opts)).
		SetNext(configurations.NewPlace(opts)).
		SetNext(configurations.NewCategory(opts)).
		SetNext(configurations.NewStudents(opts)).
		SetNext(configurations.NewCampusGet(opts)).
		SetNext(configurations.NewCampusSet(opts)).
		SetNext(configurations.NewCampusEdit(opts)).
		SetNext(configurations.NewCampusUpdate(opts)).
		SetNext(configurations.NewCampusUpdateExec(opts)).
		SetNext(configurations.NewCampusDelete(opts))

	msgReply, err := chain.Handle(ctx, user)
	if err != nil {
		return nil, err
	}

	return msgReply, nil
}
