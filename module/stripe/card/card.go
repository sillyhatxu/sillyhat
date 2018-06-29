package main

import (
	"github.com/stripe/stripe-go/source"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/token"
	"log"
	"github.com/stripe/stripe-go/paymentsource"
)

func addCard(customerId string,cardNumber string,exeMonth string ,exeYear string,cvc string) {
	stripeToken :=getToken(cardNumber,exeMonth,exeYear,cvc)
	stripeCustomerSource, err := paymentsource.New(&stripe.CustomerSourceParams{
		Customer: customerId,
		Source: &stripe.SourceParams{Token: stripeToken.ID},
	})
	checkError(err)
	log.Println(stripeCustomerSource)
}

func stripeKey() {
	stripe.Key = "sk_test_Nrdfopidhtl6rzI72tkg1vQs"
}

func checkError(err error){
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
				log.Printf("Error -----> Card was declined with code: %v\n", cardErr.DeclineCode)
			} else {
				log.Printf("Error -----> Other Stripe error occurred: %v\n", stripeErr.Error())
			}
		} else {
			log.Printf("Error -----> Other error occurred: %v\n", err.Error())
		}
	}
}

func getById(sourceId string)  {
	source, err := source.Get(sourceId, nil)
	checkError(err)
	log.Println(source)
}


func getToken(cardNumber string,exeMonth string ,exeYear string,cvc string) *stripe.Token {
	stripeKey()
	tokenDTO, err := token.New(&stripe.TokenParams{
		Card: &stripe.CardParams{
			Number: cardNumber,
			Month:  exeMonth,
			Year:   exeYear,
			CVC:    cvc,
		},
	})
	checkError(err)
	return tokenDTO
}

func main() {
	//stripe.Key = "sk_test_Nrdfopidhtl6rzI72tkg1vQs"
	//getById("src_1C4kYiJ9z6Ktianxflu4fBEE")
	addCard("cus_CgnZlfwLCopI7A","4000056655665556","12","22","111")
}
