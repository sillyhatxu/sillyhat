package logger

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"net"
	"github.com/bshuster-repo/logrus-logstash-hook"
)

const (
	format_time  = "2006-01-02 15:04:05"
	field_key_msg   = "message"
	field_key_level = "level"
	field_key_time  = "timestamp"
	field_module_name  = "module_name"
	field_file  = "file"
	field_line  = "line"
	field_func  = "func"
)


var moduleName string

func Initial(name string){
	moduleName = name
}
func Log() *logrus.Entry {
	logNew := logrus.New()


	hook := logrustash.New("", logrustash.DefaultFormatter(logrus.Fields{"type": "myappName"}))
	logNew.Hooks.Add(hook)

	logNew.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  field_key_time,
			logrus.FieldKeyLevel: field_key_level,
			logrus.FieldKeyMsg:   field_key_msg,
		},
		TimestampFormat: format_time,
	}
	logNew.Level = logrus.DebugLevel
	if pc, file, line, ok := runtime.Caller(1); ok {
		fName := runtime.FuncForPC(pc).Name()
		return logNew.WithField(field_file, file).WithField(field_line, line).WithField(field_func, fName).WithField(field_module_name,moduleName)
	}
	return logNew.WithField(field_module_name,moduleName)
}
