package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"bytes"
	"time"
)

type TestKafka struct {
	theadName string   `json:"thead_name"`
	value  int   `json:"value"`
}

func doPut(url string,theadName string,value int) {
	client := &http.Client{}
	//testKafka := &TestKafka{theadName: "Rocky",value: 5454}
	//testKafka := new(TestKafka)
	//testKafka.theadName = "test"
	//testKafka.value = 0
	params := fmt.Sprintf("{\"thead_name\":\"%v\",\"value\":%v}", theadName,value)
	var jsonStr = []byte(params)
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
		//fmt.Println("   ", response.StatusCode)
		if(response.StatusCode == 200){
			fmt.Println(string(contents))
		}else{
			log.Println("error")
		}
		//hdr := response.Header
		//for key, value := range hdr {
		//	fmt.Println("   ", key, ":", value)
		//}

	}
}

func loop(theadName string,url string) {
	for i := 1;i <= 100000 ;i++  {
		log.Printf("theadName : %v; value : %v \n",theadName,i)
		doPut(url,theadName,i)
	}
}

func main() {
	go loop("AAAAA","http://localhost:12100/client/test_kafka")
	go loop("BBBBB","http://localhost:12100/client/test_kafka")
	go loop("CCCCC","http://localhost:12100/client/test_kafka")
	go loop("DDDDD","http://localhost:12100/client/test_kafka")
	go loop("EEEEE","http://localhost:12100/client/test_kafka")
	//curl -X PUT --header 'Content-Type: application/json' --header 'Accept: application/json' -d '{ "thead_name": "string", "value": 0 }' 'http://localhost:12100/client/test_kafka'
	time.Sleep(time.Second*60) // 停顿一秒
}
