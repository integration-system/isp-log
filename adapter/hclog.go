package adapter

import (
	"fmt"
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
	isplog "github.com/integration-system/isp-log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

var convertHcLogLevel = map[hclog.Level]logrus.Level{
	hclog.Trace:   logrus.TraceLevel,
	hclog.Debug:   logrus.DebugLevel,
	hclog.Info:    logrus.InfoLevel,
	hclog.Warn:    logrus.WarnLevel,
	hclog.Error:   logrus.ErrorLevel,
	hclog.NoLevel: logrus.InfoLevel,
}

func NewHcLogger(name string, code int) hclog.Logger {
	return &hcLogger{name: name, code: code}
}

type hcLogger struct {
	code    int
	name    string
	implied []interface{}
}

func (l *hcLogger) Log(level hclog.Level, msg string, args ...interface{}) {
	l.log(convertHcLogLevel[level], msg, args...)
}

func (l *hcLogger) Trace(msg string, args ...interface{}) {
	l.log(logrus.TraceLevel, msg, args...)
}

func (l *hcLogger) Debug(msg string, args ...interface{}) {
	l.log(logrus.DebugLevel, msg, args...)
}

func (l *hcLogger) Info(msg string, args ...interface{}) {
	l.log(logrus.InfoLevel, msg, args...)
}

func (l *hcLogger) Warn(msg string, args ...interface{}) {
	l.log(logrus.WarnLevel, msg, args...)
}

func (l *hcLogger) Error(msg string, args ...interface{}) {
	l.log(logrus.ErrorLevel, msg, args...)
}

func (l *hcLogger) IsTrace() bool {
	return isplog.IsLevelEnabled(logrus.TraceLevel)
}

func (l *hcLogger) IsDebug() bool {
	return isplog.IsLevelEnabled(logrus.DebugLevel)
}

func (l *hcLogger) IsInfo() bool {
	return isplog.IsLevelEnabled(logrus.InfoLevel)
}

func (l *hcLogger) IsWarn() bool {
	return isplog.IsLevelEnabled(logrus.WarnLevel)
}

func (l *hcLogger) IsError() bool {
	return isplog.IsLevelEnabled(logrus.ErrorLevel)
}

func (l *hcLogger) ImpliedArgs() []interface{} {
	return l.implied
}

func (l *hcLogger) With(args ...interface{}) hclog.Logger {
	sl := *l
	sl.implied = append(sl.implied, args...)
	return &sl
}

func (l *hcLogger) Name() string {
	return l.name
}

func (l *hcLogger) Named(name string) hclog.Logger {
	sl := *l
	if sl.name != "" {
		sl.name = sl.name + "." + name
	} else {
		sl.name = name
	}
	return &sl
}

func (l *hcLogger) ResetNamed(name string) hclog.Logger {
	sl := *l
	sl.name = name
	return &sl
}

// Skip
func (l *hcLogger) SetLevel(level hclog.Level) {}

func (l *hcLogger) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	return log.New(isplog.GetOutput(), "", 0)
}

func (l *hcLogger) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	return isplog.GetOutput()
}

func (l *hcLogger) log(level logrus.Level, msg string, args ...interface{}) {
	args = append(l.implied, args...)
	if l.name != "" {
		msg = fmt.Sprintf("%s: %s", l.name, msg)
	}
	if len(args) != 0 {
		metadata := make(map[string]interface{}, len(args)/2)
		for i := 0; i < len(args)-1; i += 2 {
			k := cast.ToString(args[i])
			metadata[k] = args[i+1]
		}
		isplog.WithMetadata(metadata).Log(level, l.code, msg)
	} else {
		isplog.Log(level, l.code, msg)
	}
}
