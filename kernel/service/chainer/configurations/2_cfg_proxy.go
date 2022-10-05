package configurations

import (
	"context"
	"strconv"
	"strings"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type cfgProxy struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewCfgProxy(opts *chainer.Opts) chainer.Chainer {
	return &cfgProxy{opts: opts}
}

func (r *cfgProxy) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r cfgProxy) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.CfgProxyItemsStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	proxyAndCmd := strings.Split(r.opts.Update.CallbackQuery.Data, "$")
	step, _ := strconv.Atoi(proxyAndCmd[1])

	var ch chainer.Chainer
	switch step {
	case int(chainer.CfgCampusStep):
		ch = NewCampus(r.opts)
		user.HandleStep = int(chainer.CfgCampusStep)
	case int(chainer.CfgCategoryStep):
		ch = NewCategory(r.opts)
		user.HandleStep = int(chainer.CfgCategoryStep)
	case int(chainer.CfgInventoryStep):
		ch = NewInventory(r.opts)
		user.HandleStep = int(chainer.CfgInventoryStep)
	case int(chainer.CfgPlaceStep):
		ch = NewPlace(r.opts)
		user.HandleStep = int(chainer.CfgPlaceStep)
	case int(chainer.CfgStudentsStep):
		ch = NewStudents(r.opts)
		user.HandleStep = int(chainer.CfgStudentsStep)
	}

	return ch.Handle(ctx, user)
}
