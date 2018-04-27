package backups

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func get_content3() {
	// json data
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
	fmt.Printf("Results: %v\n", data)
	os.Exit(0)
}

func main() {
	get_content3()
}