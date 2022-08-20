package register

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/gomail.v2"
)

type sendConfirmURL struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewSendConfirmURL(opts *chainer.Opts) chainer.Chainer {
	return &sendConfirmURL{opts: opts}
}

func (r *sendConfirmURL) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

var validEmail = regexp.MustCompile(`^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@((21-school|student.21-school).ru)$`)

func (r sendConfirmURL) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.StartSendConfirmCodeStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}
	userEmail := r.opts.Update.Message.Text
	var msgReply tg.MessageConfig
	if validEmail.Match([]byte(userEmail)) {
		msgReply.Text = "На твою почту в течении 5 минут придет письмо с кодом подтверждения, для продолжения регистрации введи код сюда."
	} else {
		msgReply.Text = "Невалидный адрес почты."
		return &msgReply, nil
	}
	rand.Seed(time.Now().UnixNano())
	userCode := rand.Intn(9999-1000) + 1000

	session := &models.Session{
		UserID: user.ID,
		Code:   userCode,
	}

	err := r.opts.SessionRepo.Create(ctx, session)
	if err != nil {
		return nil, err
	}

	var confirmText = fmt.Sprintf(`Привет!
Твой код подтверждения:<br>
<code>%d</code>`, userCode)

	user.HandleStep = int(chainer.StartCheckConfirmCodeStep)

	err = sendMail(confirmText, userEmail)
	if err != nil {
		return nil, err
	}

	user.Email = userEmail
	err = r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}

const (
	mailUser     = "EMAIL"
	mailPassword = "EMAIL_PASSWORD"
)

func sendMail(body string, to ...string) error {
	body += "<p>С уважением, BookBot!"
	m := gomail.NewMessage()
	m.SetAddressHeader("From", os.Getenv(mailUser), "Book Bot")
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "Авторизация.")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.yandex.ru", 465, os.Getenv(mailUser), os.Getenv(mailPassword))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
