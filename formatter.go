package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	FullDateFormat = "2006-01-02T15:04:05.000-07:00"
	Code           = "$$code"
)

const (
	red       = 31
	redLight  = 91
	yellow    = 33
	blue      = 36
	gray      = 37
	green     = 21
	cyan      = 36
	cyanLight = 96
)

var LevelsString = map[logrus.Level]string{
	logrus.PanicLevel: "PANIC",
	logrus.FatalLevel: "FATAL",
	logrus.ErrorLevel: "ERROR",
	logrus.WarnLevel:  "WARN ",
	logrus.InfoLevel:  "INFO ",
	logrus.DebugLevel: "DEBUG",
	logrus.TraceLevel: "TRACE",
}

var LevelsStringColored = map[logrus.Level]string{
	logrus.PanicLevel: fmt.Sprintf("\x1b[%dm%s\x1b[0m", red, "PANIC"),
	logrus.FatalLevel: fmt.Sprintf("\x1b[%dm%s\x1b[0m", redLight, "FATAL"),
	logrus.ErrorLevel: fmt.Sprintf("\x1b[%dm%s\x1b[0m", red, "ERROR"),
	logrus.WarnLevel:  fmt.Sprintf("\x1b[%dm%s\x1b[0m", yellow, "WARN "),
	logrus.InfoLevel:  fmt.Sprintf("\x1b[%dm%s\x1b[0m", cyanLight, "INFO "),
	logrus.DebugLevel: fmt.Sprintf("\x1b[%dm%s\x1b[0m", cyan, "DEBUG"),
	logrus.TraceLevel: fmt.Sprintf("\x1b[%dm%s\x1b[0m", cyan, "TRACE"),
}

type Formatter struct {
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	// [time] [level] [code] [message] [metadata]
	// example:
	// [2015-03-26T01:27:38-04:00] [INFO] [0011] [module is up] [moudule_name="test" hi="hi"]
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	code := entry.Data[Code]
	delete(entry.Data, Code)

	var level string
	if isColored {
		level = LevelsStringColored[entry.Level]
	} else {
		level = LevelsString[entry.Level]
	}
	metadata := formatMap(entry.Data)

	_, _ = fmt.Fprintf(b, "[%v] [%s] [%04d] [%s] [%s]\n",
		entry.Time.Format(FullDateFormat), level, code, entry.Message, metadata)
	return b.Bytes(), nil
}

func formatMap(data map[string]interface{}) string {
	if len(data) == 0 {
		return ""
	}
	var builder strings.Builder
	for k, v := range data {
		builder.WriteString(fmt.Sprintf(`%v="%v" `, k, v))
	}
	return builder.String()[:builder.Len()-1]
}
