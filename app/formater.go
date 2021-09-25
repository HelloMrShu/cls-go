package app

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

type LogFormatter struct {

}

func (m *LogFormatter) Format(entry *logrus.Entry) ([]byte, error){
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("%s %s message:\"%s\" ", timestamp, entry.Level, entry.Message)

	if len(entry.Data) > 0 {
		b.WriteString("data:")
		var data []string
		for key, value := range entry.Data {
			str := fmt.Sprintf("%s=%v", key, value)
			data = append(data, str)
		}
		fieldsData := strings.Join(data, ",")
		message += fieldsData
	}
	b.WriteString(message + "\n")

	return b.Bytes(), nil
}