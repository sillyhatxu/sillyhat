package backups

import (
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"log"
	"bytes"
	"encoding/json"
)

func main() {
	u, _ := url.Parse("http://cloud-dt.deja.fashion/eureka/health")
	//{"description":"Spring Cloud Eureka Discovery Client","status":"UP"}
	//q := u.Query()
	//q.Set("username", "user")
	//q.Set("password", "passwd")
	//u.RawQuery = q.Encode()
	res, err := http.Get(u.String());
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	fmt.Println("1:",bytes.NewBuffer(result).String())
	res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("3:",result)
	fmt.Println("--------------")
	fmt.Printf("%s", result)
	fmt.Println("--------------")

	fmt.Println(decoder)

}
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/zuul/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/eureka/health'


//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/app-config/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/auth/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/cashback/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/customer/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/favourite/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/id-generator/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/inventory/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/invoice/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/legacy-db/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/message/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/ocb-syncer/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/ocr/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/order/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/payment/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/scheduler/health'
//curl -X GET --header 'Accept: application/json' 'http://cloud-dt.deja.fashion/shop/health'