package main

import (
	"fmt"
	"time"
)

// A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
// we can acheive object oriented programming using struct
// Imagine struct as a class
// order struct

// inheritance or called struct embbeding in Go
type customer struct {
	name string
	age  int
}

type Animal struct {
	name  string
	sound string
}

func (d *Animal) soundChange(sound string) {
	d.sound = sound
}

func (d Animal) Speak() {
	fmt.Println("dog speak", d.sound)
}

type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time //nanosecond precision
	customer
}

// Constructor in Go, we can crate like below func

func newOrder(id string, amount float64, status string) *order {
	// initial setup goes here....
	myOrder2 := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &myOrder2
}

// receiver type
// Change status is method of that order struct,
// if we pass pointer beside order, then only we can able to change the value of the instance itself
// for setting the value we can use pointers but for getting the value, it's not needed.

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float64 {
	return o.amount
}

func main() {

	// Inline structs, if we want to use struct one time
	//language := struct {
	//	name   string
	//	age    int
	//	isGood bool
	//}{"Golang", 12, true}
	//fmt.Println(language)
	//newCustomer := customer{
	//	name: "John Doe",
	//	age:  42,
	//}
	//myOrder := order{
	//	id:     "1",
	//	amount: 42,
	//	status: "pending",
	// 	customer: newCustomer
	//}
	//
	//myOrder.createdAt = time.Now()
	//myOrder.customer.name = "sambit" // to change the customer name using struct embbeding
	//fmt.Println(myOrder.status)
	//fmt.Println(myOrder)

	// if you don't set any value, it will take the default values, which is
	// int--->0, float --> 0, string----> "", bool ---> false

	//myOrder2 := order{
	//	id:        "2",
	//	amount:    142,
	//	status:    "received",
	//	createdAt: time.Now(),
	//}
	//
	//myOrder2.changeStatus("delivered")
	//fmt.Println(myOrder2.getAmount())
	//fmt.Println(myOrder2)
	//newCustomer := customer{
	//	name: "John Doe",
	//	age:  42,
	//}
	//myOrder := newOrder("12345", 1.0, "Paid")
	//myOrder.customer = newCustomer
	//fmt.Println(*myOrder)            //{12345 1 Paid {13980015405125190560 166242 0x5f43480}}
	//fmt.Println(myOrder.getAmount()) // 1
	//fmt.Println(myOrder)             //&{12345 1 Paid {13980015405125190560 166242 0x5f43480}}

	d := Animal{
		name:  "Dog",
		sound: "Woof",
	}
	d.soundChange("bouuuu")
	d.Speak()
	fmt.Println(d)

}

/*
	Struct:- struct as a blueprint for creating objects, just like a class in OOP

				Go										Ruby/JS/Python
	1.  Struct											1. Class
	2. We can add behaviors via Methods with Receivers  2.Methods inside class
	3. Inheritance not supported						3. Inheritance Supported
	4.	Composition Supported 							4. Same as go

Declaring Struct
=================
	type Order struct {
		name 	string
		amount	 float64
		productStatus  string
	}

	//Setter method, now we can modify the struct fields
	func (o *Order) changeStatus(status string) {
		o.productStatus = status
	}

	//getter method
	func (o order) status() {
		fmt.Println("helloooo ", o.status)
	}

	p := Person { name: "Sambit", amount: 300.0, productStatus: "Delivered"}
	p.changeStatus("Paid")
	p.status() /// heloooo paid

		or

	var p Person
	p.name = "sambit"
	p.amount = 300.0
	p.productStatus = "Delivered"


*/
