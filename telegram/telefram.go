package telegram

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const (
	CHAT_ID      = -1001560843782
	BOT_TOKEN    = "1928876840:AAETQS2TNBn4irEnAWiH5xEI6Eip4Yy3lM0"
	TELEGRAM_URL = "https://api.telegram.org/bot"
)

type BotSendMessageId struct {
	Result struct {
		Message_id int
	}
}

func SendMessage(text string) {
	textAll := fmt.Sprintf("%s", text)
	fmt.Println(textAll)
	data := []byte(fmt.Sprintf(`{"chat_id":%d, "text":"%s", "parse_mode":"HTML", "disable_web_page_preview": "true"}`, CHAT_ID, textAll))

	// https://api.telegram.org/bot1928876840:AAETQS2TNBn4irEnAWiH5xEI6Eip4Yy3lM0/sendMessage?chat_id=-575227968&text=123
	fmt.Println(fmt.Sprintf("%s%s/sendMessage", TELEGRAM_URL, BOT_TOKEN))
	body := bytes.NewReader(data)
	_, err := http.Post(fmt.Sprintf("%s%s/sendMessage", TELEGRAM_URL, BOT_TOKEN), "application/json", body)
	//	_, err := http.Get(fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=%s&parse_mode=HTML", TELEGRAM_URL, BOT_TOKEN, CHAT_ID, textAll))
	if err != nil {
		log.Fatalf(err.Error())
	}
}
