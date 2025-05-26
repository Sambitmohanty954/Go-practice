package main

import "fmt"

// enumerated types
// type orderStatus int
type orderStatus string

// this will give me INtger output
//const (
//	Received orderStatus = iota
//	Prepared
//	Delivered
//	Done
//)

const (
	Received  orderStatus = "received"
	Prepared              = "prepared"
	Delivered             = "delivered"
	Done                  = "done"
)

func changedStatus(status orderStatus) {
	fmt.Println("order is in", status, "state")
}
func main() {
	changedStatus(Done)
}
