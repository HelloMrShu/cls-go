package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func InitLogger() {
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&LogFormatter{})
	Logger.SetOutput(os.Stdout)
	file, _ := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	Logger.SetOutput(file)
	Logger.SetLevel(logrus.InfoLevel)
}