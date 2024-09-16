package main

import "fmt"

func main() {
	var score = 80

	if score == 100 {
		fmt.Println("Perfect!")
	} else if score >= 90 {
		fmt.Println("Nice!")
	} else if score >= 80 {
		fmt.Println("Good!")
	} else {
		fmt.Println("Bad!")
	}

	// if statement with temporary variable
	if temp := score / 10; temp == 10 {
		fmt.Println("Perfect!")
	} else if temp >= 9 {
		fmt.Println("Nice!")
	} else if temp >= 8 {
		fmt.Println("Good!")
	} else {
		fmt.Println("Bad!")
	}

	// using switch
	var point = 9

	switch point {
	case 10:
		fmt.Println("Perfect!")
	case 9, 8, 7:
		fmt.Println("Nice!")
		fallthrough
	case 6, 5, 4:
		fmt.Println("Good!")
	default:
		fmt.Println("Bad!")
	}

	// nested selection condition
	var point2 = 3

	if point2 > 7 {
		switch point2 {
		case 10:
			fmt.Println("Perfect!")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if point2 == 5 {
			fmt.Println("Good!")
		} else {
			fmt.Println("Bad!")
		}
	}
}
