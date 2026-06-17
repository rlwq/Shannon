package main

import (
	"Shannon/app"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func main() {
	godotenv.Load()
	app := app.NewApp(os.Getenv("BOT_TOKEN"), "./db/file.db")
	app.Run()
}
