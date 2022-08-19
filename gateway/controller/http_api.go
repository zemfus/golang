package controller

import (
	"context"

	"boobot/kernel/service"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

type httpAPI struct {
	bot      *tg.BotAPI
	connPool *pgxpool.Pool
	services map[string]service.NewServiceFunc
}

func NewHTTPApi(bot *tg.BotAPI, connPool *pgxpool.Pool) Controller {
	//TODO:: init services
	return &httpAPI{
		bot:      bot,
		connPool: connPool,
	}
}

func (h httpAPI) Process(ctx context.Context, opts *Opts) error {

	return nil
}
