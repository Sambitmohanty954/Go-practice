package main

import "fmt"

func main() {
	//simple switch

	//i := 5
	//switch i {
	//case 1:
	//	fmt.Println("one")
	//case 2:
	//	fmt.Println("two")
	//case 3:
	//	fmt.Println("three")
	//default:
	//	fmt.Println("Other")
	//}

	//Multiple condition switch
	//switch time.Now().Weekday() {
	//case time.Saturday, time.Sunday:
	//	fmt.Println("It's weekend", time.Now().Weekday())
	//default:
	//	fmt.Println("It's a weekday")
	//}

	//type switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		case float64:
			fmt.Println("I'm a float64")
		default:
			fmt.Println("I'm an interface", t)
		}
	}

	whoAmI(3.13)
}
