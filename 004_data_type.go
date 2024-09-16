package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	var decimalNumber = 2.62
	var isExist bool = true
	var message string = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Printf("Positive Number: %d\n", positiveNumber)
	fmt.Printf("Negative Number: %d\n", negativeNumber)
	fmt.Println("Decimal Number:", decimalNumber)
	fmt.Println("Is Exist:", isExist)
	fmt.Println("Message:", message)
}
