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
	// 実行ファイルの絶対パスを取得
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 取得したパスのディレクトリを取得
	executableDir := filepath.Dir(executablePath)

	return executableDir
}

func GetWebhookUrl() *SlackWebhookUlr {
	url := new(SlackWebhookUlr)

	// 環境変数から取得
	url.WebhookUrl = os.Getenv(ENV_SLACK_WEBHOOK_URL)

	return url
}
