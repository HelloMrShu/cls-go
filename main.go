package main

import (
	"finance/app"
	"finance/cls"
	"strings"
	"time"
)

func init() {
	app.NewCycle()
}

func main() {
	ch := make(chan int64, 1)
	ch <- time.Now().Unix() - 10

	for {
		if cls.CheckMoment() {
			categories := strings.Split(app.Conf.Cls.Hot.Categories, ",")
			for _, cat := range categories {
				cls.HotRequest(cat)
			}
		}

		ts := <-ch
		newTs := cls.NewsRequest(ts)
		ch <- newTs

		time.Sleep(time.Duration(60) * time.Second)
	}
}
