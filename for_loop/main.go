package main

import "fmt"

func main() {
	//While loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// traditional for loop
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	/*
		infinite loop
		for {
			fmt.Println("hello")
		}
	*/

	cards := []string{"sambir", "sam2", "sam3"}
	// Go For loop using range
	for k, car := range cards {
		fmt.Println(k, car)
	}

	// if we don't want to use index, use _
	for _, car := range cards {
		fmt.Println(car)
	}

	// 1.22 range
	for i := range 10 {
		fmt.Println(i)
	}
}
