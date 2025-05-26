package main

import "fmt"

type paymenter interface {
	pay(amount float64)
	refund(amount float64, account string)
}

type Payment struct {
	gateway paymenter
}

func (p Payment) makePayment(amount float64) {
	//razorpayPaymentGW := Razorpay{}
	//razorpayPaymentGW.pay(amount)
	//stripeGw := stripe{}
	//stripeGw.makeStripe(amount)
	p.gateway.pay(amount)
}

type Razorpay struct {
}

type stripe struct{}

func (s stripe) pay(amount float64) {
	fmt.Println("making payment using Stripe", amount)
}

func (r Razorpay) pay(amount float64) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

type faKePayment struct{}

func (f faKePayment) pay(amount float64) {
	fmt.Println("making payment using faKePayment", amount)
}

func (f faKePayment) refund(amount float64, account string) {

}
func main() {
	//stripeGw := stripe{}
	//razorpayPaymentGW := Razorpay{}
	fakePW := faKePayment{}
	newPayment := Payment{
		gateway: fakePW,
	}
	newPayment.makePayment(100)
}
