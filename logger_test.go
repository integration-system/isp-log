package log

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLog(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Info(3, 4, "mes sage")
	suffix := `[INFO] [0003][0004] [mes sage] []`
	res := strings.TrimSpace(buf.String())

	assert.True(t, strings.HasSuffix(res, suffix))
}

func TestLogf(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Infof(3, 4, "mes %v", 45)
	suffix := `[INFO] [0003][0004] [mes 45] []`
	res := strings.TrimSpace(buf.String())

	assert.True(t, strings.HasSuffix(res, suffix))

}

func TestWithMetadata(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	WithMetadata(Metadata{
		"integer": 456,
	}).Infof(3, 4, "mes %v", 45)
	suffix := `[INFO] [0003][0004] [mes 45] [integer="456"]`
	res := strings.TrimSpace(buf.String())

	assert.True(t, strings.HasSuffix(res, suffix))

}
