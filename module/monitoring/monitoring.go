package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"log"
)

type ModuleStatus struct {
	description string
	status      string
}

func main() {

	url := "http://cloud-dt.deja.fashion/app-config/health"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var data interface{} // TopTracks
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	log.Println(data)
	fmt.Printf("Results: %v\n", data)
	os.Exit(0)

}