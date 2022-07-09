package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	ENV_SLACK_WEBHOOK_URL = string("SLACK_WEBHOOK_URL")
)

type SlackWebhookUlr struct {
	WebhookUrl string `json:"slackWebhookUrl"`
}

func GetAbsPathOfExecutable() string {
	relativePath := filepath.Dir(os.Args[0])
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return absPath
}

func GetWebhookUrl() *SlackWebhookUlr {
	url := new(SlackWebhookUlr)

	// 環境変数から取得
	url.WebhookUrl = os.Getenv(ENV_SLACK_WEBHOOK_URL)

	return url
}
