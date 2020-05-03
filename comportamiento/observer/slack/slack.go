package slack

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const (
	apiURL  = "https://hooks.slack.com/services"
	webhook = "/T207WP78B/B012XGPKUSZ/cKhZjr5JhWRmkp2bGX9qPQQ0"
)

// Slack .
type Slack struct{}

// Notify .
func (s *Slack) Notify(data string) {
	sendMessage(data)
}

func sendMessage(data string) {
	msg := fmt.Sprintf(`{"text":%q}`, data)
	req, err := http.NewRequest(http.MethodPost, apiURL+webhook, bytes.NewReader([]byte(msg)))
	if err != nil {
		log.Fatalf("no se pudo crear el request: %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("no se pudo enviar el mensaje a slack: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("se obtuvo un código de respuesta no esperado: %d", resp.StatusCode)
	}

	fmt.Println("mensaje enviado a slack")
}
