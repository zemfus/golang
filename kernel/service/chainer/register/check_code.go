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
	campuses, err := r.opts.RootRepo.GetAllCampuses(ctx)
	if err != nil {
		return nil, err
	}

	if len(campuses) == 0 {
		msgReply.Text = "Не определен кампус, обращайся к администраторам своего кампуса."
		return &msgReply, nil
	}

	rowsCampuses := make([][]tg.InlineKeyboardButton, 0, len(campuses))
	for _, campus := range campuses {
		row := tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(campus.Name, strconv.Itoa(campus.ID)),
		)
		rowsCampuses = append(rowsCampuses, row)
	}
	msgReply.ReplyMarkup = tg.NewInlineKeyboardMarkup(rowsCampuses...)

	user.Nickname = strings.Split(user.Email, "@")[0]
	user.HandleStep = int(chainer.StartSetCampusStep)

	err = r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
