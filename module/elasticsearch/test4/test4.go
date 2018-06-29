package main

import (
	"context"
	"github.com/olivere/elastic"
	"time"
	"os"
	"log"
	"reflect"
	"encoding/json"
)

var ctx context.Context
var client *elastic.Client

const (
	URL     = "http://172.28.2.22:9200"
	INDEX   = "deja_products"
	MAPPING = "tags"
)

func init() {
	ctx = context.Background()
	// Obtain a client and connect to the default Elasticsearch installation on 127.0.0.1:9200. Of course you can configure your client to connect to other hosts and configure it in various other ways.
	thisClient, err := elastic.NewClient(
		elastic.SetURL(URL),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		// Handle error
		panic(err)
	}
	client = thisClient
}

func search(){
	termQuery := elastic.NewTermQuery("brand_id", 2007)
	//termQuery := elastic.NewTermQuery("product_name", "ELAINE DRESS")
	//termQuery := elastic.NewTermQuery("product_id", "5429755")
	searchResult, err := client.Search().
		Index(INDEX).   // search in index "twitter"
		Query(termQuery).   // specify the query
		Sort("product_id", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	log.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	log.Println("Document count:", searchResult.Hits.TotalHits)

	//query := client.Search().Index(INDEX).Type(MAPPING)
	//result, err := query.Do(ctx)
	//if err != nil {
	//	panic(err)
	//}

	var ttyp Product
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Product); ok {
			log.Printf("Product by productId : %v ; productCode : %v\n", t.ProductId, t.ProductCode)
		}
	}
	// TotalHits is another convenience function that works even when something goes wrong.
	log.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	// Here's how you iterate through results with full control over each step.
	if searchResult.Hits.TotalHits > 0 {
		log.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var t Product
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			log.Printf("Tweet by %s: %s\n", t.ProductId, t.ProductCode)
		}
	} else {
		// No hits
		log.Print("Found no tweets\n")
	}
}

func main() {
	search()
}


//@ElasticSearchField(primaryKey = true)
//private String id;
//
//@ElasticSearchField(value = "product_id")
//private Long productId;
//
//@ElasticSearchField(value = "product_code")
//private String productCode;
//
//@ElasticSearchField(value = "product_group_id")
//private String productGroupId;
//
//@ElasticSearchField(value = "product_name")
//private String productName;
//
//@ElasticSearchField(value = "brand_id")
//private Long brandId;
//
//@ElasticSearchField(value = "brand_name")
//private String brandName;
//
//@ElasticSearchField(value = "category")
//private Integer category;
//
//@ElasticSearchField(value = "subcategory")
//private Integer subcategory;
//
//@ElasticSearchField(value = "color")
//private String color;
//
//@ElasticSearchField(value = "color_and_pattern")
//private Integer colorAndPattern;
//
//@ElasticSearchField(value = "ocb")
//private boolean ocb;
//
//@ElasticSearchField(value = "is_discount")
//private boolean isDiscount;
//
//@ElasticSearchField(value = "is_new_arrival")
//private boolean isNewArrival;
//
//@ElasticSearchField(value = "all_size")
//private boolean allSize;
//
//@ElasticSearchField(value = "is_purchasable")
//private boolean isPurchasable;
//
//@ElasticSearchField(value = "is_recommend")
//private boolean isRecommend;
//
//@ElasticSearchField(value = "image_url")
//private String imageUrl;
//
//@ElasticSearchField(value = "height")
//private Long height;
//
//@ElasticSearchField(value = "width")
//private Long width;
//
//@ElasticSearchField(value="original_price")
//private Long originalPrice;
//
//@ElasticSearchField(value="current_price")
//private Long currentPrice;
//
//@ElasticSearchField(value="deja_price")
//private Long deja_price;
//
//@ElasticSearchField(value="currency")
//private String currency;
//
//@ElasticSearchField(value = "recommend_reason")
//private String recommendReason;
//
//@ElasticSearchField(value = "weight")
//private Integer weight;
//
//@ElasticSearchField(value = "update_time")
//private Date updateTime;
type Product struct {
	ProductId     int64  `json:"product_id"`
	ProductCode   string `json:"product_code"`
	ProductNAME   string `json:"product_name"`
	CurrentPrice  int64  `json:"currentPrice"`
	Category      int    `json:"category"`
	IsPurchasable bool   `json:"is_purchasable"`
	Is_recommend  bool   `json:"is_recommend"`
	UpdateTime    int64  `json:"update_time"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"image":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`