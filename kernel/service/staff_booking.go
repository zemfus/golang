package service

import (
	"context"

	"boobot/kernel/domain/btn"
	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	staffBkg "boobot/kernel/service/chainer/staff_booking"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type staffBooking struct {
	opts  *Opts
	chain chainer.Chainer
}

func NewStaffBooking(opts *Opts) (Service, error) {
	//todo: check opts
	return &staffBooking{
		opts: opts,
	}, nil
}

func (s staffBooking) Execute(ctx context.Context, user *models.User) (tg.Chattable, error) {
	user.HandleStep = chainer.CheckStepHandle(user.HandleStep, chainer.StaffShowBtnBookingsStep,
		chainer.StaffBookingSteps...)

	if s.opts.Update.Message != nil && s.opts.Update.Message.Text == btn.Booking {
		user.HandleStep = int(chainer.StaffShowBtnBookingsStep)
	}

	opts := &chainer.Opts{
		UserRepo:    s.opts.UserRepo,
		Update:      s.opts.Update,
		SessionRepo: s.opts.SessionRepo,
		RootRepo:    s.opts.RootRepo,
	}

	chain := staffBkg.NewShowBtn(opts)
	chain.SetNext(staffBkg.NewProxyCreateVSShow(opts)).
		SetNext(staffBkg.NewChangeType(opts)).
		SetNext(staffBkg.NewChangeCategory(opts)).
		SetNext(staffBkg.NewChangeObj(opts)).
		SetNext(staffBkg.NewChangeDate(opts)).
		SetNext(staffBkg.NewChangeTime(opts))

	msgReply, err := chain.Handle(ctx, user)
	if err != nil {
		return nil, err
	}

	return msgReply, nil
}
