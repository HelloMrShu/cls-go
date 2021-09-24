package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func InitLogger() {
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&LogFormatter{})
	Logger.SetOutput(os.Stdout)

	logPath := "./logs/app.log"
	_, err := os.Stat(logPath)
	if err != nil {
		ret := os.MkdirAll(logPath, 0755)
		if ret != nil {
			fmt.Println("can't create file ", ret)
			return
		}
	}

	file, err := os.OpenFile("", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("open ", err)
	}
	Logger.SetOutput(file)
	Logger.SetLevel(logrus.InfoLevel)
}