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
	Bearer  string
	Dataset string
}

func (s *GigaChatService) SendRequest(text string) (models.GigChatResp, error) {

	requestBody := models.GigChatReq{
		Model: s.Model,
		Mes: []models.Message{
			{Role: "system", Content: s.Dataset},
			{Role: s.Role, Content: text},
		},
		Temperature: 1,
		Stream:      false,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return models.GigChatResp{}, fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", s.BaseURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return models.GigChatResp{}, fmt.Errorf("error creating request: %w", err)

	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.Bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.GigChatResp{}, fmt.Errorf("error sending request: %w", err)

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.GigChatResp{}, fmt.Errorf("error reading response: %w", err)

	}

	fmt.Println("Status Code:", resp.Status)
	var jsonResp models.GigChatResp
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return models.GigChatResp{}, fmt.Errorf("error Unmarshal JSON: %w", err)
	}
	fmt.Println(jsonResp.Choices[0].Message.Content)

	return jsonResp, nil

}
