package logger

import (
	"fmt"
	"io"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	defaultTimestampFormat = "2006-01-02 15:04:05"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type MyFormatter struct {
	TimestampFormat string
	enableColor     bool
	errorWrite      io.Writer
}

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}
	timestamp := entry.Time.Format(timestampFormat)

	level := entry.Level
	msg := fmt.Sprintf("[%s] [%s] %s\n", timestamp, strings.ToUpper(level.String()), entry.Message)

	if f.errorWrite != nil && level == logrus.ErrorLevel {
		f.errorWrite.Write([]byte(msg))
	}

	if f.enableColor {
		var levelColor int
		switch level {
		case logrus.DebugLevel, logrus.TraceLevel:
			levelColor = gray
		case logrus.WarnLevel:
			levelColor = yellow
		case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
			levelColor = red
		default:
			levelColor = blue
		}
		msg = fmt.Sprintf("\x1b[%dm%s\x1b[0m", levelColor, msg)
	}

	return []byte(msg), nil
}
