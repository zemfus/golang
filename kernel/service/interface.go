package service

import (
	"context"

	"boobot/dal/repo"
	"boobot/kernel/domain/models"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Opts struct {
	UserRepo    repo.User
	SessionRepo repo.Session
	Update      *tg.Update
}

type NewServiceFunc func(opts *Opts) (Service, error)
type Func func(ctx context.Context, user *models.User) (*tg.MessageConfig, error)

type Service interface {
	Execute(ctx context.Context, user *models.User) (*tg.MessageConfig, error)
}
