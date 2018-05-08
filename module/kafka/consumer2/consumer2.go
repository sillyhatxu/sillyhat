package main

import (
	"fmt"
	"strings"
	"time"
	"github.com/bsm/sarama-cluster" //support automatic consumer-group rebalancing and offset tracking
	"github.com/golang/glog"
	"github.com/Shopify/sarama"
)


func main() {
	groupID:= "ocb-syncer"
	topicList:= "MessageCommandLocal"
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	//config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	c, err := cluster.NewConsumer(strings.Split("172.28.2.22:9092,172.28.2.22:9091,172.28.2.22:9090", ","), groupID, strings.Split(topicList, ","), config)
	if err != nil {
		glog.Errorf("Failed open consumer: %v", err)
		return
	}
	defer c.Close()
	go func() {
		for err := range c.Errors() {
			glog.Errorf("Error: %s\n", err.Error())
		}
	}()

	go func() {
		for note := range c.Notifications() {
			glog.Infof("Rebalanced: %+v\n", note)
		}
	}()

	for msg := range c.Messages() {
		fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		fmt.Println()
		c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}
}

