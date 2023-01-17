package main

import "fmt"

func main() {
	arr := make([]int, 5)

	for i, item := range arr {
		fmt.Println(i, "->", item)
	}

	x := "hello"
	for _, c := range x {
		fmt.Println(c)
	}

	for _, c := range x {
		fmt.Printf("%c ", c)
	}
}
