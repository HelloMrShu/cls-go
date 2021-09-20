package message

import (
	"finance/app"
	"log"
	"net/http"
	"strings"
)

func SendNotice(notice string) bool {
	webhook := app.Conf.Webhook //此处填写企业微信的webhook地址

	req, _ := http.NewRequest("POST", webhook, strings.NewReader(notice))

	req.Header.Set("Content-Type", app.Conf.ContentType)
	req.Header.Set("Charset", app.Conf.Charset)

	resp, err := (&http.Client{}).Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}
