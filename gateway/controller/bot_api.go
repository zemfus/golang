package controller

import (
	"context"
	"log"
	"strconv"
	"strings"

	"boobot/dal/repo"
	"boobot/kernel/domain/btn"
	"boobot/kernel/service"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

type botAPI struct {
	bot      *tg.BotAPI
	connPool *pgxpool.Pool
	services map[string]service.NewServiceFunc
}

func NewBotAPI(bot *tg.BotAPI, connPool *pgxpool.Pool) Controller {
	return &botAPI{
		bot:      bot,
		connPool: connPool,
		services: registerServices(),
	}
}

func (c botAPI) Process(ctx context.Context, opts *Opts) error {
	// check exists serv
	var text string
	var userID int

	userRepo := repo.NewUser(c.connPool)
	sessionRepo := repo.NewSession(c.connPool)
	rootRepo := repo.NewRoot(c.connPool)
	//TODO:: CHECK USER
	user, err := userRepo.GetByID(ctx, getUserID(opts.Update))
	if err != nil {
		return err
	}

	if opts.Update.Message != nil {
		userID = int(opts.Update.Message.From.ID)
		if btn.AllCmds[opts.Update.Message.Text] {
			text = opts.Update.Message.Text
		} else {
			text = strconv.Itoa(user.HandleStep)
		}
	} else {
		userID = int(opts.Update.CallbackQuery.From.ID)
		text = strings.Split(opts.Update.CallbackQuery.Data, "$")[0]
		if text == "0" {
			return nil
		}
	}
	//TODO:: check cmd
	//TODO:: check auth

	if opts.Update.CallbackQuery != nil {
		if handleStep, err := strconv.Atoi(text); err == nil {
			user.HandleStep = handleStep
		} else {
			_, err = c.bot.Send(tg.NewMessage(int64(userID), "Что-то пошло не так, разрабы устали."))
			if err != nil {
				log.Println(err)
			}
		}
	}

	srv, ok := c.services[text]
	if !ok {
		_, err = c.bot.Send(tg.NewMessage(int64(userID), "Ку-ку, не понял, что ты хочешь, выбери команду."))
		if err != nil {
			log.Println(err)
		}
		return nil
	}

	s, err := srv(&service.Opts{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
		Update:      opts.Update,
		RootRepo:    rootRepo,
	})
	if err != nil {
		return err
	}

	msgReply, err := s.Execute(ctx, user)
	if err != nil {
		return err
	}

	if opts.Update.Message != nil {
		switch ms := msgReply.(type) {
		case *tg.MessageConfig:
			ms.ChatID = opts.Update.Message.From.ID
		}
	} else {
		switch ms := msgReply.(type) {
		case *tg.MessageConfig:
			ms.ChatID = opts.Update.CallbackQuery.From.ID
		}
	}

	// todo retry
	_, err = c.bot.Send(msgReply)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func getUserID(update *tg.Update) int {
	if update.Message != nil {
		return int(update.Message.From.ID)
	}
	return int(update.CallbackQuery.From.ID)
}

func registerServices() map[string]service.NewServiceFunc {
	return map[string]service.NewServiceFunc{
		btn.Start:   service.NewStart,
		btn.Booking: service.NewStaffBooking,

		strconv.Itoa(int(chainer.StartSendConfirmCodeStep)):  service.NewStart,
		strconv.Itoa(int(chainer.StartCheckConfirmCodeStep)): service.NewStart,
		strconv.Itoa(int(chainer.StartChangeCampusStep)):     service.NewStart,
		strconv.Itoa(int(chainer.StartSetCampusStep)):        service.NewStart,

		strconv.Itoa(int(chainer.StaffShowBtnBookingsStep)): service.NewStaffBooking,
		strconv.Itoa(int(chainer.StaffProxyCreateVSShow)):   service.NewStaffBooking,
		strconv.Itoa(int(chainer.StaffChangeTypeStep)):      service.NewStaffBooking,
		strconv.Itoa(int(chainer.StaffChangeCategoryStep)):  service.NewStaffBooking,
		strconv.Itoa(int(chainer.StaffChangeObjectStep)):    service.NewStaffBooking,
		strconv.Itoa(int(chainer.StaffChangeDateStep)):      service.NewStaffBooking,
		strconv.Itoa(int(chainer.StaffChangeTimeStep)):      service.NewStaffBooking,
	}
}
