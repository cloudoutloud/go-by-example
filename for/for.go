package main

import "fmt"

func main() {

	// will print 1 to 5
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	// will print 7 to 9
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// will break loop with text defined
	for {
		fmt.Println("break loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}