package testpublic

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

const (
	format_time  = "2006-01-02 15:04:05"
	field_key_msg   = "message"
	field_key_level = "level"
	field_key_time  = "timestamp"
)


var moduleName string

func decorateRuntimeContext(logger *logrus.Entry,moduleName string) *logrus.Entry {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fName := runtime.FuncForPC(pc).Name()
		return logger.WithField("file", file).WithField("line", line).WithField("func", fName).WithField("module_name",moduleName)
	} else {
		return logger
	}
}

func Initial(name string){
	moduleName = name
}
func logrusConfig() *logrus.Entry {
	logNew := logrus.New()
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
		return logNew.WithField("file", file).WithField("line", line).WithField("func", fName).WithField("module_name",moduleName)
	}
	return logNew.WithField("module_name",moduleName)
}



func Debug(args ...interface{}) {
	logrusConfig().Debug(args...)
}

func Print(args ...interface{}) {
	logrusConfig().Print(args...)
}

func Info(args ...interface{}) {
	logrusConfig().Info(args...)
}

func Warn(args ...interface{}) {
	logrusConfig().Warn(args...)
}

func Warning(args ...interface{}) {
	logrusConfig().Warning(args...)
}

func Error(args ...interface{}) {
	logrusConfig().Error(args...)
}

func Fatal(args ...interface{}) {
	logrusConfig().Fatal(args...)
}

func Panic(args ...interface{}) {
	logrusConfig().Panic(args...)
}

//@timestamp	       	June 4th 2018, 22:16:23.303
//t  level	       	INFO
//t  message	       	2018-06-04 22:16:21.290 - INFO [inventory][io-12060-exec-1] f.d.i.controller.InventoryController     [   109]: URI(queryShopItemInventoryByShopItemGroupId) return ShopItemInventoryList : [InventoryItemDTO(id=115701, shopItemId=5457598, shopItemGroupId=lzzie_3996, size=S, quantity=10), InventoryItemDTO(id=115702, shopItemId=5457599, shopItemGroupId=lzzie_3996, size=S, quantity=10), InventoryItemDTO(id=181236, shopItemId=5457598, shopItemGroupId=lzzie_3996, size=M, quantity=10), InventoryItemDTO(id=181237, shopItemId=5457599, shopItemGroupId=lzzie_3996, size=M, quantity=10), InventoryItemDTO(id=246771, shopItemId=5457598, shopItemGroupId=lzzie_3996, size=L, quantity=10), InventoryItemDTO(id=246772, shopItemId=5457599, shopItemGroupId=lzzie_3996, size=L, quantity=10)]



//t  @version	       	1
//t  TIMESTAP	       	2018-06-04 22:16:21.290
//t  _id	      	LAEny2MBOaSwOwOS9IgU
//t  _index	      	logstash-2018.06.04
//#  _score	    	 -
//t  _type	      	doc
//t  host	       	10.255.0.2
//t  java_file	       	f.d.i.controller.InventoryController

//t  line	       	   109
//t  message_format	       	URI(queryShopItemInventoryByShopItemGroupId) return ShopItemInventoryList : [InventoryItemDTO(id=115701, shopItemId=5457598, shopItemGroupId=lzzie_3996, size=S, quantity=10), InventoryItemDTO(id=115702, shopItemId=5457599, shopItemGroupId=lzzie_3996, size=S, quantity=10), InventoryItemDTO(id=181236, shopItemId=5457598, shopItemGroupId=lzzie_3996, size=M, quantity=10), InventoryItemDTO(id=181237, shopItemId=5457599, shopItemGroupId=lzzie_3996, size=M, quantity=10), InventoryItemDTO(id=246771, shopItemId=5457598, shopItemGroupId=lzzie_3996, size=L, quantity=10), InventoryItemDTO(id=246772, shopItemId=5457599, shopItemGroupId=lzzie_3996, size=L, quantity=10)]
//t  module_name	       	inventory
//#  port	       	57,902
//t  thread_name	       	io-12060-exec-1
//t  type	       	syslog