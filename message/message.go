package message

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"finance/app"
)

func SendNotice(notice string) bool {
    webhook := app.Conf.Webhook //此处填写企业微信的webhook地址

	req, _ := http.NewRequest("POST", webhook, strings.NewReader(notice))

	req.Header.Set("Content-Type","application/json")
    req.Header.Set("Charset","utf8")

    resp, err := (&http.Client{}).Do(req)

	fmt.Println(resp, err)
	os.Exit(0)
    if err != nil {
		return false
    }

    defer resp.Body.Close()
    return true 
}

