package log

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	var buf bytes.Buffer
	SetLevel("TRACE")
	SetOutput(&buf)
	Debug(4, "mes sage")
	suffix := `[DEBUG] [0004] [mes sage] []`
	res := strings.TrimSpace(buf.String())

	assert.True(t, strings.HasSuffix(res, suffix))
}

func TestLog2(t *testing.T) {
	var buf bytes.Buffer
	SetLevel("TRACE")
	SetOutput(&buf)
	Trace(4, "mes sage")
	suffix := `[TRACE] [0004] [mes sage] []`
	res := strings.TrimSpace(buf.String())

	assert.True(t, strings.HasSuffix(res, suffix))
}

func TestLog3(t *testing.T) {
	var buf bytes.Buffer
	SetLevel("TRACE")
	SetOutput(&buf)
	assert.True(t, true)
	Error(4, "mes sage")
	suffix := `[ERROR] [0004] [mes sage] []`
	res := strings.TrimSpace(buf.String())

	assert.True(t, strings.HasSuffix(res, suffix))
}

func TestLogf(t *testing.T) {
	var buf bytes.Buffer
	SetLevel("TRACE")
	SetOutput(&buf)
	Infof(3, "mes %v", 45)
	suffix := `[INFO ] [0003] [mes 45] []`
	res := strings.TrimSpace(buf.String())

	assert.True(t, strings.HasSuffix(res, suffix))
}

func TestWithMetadata(t *testing.T) {
	var buf bytes.Buffer
	SetLevel("TRACE")
	SetOutput(&buf)
	WithMetadata(Metadata{
		"integer": 456,
	}).Warnf(39, "mes %v", 45)
	suffix := `[WARN ] [0039] [mes 45] [integer="456"]`
	res := strings.TrimSpace(buf.String())

	assert.True(t, strings.HasSuffix(res, suffix))
}
