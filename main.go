package main

import (
	"Shannon/bot"
	"Shannon/repository"
	"Shannon/shannon"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func main() {
	repo := repository.NewRepository("./db/file.db")
	repo.WriteProfile(&shannon.Profile{UserID: 777, Name: "Arima", Bio: "Ghoul hunter"})
	godotenv.Load()

	tg_bot, err := bot.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	tg_bot.Start()
}
