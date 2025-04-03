package main

import (
	"fmt"

	"github.com/ilho-tiger/slack-notifier/pkg/config"
	"github.com/ilho-tiger/slack-notifier/pkg/slack"
)

func main() {
	c := config.InitConfig()
	c.Add("SLACK_WEBHOOK", "slack-webhook", "https://testing.com", "slack webhook url")
	c.Parse()

	for _, configValue := range c.Configuration() {
		fmt.Println(configValue)
	}

	// Initialize the Slack webhook with the URL from the config
	webhookURL, ok := c.Get("SLACK_WEBHOOK")
	if !ok {
		fmt.Println("SLACK_WEBHOOK not found in config")
		return
	}
	// Create a new Slack webhook instance
	webhook := slack.NewWebhook(webhookURL)
	// Send a test message to the Slack webhook
	if err := webhook.Send("Hello, Slack!"); err != nil {
		fmt.Printf("Error sending message to Slack: %v\n", err)
	} else {
		fmt.Println("Message sent to Slack successfully!")
	}
}
