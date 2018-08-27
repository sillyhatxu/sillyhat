package main

import (
"encoding/json"
"fmt"
"log"
)

type ResponseControllerList struct {
	Code        int         `json:"code"`
	ApiStatus   int         `json:"api_status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data,omitempty"`
	TotalRecord interface{} `json:"total_record,omitempty"`
}

func createResponse(v interface{}) (*ResponseControllerList, error) {
	ratingsCount, ok := v.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("want type map[string]interface{};  got %T", v)
	}
	return &ResponseControllerList{
		200,
		1,
		"success",
		nil,
		ratingsCount["count"],
	}, nil
}

func main() {
	resp, err := createResponse(map[string]interface{}{"count": 1})
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
