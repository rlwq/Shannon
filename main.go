package main

import (
	"Shannon/bot"
	"database/sql"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
	"os"
)

func main() {
	println("123")
	db, err := sql.Open("sqlite", "./db/file.db")
	println(err)
	createTableSQL := `CREATE TABLE IF NOT EXISTS PROFILES (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER UNIQUE NOT NULL,
		username TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	db.Exec(createTableSQL)

	godotenv.Load()

	tg_bot, err := bot.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	tg_bot.Start()
}
