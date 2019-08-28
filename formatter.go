package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	FullDateFormat = "2006-01-02T15:04:05.999-07:00"
	Code           = "$$code"
)

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
	level := strings.ToUpper(entry.Level.String())
	metadata := formatMap(entry.Data)

	fmt.Fprintf(b, "[%v] [%s] [%04d] [%s] [%s]\n",
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
