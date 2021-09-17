package cls

import (
	"encoding/json"
	"finance/app"
	msg "finance/message"
	util "finance/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// HotUrl cls hot plate struct
type HotUrl struct {
    host string
    path string
    app string
    category string
    way  string
    sign string
}

type Plate struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data []Industry `json:"data"`
}
type UpStock struct {
	SecuCode string `json:"secu_code"`
	SecuName string `json:"secu_name"`
	Change float64 `json:"change"`
}
type Industry struct {
	SecuCode string `json:"secu_code"`
	SecuName string `json:"secu_name"`
	Change float64 `json:"change"`
	MainFundDiff int `json:"main_fund_diff"`
	UpStock []UpStock `json:"up_stock"`
}

func HotRequest() bool {
    hotHost := "https://x-quote.cls.cn"
    hotPath := "/web_quote/plate/hot_plate"

    hp := HotUrl{hotHost, hotPath, "CailianpressWeb", "industry", "change", util.GenRandStrings(16)}
    url := fmt.Sprintf("%s%s?app=%s&type=%s&way=%s&sign=%s", hp.host, hp.path, hp.app, hp.category, hp.way, hp.sign)

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Content-Type","application/json")
    req.Header.Set("Charset","utf8")
    req.Header.Set("Referer","https://www.cls.cn")
    req.Header.Set("User-Agent", ua) 

    resp, err := (&http.Client{}).Do(req)
    if err != nil {
        fmt.Println(err)
        return false
    }   
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

	rs := Plate{}
	err2 := json.Unmarshal(body, &rs)

	if err2 != nil || rs.Code != 200 || len(rs.Data) == 0 {
		return false	
	}

	plates := make([]string, 0)
	plates = append(plates, "当前热门板块: \n\n")
	for _, v := range rs.Data {
		change := strconv.FormatFloat(v.Change * 100, 'g', 2, 64)
		plate := v.SecuName + "(" + change + "%)"
		plates = append(plates, plate) 
	}

	res := strings.Join(plates, "\n")
	text := GenNewsMessage(res)
	msg.SendNotice(text)

    return true
}
