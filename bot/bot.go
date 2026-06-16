package bot

import (
	"Shannon/shannon"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api        *tg.BotAPI
	fsm        *FSM
	unfinished map[int64]*shannon.Profile
}

func NewBot(token string) (*Bot, error) {
	api, err := tg.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{
		api:        api,
		fsm:        NewFSM(),
		unfinished: (map[int64]*shannon.Profile{}),
	}, nil
}

func (bot *Bot) SendMessage(chat int64, text string) {
	msg := tg.NewMessage(chat, text)
	bot.api.Send(msg)
}

func (bot *Bot) handleUpdate(update tg.Update) {
	if update.Message == nil {
		return
	}

	from := update.Message.From.ID
	text := update.Message.Text

	state, ok := bot.fsm.states[from]

	if !ok {
		state = StateUnknown
		bot.fsm.states[from] = StateUnknown
	}

	switch state {
	case StateUnknown:
		bot.SendMessage(from, "Welcome!")

		bot.SendMessage(from, "Enter your name: ")
		bot.fsm.states[from] = StateWaitName

		bot.unfinished[from] = &shannon.Profile{}

	case StateWaitName:
		bot.SendMessage(from, "Cool name!")

		bot.SendMessage(from, "Enter your bio: ")
		bot.fsm.states[from] = StateWaitBio

		bot.unfinished[from].Name = text

	case StateWaitBio:
		bot.SendMessage(from, "Cool bio!")

		bot.SendMessage(from, "Start browsing!")
		bot.fsm.states[from] = StateBrowsing

		bot.unfinished[from].Bio = text

		bot.SendMessage(
			from,
			"Your profile:\n"+
				bot.unfinished[from].Name+
				"\n"+bot.unfinished[from].Bio)

		delete(bot.unfinished, from)
	}
}

func (bot *Bot) Start() {
	updateConfig := tg.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.api.GetUpdatesChan(updateConfig)

	for update := range updates {
		bot.handleUpdate(update)
	}
}
