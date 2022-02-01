package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	Logger = logrus.New()
	err error
)

func InitLogger() {
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&LogFormatter{})
	Logger.SetOutput(os.Stdout)

	logDir := "./logs/"
	logFile := logDir + "app.log"
	_, err = os.Stat(logDir)
	if err != nil {
		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			Logger.Fatal("create log directory error: ", err)
			os.Exit(1)
		}
		_, err = os.Stat(logFile)
		if err != nil {
			_, err := os.Create(logDir + logFile)
			if err != nil {
				Logger.Fatal("create log file error: ", err)
				os.Exit(1)
			}
		}
	}

	src, fErr := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if fErr != nil {
		fmt.Println("open log file err ", fErr)
	}
	Logger.SetOutput(src)
	Logger.SetLevel(logrus.InfoLevel)
}