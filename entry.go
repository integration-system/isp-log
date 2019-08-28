package log

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Time    time.Time
	Level   string
	Code    int
	Message string
	Data    map[string]string
}

var logRe = regexp.MustCompile(`(?mU)\[(.*)\]`)
var logDataRe = regexp.MustCompile(`(?mU)(.+)="(.+)"`)

func ParseLog(log string) (Entry, error) {
	var entry Entry
	params := logRe.FindAllStringSubmatch(log, -1)
	if len(params) != 5 {
		return Entry{}, errors.New("invalid string format")
	}
	ti, err := time.Parse(FullDateFormat, params[0][1])
	if err != nil {
		return Entry{}, err
	}
	entry.Time = ti
	entry.Level = params[1][1]
	i, err := strconv.Atoi(params[2][1])
	if err != nil {
		return Entry{}, err
	}
	entry.Code = i
	entry.Message = params[3][1]

	entry.Data = map[string]string{}
	for _, match := range logDataRe.FindAllStringSubmatch(params[4][1], -1) {
		key := strings.TrimSpace(match[1])
		entry.Data[key] = match[2]
	}
	return entry, nil
}
