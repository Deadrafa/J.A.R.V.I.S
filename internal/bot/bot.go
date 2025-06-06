package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler struct {
	Bot *tgbotapi.BotAPI
}

func NewBotHandler(token string) (*BotHandler, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &BotHandler{Bot: bot}, nil
}

func (h *BotHandler) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := h.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Voice != nil {
			HandleAudio(h.Bot, update.Message, update.Message.Voice.FileID)
		}

	}
}
