package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Metadata logrus.Fields

var logger *logrus.Logger

func init() {
	logger = &logrus.Logger{
		Out:          os.Stdout,
		Formatter:    new(Formatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
}

func SetLevel(level string) error {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logger.SetLevel(lvl)
	return nil
}

func SetOutput(output io.Writer) {
	logger.SetOutput(output)
}

func WithMetadata(metadata Metadata) *Metadata {
	return &metadata
}

func (m *Metadata) Log(level logrus.Level, group, code int, message string) {
	fields := logrus.Fields(*m)
	fields[Group] = group
	fields[Code] = code
	logger.WithFields(fields).Log(level, message)
}

func (m *Metadata) Logf(level logrus.Level, group, code int, format string, args ...interface{}) {
	fields := logrus.Fields(*m)
	fields[Group] = group
	fields[Code] = code
	logger.WithFields(fields).Logf(level, format, args...)
}

func (m *Metadata) Trace(group, code int, message string) {
	m.Log(logrus.TraceLevel, group, code, message)
}

func (m *Metadata) Tracef(group, code int, format string, args ...interface{}) {
	m.Logf(logrus.TraceLevel, group, code, format, args...)
}

func (m *Metadata) Debug(group, code int, message string) {
	m.Log(logrus.DebugLevel, group, code, message)
}

func (m *Metadata) Debugf(group, code int, format string, args ...interface{}) {
	m.Logf(logrus.DebugLevel, group, code, format, args...)
}

func (m *Metadata) Info(group, code int, message string) {
	m.Log(logrus.InfoLevel, group, code, message)
}

func (m *Metadata) Infof(group, code int, format string, args ...interface{}) {
	m.Logf(logrus.InfoLevel, group, code, format, args...)
}

func (m *Metadata) Warn(group, code int, message string) {
	m.Log(logrus.WarnLevel, group, code, message)
}

func (m *Metadata) Warnf(group, code int, format string, args ...interface{}) {
	m.Logf(logrus.WarnLevel, group, code, format, args...)
}

func (m *Metadata) Error(group, code int, message string) {
	m.Log(logrus.ErrorLevel, group, code, message)
}

func (m *Metadata) Errorf(group, code int, format string, args ...interface{}) {
	m.Logf(logrus.ErrorLevel, group, code, format, args...)
}

func (m *Metadata) Panic(group, code int, message string) {
	m.Log(logrus.PanicLevel, group, code, message)
}

func (m *Metadata) Panicf(group, code int, format string, args ...interface{}) {
	m.Logf(logrus.PanicLevel, group, code, format, args...)
}

func (m *Metadata) Fatal(group, code int, message string) {
	m.Log(logrus.FatalLevel, group, code, message)
}

func (m *Metadata) Fatalf(group, code int, format string, args ...interface{}) {
	m.Logf(logrus.FatalLevel, group, code, format, args...)
}

func Log(level logrus.Level, group, code int, message string) {
	logger.WithFields(logrus.Fields{
		Group: group,
		Code:  code,
	}).Log(level, message)
}

func Logf(level logrus.Level, group, code int, format string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		Group: group,
		Code:  code,
	}).Logf(level, format, args...)
}

func Trace(group, code int, message string) {
	Log(logrus.TraceLevel, group, code, message)
}

func Tracef(group, code int, format string, args ...interface{}) {
	Logf(logrus.TraceLevel, group, code, format, args...)
}

func Debug(group, code int, message string) {
	Log(logrus.DebugLevel, group, code, message)
}

func Debugf(group, code int, format string, args ...interface{}) {
	Logf(logrus.DebugLevel, group, code, format, args...)
}

func Info(group, code int, message string) {
	Log(logrus.InfoLevel, group, code, message)
}

func Infof(group, code int, format string, args ...interface{}) {
	Logf(logrus.InfoLevel, group, code, format, args...)
}

func Warn(group, code int, message string) {
	Log(logrus.WarnLevel, group, code, message)
}

func Warnf(group, code int, format string, args ...interface{}) {
	Logf(logrus.WarnLevel, group, code, format, args...)
}

func Error(group, code int, message string) {
	Log(logrus.ErrorLevel, group, code, message)
}

func Errorf(group, code int, format string, args ...interface{}) {
	Logf(logrus.ErrorLevel, group, code, format, args...)
}

func Panic(group, code int, message string) {
	Log(logrus.PanicLevel, group, code, message)
}

func Panicf(group, code int, format string, args ...interface{}) {
	Logf(logrus.PanicLevel, group, code, format, args...)
}

func Fatal(group, code int, message string) {
	Log(logrus.FatalLevel, group, code, message)
}

func Fatalf(group, code int, format string, args ...interface{}) {
	Logf(logrus.FatalLevel, group, code, format, args...)
}
