package main

import "fmt"

func main() {
	//Now a is an array of 5 integers.
	// var arr [5]int
	// fmt.Println(arr[1])

	arr1 := [5]int{0, 2, 4, 6, 8}
	fmt.Println(arr1[1])

	fmt.Println(arr1[2:4])

	arr2 := make([]int, 5)
	fmt.Println(arr2)
	arr2 = append(arr2, 8)
	fmt.Println(arr2)
}
