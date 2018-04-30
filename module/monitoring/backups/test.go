package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://cloud-dt.deja.fashion/eureka/health"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}