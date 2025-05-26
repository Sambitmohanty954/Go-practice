package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count += 1
		return count
	}
}

//A closure is a function that has access to the parent scope, after the parent function has closed.
//Closures has historically been used to:
//Create private variables
//Preserve state between function calls
//Simulate block-scoping before let and const existed
//Implement certain design patterns like currying and memoization

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
