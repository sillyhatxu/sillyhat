package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strconv"
	"os"
	"fmt"
)

func getValue(threadName string,i int) string {
	return `{"commandName":"sendSlack","commandBody":"{\"slack\":{\"pretext\":\"` + threadName+`---`+strconv.Itoa(i) + `\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"time\":1526630475,\"color\":\"#36a64f\",\"slackDetailList\":[{\"title\":\"title1\",\"value\":\"value1\"},{\"title\":\"title2\",\"value\":\"value2\"},{\"title\":\"title3\",\"value\":\"value3\"},{\"title\":\"title4\",\"value\":\"value4\"}]}}"}`
}

func main() {

	//if len(os.Args) != 3 {
	//	fmt.Fprintf(os.Stderr, "Usage: %s <broker> <topic>\n",
	//		os.Args[0])
	//	os.Exit(1)
	//}

	broker := "172.28.2.22:9092,172.28.2.22:9091,172.28.2.22:9090"
	//topic := "MessageCommandLocal"

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created Producer %v\n", p)
	//
	//doneChan := make(chan bool)
	//
	//go func() {
	//	defer close(doneChan)
	//	for e := range p.Events() {
	//		switch ev := e.(type) {
	//		case *kafka.Message:
	//			m := ev
	//			if m.TopicPartition.Error != nil {
	//				fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	//			} else {
	//				fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
	//					*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	//			}
	//			return
	//
	//		default:
	//			fmt.Printf("Ignored event: %s\n", ev)
	//		}
	//	}
	//}()
	//
	//value := getValue("testA",1)
	//p.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(value)}
	//
	//// wait for delivery report goroutine to finish
	//_ = <-doneChan
	//
	//p.Close()
}