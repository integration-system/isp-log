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
	Group   int
	Code    int
	Message string
	Data    map[string]string
}

func ParseLog(log string) (Entry, error) {
	var entry Entry
	var re = regexp.MustCompile(`(?mU)\[(.*)\]`)
	params := re.FindAllStringSubmatch(log, -1)
	if len(params) != 6 {
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
	entry.Group = i
	i, err = strconv.Atoi(params[3][1])
	if err != nil {
		return Entry{}, err
	}
	entry.Code = i
	entry.Message = params[4][1]

	entry.Data = map[string]string{}
	var dataRe = regexp.MustCompile(`(?mU)(.+)="(.+)"`)

	for _, match := range dataRe.FindAllStringSubmatch(params[5][1], -1) {
		key := strings.TrimSpace(match[1])
		entry.Data[key] = match[2]
	}
	return entry, nil
}
