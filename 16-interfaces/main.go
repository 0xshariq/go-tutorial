package main

import "fmt"

// defining an interface
type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

// defining a struct that uses the interface
type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

func (p payment) makeRefund(amount float32, account string) {
	p.gateway.refund(amount, account)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}
func (r razorpay) refund(amount float32, account string) {
	fmt.Println("refunding amount using razorpay", amount, "to account", account)
}

func main() {
	razorpayPaymentGw := razorpay{}
	newPayment := payment{
		gateway: razorpayPaymentGw,
	}
	newPayment.makePayment(100)
	newPayment.makeRefund(50, "user_account_123")
}
