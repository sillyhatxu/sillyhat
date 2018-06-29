package main

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"fmt"
)

func main() {
	stripe.Key = "sk_test_Nrdfopidhtl6rzI72tkg1vQs"
	customer, err := card.Del(
		"card_1CRubFJ9z6KtianxfI9zDmDF",
		&stripe.CardParams{Customer: "cus_CqBBaHkj5iRYIM"},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
