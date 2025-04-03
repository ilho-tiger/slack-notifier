package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Webhook represents a Slack webhook configuration.
// It contains the URL of the webhook endpoint used to send messages to Slack.
type Webhook struct {
	url string
}

// NewWebhook creates a new instance of Webhook with the specified URL.
// The URL is used to send messages to a Slack webhook endpoint.
//
// Parameters:
//   - url: The Slack webhook URL.
//
// Returns:
//
//	A pointer to a Webhook instance initialized with the provided URL.
func NewWebhook(url string) *Webhook {
	return &Webhook{url: url}
}

// Send sends a message to the Slack webhook URL configured in the Webhook instance.
// It takes a string message as input, marshals it into a JSON payload, and sends it
// as an HTTP POST request with a "text" field containing the message.
//
// If the HTTP request fails or the response status code is not 200 OK, it returns an error
// with details about the failure. If the operation is successful, it returns nil.
//
// Parameters:
//   - message: The message string to be sent to the Slack webhook.
//
// Returns:
//   - error: An error if the message could not be sent or if the response from Slack
//     indicates a failure; otherwise, nil.
func (w *Webhook) Send(message string) error {
	payload := map[string]string{"text": message}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(w.url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error sending message to Slack: %s", string(body))
	}

	return nil
}
