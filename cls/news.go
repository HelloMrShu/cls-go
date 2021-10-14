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

type Url struct {
	host               string
	path               string
	app                string
	hasFirstVipArticle int
	lastTime           int64
	sign               string
}

type Result struct {
	Error int  `json:"error"`
	Data  Data `json:"data"`
}

type Data struct {
	RollData  []News `json:"roll_data"`
	UpdateNum int    `json:"update_num"`
}

type Subjects struct {
	SubjectID   int    `json:"subject_id"`
	SubjectName string `json:"subject_name"`
}

type News struct {
	Level      string        `json:"level"`
	Content    string        `json:"content"`
	InRoll     int           `json:"in_roll"`
	Recommend  int           `json:"recommend"`
	Confirmed  int           `json:"confirmed"`
	UserID     int           `json:"user_id"`
	IsTop      int           `json:"is_top"`
	Brief      string        `json:"brief"`
	ID         int           `json:"id"`
	Ctime      int           `json:"ctime"`
	Type       int           `json:"type"`
	Title      string        `json:"title"`
	Bold       int           `json:"bold"`
	SortScore  int           `json:"sort_score"`
	CommentNum int           `json:"comment_num"`
	Status     int           `json:"status"`
	Category   string        `json:"category"`
	ShareNum   int           `json:"share_num"`
	StockList  []interface{} `json:"stock_list"`
	IsAd       int           `json:"is_ad"`
	AudioURL   []string      `json:"audio_url"`
	Subjects   []Subjects    `json:"subjects"`
}

type Txt struct {
	Content string `json:"content"`
}

type Notice struct {
	MsgType string `json:"msgtype"`
	Text    Txt    `json:"text"`
}

func NewsRequest(lt int64) int64 {
	newsConf := app.Conf.Cls.News

	st := Url{newsConf.Host, newsConf.Path, "CailianpressWeb", 1, lt, util.GenRandStrings(16)}
	url := fmt.Sprintf("%s%s?app=%s&hasFirstVipArticle=%d&lastTime=%d&sign=%s", st.host, st.path, st.app, st.hasFirstVipArticle, st.lastTime, st.sign)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", app.Conf.ContentType)
	req.Header.Set("Charset", app.Conf.Charset)
	req.Header.Set("Referer", newsConf.Refer)
	req.Header.Set("User-Agent", app.Conf.Ua)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		app.Logger.WithFields(logrus.Fields{"error": err}).Error("请求news接口异常")
		return lt
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	r := Result{}
	err1 := json.Unmarshal(body, &r)
	if r.Error != 0 {
		app.Logger.Error("cls request error")
		return lt
	}

	if err1 != nil {
		app.Logger.Error("解析body时，发生错误")
		return lt
	}

	updateNum := r.Data.UpdateNum
	news := r.Data.RollData

	curTs := time.Now().Unix()
	if updateNum == 0 {
		return curTs
	}

	app.Logger.Info("新闻消息推送条数: " + strconv.Itoa(updateNum))
	for _, v := range news {
		keywords := extractSubjects(v.Subjects, v.Recommend)
		msgInfo := GenNewsMessage(v.Brief + keywords)
		msg.SendNotice(msgInfo)
		curTs = int64(v.SortScore)
	}
	return curTs
}

func extractSubjects(subjects []Subjects, recommend int) string {

	if len(subjects) <= 0 {
		return ""
	}

	var items []string
	for _, v := range subjects {
		items = append(items, v.SubjectName)
	}

	keywords := strings.Join(items, ",")
	if recommend > 0 {
		keywords = keywords + " →推荐"
	}

	return "\n\n关键词：" + keywords
}

func GenNewsMessage(data string) string {
	nt := new(Notice)
	nt.MsgType = "text"
	nt.Text.Content = data

	notice, _ := json.Marshal(nt)
	return string(notice)
}

// CheckNewsMoment 新闻发送时间段
func CheckNewsMoment() bool {
	now := time.Now()
	hour, err := strconv.Atoi(now.Format("15")) //小时

	if err != nil {
		return false
	}

	if hour >= 23 || hour <= 7 { //早8点-晚10点发送
		return false
	}
	return true
}
