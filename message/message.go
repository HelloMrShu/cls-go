package message

import (
	"strings"
	"net/http"
)

func SendNotice(notice string) bool {
    webhook := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=f388073a-cdf7-4738-b43b-ede3abb4d692"

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

