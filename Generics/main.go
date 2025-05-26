package main

import "fmt"

// if we want to re-use print Slice functionality for string then we can use generic [T any] or we can add [T int | string | bool]
// If we want to scope the func "func PrintSLice[T any](items []T) "
// or instead of any we can use comparable
//comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types).

func PrintSLice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// we can pass generics in struct as well
type Stack[T any] struct {
	items []T
}

func main() {
	//items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//items := []string{"sa", "mb", "it"}
	myStack := Stack[string]{
		items: []string{"sa", "mb", "it"},
	}
	fmt.Println(myStack)
	//PrintSLice(myStack)
}
