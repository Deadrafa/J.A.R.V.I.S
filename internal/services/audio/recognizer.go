package audio

import (
	"os"
	"os/exec"
)

type SpeechRecognitionService struct {
}

func (s *SpeechRecognitionService) Recognize(audioPath string) (string, error) {
	cmd := exec.Command("whisper", audioPath, "--language", "ru", "--model", "base", "--output_format", "txt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	txtPath := audioPath[:len(audioPath)-4] + ".txt"
	content, err := os.ReadFile(txtPath)
	if err != nil {
		return "", err
	}
	defer os.Remove(txtPath)
	defer os.Remove(audioPath)

	return string(content), nil
}
