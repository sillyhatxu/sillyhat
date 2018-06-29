package main

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/source"
	"fmt"
)

func main() {
	stripe.Key = "sk_test_Nrdfopidhtl6rzI72tkg1vQs"
	//customerId := "cus_CqBBaHkj5iRYIM"
	sourceId := "src_1CQUr0J9z6Ktianx2AmJm5CT"
	//expMonth := "12"
	//expYear := "23"
	//cvc := "888"
	params := &stripe.SourceObjectParams{}

	//card.put("exp_month", request.getExpMonth());
	//card.put("exp_year", request.getExpYear());
	//card.put("cvc", request.getCvc());
	//cardParams.put("card", card);


	source, err := source.Update(sourceId, params)
	if err != nil{

	}
	fmt.Println(source)
}
