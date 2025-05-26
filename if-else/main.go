package main

import "fmt"

func main() {
	//age := 13
	//if age >= 18 {
	//	fmt.Println("Person is adult")
	//} else if age >= 12 || age < 18 {
	//	fmt.Println("Person is teenage")
	//} else {
	//	fmt.Println("Person is kid")
	//}

	//Go doesn't have ternary operator, you will have to use normal if-else
	if age := 15; age <= 18 {
		fmt.Println("Person is teenage")
	}
}
