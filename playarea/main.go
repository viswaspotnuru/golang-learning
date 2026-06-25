package main

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64) bool
}

func ProcessPayment(p PaymentProcessor, amount float64) {
	success := p.Pay(amount)
	if success {
		fmt.Println("Payment successful!")
	} else {
		fmt.Println("Payment failed!")
	}
}

type Stripe struct {
	APIKey string
}

type Paypal struct {
	Email string
}

func (s Stripe) Pay(amount float64) bool {
	fmt.Printf("Stripe using key %s to pay $%.2f\n", s.APIKey, amount)
	return true
}

func (p Paypal) Pay(amount float64) bool {
	fmt.Printf("Paypal using email %s to pay $%.2f\n", p.Email, amount)
	return true
}

func main() {
	stripePayment := Stripe{APIKey: "sk_test_123"}
	paypalPayment := Paypal{Email: "randomemail@gmail.com"}

	ProcessPayment(stripePayment, 100.0)
	ProcessPayment(paypalPayment, 200.0)

}
