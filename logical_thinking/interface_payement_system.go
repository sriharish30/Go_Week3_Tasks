package main

import "fmt"

type Payment interface {
	Pay(amount float64)
}

type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card\n", amount)
}

type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal\n", amount)
}

func main() {

	creditCard := CreditCard{CardNumber: "1234-5678-9012-3456"}
	payPal := PayPal{Email: "user@example.com"}

	payments := []Payment{creditCard, payPal}
	for _, payment := range payments {
		payment.Pay(100)
		payment.Pay(789)
	}

}
