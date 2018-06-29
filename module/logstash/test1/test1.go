package main

import (
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	logNew := logrus.New()
	logNew.Formatter = &logrus.JSONFormatter{}
	conn, err := net.Dial("tcp", "172.16.99.130:5000")
	if err != nil {
		logNew.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "myappName"}))

	logNew.Hooks.Add(hook)


	ctx := logNew.WithFields(logrus.Fields{
		"method": "main",
		"level": "info",
		"module_name": "test",
		"thread_name": "test thread_name",
	})
	ctx.Info("Hello World!")
}

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