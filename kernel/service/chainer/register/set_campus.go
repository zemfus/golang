package register

import (
	"context"
	"fmt"
	"strings"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type setCampus struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewSetCampus(opts *chainer.Opts) chainer.Chainer {
	return &setCampus{opts: opts}
}

func (r *setCampus) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r setCampus) Handle(ctx context.Context, user *models.User) (*tg.MessageConfig, error) {
	if int(chainer.StartSetCampusStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	userNickname := strings.Split(user.Email, "@")[0]

	var msgReply tg.MessageConfig
	msgReply.Text = fmt.Sprintf("%s, ты успешно зарегистрировался, теперь тебе доступны все функциональности бота.", userNickname)

	user.Nickname = userNickname
	user.HandleStep = int(chainer.NonStep)

	err := r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
