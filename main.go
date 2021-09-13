package main

import (
	cls "finance/cls"
	util "finance/utils"
	"fmt"
	"math/rand"
	"time"
)

func getHotSendHours() map[int]bool {

	hotHours := make(map[int]bool)

	hotHours[10] = true
	hotHours[11] = true
	hotHours[14] = true
	hotHours[15] = true

	return hotHours
}


func main() {

	ch := make(chan int64, 1)
	ch <- time.Now().Unix()

	count := 1

	for {
		fmt.Println("当前请求时间：", util.GetCurDate())

		hour := time.Now().Hour()
		hotHours := getHotSendHours()
		_, ok := hotHours[hour]
		if ok {
			fmt.Println("热门板块消息推送...")
			cls.HotRequest()
		}

		if hour >= 22 || hour <= 8 {
			fmt.Println("新闻消息推送服务暂停中...")
			time.Sleep(time.Duration(3600) * time.Second)
		}  

		ts := <- ch
		fmt.Println("新闻消息上次投递时间：", util.GetTsToDate(ts))
	
		new_ts := cls.NewsRequest(ts + 10)
		ch <- new_ts
		
		if count < 20 {
			count = count + 1
		} else {
			count = 1
		}

		delay := count * (rand.Intn(15) + 15) 
		time.Sleep(time.Duration(delay) * time.Second)
	}
}
