package messeage

import (
	"os"
)

func GetMessage() string {
	var messeage string

	// コマンドライン引数で渡された場合はその文字文字列を受け取る
	if len(os.Args) >= 1 {
		messeage = os.Args[1]
	}

	return messeage
}
