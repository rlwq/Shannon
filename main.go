package main

import (
	"Shannon/bot"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()

	tg_bot, err := bot.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	tg_bot.Start()
}
