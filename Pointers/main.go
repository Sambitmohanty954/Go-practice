package main

import "fmt"

// by value
//func changeNumber(number int) {
//	number = 5
//	fmt.Println("change number ", number)
//}

//by reference
// we need to pass * which is a pointer as params
// in this function number in println will show you the memory location, so to de-reference it you need to pass *, then only it will show you the value

func changeNumber(number *int) {
	*number = 5
	fmt.Println("change number ", *number)
}

// to see the memory address of the variable then we need to use "&"
//
//	A pointer holds the memory address of a value.
func main() {
	number := 1
	//changeNumber(number) // this function call will not going to change number value
	changeNumber(&number)

	fmt.Println("main function number", number)
	//fmt.Println("main function number", &number)

	// example --2
	i, _ := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
}
