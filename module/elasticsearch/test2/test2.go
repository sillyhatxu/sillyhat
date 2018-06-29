package main

import (
	"fmt"
	"log"
	"os"
	"github.com/olivere/elastic"
	"context"
)

// Connection parameters
const (
	URL     = "http://172.28.2.22:9200"
	INDEX   = "deja_products"
	MAPPING = "tags"
)

func main() {
	ctx := context.Background()
	// Configuration
	cfg := []elastic.ClientOptionFunc{
		elastic.SetURL(URL),

		elastic.SetSniff(false),

		elastic.SetInfoLog(log.New(os.Stdout, "ES-INFO: ", 0)),
		elastic.SetTraceLog(log.New(os.Stdout, "ES-TRACE: ", 0)),
		elastic.SetErrorLog(log.New(os.Stdout, "ES-ERROR: ", 0)),
	}

	// New Client
	client, err := elastic.NewClient(cfg...)
	if err != nil {
		panic(err)
	}

	query := client.Search().Index(INDEX).Type(MAPPING)

	result, err := query.Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Document count:", result.Hits.TotalHits)
	fmt.Println()
}