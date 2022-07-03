package messeage

import (
	"fmt"
	"os"
	"time"
)

type MessageData struct {
	HostName     string
	DataTimeText string
	Message      string
}

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

func GetMessage() string {
	md := NewMessageData("")

	// コマンドライン引数で渡された場合はその文字文字列を受け取る
	if len(os.Args) >= 2 {
		md.Message = os.Args[1]
	}

	return fmt.Sprint(md)
}
