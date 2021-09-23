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

	log.SetReportCaller(true)
	log.SetFormatter(&app.LogFormatter{})
	log.SetOutput(os.Stdout)
	file, _ := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	log.SetOutput(file)
	log.SetLevel(logrus.InfoLevel)
}

func main() {
	ch := make(chan int64, 1)
	ch <- time.Now().Unix() - 10

	log.WithFields(logrus.Fields{
		"name": "gaoqi",
		"age" : 10,
	}).Info("test log1")
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
