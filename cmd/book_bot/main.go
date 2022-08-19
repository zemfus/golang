package main

import (
	"context"

	"boobot/internal/app"
)

func main() {
	ctx := context.Background()
	bot := app.New(ctx)

	bot.Run(ctx)
}
