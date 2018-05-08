package main

import (
	"fmt"
	"encoding/json"
)

type Slack struct {

	Pretext string `json:"pretext"`

	MessageUrl string `json:"messageUrl"`

	Time int64 `json:"time"`

	Color string `json:"color"`

	SlackDetailArray []SlackDetail
}

type SlackDetail struct {

	Title string

	Value string

}
func main() {
	var slack Slack
	jsonSrcSrc := "{\"pretext\":\"hello world,hahahahahahahahahahahahaha\",\"messageUrl\":\"https://hooks.slack.com/services/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\",\"time\":54645645645,\"color\":\"#FF0000\",\"SlackDetailArray\":[{\"Title\":\"a\",\"Value\":\"multiple\"},{\"Title\":\"b\",\"Value\":\"b-option\"}]}"
	fmt.Println(jsonSrcSrc)

	json.Unmarshal([]byte(jsonSrcSrc), &slack)
	fmt.Println(slack)
	fmt.Println(slack.Pretext)
	fmt.Println(slack.Color)
	fmt.Println(slack.MessageUrl)
	fmt.Println(slack.Time)
	for i:= 0;i < len(slack.SlackDetailArray);i++{
		slackDetail := slack.SlackDetailArray[i]
		fmt.Println(slackDetail.Title + "----" + slackDetail.Value)
	}

}
