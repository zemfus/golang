package controller

import (
	"context"

	"boobot/kernel/domain/models"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Opts struct {
	User   *models.User
	Update *tg.Update
}

// Controller проверяет права на команды и нас основе этого перенаправляет в нужный сервис
type Controller interface {
	Process(ctx context.Context, opts *Opts) error
}
