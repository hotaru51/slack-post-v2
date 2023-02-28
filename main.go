package main

import (
	"fmt"

	"github.com/hotaru51/slack-post-v2/messeage"
	"github.com/hotaru51/slack-post-v2/config"
)

func main() {
	fmt.Println(messeage.GetMessage().Message)
	fmt.Println(config.GetAbsPathOfExecutable())
	fmt.Println(config.GetWebhookUrl().WebhookUrl)
}
