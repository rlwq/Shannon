package bot

import (
	"Shannon/service"
	"Shannon/shannon"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var browsingKeyboard = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("👍"),
		tg.NewKeyboardButton("👎"),
	),
)

type Bot struct {
	api        *tg.BotAPI
	fsm        *FSM
	unfinished map[int64]*shannon.Profile

	service *service.Service
}

func NewBot(token string, service *service.Service) (*Bot, error) {
	api, err := tg.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{
		api:        api,
		fsm:        NewFSM(),
		unfinished: (map[int64]*shannon.Profile{}),
		service:    service,
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

	state := bot.fsm.GetState(from)

	if state == StateUnknown && bot.service.DoesProfileExist(from) {
		state = StateSleep
	}

	switch state {
	case StateUnknown:
		bot.SendMessage(from, "Welcome!")

		bot.SendMessage(from, "Enter your name: ")
		bot.fsm.SetState(from, StateWaitName)

		bot.unfinished[from] = &shannon.Profile{UserID: from}

	case StateSleep:
		bot.SendMessage(from, "Welcome back!")
		bot.fsm.SetState(from, StateBrowsing)

	case StateWaitName:
		bot.SendMessage(from, "Cool name!")

		bot.SendMessage(from, "Enter your bio: ")
		bot.fsm.SetState(from, StateWaitBio)

		bot.unfinished[from].Name = text

	case StateWaitBio:
		bot.SendMessage(from, "Cool bio!")

		bot.SendMessage(from, "Start browsing!")
		bot.fsm.SetState(from, StateBrowsing)

		bot.unfinished[from].Bio = text

		bot.SendMessage(
			from,
			"Your profile:\n"+
				bot.unfinished[from].Name+
				"\n"+bot.unfinished[from].Bio)

		bot.service.CreateProfile(*bot.unfinished[from])
		delete(bot.unfinished, from)

		profile := bot.service.NextProfileFor(from)
		bot.SendProfile(from, profile)

	case StateBrowsing:
		profile := bot.service.NextProfileFor(from)
		bot.SendProfile(from, profile)
	}
}

func (bot *Bot) SendProfile(chat int64, profile shannon.Profile) {
	text := profile.Name + "\n" + profile.Bio
	msg := tg.NewMessage(chat, text)
	msg.ReplyMarkup = browsingKeyboard
	bot.api.Send(msg)
}

func (bot *Bot) Run() {
	updateConfig := tg.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.api.GetUpdatesChan(updateConfig)

	for update := range updates {
		bot.handleUpdate(update)
	}
}
