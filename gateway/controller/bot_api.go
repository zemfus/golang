package controller

import (
	"context"
	"log"

	"boobot/dal/repo"
	"boobot/kernel/domain/cmd"
	"boobot/kernel/service"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

type botAPI struct {
	bot          *tg.BotAPI
	connPool     *pgxpool.Pool
	services     map[string]service.NewServiceFunc
	servicesPost map[chainer.StepHandle]service.NewServiceFunc
}

func NewBotAPI(bot *tg.BotAPI, connPool *pgxpool.Pool) Controller {
	srvs := map[string]service.NewServiceFunc{
		cmd.Start: service.NewStart,
	}
	srvsPost := map[chainer.StepHandle]service.NewServiceFunc{
		chainer.StartSendConfirmURLStep: service.NewStart,
	}
	return &botAPI{
		bot:          bot,
		connPool:     connPool,
		services:     srvs,
		servicesPost: srvsPost,
	}
}

func (c botAPI) Process(ctx context.Context, opts *Opts) error {
	// check exists serv
	var text string
	var userID int
	if opts.Update.Message != nil {
		userID = int(opts.Update.Message.From.ID)
		text = opts.Update.Message.Text
	} else {
		userID = int(opts.Update.CallbackQuery.Message.From.ID)
		text = opts.Update.CallbackQuery.Message.Text
	}

	userRepo := repo.NewUser(c.connPool)
	sessionRepo := repo.NewSession(c.connPool)
	//TODO:: CHECK USER
	user, err := userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	//TODO:: check cmd
	//TODO:: check auth

	srv, ok := c.services[text]
	if !ok {
		srv, ok = c.servicesPost[chainer.StepHandle(user.HandleStep)]
		if !ok {
			_, err = c.bot.Send(tg.NewMessage(int64(userID), "Ты не выбрал команду."))
			if err != nil {
				log.Println(err)
			}
			return nil
		}
	}

	s, err := srv(&service.Opts{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
		Update:      opts.Update,
	})
	if err != nil {
		return err
	}

	msgReply, err := s.Execute(ctx, user)
	if err != nil {
		return err
	}

	// todo retry
	_, err = c.bot.Send(msgReply)
	if err != nil {
		log.Println(err)
	}

	return nil
}
