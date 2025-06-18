package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Deadrafa/J.A.R.V.I.S/pkg/models"
)

type GigaChatService struct {
	BaseURL string
	Token   string
	Model   string
	Role    string
}

func (s *GigaChatService) SendRequest(text string) error {
	requestBody := models.GigChatResp{
		Model: s.Model,
		Mes: []models.Messages{
			{Role: s.Role, Content: text},
		},
		Temperature: 1,
		Stream:      false,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", s.BaseURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)

	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer secret123")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)

	}

	fmt.Println("Status Code:", resp.Status)
	fmt.Println("Response Body:", string(body))
	return nil

}
