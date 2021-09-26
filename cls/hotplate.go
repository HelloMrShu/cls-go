package cls

import (
	"encoding/json"
	"finance/app"
	msg "finance/message"
	util "finance/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// HotUrl cls hot plate struct
type HotUrl struct {
	host     string
	path     string
	app      string
	category string
	way      string
	sign     string
}

type Plate struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data []Industry `json:"data"`
}
type UpStock struct {
	SecuCode string  `json:"secu_code"`
	SecuName string  `json:"secu_name"`
	Change   float64 `json:"change"`
}
type Industry struct {
	SecuCode     string    `json:"secu_code"`
	SecuName     string    `json:"secu_name"`
	Change       float64   `json:"change"`
	MainFundDiff int       `json:"main_fund_diff"`
	UpStock      []UpStock `json:"up_stock"`
}

// HotRequest 板块信息推送
func HotRequest(category string) bool {
	hotConf := app.Conf.Cls.Hot
	hp := HotUrl{hotConf.Host, hotConf.Path, "CailianpressWeb", category, "change", util.GenRandStrings(16)}
	url := fmt.Sprintf("%s%s?app=%s&type=%s&way=%s&sign=%s", hp.host, hp.path, hp.app, hp.category, hp.way, hp.sign)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", app.Conf.ContentType)
	req.Header.Set("Charset", app.Conf.Charset)
	req.Header.Set("Referer", hotConf.Refer)
	req.Header.Set("User-Agent", app.Conf.Ua)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		app.Logger.WithFields(logrus.Fields{"error": err}).Error("请求hot plate接口异常")
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	rs := Plate{}
	err2 := json.Unmarshal(body, &rs)

	if err2 != nil || rs.Code != 200 || len(rs.Data) == 0 {
		return false
	}

	title := convertName(category) + "：\n"

	plates := make([]string, 0)
	plates = append(plates, title)
	for _, v := range rs.Data {
		change := strconv.FormatFloat(v.Change*100, 'g', 2, 64)
		plate := v.SecuName + "(" + change + "%)"
		plates = append(plates, plate)
	}

	res := strings.Join(plates, "\n")
	text := GenNewsMessage(res)
	msg.SendNotice(text)
	return true
}

// CheckMoment 检查热门板块发送时刻
func CheckMoment() bool {
	now := time.Now()
	weekday := now.Weekday().String()
	if weekday == "Sunday" || weekday == "Saturday" {
		return false
	}

	hourMinute := now.Format("15:04") //时分
	moment := util.GetHotPlateMoment()
	_, ok := moment[hourMinute]
	if ok {
		return true
	}
	return false
}

// 板块中英文转换
func convertName(cat string) string {
	mp := make(map[string]string)
	mp["industry"] = "行业板块"
	mp["concept"] = "概念板块"

	zh, ok := mp[cat]

	if ok {
		return zh
	}
	return ""
}
