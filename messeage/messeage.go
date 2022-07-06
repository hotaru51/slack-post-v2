package messeage

import (
	"fmt"
	"os"
	"io/ioutil"
	"time"

	"golang.org/x/term"
)

type MessageData struct {
	HostName		string
	DataTimeText	string
	Message			string
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

	if !term.IsTerminal(0) {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		md.Message = string(data)
	} else if len(os.Args) >= 2 {
		md.Message = os.Args[1]
	} else {
		fmt.Println("no text specified.")
			os.Exit(1)
	}

	return fmt.Sprint(md)
}
