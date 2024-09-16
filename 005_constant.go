package main

import "fmt"

func main() {
	const a int = 10
	const b string = "golang"
	const c float64 = 3.14

	const (
		d          = 10
		e          = "golang"
		lightSpeed = 299792458
		pi         = 3.14159265359
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(lightSpeed)
	fmt.Println(pi)
}
