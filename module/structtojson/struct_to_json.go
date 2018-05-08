package main

import (
	"encoding/json"
	"fmt"
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
//&SlackDetail{Title: "",Value: ""}
func main() {
	slack := &Slack{Pretext: "hello world,hahahahahahahahahahahahaha",MessageUrl: "https://hooks.slack.com/services/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",Time: 54645645645,Color: "#FF0000",SlackDetailArray: []SlackDetail {
		{Title: "a",Value: "multiple",}, {Title: "b", Value: "b-option",},
	}}
	e, err := json.Marshal(slack)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
}