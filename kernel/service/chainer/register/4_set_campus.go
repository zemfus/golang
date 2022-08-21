package register

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"boobot/kernel/domain/btn"
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

func (r setCampus) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.StartSetCampusStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	userNickname := strings.Split(user.Email, "@")[0]

	var msgReply tg.MessageConfig
	if user.Role == models.Staff || user.ID == 234899515 {
		msgReply.Text = fmt.Sprintf("%s, ты успешно зарегистрировался, теперь тебе доступны все функциональности для стафф.", userNickname)
		msgReply.ReplyMarkup = btn.Staff
	} else {
		msgReply.Text = fmt.Sprintf("%s, ты успешно зарегистрировался, теперь тебе доступны все функциональности для студента.", userNickname)
		msgReply.ReplyMarkup = btn.Student
	}

	stepAndCampus := strings.Split(r.opts.Update.CallbackQuery.Data, "$")
	campusID, _ := strconv.Atoi(stepAndCampus[1])

	user.Nickname = userNickname
	user.CampusID = &campusID
	user.HandleStep = int(chainer.NonStep)

	err := r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
