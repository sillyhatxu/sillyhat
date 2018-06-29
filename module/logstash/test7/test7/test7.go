package test7

import (
	"github.com/sirupsen/logrus"
	"sillyhat/module/logstash/test7/lfshook"
)

var Log *logrus.Logger

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "/Users/xushikuan/golang-log/info.log",
		logrus.ErrorLevel: "/Users/xushikuan/golang-log/error.log",
	}

	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return Log
}

