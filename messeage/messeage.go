package messeage

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"time"
	"strings"

	"golang.org/x/term"
)

type MessageData struct {
	HostName     string
	DataTimeText string
	Message      string
}

// 現在日時、ホスト名を反映したMessageDataを生成
func NewMessageData(message string) *MessageData {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}

	now := time.Now()
	timeText := fmt.Sprintf(
		"%02d-%02d-%02d %02d:%02d:%02d",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
	)

	md := &MessageData{
		HostName: hostname,
		DataTimeText: timeText,
		Message: message,
	}

	return md
}

func (m *MessageData) String() string {
	return fmt.Sprintf("HostName: %s, Time: %s, Message: %s", m.HostName, m.DataTimeText, m.Message)
}

/*
 * postするrequest bodyのイメージ
 * {
 *   "text": "message",
 *   "blocks": [
 *     {
 *       "type": "section",
 *       "text": {
 *         "type": "mrkdwn",
 *         "text": "message"
 *       }
 *     }
 *   ]
 * }
 */

type BlockText struct {
	TextType string `json:"type"`
	Text     string `json:"text"`
}

func NewBlockText(textType string, text string) *BlockText {
	return &BlockText{
		TextType: textType,
		Text: text,
	}
}

type Block struct {
	BlockType string     `json:"type"`
	BlockText *BlockText `json:"text"`
}

func NewBlock(blockType string, blockText *BlockText) *Block {
	return &Block{
		BlockType: blockType,
		BlockText: blockText,
	}
}

type MessageBody struct {
	Text   string   `json:"text"`
	Blocks []*Block `json:"blocks"`
}

func NewMessageBody(text string, blocks []*Block) *MessageBody {
	return &MessageBody{
		Text: text,
		Blocks: blocks,
	}
}

// コマンドライン引数、またはパイプで渡されたメッセージを取得して返す
func GetMessage() *MessageData {
	md := NewMessageData("")

	if !term.IsTerminal(0) { // パイプで渡された場合
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		md.Message = string(data)
	} else if len(os.Args) >= 2 { // コマンドライン引数で渡された場合
		md.Message = os.Args[1]
	} else {
		fmt.Println("no text specified.")
			os.Exit(1)
	}

	md.Message = strings.TrimSpace(md.Message)

	// メッセージテンプレート
	tmplText := "*host:* `{{ .HostName }}`\n"
	tmplText += "*time:* `{{ .DataTimeText }}`\n"
	tmplText += "*message:*\n"
	tmplText += "```\n"
	tmplText += "{{ .Message }}\n"
	tmplText += "```"

	// テンプレートにMessageDataの値を反映してテキスト生成
	tmpl, err := template.New("").Parse(tmplText)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	buf := bytes.NewBuffer(make([]byte, 0))
	if err = tmpl.Execute(buf, md); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	md.Message = string(buf.Bytes())

	return md
}
