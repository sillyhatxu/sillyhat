package main

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

const format_time  = "2006-01-02 15:04:05"

func DecorateRuntimeContext(logger *logrus.Entry) *logrus.Entry {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fName := runtime.FuncForPC(pc).Name()
		return logger.WithField("file", file).WithField("line", line).WithField("func", fName)
	} else {
		return logger
	}
}

func main() {
	Log := DecorateRuntimeContext // so I don't have to type it out every time
	Log(logrus.WithFields(logrus.Fields{
		"key": "value",
	})).Info("message")


	Log(logrus.WithFields(logrus.Fields{
		"key": "value",
	})).Info("11111")



	Log(logrus.WithFields(logrus.Fields{
		"key": "value",
	})).Info("222222")



	logNew := logrus.New()
	logNew.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "TIMESTAMP",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: format_time,
	}
	ctx := logNew.WithFields(logrus.Fields{
		"method": "main",
		"module_name": "test",
		"thread_name": "test thread_name",
	})
	ctx.Info("Hello World!")
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



//t  TIMESTAP	       	2018-06-01 15:32:45.954
//t  host	       	10.255.0.2
//t  java_file	       	f.d.l.m.h.e.DeviceUpdatedEventHandler
//t  level	       	INFO
//t  line	       	    35
//t  message	       	2018-06-01 15:32:45.954 - INFO [legacy-db][          -L-25] f.d.l.m.h.e.DeviceUpdatedEventHandler    [    35]: DeviceUpdatedEventHandler body : DeviceUpdatedEvent(deviceMessageDTO=DeviceMessageDTO(uid=83032, deviceId=a7c678d3380d547113569f7c37ec71af, rawUa=iOS/11.1.2 CiOS/1804181055 Encoding/UTF-8 Lang/zh-Hans-SG Morange/6.4.4 Caps/0 PI/a7c678d3380d547113569f7c37ec71af Domain/(null) DeviceBrand/Apple DeviceModel/iPhone_6 DeviceVersion/11.1.2 ClientType/CiOS ClientBuild/6.4.4.1804181055 appID/com.dejafashion.test ScreenWidth/750 ScreenHeight/1334 Mcc/525 Mnc/01, source=CiOS, appType=5, buildVersion=6.4.4.1804181055, ip=122.11.173.111))
//t  message_format	       	DeviceUpdatedEventHandler body : DeviceUpdatedEvent(deviceMessageDTO=DeviceMessageDTO(uid=83032, deviceId=a7c678d3380d547113569f7c37ec71af, rawUa=iOS/11.1.2 CiOS/1804181055 Encoding/UTF-8 Lang/zh-Hans-SG Morange/6.4.4 Caps/0 PI/a7c678d3380d547113569f7c37ec71af Domain/(null) DeviceBrand/Apple DeviceModel/iPhone_6 DeviceVersion/11.1.2 ClientType/CiOS ClientBuild/6.4.4.1804181055 appID/com.dejafashion.test ScreenWidth/750 ScreenHeight/1334 Mcc/525 Mnc/01, source=CiOS, appType=5, buildVersion=6.4.4.1804181055, ip=122.11.173.111))
//t  module_name	       	legacy-db
//t  thread_name	       	          -L-25
//
//
//t  host	       	10.255.0.2
//t  message	       	{"@timestamp":"2018-06-01T15:35:36+08:00","@version":"1","level":"info","message":"Hello World!","method":"main","type":"myappName"}
//t  tags	       	_grokparsefailure