package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"encoding/json"
)

const (
	ENV_SLACK_WEBHOOK_URL = string("SLACK_WEBHOOK_URL")
	CONFIG_FILE_NAME = string("config.json")
)

type SlackWebhookUrl struct {
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

func GetWebhookUrl() *SlackWebhookUrl {
	url := new(SlackWebhookUrl)

	// 環境変数から取得
	if fromEnv := os.Getenv(ENV_SLACK_WEBHOOK_URL); len(fromEnv) >= 1 {
		url.WebhookUrl = fromEnv
	} else {
		url.WebhookUrl = ReadConfigJson()
	}

	if len(url.WebhookUrl) <= 0 {
		fmt.Println("please specify a valid webhook URL.(config.json or environment variable SLACK_WEBHOOK_URL)")
		os.Exit(1)
	}

	return url
}

func ReadConfigJson() string {
	jsonFilePath := GetAbsPathOfExecutable() + "/" + CONFIG_FILE_NAME
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	byteArr, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var webhookUrl SlackWebhookUrl
	err = json.Unmarshal(byteArr, &webhookUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return webhookUrl.WebhookUrl
}
