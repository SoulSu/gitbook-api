package log

import (
	"fmt"

	"github.com/SoulSu/gitbook-api/libs/env"
	"github.com/sirupsen/logrus"
)

const ConfigFile = "log.yaml"

type LogConfig struct {
	Formatter string
	Level     string
}

var l *logrus.Logger

func init() {
	l = logrus.New()
	cfg := LogConfig{
		Formatter: "",
		Level:     "debug",
	}
	if err := env.Env.LoadConfig(ConfigFile, &cfg); err != nil {
		fmt.Println("load file:" + err.Error())
		fmt.Println("use default")
	}
	logLevel, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		logLevel = logrus.TraceLevel
	}
	var formatter logrus.Formatter
	formatter = &logrus.TextFormatter{}
	if cfg.Formatter == "json" {
		formatter = nil
		formatter = &logrus.JSONFormatter{}
	}
	l.SetFormatter(formatter)
	l.SetLevel(logLevel)
}

func Debug(v ...interface{}) {
	l.Debug(v...)
}
func Debugf(format string, v ...interface{}) {
	l.Debugf(format, v...)
}
func Error(v ...interface{}) {
	l.Error(v...)
}
func Errorf(format string, v ...interface{}) {
	l.Errorf(format, v...)
}
func Info(v ...interface{}) {
	l.Info(v...)
}
func Infof(format string, v ...interface{}) {
	l.Infof(format, v...)
}
func Warning(v ...interface{}) {
	l.Warning(v...)
}
func Warningf(format string, v ...interface{}) {
	l.Warningf(format, v...)
}
func Fatal(v ...interface{}) {
	l.Fatal(v...)
}
func Fatalf(format string, v ...interface{}) {
	l.Fatalf(format, v...)
}
func Panic(v ...interface{}) {
	l.Panic(v...)
}
func Panicf(format string, v ...interface{}) {
	l.Panicf(format, v...)
}
