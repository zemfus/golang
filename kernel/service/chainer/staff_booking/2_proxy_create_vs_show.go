package staffBooking

import (
	"context"
	"strconv"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type proxyCreateVSShow struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewProxyCreateVSShow(opts *chainer.Opts) chainer.Chainer {
	return &proxyCreateVSShow{opts: opts}
}

func (r *proxyCreateVSShow) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r proxyCreateVSShow) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.StaffProxyCreateVSShow) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	step, _ := strconv.Atoi(r.opts.Update.CallbackQuery.Data)
	var ch chainer.Chainer

	if int(chainer.StaffChangeTypeStep) == step {
		ch = NewChangeType(r.opts)
		user.HandleStep = int(chainer.StaffChangeTypeStep)
	} else {
		//todo prosmotr
	}
	return ch.Handle(ctx, user)
}
