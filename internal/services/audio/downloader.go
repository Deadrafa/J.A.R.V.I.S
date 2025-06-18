package audio

import (
	"fmt"
	"io"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAudioDownloader struct {
	Bot *tgbotapi.BotAPI
}

func (d *TelegramAudioDownloader) Download(fileID string) (string, error) {
	fileURL, err := d.Bot.GetFileDirectURL(fileID)
	if err != nil {
		return "", fmt.Errorf("error Bot.GetFileDirectURL(): %w", err)
	}

	resp, err := http.Get(fileURL)
	if err != nil {
		return "", fmt.Errorf("error http.Get(): %w", err)
	}
	defer resp.Body.Close()

	audioPath := "audio_message.ogg"
	outFile, err := os.Create(audioPath)
	if err != nil {
		return "", fmt.Errorf("error os.Create(): %w", err)
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, resp.Body); err != nil {
		return "", fmt.Errorf("error io.Copy(): %w", err)
	}

	return audioPath, nil
}
