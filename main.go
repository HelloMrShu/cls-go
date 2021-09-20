package main

import (
	"finance/app"
	"finance/cls"
	"flag"
	"strings"
	"time"
)

var (
	env string
)

func init() {
	flag.StringVar(&env, "c", "local", "conf path")
	flag.Parse()
	config := "./conf/" + env + ".json"

	app.NewCycle(config)

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
