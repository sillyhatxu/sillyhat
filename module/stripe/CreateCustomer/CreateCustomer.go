package main

import (
	"fmt"
	"github.com/stripe/stripe-go"
	//"github.com/stripe/stripe-go/currency"
	"github.com/stripe/stripe-go/customer"
	"log"
)

func check(err error) {
	if err != nil {
		panic(err)
		log.Println(err)
	}
}

func main() {
	//cus_CpQKOpV3sYx8g9
	//cus_CpQLXLXMNC4M0W
	stripe.Key = "sk_test_Nrdfopidhtl6rzI72tkg1vQs"

	customerParams := &stripe.CustomerParams{
		Desc: "Customer for liam.davis@example.com",
	}
	//customerParams.SetSource("tok_mastercard") // obtained with Stripe.js
	c, err := customer.New(customerParams)
	check(err)
	fmt.Println(c)
}
