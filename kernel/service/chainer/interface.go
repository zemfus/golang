package chainer

import (
	"context"

	"boobot/dal/repo"
	"boobot/kernel/domain/models"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Opts struct {
	UserRepo    repo.User
	Update      *tg.Update
	SessionRepo repo.Session
	RootRepo    repo.Root
}

type Chainer interface {
	SetNext(Chainer) Chainer
	Handle(context.Context, *models.User) (tg.Chattable, error)
}
