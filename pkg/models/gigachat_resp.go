package models

type GigChatReq struct {
	Model       string    `json:"model"`
	Mes         []Message `json:"messages"`
	Temperature int       `json:"temperature"`
	Stream      bool      `json:"stream"`
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type GigChatResp struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message      `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
