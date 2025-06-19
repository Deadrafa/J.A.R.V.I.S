package instructions

import (
	"fmt"
	"os"
)

func UnloadingDataset(Filepath string) (string, error) {
	content, err := os.ReadFile(Filepath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(content), nil

}
