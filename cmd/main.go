package main

import (
	"flag"
	"log"

	"github.com/Deadrafa/J.A.R.V.I.S/internal/bot"
	"github.com/Deadrafa/J.A.R.V.I.S/internal/config"
)

func main() {
	cfgPath := flag.String("cfg", "internal/config/default.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.NewConfig(*cfgPath)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфига: %v", err)
	}

	botHandler, err := bot.NewBotHandler(cfg.Token)
	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %v", err)
	}

	log.Println("Бот успешно запущен!")

	botHandler.Start()

}
