package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/hotaru51/slack-post-v2/config"
	"github.com/hotaru51/slack-post-v2/messeage"
)

func main() {
	webhookUrl := config.GetWebhookUrl()
	body := messeage.GenerateMessageJson()

	resp, err := http.Post(webhookUrl.WebhookUrl, "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(body)
	fmt.Println(resp.Status)
}
