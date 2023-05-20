package discord

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// TODO
type WebhookSend struct {
	Content   string `json:"content,omitempty"`
	Username  string `json:"username,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Embeds    []any  `json:"embeds,omitempty"`
}

func SendWebhook(data WebhookSend) error {
	body, err := json.Marshal(data)
	var resp *http.Response

	resp, err = http.Post("", "application/json", bytes.NewReader(body))
	bodyBytes, err := io.ReadAll(resp.Body)
	print(string(bodyBytes))
	return err
}
