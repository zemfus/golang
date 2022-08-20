package app

import (
	"context"
	"flag"
	"log"
	"os"

	"boobot/gateway/controller"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var (
	config = flag.String("config", "./configs/.env", "configurations for kind bot")
)

const (
	botToken = "BOT_TOKEN"
	dbURL    = "DB_URL"
)

func init() {
	flag.Parse()

	println(*config)

	if err := godotenv.Load(*config); err != nil {
		log.Fatalf("No .env file found: %s", *config)
	}
}

type app struct {
	bot      *tg.BotAPI
	connPool *pgxpool.Pool
	ctrl     controller.Controller
}

func New(ctx context.Context) *app {
	pool, err := pgxpool.Connect(ctx, os.Getenv(dbURL))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// init bot api
	bot, err := tg.NewBotAPI(os.Getenv(botToken))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	ctrl := controller.NewBotAPI(bot, pool)

	return &app{
		connPool: pool,
		bot:      bot,
		ctrl:     ctrl,
	}
}

func (a *app) Run(ctx context.Context) {
	u := tg.NewUpdate(0)
	u.Timeout = 60
	updates := a.bot.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			a.Stop(ctx)
			return
		case update := <-updates:
			opts := &controller.Opts{
				Update: &update,
			}
			err := a.ctrl.Process(ctx, opts)
			if err != nil {
				log.Println(err, update)
			}
		}
	}
}

func (a *app) Stop(_ context.Context) {
	a.connPool.Close()
}
