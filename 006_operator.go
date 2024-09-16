package main

import "fmt"

func main() {
	var value = 10 + 15 - 5
	var isEqual = value == 20

	var left = false
	var right = true

	fmt.Println(value)
	fmt.Println(isEqual)
	fmt.Println(left && right)
	fmt.Println(left || right)
	fmt.Println(!left)
	fmt.Print(!right)
}
