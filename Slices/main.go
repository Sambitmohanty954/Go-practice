package main

import "fmt"

// Slices ---> Dynamic
// Most used constructor in go
func main() {
	// unintialized slice is nil
	//var nums []int
	//fmt.Println(nums == nil) //true
	//fmt.Println(len(nums) == 0) //true

	// Here [0,0,0] created with 5 capacity
	// Maximum numbers an element can fit
	// we can check cap fmt.Println(cap(nums))
	//var nums = make([]int, 3, 5) //we can declare like this as well
	//var nums = []int{1} // we can declare slice like this as well
	//fmt.Println(nums)
	//nums = append(nums, 1) // append a value
	//nums[1] = 2
	//nums = append(nums, 4, 5, 6, 7, 8, 9, 10) // append multiple values
	//fmt.Println(nums)
	//fmt.Println(cap(nums))

	//copy function
	//var nums = make([]int, 0, 5)
	//nums = append(nums, 1, 2, 3, 4, 5, 5)
	//var nums2 = make([]int, len(nums))
	//copy(nums2, nums)
	//fmt.Println(nums, nums2)

	// slice operator
	//var nums = []int{1, 2, 3, 4, 5}
	//fmt.Println(nums[3:]) // slice operator output 4,5
	//fmt.Println(nums[:])  // 1,2,3,4,5
	//fmt.Println(nums[:3]) // 1,2,3, 0 to number of element from 0

	// compare slices (dynamic array) slices package
	//var nums1 = []int{1, 2, 3, 4, 5}
	//var nums2 = []int{1, 2, 3, 4, 5}
	//fmt.Println(slices.Equal(nums1, nums2))

	// 2D array in Slices
	var nums = [][]int{{1, 2, 3, 4}, {2, 3, 4, 5}}
	fmt.Println(nums)
}
