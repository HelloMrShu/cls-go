package main

import (
	"finance/app"
	"finance/cls"
	util "finance/utils"
	"fmt"
	"time"
)

func init() {
	app.NewCycle()
}

func main() {
	ch := make(chan int64, 1)
	ch <- time.Now().Unix() - 10

	for {
		hour := time.Now().Hour()
		hotHours := util.GetHotPlateSendHours()
		_, ok := hotHours[hour]
		if ok {
			fmt.Println("热门板块消息推送...", util.GetCurDate())
			cls.HotRequest()
		}

		ts := <-ch
		fmt.Println("新闻消息上次投递时间：", util.GetTsToDate(ts))

		newTs := cls.NewsRequest(ts)
		ch <- newTs

		time.Sleep(time.Duration(60) * time.Second)
	}
}
