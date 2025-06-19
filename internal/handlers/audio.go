package handlers

import (
	"log"

	"github.com/Deadrafa/J.A.R.V.I.S/internal/controllers/event"
	"github.com/Deadrafa/J.A.R.V.I.S/internal/services/ai"
	"github.com/Deadrafa/J.A.R.V.I.S/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AudioHandler struct {
	Bot         *tgbotapi.BotAPI
	Downloader  repository.AudioDownloader
	Recognizer  repository.SpeechRecognizer
	GigaService ai.GigaChatService
	EventRouter event.EventRouter
}

func (h *AudioHandler) HandleAudio(msg *tgbotapi.Message, fileID string) {
	h.Bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Аудио получено! Обрабатываю..."))

	audio, err := h.Downloader.Download(fileID)
	if err != nil {
		log.Printf("Ошибка загрузки: %v", err)
		return
	}

	transcript, err := h.Recognizer.Recognize(string(audio))
	if err != nil {
		log.Printf("Ошибка распознавания: %v", err)
		h.Bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Ошибка распознавания речи"))
		return
	}

	h.Bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Распознанный текст:\n"+transcript))

	respBody, err := h.GigaService.SendRequest(transcript)
	if err != nil {
		log.Printf("Ошибка отправки в GigaChat: %v", err)
		return
	}
	if err = h.EventRouter.RouteEvent(&respBody); err != nil {
		log.Printf("Ошибка в контроллере: %v", err)
		return
	}
}
