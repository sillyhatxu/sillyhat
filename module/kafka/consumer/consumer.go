package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	wg     sync.WaitGroup
	logger = log.New(os.Stderr, "[srama]", log.LstdFlags)
)

func main() {
	topic:="MessageCommandLocal"
	sarama.Logger = logger

	//groupID := "ocb-syncer"

	//leaveGroupRequest := &sarama.LeaveGroupRequest{
	//	GroupId:"",
	//}
	//sarama.LeaveGroupRequest{
	//	GroupId:"",
	//}
	config := sarama.NewConfig()
	//config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	//sarama.JoinGroupRequest{
	//
	//}

	//consumer, err := sarama.NewConsumer(strings.Split("172.28.2.22:9092,172.28.2.22:9091,172.28.2.22:9090", ","), groupID, strings.Split(topics, ","), config)

	consumer, err := sarama.NewConsumer(strings.Split("172.28.2.22:9092,172.28.2.22:9091,172.28.2.22:9090", ","), config)
	if err != nil {
		logger.Println("Failed to start consumer: %s", err)
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		logger.Println("Failed to get the list of partitions: ", err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			logger.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
		}
		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}
		}(pc)
	}
	wg.Wait()

	logger.Printf("Done consuming topic %v\n",topic)
	consumer.Close()
}