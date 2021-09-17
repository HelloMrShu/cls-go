package message

import (
	"finance/app"
	"net/http"
	"strings"
)

func SendNotice(notice string) bool {
    webhook := app.Conf.Webhook //此处填写企业微信的webhook地址

	req, _ := http.NewRequest("POST", webhook, strings.NewReader(notice))

	req.Header.Set("Content-Type","application/json")
    req.Header.Set("Charset","utf8")

    resp, err := (&http.Client{}).Do(req)
	defer resp.Body.Close()

	if err != nil {
		return false
    }

    return true
}

