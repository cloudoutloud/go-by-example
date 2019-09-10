package main

import "fmt"

func main() {

	// When 7 divided by 2 is 0, will print odd 
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// When 8 divied by 2 is 0, will print even
	if 8%4 == 0 {
		fmt.Println("8 is even")
	} else {
		fmt.Println("8 is odd")
	}

	if 10%5 == 0 {
		fmt.Println("10 is even")
	} else {
		fmt.Println("10 is odd")
	}

	if 20%10 == 0 {
		fmt.Println("20 is divisiable by 10")
	}

	// Will look at value "num" and provide outcome
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else{
		fmt.Println(num, "has multiple digits")
	}

}
