package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseLog(t *testing.T) {
	s := `[2019-10-10T12:24:08.113+03:00] [INFO] [0004] [mes sage] [times="2019-08-27 16:31:12.680299561 +0300 MSK m=+0.001004407" name="test" integer="456"]`
	ti, _ := time.Parse(FullDateFormat, "2019-10-10T12:24:08.113+03:00")
	expected := Entry{
		Time:    ti,
		Level:   "INFO",
		Code:    4,
		Message: "mes sage",
		Data: map[string]string{
			"times":   "2019-08-27 16:31:12.680299561 +0300 MSK m=+0.001004407",
			"name":    "test",
			"integer": "456",
		},
	}
	actual, err := ParseLog(s)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}
