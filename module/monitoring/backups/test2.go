package backups

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

type Tracks struct {
	Toptracks Toptracks_info
}

type Toptracks_info struct {
	Track []Track_info
	Attr  Attr_info `json: "@attr"`
}

type Track_info struct {
	Name       string
	Duration   string
	Listeners  string
	Mbid       string
	Url        string
	Streamable Streamable_info
	Artist     Artist_info
	Attr       Track_attr_info `json: "@attr"`
}
type Streamable_info struct {
	Text      string `json: "#text"`
	Fulltrack string
}

type Artist_info struct {
	Name string
	Mbid string
	Url  string
}

type Track_attr_info struct {
	Rank string
}
type Attr_info struct {
	Country    string
	Page       string
	PerPage    string
	TotalPages string
	Total      string
}

func get_content() {
	// json data
	url := "http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&api_key=c1572082105bd40d247836b5c1819623&format=json&country=Netherlands"
	url += "&limit=1" // limit data for testing
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
	get_content()
}