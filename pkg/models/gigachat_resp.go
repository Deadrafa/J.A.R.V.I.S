package models

type GigChatResp struct {
	Model       string     `json:"model"`
	Mes         []Messages `json:"messages"`
	Temperature int        `json:"temperature"`
	Stream      bool       `json:"stream"`
}

type Messages struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}
