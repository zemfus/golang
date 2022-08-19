package register

import (
	"context"
	"strconv"
	"strings"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type checkCode struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewCheckCode(opts *chainer.Opts) chainer.Chainer {
	return &checkCode{opts: opts}
}

func (r *checkCode) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

func (r checkCode) Handle(ctx context.Context, user *models.User) (*tg.MessageConfig, error) {
	if int(chainer.StartCheckConfirmCodeStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	var msgReply tg.MessageConfig

	code, err := strconv.Atoi(r.opts.Update.Message.Text)
	if err != nil {
		msgReply.Text = "Не правильный код."
		return &msgReply, nil
	}

	ok, err := r.opts.SessionRepo.ExistsByCodeAndUserID(ctx, user.ID, code)
	if err != nil {
		return nil, err
	}
	if !ok {
		msgReply.Text = "Не правильный код."
		return &msgReply, nil
	}

	if strings.HasSuffix(user.Email, "@21-school.ru") {
		user.Role = models.Staff
	} else {
		user.Role = models.Student
	}

	msgReply.Text = "Выбери кампус:"
	user.HandleStep = int(chainer.StartChangeCampusStep)

	//todo : getall campuses

	//err := r.opts.UserRepo.Create(ctx, user)
	//if err != nil {
	//	return nil, err
	//}

	return &msgReply, nil
}
