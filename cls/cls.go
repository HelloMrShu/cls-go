package cls

import (
	"fmt"
    "net/http"
    "time"
	"encoding/json"
	"io/ioutil"
	msg "finance/message"
	util "finance/utils"
)

type ClsUrl struct {
  	host string
   	path string
   	app string
    hasFirstVipArticle int
    lastTime int64
	sign string
}

type Result struct {
    Error int `json:"error"`
  	Data Data `json:"data"`
   	VipData []interface{} `json:"vipData"`
}

type Data struct {
	RollData []News `json:"roll_data"` 
	UpdateNum int `json:"update_num"`
}

type Ad struct {
        ID int `json:"id"`
        Title string `json:"title"`
        Img string `json:"img"`
        URL string `json:"url"`
        MonitorURL string `json:"monitorUrl"`
        VideoURL string `json:"video_url"`
        AdTag string `json:"adTag"`
        FullScreen int `json:"fullScreen"`
}

type Subjects struct {
        ArticleID int `json:"article_id"`
        SubjectID int `json:"subject_id"`
        SubjectName string `json:"subject_name"`
        SubjectImg string `json:"subject_img"`
        SubjectDescription string `json:"subject_description"`
        CategoryID int `json:"category_id"`
        AttentionNum int `json:"attention_num"`
        IsAttention bool `json:"is_attention"`
}

type News struct {
        AuthorExtends string `json:"author_extends"`
        DepthExtends string `json:"depth_extends"`
        DenyComment int `json:"deny_comment"`
        Level string `json:"level"`
        ReadingNum int `json:"reading_num"`
        Content string `json:"content"`
        InRoll int `json:"in_roll"`
        Recommend int `json:"recommend"`
        Confirmed int `json:"confirmed"`
        Jpush int `json:"jpush"`
        Img string `json:"img"`
        UserID int `json:"user_id"`
        IsTop int `json:"is_top"`
        Brief string `json:"brief"`
        ID int `json:"id"`
        Ctime int `json:"ctime"`
        Type int `json:"type"`
        Title string `json:"title"`
        Bold int `json:"bold"`
        SortScore int `json:"sort_score"`
        CommentNum int `json:"comment_num"`
        ModifiedTime int `json:"modified_time"`
        Status int `json:"status"`
        Collection int `json:"collection"`
        HasImg int `json:"has_img"`
        Category string `json:"category"`
        Shareurl string `json:"shareurl"`
        ShareImg string `json:"share_img"`
        ShareNum int `json:"share_num"`
        SubTitles []interface{} `json:"sub_titles"`
        Tags []interface{} `json:"tags"`
        Imgs []interface{} `json:"imgs"`
        Images []interface{} `json:"images"`
        ExplainNum int `json:"explain_num"`
        StockList []interface{} `json:"stock_list"`
        IsAd int `json:"is_ad"`
        Ad Ad `json:"ad"`
		Subjects []Subjects `json:"subjects"`
        AudioURL []string `json:"audio_url"`
        Author string `json:"author"`
}

type Txt struct {
    Content string `json:"content"`
}

type Notice struct {
	Msgtype string `json:"msgtype"`
   	Text Txt `json:"text"`
}

var ua string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"

func NewsRequest(lt int64) int64 {
		clsHost := "https://www.cls.cn"
		realTech := "/nodeapi/updateTelegraphList"

        st := ClsUrl{clsHost, realTech, "CailianpressWeb",1, lt, util.GenRandStrings(16)}

        url := fmt.Sprintf("%s%s?app=%s&hasFirstVipArticle=%d&lastTime=%d&sign=%s", st.host,st.path,st.app,st.hasFirstVipArticle, st.lastTime, st.sign)

        req, _ := http.NewRequest("GET", url, nil)
        req.Header.Set("Content-Type","application/json")
        req.Header.Set("Charset","utf8")
        req.Header.Set("Referer","https://www.cls.cn")
        req.Header.Set("User-Agent", ua)

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
			return lt
		}

		new_ts := 0
		for _, v := range news {
			//id := v.ID
			text := GenNewsMessage(v.Brief)	
			msg.SendNotice(text)
			new_ts = v.SortScore 

			time.Sleep(time.Duration(2) * time.Second)
		}
		return int64(new_ts)
}

func GenNewsMessage(data string) string {
	var txt Txt
	txt.Content = data

	var nt Notice
	nt.Msgtype = "text"
	nt.Text = txt

	notice, _ := json.Marshal(nt)
	return string(notice)
}
