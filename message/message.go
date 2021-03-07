package message

import (
	"strings"
	"net/http"
)

func SendNotice(notice string) bool {
    webhook := ""//此处填写企业微信的webhook地址

    req, _ := http.NewRequest("POST", webhook, strings.NewReader(notice))
    req.Header.Set("Content-Type","application/json")
    req.Header.Set("Charset","utf8")

    resp, err := (&http.Client{}).Do(req)
    if err != nil {
		return false
    }

    defer resp.Body.Close()
    return true 
}

