package main

// func main() {
// 	cfgPath := flag.String("cfg", "internal/config/default.yaml", "path to config file")
// 	flag.Parse()

// 	cfg, err := config.NewConfig(*cfgPath)
// 	if err != nil {
// 		log.Fatalf("Ошибка загрузки конфига: %v", err)
// 	}

// 	botHandler, err := bot.NewBotHandler(cfg.Token)
// 	if err != nil {
// 		log.Fatalf("Ошибка инициализации бота: %v", err)
// 	}

// 	log.Println("Бот успешно запущен!")

// 	botHandler.Start()
// }

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	url := "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"
	method := "POST"

	payload := strings.NewReader(`{
  "model": "GigaChat",
  "messages": [
    {
      "role": "system",
      "content": "Ты — профессиональный переводчик на английский язык. Переведи точно сообщение пользователя."
    },
    {
      "role": "user",
      "content": "GigaChat — это сервис, который умеет взаимодействовать с пользователем в формате диалога, писать код, создавать тексты и картинки по запросу пользователя."
    }
  ],
  "stream": false,
  "update_interval": 0
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer secret123")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
