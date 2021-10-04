package main

import "fmt"
import "math"

func main() {

	var a = "first"
	fmt.Println(a)

	// multiple variables
	var b , c int = 1,2
	fmt.Println(b,c)

	var d = "false"
	fmt.Println(d)

	// int will always return 0
	var e int
	fmt.Println(e)

	f := "sean rigbyyyyy"
	fmt.Println(f)

	var g = "sean rigby"
	fmt.Println(g)

	// := can only be used in a func body, := can be used with var to left is newly defined
	x, y := 145.8, 348.8
	z := math.Min(x, y)
	fmt.Println("min value is", z)

}



