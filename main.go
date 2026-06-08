package main
import (
	"os"
	"github.com/joho/godotenv"
	"Shannon/bot"
)

 func main() {
    godotenv.Load()
    
    tg_bot, err := bot.NewBot(os.Getenv("BOT_TOKEN"))
    if err != nil {
        panic(err)
    }
   
    tg_bot.Start()
}
