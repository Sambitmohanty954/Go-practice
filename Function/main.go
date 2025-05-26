package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// We can return multiple values
func getLanguages() (string, string, int) {
	return "Go", "C", 3
}

// Variadic function i can pass n numbers of params, rest/spread operator
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	addition := add(3, 5)
	fmt.Println(addition)
	l1, l2, _ := getLanguages()
	fmt.Println(l1, l2)
	fmt.Println(getLanguages())

	// here we can pass n number of params
	nums := []int{1, 2, 3, 4, 5}
	// result := sum(1,2,3,4,5,6)
	result := sum(nums...) // we can call like this as well
	fmt.Println(result)
}
