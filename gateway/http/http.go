package http

import (
	context2 "context"
	"fmt"

	"boobot/kernel/domain/models"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// todo:: add sql injection validation

func Session(ctx *routing.Context, opts *Opts, userCh chan *models.User) routing.Handler {
	return func(context *routing.Context) error {
		uuid := ctx.Param("<tag>")
		if len(uuid) == 0 {
			ctx.WriteString("Bad Request")
			ctx.SetStatusCode(fasthttp.StatusBadRequest)

			return nil
		}
		id, err := ctx.QueryArgs().GetUint("id")
		if err != nil {
			ctx.WriteString("Bad Request")
			ctx.SetStatusCode(fasthttp.StatusBadRequest)

			return nil
		}

		ctx2 := context2.Background()
		session, err := opts.SessionRepo.GetByID(ctx2, id)
		// todo add logic
		session.ID = 3 // for linter, delete me

		user, err := opts.UserRepo.GetByID(ctx2, id)
		// todo add logic

		userCh <- user

		fmt.Fprintf(ctx, "Вы успешно подтвердили почту, скоро вам придет сообщения в телеграм с дальнейшими инстурциями.")
		return nil
	}
}
