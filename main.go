package main

import (
	"finance/app"
	"finance/cls"
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

var env string
var log = logrus.New()

func init() {
	flag.StringVar(&env, "c", "local", "conf path")
	flag.Parse()
	config := "./conf/" + env + ".json"
	app.NewCycle(config)

	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetOutput(os.Stdout)
	file, _ := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	//设置最低loglevel
	log.SetLevel(logrus.InfoLevel)
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
		delay := 0
		if cls.CheckNewsMoment() {
			newTs := cls.NewsRequest(ts)
			ch <- newTs
			delay = 60
		} else {
			ch <- time.Now().Unix()
			delay = 900
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}
}
