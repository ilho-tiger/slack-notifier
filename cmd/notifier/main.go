package main

import (
	"fmt"

	"github.com/ilho-tiger/slack-notifier/pkg/config"
)

func main() {
	c := config.InitConfig()
	c.Add("SLACK_WEBHOOK", "slack-webhook", "https://testing.com", "slack webhook url")
	c.Parse()

	for _, configValue := range c.Configuration() {
		fmt.Println(configValue)
	}
}
