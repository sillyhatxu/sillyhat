package main

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"fmt"
)

func main() {
	stripe.Key = "sk_test_Nrdfopidhtl6rzI72tkg1vQs"
	customerDTO, err := customer.Del(
		"cus_CpQKOpV3sYx8g9",
		&stripe.CustomerParams{},
	)
	if err != nil {
		// Try to safely cast a generic error to a stripe.Error so that we can get at
		// some additional Stripe-specific information about what went wrong.
		if stripeErr, ok := err.(*stripe.Error); ok {
			// The Code field will contain a basic identifier for the failure.
			switch stripeErr.Code {
			case stripe.IncorrectNum:
			case stripe.InvalidNum:
			case stripe.InvalidExpM:
			case stripe.InvalidExpY:
			case stripe.InvalidCvc:
			case stripe.ExpiredCard:
			case stripe.IncorrectCvc:
			case stripe.IncorrectZip:
			case stripe.CardDeclined:
			case stripe.Missing:
			case stripe.ProcessingErr:
			}
			// The Err field can be coerced to a more specific error type with a type
			// assertion. This technique can be used to get more specialized
			// information for certain errors.
			if cardErr, ok := stripeErr.Err.(*stripe.CardError); ok {
				fmt.Printf("Card was declined with code: %v\n", cardErr.DeclineCode)
			} else {
				fmt.Printf("Other Stripe error occurred: %v\n", stripeErr.Error())
			}
		} else {
			fmt.Printf("Other error occurred: %v\n", err.Error())
		}
	}
	fmt.Println(customerDTO)
}