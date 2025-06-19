package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Deadrafa/J.A.R.V.I.S/internal/config"
	"github.com/Deadrafa/J.A.R.V.I.S/internal/handlers"
	"github.com/Deadrafa/J.A.R.V.I.S/internal/services/ai"
	"github.com/Deadrafa/J.A.R.V.I.S/internal/services/ai/instructions"
	"github.com/Deadrafa/J.A.R.V.I.S/internal/services/audio"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfgPath := flag.String("cfg", "internal/config/default.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.NewConfig(*cfgPath)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфига: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("Ошибка создание бота : %v", err)
	}

	dataset, err := instructions.UnloadingDataset("internal/services/ai/instructions/event_operations.txt")
	if err != nil {
		log.Fatalf("Ошибка загрузки dataset: %v", err)
	}
	downloader := &audio.TelegramAudioDownloader{Bot: bot}
	recognizer := &audio.SpeechRecognitionService{}
	gigaService := &ai.GigaChatService{
		BaseURL: cfg.GigaChatURL,
		Token:   cfg.GigaChatToken,
		Model:   cfg.Model,
		Role:    cfg.Role,
		Bearer:  cfg.Bearer,
		Dataset: dataset,
	}

	audioHandler := &handlers.AudioHandler{
		Bot:         bot,
		Downloader:  downloader,
		Recognizer:  recognizer,
		GigaService: *gigaService,
	}

	fmt.Println("Бот запущен")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := audioHandler.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Voice != nil {
			audioHandler.HandleAudio(update.Message, update.Message.Voice.FileID)
		}

	}
}
