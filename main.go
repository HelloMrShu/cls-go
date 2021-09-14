package main

import (
	"finance/cls"
	util "finance/utils"
	"fmt"
	"math/rand"
	"time"
)

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

		fmt.Println(count)
		if count < 20 {
			count = count + 1
		} else {
			count = 1
		}

		delay := count * (rand.Intn(5) + 15)
		time.Sleep(time.Duration(delay) * time.Second)
	}
}
