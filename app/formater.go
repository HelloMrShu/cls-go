package app

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
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
	prefix := fmt.Sprintf("%s %s message:\"%s\" ", timestamp, entry.Level, entry.Message)
	b.WriteString(prefix)

	if len(entry.Data) > 0 {
		b.WriteString("data:\"")
		for key, value := range entry.Data {
			b.WriteString(key)
			b.WriteByte('=')
			fmt.Fprint(b, value)
			b.WriteByte(',')
		}
		b.WriteString("\"\n")
	}

	return b.Bytes(), nil
}