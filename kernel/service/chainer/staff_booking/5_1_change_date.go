package staffBooking

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"boobot/kernel/domain/models"
	"boobot/kernel/service/chainer"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type changeDate struct {
	opts *chainer.Opts
	next chainer.Chainer
}

func NewChangeDate(opts *chainer.Opts) chainer.Chainer {
	return &changeDate{opts: opts}
}

func (r *changeDate) SetNext(chainer chainer.Chainer) chainer.Chainer {
	r.next = chainer

	return chainer
}

var patternDate = "2/1/2006"

func (r changeDate) Handle(ctx context.Context, user *models.User) (tg.Chattable, error) {
	if int(chainer.StaffChangeDateStep) != user.HandleStep {
		return r.next.Handle(ctx, user)
	}

	chatID := r.opts.Update.CallbackQuery.From.ID
	msgID := r.opts.Update.CallbackQuery.Message.MessageID

	CmdAndBookTypeAndCatAndObj := r.opts.Update.CallbackQuery.Data
	CmdAndBookTypeAndCatAndObjSl := strings.Split(CmdAndBookTypeAndCatAndObj, "$")
	_, bookTypeAndCatAndObj, _ := strings.Cut(CmdAndBookTypeAndCatAndObj, "$")
	//bookTypeAndCatAndObjSL := strings.Split(r.opts.Update.CallbackQuery.Data, "$")
	//categoryID, _ := strconv.Atoi(bookTypeAndCatAndObjSL[1])
	//objID, _ := strconv.Atoi(bookTypeAndCatAndObjSL[2])

	now := time.Now()

	if strings.HasSuffix(bookTypeAndCatAndObj, "$next") || strings.HasSuffix(bookTypeAndCatAndObj, "$prev") {
		parsStr := CmdAndBookTypeAndCatAndObjSl[len(CmdAndBookTypeAndCatAndObjSl)-2]
		now, _ = time.Parse(time.ANSIC, parsStr)
		println(now.String())
		if now.Year() == time.Now().Year() && now.Month() == time.Now().Month() {
			now = time.Now()
		}
	}

	rows := make([][]tg.InlineKeyboardButton, 0, 6)
	rows = append(rows, tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(now.Month().String(), "0")))
	// for navigation
	rows = append(rows, tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(now.Month().String(), "0")))
	//rows = append(rows, row)

ro:
	for i := 0; i < 6; i++ {
		row := make([]tg.InlineKeyboardButton, 0, 7)
		for j := 0; j < 7; j++ {
			switch now.Weekday() {
			case time.Monday:
				if j != 0 {
					row = append(row, tg.NewInlineKeyboardButtonData(" ", "0"))
					continue
				}
			case time.Tuesday:
				if j != 1 {
					row = append(row, tg.NewInlineKeyboardButtonData(" ", "0"))
					continue
				}
			case time.Wednesday:
				if j != 2 {
					row = append(row, tg.NewInlineKeyboardButtonData(" ", "0"))
					continue
				}
			case time.Thursday:
				if j != 3 {
					row = append(row, tg.NewInlineKeyboardButtonData(" ", "0"))
					continue
				}
			case time.Friday:
				if j != 4 {
					row = append(row, tg.NewInlineKeyboardButtonData(" ", "0"))
					continue
				}
			case time.Saturday:
				if j != 5 {
					row = append(row, tg.NewInlineKeyboardButtonData(" ", "0"))
					continue
				}
			case time.Sunday:
				if j != 6 {
					row = append(row, tg.NewInlineKeyboardButtonData(" ", "0"))
					continue
				}
			}
			row = append(row, tg.NewInlineKeyboardButtonData(strconv.Itoa(now.Day()), fmt.Sprintf("%d%s$%s",
				chainer.StaffChangeTimeStep,
				bookTypeAndCatAndObj,
				now.Format(patternDate))),
			)
			now = now.Add(time.Hour * 24)

			if now.Day() == 1 {
				prev := now
				for now.Weekday() != time.Monday {
					row = append(row, tg.NewInlineKeyboardButtonData(" ", "0"))
					now = now.Add(time.Hour * 24)
				}
				rows = append(rows, row)
				now = prev
				break ro
			}
		}

		rows = append(rows, row)
		if now.Day() == 1 {
			break
		}
	}
	if now.Day() != 1 {
		now = now.Add(time.Hour * 24)
	}
	rowNav := make([]tg.InlineKeyboardButton, 0, 2)
	prevTime := now.AddDate(0, -2, 0)
	if strings.HasSuffix(bookTypeAndCatAndObj, "$next") {
		CmdAndBookTypeAndCatAndObj = strings.Join(CmdAndBookTypeAndCatAndObjSl[:len(CmdAndBookTypeAndCatAndObjSl)-2], "$")
		rowNav = append(rowNav,
			tg.NewInlineKeyboardButtonData("prev", fmt.Sprintf("%s$%s$prev", CmdAndBookTypeAndCatAndObj, prevTime.Format(time.ANSIC))),
			tg.NewInlineKeyboardButtonData("next", fmt.Sprintf("%s$%s$next", CmdAndBookTypeAndCatAndObj, now.Format(time.ANSIC))),
		)
	} else if strings.HasSuffix(bookTypeAndCatAndObj, "$prev") && now.AddDate(0, -1, 0).After(time.Now()) {
		CmdAndBookTypeAndCatAndObj = strings.Join(CmdAndBookTypeAndCatAndObjSl[:len(CmdAndBookTypeAndCatAndObjSl)-2], "$")
		rowNav = append(rowNav,
			tg.NewInlineKeyboardButtonData("prev", fmt.Sprintf("%s$%s$prev", CmdAndBookTypeAndCatAndObj, prevTime.Format(time.ANSIC))),
			tg.NewInlineKeyboardButtonData("next", fmt.Sprintf("%s$%s$next", CmdAndBookTypeAndCatAndObj, now.Format(time.ANSIC))),
		)
	} else {
		if strings.HasSuffix(bookTypeAndCatAndObj, "$prev") {
			CmdAndBookTypeAndCatAndObj = strings.Join(CmdAndBookTypeAndCatAndObjSl[:len(CmdAndBookTypeAndCatAndObjSl)-2], "$")
		}
		rowNav = append(rowNav,
			tg.NewInlineKeyboardButtonData("next", fmt.Sprintf("%s$%s$next", CmdAndBookTypeAndCatAndObj, now.Format(time.ANSIC))),
		)
	}
	rows[1] = rowNav

	msgReply := tg.NewEditMessageTextAndMarkup(chatID, msgID, "Выбери число:", tg.NewInlineKeyboardMarkup(rows...))

	user.HandleStep = int(chainer.StaffChangeTimeStep)
	err := r.opts.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &msgReply, nil
}
