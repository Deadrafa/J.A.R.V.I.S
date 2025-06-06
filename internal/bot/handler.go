package bot

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleAudio(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, fileID string) {
	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Аудио получено! Обрабатываю..."))

	fileURL, err := bot.GetFileDirectURL(fileID)
	if err != nil {
		log.Printf("Ошибка получения URL: %v", err)
		return
	}

	resp, err := http.Get(fileURL)
	if err != nil {
		log.Printf("Ошибка скачивания: %v", err)
		return
	}
	defer resp.Body.Close()

	audioPath := "audio_message.ogg"
	outFile, err := os.Create(audioPath)
	if err != nil {
		log.Printf("Ошибка создания файла: %v", err)
		return
	}
	defer outFile.Close()
	defer os.Remove(audioPath)

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		log.Printf("Ошибка сохранения: %v", err)
		return
	}

	wavPath := "audio_message.wav"
	defer os.Remove(wavPath)
	if err := convertOggToWav(audioPath, wavPath); err != nil {
		log.Printf("Ошибка конвертации: %v", err)
		return
	}

	transcript, err := TranscribeAudio(wavPath)
	if err != nil {
		log.Printf("Ошибка распознавания: %v", err)
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Произошла ошибка при распознавании."))
		return
	}

	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Распознанный текст:\n"+transcript))
}

func convertOggToWav(oggPath, wavPath string) error {
	cmd := exec.Command("ffmpeg", "-i", oggPath, "-ar", "16000", "-ac", "1", wavPath)
	return cmd.Run()
}

func TranscribeAudio(path string) (string, error) {
	cmd := exec.Command("whisper", path, "--language", "ru", "--model", "base", "--output_format", "txt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	txtPath := path[:len(path)-4] + ".txt"
	content, err := os.ReadFile(txtPath)
	if err != nil {
		return "", err
	}
	defer os.Remove(txtPath)

	return string(content), nil
}
