package main

import "fmt"

func main() {
	// deklarasi variabel string
	var firstName string = "John"
	var lastName string

	// deklasi type inference
	salutation := "Mr."

	// deklasi multi variabel
	var berat, tinggi int = 70, 170

	// membuat variable underscore
	_ = "belajar golang"

	// deklarasi menggunakan variable new
	var className = new(string)

	// menampilkan output ke layar
	lastName = "Doe"
	*className = "Golang"

	fmt.Printf("Hello %s %s %s!\n", salutation, firstName, lastName)
	fmt.Printf("Berat badan: %d kg\n", berat)
	fmt.Println("Tinggi badan:", tinggi, "cm")
	fmt.Println("Class Name:", *className)
	fmt.Println("Class Name:", className) //pointer
}
