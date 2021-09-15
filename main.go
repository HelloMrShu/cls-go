package main

import (
	"encoding/json"
	"finance/cls"
	util "finance/utils"
	"fmt"
	"math/rand"
	"time"
	"io/ioutil"
)

var conf *Conf

type HotPlate struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

type NewsPlate struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

type ClsInfo struct {
	Hot HotPlate `json:"hot"`
	News NewsPlate `json:"news"`
}

type Conf struct {
	Webhook string `json:"webhook"`
	Cls ClsInfo `json:"cls"`
}

func init()  {
	filePath := "./conf/dev.json"

	var (
		data []byte
		err error
	)
	data, err = ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read conf file err", err)
		return
	}

	er := json.Unmarshal(data, &conf)
	if er != nil {
		fmt.Println("json unMarsha1 conf file err", er)
		return
	}
}

func main() {
	ch := make(chan int64, 1)
	ch <- time.Now().Unix()

	count := 1

	for {
		fmt.Println("热门板块消息推送...", util.GetCurDate())
		cls.HotRequest()

		ts := <-ch
		fmt.Println("新闻消息上次投递时间：", util.GetTsToDate(ts))

		newTs := cls.NewsRequest(ts + 10)
		ch <- newTs

		if count < 20 {
			count = count + 1
		} else {
			count = 1
		}

		delay := count * (rand.Intn(5) + 15)
		time.Sleep(time.Duration(delay) * time.Second)
	}
}
