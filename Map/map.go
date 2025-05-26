package main

import (
	"fmt"
	"maps"
)

func main() {
	/*
		// creating Map/Hash
		// [string] ---> key type, string ---> value type
		// If key doesn't exist in the map, then it returns 0 value
		m := make(map[string]string)

		// Setting an element
		m["a"] = "a"
		m["name"] = "golang"

		//get an element
		fmt.Println(m["name"])

		//Length of a hash/map
		fmt.Println(len(m))

		//delete an element
		delete(m, "a")

		//Clear all value from map
		//clear(m)
		fmt.Println(m)
	*/

	// Declaring map without make
	ma := map[string]int{"weight": 40, "age": 30}
	fmt.Println(ma)

	////if the key exist in the map or not
	//value, ok := ma["weight"]
	//fmt.Println(value, ok)
	//if ok {
	//	fmt.Println("all ok")
	//} else {
	//	fmt.Println("not ok")
	//}

	// Comparing hash/map using maps package
	ma1 := map[string]int{"weight": 40, "age": 30}
	ma2 := map[string]int{"weight": 40, "age": 30}
	fmt.Println(maps.Equal(ma1, ma2))

}
