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

func (m *Metadata) Log(level logrus.Level, code int, message string) {
	fields := logrus.Fields(*m)
	fields[Code] = code
	logger.WithFields(fields).Log(level, message)
}

func (m *Metadata) Logf(level logrus.Level, code int, format string, args ...interface{}) {
	fields := logrus.Fields(*m)
	fields[Code] = code
	logger.WithFields(fields).Logf(level, format, args...)
}

func (m *Metadata) Trace(code int, message string) {
	m.Log(logrus.TraceLevel, code, message)
}

func (m *Metadata) Tracef(code int, format string, args ...interface{}) {
	m.Logf(logrus.TraceLevel, code, format, args...)
}

func (m *Metadata) Debug(code int, message string) {
	m.Log(logrus.DebugLevel, code, message)
}

func (m *Metadata) Debugf(code int, format string, args ...interface{}) {
	m.Logf(logrus.DebugLevel, code, format, args...)
}

func (m *Metadata) Info(code int, message string) {
	m.Log(logrus.InfoLevel, code, message)
}

func (m *Metadata) Infof(code int, format string, args ...interface{}) {
	m.Logf(logrus.InfoLevel, code, format, args...)
}

func (m *Metadata) Warn(code int, message string) {
	m.Log(logrus.WarnLevel, code, message)
}

func (m *Metadata) Warnf(code int, format string, args ...interface{}) {
	m.Logf(logrus.WarnLevel, code, format, args...)
}

func (m *Metadata) Error(code int, message string) {
	m.Log(logrus.ErrorLevel, code, message)
}

func (m *Metadata) Errorf(code int, format string, args ...interface{}) {
	m.Logf(logrus.ErrorLevel, code, format, args...)
}

func (m *Metadata) Panic(code int, message string) {
	m.Log(logrus.PanicLevel, code, message)
}

func (m *Metadata) Panicf(code int, format string, args ...interface{}) {
	m.Logf(logrus.PanicLevel, code, format, args...)
}

func (m *Metadata) Fatal(code int, message string) {
	m.Log(logrus.FatalLevel, code, message)
}

func (m *Metadata) Fatalf(code int, format string, args ...interface{}) {
	m.Logf(logrus.FatalLevel, code, format, args...)
}

func Log(level logrus.Level, code int, message string) {
	logger.WithFields(logrus.Fields{
		Code: code,
	}).Log(level, message)
}

func Logf(level logrus.Level, code int, format string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		Code: code,
	}).Logf(level, format, args...)
}

func Trace(code int, message string) {
	Log(logrus.TraceLevel, code, message)
}

func Tracef(code int, format string, args ...interface{}) {
	Logf(logrus.TraceLevel, code, format, args...)
}

func Debug(code int, message string) {
	Log(logrus.DebugLevel, code, message)
}

func Debugf(code int, format string, args ...interface{}) {
	Logf(logrus.DebugLevel, code, format, args...)
}

func Info(code int, message string) {
	Log(logrus.InfoLevel, code, message)
}

func Infof(code int, format string, args ...interface{}) {
	Logf(logrus.InfoLevel, code, format, args...)
}

func Warn(code int, message string) {
	Log(logrus.WarnLevel, code, message)
}

func Warnf(code int, format string, args ...interface{}) {
	Logf(logrus.WarnLevel, code, format, args...)
}

func Error(code int, message string) {
	Log(logrus.ErrorLevel, code, message)
}

func Errorf(code int, format string, args ...interface{}) {
	Logf(logrus.ErrorLevel, code, format, args...)
}

func Panic(code int, message string) {
	Log(logrus.PanicLevel, code, message)
}

func Panicf(code int, format string, args ...interface{}) {
	Logf(logrus.PanicLevel, code, format, args...)
}

func Fatal(code int, message string) {
	Log(logrus.FatalLevel, code, message)
}

func Fatalf(code int, format string, args ...interface{}) {
	Logf(logrus.FatalLevel, code, format, args...)
}
