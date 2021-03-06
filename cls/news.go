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
	Shareurl   string        `json:"shareurl"`
	ShareNum   int           `json:"share_num"`
	SubTitles  []interface{} `json:"sub_titles"`
	StockList  []interface{} `json:"stock_list"`
	IsAd       int           `json:"is_ad"`
	AudioURL   []string      `json:"audio_url"`
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
		fmt.Println(err)
		return lt
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	r := Result{}
	err1 := json.Unmarshal(body, &r)
	if err1 != nil {
		fmt.Println("json to struct error")
	}

	if r.Error != 0 {
		fmt.Println("cls request error")
	}

	updateNum := r.Data.UpdateNum
	news := r.Data.RollData

	if updateNum == 0 {
		return time.Now().Unix()
	}

	fmt.Println("???????????????????????????", updateNum)
	newTs := time.Now().Unix()
	for _, v := range news {
		msgInfo := GenNewsMessage(v.Brief)
		msg.SendNotice(msgInfo)
		newTs = int64(v.SortScore)
	}
	return newTs
}

func GenNewsMessage(data string) string {
	var txt Txt
	txt.Content = data

	var nt Notice
	nt.MsgType = "text"
	nt.Text = txt

	notice, _ := json.Marshal(nt)
	return string(notice)
}

func CheckNewsMoment() bool {
	now := time.Now()
	hour, err := strconv.Atoi(now.Format("15")) //??????

	if err != nil {
		return false
	}

	if hour > 22 && hour < 7 {
		return false
	}
	return true
}
